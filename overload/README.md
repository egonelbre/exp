POC: An example how to quickly overload operators using [eg](https://godoc.org/golang.org/x/tools/refactor/eg).

Run with:

    eg -t add.template example.go

And then replace template import.

*When you want to use something like this a lot in practice, you probably should implement it properly, instead of hacking it together from things that weren't meant for it.*