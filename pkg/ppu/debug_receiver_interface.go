package ppu

import "nes-emulator/pkg/visuals"

type DebugReceiver interface {
	RenderPatternTables(uint16, *visuals.Sprite)
}
