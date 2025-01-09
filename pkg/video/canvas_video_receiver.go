//go:build wasm
// +build wasm

package video

import (
	"nes-emulator/pkg/visuals"
	"syscall/js"
)

const (
	gameWidth  = 256
	gameHeight = 240

	debugWidth  = 256
	debugHeight = 128
)

type CanvasVideoReceiver struct {
	ctx     js.Value
	imgData js.Value

	patternTablesCtx     js.Value
	patternTable1ImgData js.Value
	patternTable2ImgData js.Value
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

		patternTablesCtx:     debugCtx,
		patternTable1ImgData: debugCtx.Call("createImageData", debugWidth/2, debugHeight),
		patternTable2ImgData: debugCtx.Call("createImageData", debugWidth/2, debugHeight),
	}
}

func (cvr *CanvasVideoReceiver) Render(pixelData []uint8) {
	data := cvr.imgData.Get("data")

	js.CopyBytesToJS(data, pixelData)

	cvr.ctx.Call("putImageData", cvr.imgData, 0, 0)
}

func (cvr *CanvasVideoReceiver) RenderPatternTables(i uint16, sprite *visuals.Sprite) {
	var imgData js.Value
	if i == 0 {
		imgData = cvr.patternTable1ImgData.Get("data")
	} else {
		imgData = cvr.patternTable2ImgData.Get("data")
	}

	js.CopyBytesToJS(imgData, sprite.GetPixelsData())

	if i == 0 {
		cvr.patternTablesCtx.Call("putImageData", cvr.patternTable1ImgData, 0, 0)
	} else {
		cvr.patternTablesCtx.Call("putImageData", cvr.patternTable2ImgData, debugWidth/2, 0)
	}
}
