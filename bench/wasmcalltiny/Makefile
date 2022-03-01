
wasm:
	docker run --platform linux/amd64 --rm -v $(pwd):/src -w /src tinygo/tinygo:0.22.0 tinygo build -o wasm.wasm -target=wasm --no-debug ./main.go
	docker run --platform linux/amd64 --rm -v $(pwd):/src -w /src tinygo/tinygo:0.22.0 /bin/bash -c "cp /usr/local/tinygo/targets/wasm_exec.js wasm_exec.js"
