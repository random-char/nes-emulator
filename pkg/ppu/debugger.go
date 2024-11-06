package ppu

import "nes-emulator/pkg/debugger"

func (ppu *Ricoh2c02) SetDebugger(debugger *debugger.Debugger) {
	ppu.debugger = debugger
}
