copy-wasm-go:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/assets/js

build-wasm:
	GOOS=js GOARCH=wasm go build -o nes ./wasm/.
	mv nes web/public/assets/wasm/nes.wasm
