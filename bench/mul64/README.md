# mul64

Microbenchmarks comparing `bits.Mul64(x, K)` (a 64×64 → 128 unsigned multiply)
against equivalent shift/add/sub sequences for constants `K` with a sparse
signed-digit form. The point is to decide, on a given CPU, whether a compiler
should strength-reduce `bits.Mul64(x, K)` to shifts.

Motivation: word-by-word Montgomery reduction by NIST/Solinas moduli (P-256,
P-384, P-521) calls `bits.Mul64(x, K)` many times with sparse `K` such as
`0xFFFFFFFFFFFFFFFF` and `0xFFFFFFFF00000000`. Replacing these calls with
shifts and a borrow sequence (a "strength reduction") can save MUL throughput
but adds ALU pressure and serial carry-chain latency. Whether it wins is
microarchitecture-dependent.

## Running

```
go test ./...                              # verify the shift variants match
go test -bench . -benchtime=2s -count=6    # bench all variants
```

For comparing variants, redirect to a file and use `benchstat`:

```
go test -bench BenchmarkLatency -count=10 > latency.txt
benchstat -filter '.name:/Mul$' latency.txt latency.txt
```

(or just eyeball the `*/Mul` vs `*/Sft` pairs).

## What's measured

The bench file runs each variant in three regimes:

- **`BenchmarkLatency`** — each iteration's input depends on the previous
  iteration's output. Measures the *critical-path latency* of one
  mul-or-shift sequence. Relevant when calls cannot be parallelized by the CPU.

- **`BenchmarkThroughput`** — 8 independent calls per iteration with no
  carried dependency. Measures *throughput* — how many sequences per cycle
  the backend can sustain. Relevant when many multiplications happen back-to-back
  (limb-by-limb field/scalar arithmetic).

- **`BenchmarkReduction`** — 4 sequentially dependent calls per iteration with
  an interleaved carry-fold. Closer to the shape of one Montgomery reduction
  round.

## Adding patterns

Append entries to the `variants` table in `mul64.go` and add the corresponding
`mul<Name>` / `sft<Name>` functions. The benchmark runners pick them up
automatically.

## Interpreting results

A real win shows up as `*/Sft` being faster than `*/Mul` across all three
regimes — or at least the regime that matches your real workload. If the
results are tied, strength reduction is unlikely to help (and may hurt by
adding scheduling pressure to surrounding code).

### Apple M4 Max (Go 1.27 dev, 2026-05)

```
Latency  AllOnes:        Mul 1.14 ns/op   Sft 1.17 ns/op   ~tied
Latency  NegPow32:       Mul 1.40         Sft 1.41         ~tied
Latency  Pow32:          Mul 0.76         Sft 0.75         ~tied
Latency  PlusOne:        Mul 1.42         Sft 1.42         ~tied
Throughput AllOnes:      Mul 3.57 ns/8op  Sft 3.58         ~tied
Throughput NegPow32:     Mul 4.16 ns/8op  Sft 4.22         Mul wins ~1%
```

On M4 Max none of these patterns are a meaningful win — MUL+UMULH issues fast
enough through the dedicated multiplier units that swapping it for 3–5 ALU ops
is essentially a wash. The slight P-256 ord speedup observed in `crypto/ecdsa`
benchmarks from a source-level rewrite comes from *scheduling effects*
(fewer multipliers in flight at once gives the OoO engine more room),
not from saved cycles in the rewritten instructions themselves.

Other architectures (notably amd64 with its 3-cycle, 1-per-cycle MUL throughput,
and in-order ARMv7/MIPS cores) may tip the balance the other way. Run the
bench on the target hardware to find out.
