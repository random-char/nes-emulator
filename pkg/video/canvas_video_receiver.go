//go:build wasm
// +build wasm

package video

import (
	"nes-emulator/pkg/ppu/visuals"
	"syscall/js"
)

const (
	gameWidth  = 256
	gameHeight = 240

	debugWidth  = 128
	debugHeight = 256
)

type CanvasVideoReceiver struct {
	ctx     js.Value
	imgData js.Value

	debugCtx     js.Value
	debugImgData js.Value
}

func NewCanvasVideoReceiver(
	canvas js.Value,
	debugCanvas js.Value,
) *CanvasVideoReceiver {
	ctx := canvas.Call(
		"getContext",
		"2d",
	)

	debugCtx := debugCanvas.Call(
		"getContext",
		"2d",
	)

	return &CanvasVideoReceiver{
		ctx:     ctx,
		imgData: ctx.Call("createImageData", gameWidth, gameHeight),

		debugCtx:     debugCtx,
		debugImgData: debugCtx.Call("createImageData", debugWidth, debugHeight),
	}
}

func (cvr *CanvasVideoReceiver) RenderFrame(pixels []*visuals.Pixel) {
	data := cvr.imgData.Get("data")

	currentIndex := 0
	for _, p := range pixels {
		data.SetIndex(currentIndex, p.R)
		data.SetIndex(currentIndex+1, p.G)
		data.SetIndex(currentIndex+2, p.B)
		data.SetIndex(currentIndex+3, 255)
		currentIndex += 4
	}

	cvr.ctx.Call(
		"putImageData",
		cvr.imgData,
		0, 0,
	)
}

func (cvr *CanvasVideoReceiver) RenderPatternTables(i int, sprite *visuals.Sprite) {
	data := cvr.debugImgData.Get("data")

	var p *visuals.Pixel
	currentIndex := i*(128*128)
	var x, y int16
	for x = 0; x < sprite.GetWidth(); x++ {
		for y = 0; y < sprite.GetHeight(); y++ {
			p = sprite.GetPixel(x, y)

			data.SetIndex(currentIndex, p.R)
			data.SetIndex(currentIndex+1, p.G)
			data.SetIndex(currentIndex+2, p.B)
			data.SetIndex(currentIndex+3, 255)

			currentIndex += 4
		}
	}

	cvr.debugCtx.Call(
		"putImageData",
		cvr.debugImgData,
		0, 0,
	)
}
