copy-wasm-go:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/public/assets/js

build-wasm:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o nes.wasm ./wasm/.
	mv nes.wasm web/public/assets/wasm/

copy-wasm-tinygo:
	cp "$$(tinygo env TINYGOROOT)/targets/wasm_exec.js" web/public/assets/js/tinygo_wasm_exec.js

build-wasm-tinygo:
	GOOS=js GOARCH=wasm tinygo build -o nes_tinygo.wasm ./wasm/.
	mv nes_tinygo.wasm web/public/assets/wasm/
