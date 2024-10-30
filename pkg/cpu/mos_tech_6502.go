package cpu

import "nes-emulator/pkg/debugger"

type MOSTechnology6502 struct {
	//internals
	a      uint8  // accumulator
	x      uint8  // x
	y      uint8  // y
	stkp   uint8  // stack pointer
	pc     uint16 // program counter
	status uint8  // status

	fetched uint8

	addrAbs uint16
	addrRel uint16

	opcode     uint8
	cycles     uint8
	clockCount uint16

	bus      bus
	debugger *debugger.Debugger
}

func New(
	bus bus,
	debugger *debugger.Debugger,
) *MOSTechnology6502 {
	cpu := &MOSTechnology6502{
		a: InitialState.a,
		x: InitialState.x,
		y: InitialState.y,

		stkp:   InitialState.stkp,
		pc:     InitialState.pc,
		status: InitialState.status,

		fetched: 0,

		addrAbs: 0,
		addrRel: 0,

		opcode:     0x00,
		cycles:     0,
		clockCount: 0,

		bus:      bus,
		debugger: debugger,
	}

	initLookupFor(cpu)

	return cpu
}

func (cpu *MOSTechnology6502) GetDebugger() *debugger.Debugger {
	return cpu.debugger
}

func (cpu *MOSTechnology6502) fetch() uint8 {
	// can't compare functions
	if lookup[cpu.opcode].AMName != am_IMP {
		cpu.fetched = cpu.read(cpu.addrAbs)
	}

	return cpu.fetched
}
