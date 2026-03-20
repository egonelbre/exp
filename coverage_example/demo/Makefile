.PHONY: test
test:
	go test -coverprofile demo.coverprofile -coverpkg "./...,ext,ext/sub"

.PHONY: covdata
covdata:
	mkdir -p covdata
	go test -c -cover -coverpkg "./...,ext,ext/sub" -o covdata/test.bin
	cd covdata && GOCOVERDIR=. ./test.bin -test.v
	rm covdata/test.bin
