package ppu

import "nes-emulator/pkg/visuals"

type VideoReceiver interface {
	RenderFrame([]*visuals.Pixel)
}
