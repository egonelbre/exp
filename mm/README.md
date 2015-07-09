Experimental memory-allocator generator, based on https://www.youtube.com/watch?v=mCrVYYlFTrA.

**NB: This is only a proof-of-concept implementation, nothing else. It has bugs, problems, it could be smarter, not as nice to use etc... So don't use this in production nor anywhere... it's just a proof-of-concept, nothing else.**

To generate an example: `go run gen.go`.

*Notes: It rarely makes sense to use such custom-allocators in Go. If you need to use it, you might be better off implementing that thing in some other language (e.g. C++, D, Rust, C etc.).*