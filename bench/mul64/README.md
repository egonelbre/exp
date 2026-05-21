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

The bench file runs each variant in several regimes, from least to most
representative of real code:

- **`BenchmarkLatency`** — each iteration's input depends on the previous.
  Measures the *critical-path latency* of one mul-or-shift sequence.
  Operations measured in isolation. Easy to interpret but rarely matches
  real-code behavior.

- **`BenchmarkThroughput`** — 8 independent calls per iteration. Measures
  *throughput* of one sequence type when surrounded by copies of itself.

- **`BenchmarkReduction`** — 4 sequentially-dependent calls with an
  interleaved carry-fold. Closer to one Montgomery reduction round, but still
  homogeneous (all calls use the same constant).

- **`BenchmarkKernel_ReductionRound_*`** — *the most informative test.*
  One full reduction round mixing four mul-by-constants (two dense, two
  strength-reducible), four add-with-carry chains, and a `mu` computation.
  Models how the sparse muls compete with other work for execution units.

- **`BenchmarkKernel_Inline_ReductionRound_*`** — same kernel, but with the
  whole round in the bench loop body (no function call). Strips the function-
  call boundary, which on its own changes the answer (see "Inlining changes
  the verdict" below).

- **`BenchmarkKernel_FourRoundsParallel_*`** — four independent reduction
  rounds in one call. Throughput-saturation case for the kernel.

- **`BenchmarkOverhead_*` / `BenchmarkInline_Latency_*`** — calibration
  benchmarks for measuring CALL+RET cost and the inline-vs-callback gap.

## Adding patterns

Append entries to the `variants` table in `mul64.go` and add the corresponding
`mul<Name>` / `sft<Name>` functions. The benchmark runners pick them up
automatically.

## Interpreting results

A real win shows up as `*/Sft` being faster than `*/Mul` across all three
regimes — or at least the regime that matches your real workload. If the
results are tied, strength reduction is unlikely to help (and may hurt by
adding scheduling pressure to surrounding code).

### Harness overhead

The variants are `//go:noinline` and dispatched through a function-pointer
slice, both of which add overhead. To quantify:

```
go test -bench Overhead -count=5    # CALL+RET cost
go test -bench Inline   -count=5    # same ops, no callback, no noinline
```

On Apple M4 Max, direct CALL+RET costs ~0.57 ns. For ops whose body is ≥1 ns
(e.g. anything involving MUL+UMULH or a borrow chain), the out-of-order
engine hides this call cost completely — the indirect-callback bench gives
the same result as the inline version. For *very fast* ops (e.g. a pair of
shifts, ~0.46 ns inline), the harness inflates the measurement by ~60%.

The relative comparison `Mul vs Sft` for the *same* constant is preserved
either way (both variants pay the same overhead). But cross-constant
comparisons and any "is this op actually free?" question need the inline
benchmarks.

### Inlining changes the verdict

Two kernels with the *same* operation count and dependency structure can
give opposite answers depending on whether the round is inlined into the
surrounding loop or kept in its own function. On Apple M4 Max:

```
Kernel_ReductionRound        Mul 3.40 ns  Sft 3.49 ns   Sft +3% slower
Kernel_Inline_ReductionRound Mul 3.28 ns  Sft 3.18 ns   Sft -3% faster
Kernel_FourRoundsParallel    Mul 7.07 ns  Sft 7.21 ns   Sft +2% slower
```

Why? The strength-reduced variant produces more live values (the shift
result, the borrow, and the second shift) for each mul-by-constant. When
the round is in its own function, those temporaries pressure a fixed
register window and the regalloc cost outweighs the saved multiplier
cycles. When inlined, the temporaries mix with the surrounding loop state
and the compiler can interleave the strength-reduction ops with the
mul-by-non-constant work — the multiplier units stay busy with the
non-strength-reducible muls in parallel.

This matches what we see in the standard library: replacing
`bits.Mul64(x, K)` calls inline inside `p256OrdMul` (one big function) is a
modest win on P-256 ord, while the same SSA-level rewrite on the larger
`p384OrdMul` (more parallel sparse muls in flight, multipliers already
saturated) is a small loss.

### Apple M4 Max (Go 1.27 dev, 2026-05) summary

```
Isolated latency  AllOnes:        Mul 1.14 ns   Sft 1.17 ns   ~tied
Isolated latency  NegPow32:       Mul 1.40      Sft 1.41      ~tied
Isolated latency  Pow32:          Mul 0.76      Sft 0.75      ~tied
Isolated throughput (x8 NegPow32): Mul 4.16     Sft 4.22      Mul +1%

Kernel ReductionRound:                       Sft +3% slower (regalloc pressure)
Kernel Inline_ReductionRound:                Sft -3% faster (ILP with surround)
Kernel FourRoundsParallel:                   Sft +2% slower (multiplier not bottleneck)
```

**Conclusion for M4 Max:** strength reduction is workload-dependent, with
the win/loss governed by surrounding code's structure (inlined vs not,
multiplier-saturated vs not), not the cost of the strength-reduced
instructions themselves. Other architectures — notably amd64 (3-cycle,
1-per-cycle MUL) and in-order ARMv7/MIPS cores — likely tip more uniformly
toward "always worth it." Run the bench on the target hardware to find out.
