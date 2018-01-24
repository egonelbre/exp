Package benchmark contains different benchmarks for compiler optimizations.

*audio* shows two common things you might want to do in audio libraries:

* Generate a sound for different number of channel counts:
    - _Dynamic_ benchmarks represents code where the code has least repetition
    - _Switch_ benchmarks shows where each channel count has different implementation. This becomes extremely repetitive, if you need to do this for each effect and buffer type.
    - _Baseline_ benchmarks shows the ideal performance (ignoring SIMD optimizations)
* Mix two sound buffers together:
    - _Dynamic_ benchmark shows the code you would usually write in Go
    - _Baseline_ benchmarks shows the ideal performance (ignoring SIMD optimizations)