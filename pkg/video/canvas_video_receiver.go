//go:build wasm
// +build wasm

package video

import (
	"nes-emulator/pkg/ppu"
	"syscall/js"
)

type CanvasVideoReceiver struct {
	ctx js.Value

	pixels [][]*ppu.Pixel
}

func NewCanvasVideoReceiver(canvas js.Value) *CanvasVideoReceiver {
	ctx := canvas.Call(
		"getContext",
		"webgl",
	)

	ctx.Call("clearColor", 0.0, 0.0, 0.0, 1.0)
	ctx.Call("clear", ctx.Get("COLOR_BUFFER_BIT"))

	pixels := make([][]*ppu.Pixel, 341)
	for i := 0; i < 341; i++ {
		pixels[i] = make([]*ppu.Pixel, 261)
	}

	return &CanvasVideoReceiver{
		ctx:    ctx,
		pixels: pixels,
	}
}

func (cvr *CanvasVideoReceiver) SetPixel(x, y int16, pixel *ppu.Pixel) {
	if y < 0 {
		return
	}

	cvr.pixels[x][y] = pixel
}

func (cvr *CanvasVideoReceiver) DrawFrame() {
	// var pixel *ppu.Pixel
	// for x := 0; x < 341; x++ {
	// 	for y := 0; y < 261; y++ {
	// 		pixel = cvr.pixels[x][y]
	// 	}
	// }
}
