package cpu

import "nes-emulator/pkg/debugger"

func (cpu *MOSTechnology6502) SetDebugger(debugger *debugger.Debugger) {
	cpu.debugger = debugger
}
