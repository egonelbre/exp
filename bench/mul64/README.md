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

> **IMPORTANT**: do not benchmark with a Go toolchain that already has the
> strength reduction you're trying to evaluate baked in (e.g. an in-progress
> SSA change). Use a stock release that does not yet apply intrinsic-time SR
> to `bits.Mul64`. On systems where multiple toolchains coexist, force the
> system install with `GOTOOLCHAIN=local`.

```
GOTOOLCHAIN=local go test ./...                            # correctness first
GOTOOLCHAIN=local go test -bench . -benchtime=2s -count=6
```

For comparing variants, redirect to a file and use `benchstat`:

```
GOTOOLCHAIN=local go test -bench BenchmarkLatency -count=10 > latency.txt
benchstat -filter '.name:/Mul$' latency.txt latency.txt
```

(or just eyeball the `*/Raw` vs `*/Mul` vs `*/Sft` triples).

## The compiler already does partial SR

Go 1.26 emits these patterns *inline* for `bits.Mul64(x, K)` with `K` a
compile-time constant — without any additional SSA work:

| K | What Go 1.26 emits |
|---|---|
| `0xFFFFFFFFFFFFFFFF` | `NEG` + `UMULH x, -1` |
| `0xFFFFFFFF00000000` | `NEG` + `LSL` + `UMULH` |
| `0xFFFFFFFFFFFFFFFE` | `NEG` + `ADD` + `UMULH` |
| `1<<32` | `LSL` + `UMULH` |
| `0x100000001` | `ADD x<<32, x` + `UMULH` |
| `0xFFFFFFFF` | `LSL` + `SUB` + `UMULH` |

So the `mul<Name>` variants in this bench are *not* a pure-MUL baseline — they
already get partial SR (lo via shifts/adds, hi still via `UMULH`). To get a
real MUL+UMULH baseline use the `raw<Name>` / `BenchmarkRaw_*` variants, which
route the constant through a `//go:noinline` parameter so the compiler can't
fold it.

This gives a three-way comparison:

- **Raw**: full `MUL`+`UMULH` (no compiler optimization at the call site).
- **Mul**: `bits.Mul64` with the constant inline (Go 1.26 partial SR applies).
- **Sft**: explicit full SR (shifts/adds/subs, no multiplier).

The relevant question for an SSA-level full-SR change is `Mul vs Sft` — does
going beyond the existing partial SR help? `Raw vs Mul` shows how much the
existing partial SR is already buying you.

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

### Apple M4 Max (Go 1.26 stock, 2026-05) — full sweep

All numbers ns/op, average of 6 runs at `-benchtime=2s`.

Isolated latency (table-driven harness, indirect call):

| Pattern | Raw (true MUL+UMULH) | Mul (partial SR) | Sft (full SR) |
|---|---:|---:|---:|
| AllOnes | 1.26 | 1.49 | 1.64 |
| NegPow32 | 1.24 | 1.67 | 1.81 |
| NegTwo | 1.26 | 1.66 | 1.81 |
| Pow32 | 1.26 | 1.48 | **1.01** ← Sft wins |
| PlusOne | 1.26 | 1.49 | 1.49 |
| Pow32Minus1 | 1.24 | (—) | (—) |

(`Raw` looks fastest because it uses a *direct* call to `mulRaw`, while
`Mul`/`Sft` go through an indirect callback — see "Harness overhead" above.
The `Mul vs Sft` comparison is fair; the `Raw` column shows the baseline
without any SR.)

Realistic kernels (one full Montgomery reduction round, function-call):

| Kernel | Mul | Sft | Sft vs Mul |
|---|---:|---:|---:|
| ReductionRound (P-256, 2/4 sparse) | 3.39 | 3.44 | +1.5% slower |
| `Inline_`ReductionRound | 3.16 | 3.15 | ~tied |
| FourRoundsParallel | 6.65 | 7.06 | **+6.1% slower** |
| ChainedMontMul (4 sequential) | 57.6 | 57.8 | +0.3% |
| P384Round (6/6 sparse) | 8.51 | 8.81 | **+3.5% slower** |

Density sweep (4-limb kernel with N of 4 sparse multiplications):

| Sparse density | Mul | Sft | Sft vs Mul |
|---|---:|---:|---:|
| S1 (1/4) | 11.42 | 11.53 | +1.0% |
| S2 (2/4) | 11.54 | 11.59 | +0.5% |
| S3 (3/4) | 11.57 | 11.57 | ~tied |
| S4 (4/4) | 11.61 | 11.67 | +0.5% |

### Findings (M4 Max)

1. **Stock Go's partial SR is already quite good.** For `bits.Mul64(x, K)`
   with sparse `K`, it emits ~3 instructions (one `NEG`/`LSL`/`ADD` + one
   `UMULH`) — saving one MUL but keeping one to compute the high half.

2. **Going from partial SR to full SR (Mul→Sft) is rarely a win.** It's
   ~tied or slightly slower for 4-limb kernels, and meaningfully slower
   (+3-6%) for kernels with high sparse density and parallel rounds —
   matching the observed regression on `Sign/P384` and `Sign/P521`.

3. **The one clear full-SR win is `K = 2^s` (pure power of two).** Stock Go
   still emits `LSL + UMULH` here, while just `LSL + LSR` would give the
   correct (hi, lo) without using the multiplier at all. This is a tiny
   peephole opportunity, not a general SR.

4. **Parallel-rounds kernels are where full SR hurts most.** Replacing the
   `UMULH` (which uses the dedicated multiplier units) with ALU ops removes
   work from a unit with spare capacity and adds it to a unit that's
   already busy with the dense muls and carry chains.

### Other architectures

Run the bench on the target hardware to find out. amd64 (3-cycle, 1-per-cycle
MUL) and in-order ARMv7 / Cortex-A53 / MIPS cores may tip more uniformly
toward full SR being a win, but the data above suggests that **on any modern
out-of-order CPU with a dedicated multiplier unit, the existing partial SR
captures most of the available win and full SR adds risk without much
upside.**

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
