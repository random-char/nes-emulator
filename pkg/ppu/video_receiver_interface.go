package ppu

import "nes-emulator/pkg/ppu/visuals"

type VideoReceiver interface {
	RenderFrame([]*visuals.Pixel)
}

type DebugReceiver interface {
    RenderPatternTables(int, *visuals.Sprite)
}
