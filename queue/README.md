# queue

Benchmarks and test of different queues.

Each queue has been named in the following way: `[SM]P[SM]C[sw]?i?p?[rnacq]_<variant>`.

* `[SM]P`: supports either single `S` or multiple `M` concurrent producers
* `[SM]C`: supports either single `S` or multiple `M` concurrent consumers
* `[sw]?`: `s` when it is a spinning implementation, `w` when it is partially spinning and partially waiting
* `i?`: `i` when it is an intrusive implementation
* `p?`: `p` when values are padded in the internal queue
* `[rnacq]`: internal implementation, `r` dynamically sized ring buffer, `n` node based, `a` fixed size array based, `c` channel based, `q` dynamically sized ring buffer with sequence number.
* `<variant>`: special variant identifier for a particular implementation, which indicates either base implementation author / paper / code.

Filenames are in a different order: `<variant>_[rnacq][SM]P[SM]Cs?i?p?`.