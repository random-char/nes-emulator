package cpu

import "nes-emulator/pkg/device"

const (
	FLAGS_6502_C = uint8(1 << 0) // carry bit
	FLAGS_6502_Z = uint8(1 << 1) // zero
	FLAGS_6502_I = uint8(1 << 2) // disable interrupts
	FLAGS_6502_D = uint8(1 << 3) // decimal mode
	FLAGS_6502_B = uint8(1 << 4) // break
	FLAGS_6502_U = uint8(1 << 5) // unused
	FLAGS_6502_V = uint8(1 << 6) // overflow
	FLAGS_6502_N = uint8(1 << 7) // negative
)

var lookup []Instruction

type Olc6502 struct {
	//registers
	a      uint8  // accumulator
	x      uint8  // x
	y      uint8  // y
	stkp   uint8  // stack pointer
	pc     uint16 // program counter
	status uint8  // status

	fetched uint8

	addrAbs uint16
	addrRel uint16

	opcode uint8
	cycles uint8

	bus devices.Bus
}

func New(bus devices.Bus) *Olc6502 {
	olc := &Olc6502{
		a:      0,
		x:      0,
		y:      0,
		stkp:   0,
		pc:     0,
		status: 0,

		fetched: 0,

		addrAbs: 0,
		addrRel: 0,

		opcode: 0,
		cycles: 0,

		bus: bus,
	}

	lookup = newLookupFor(olc)

	return olc
}

func (olc *Olc6502) Read(addr uint16) uint8 {
	return olc.bus.Read(addr, true)
}

func (olc *Olc6502) Write(addr uint16, data uint8) {
	olc.bus.Write(addr, data)
}

func (olc *Olc6502) SetFlag(flag uint8, value bool) {
	if value {
		olc.status |= flag
	} else {
		olc.status &= ^flag
	}
}

func (olc *Olc6502) GetFlag(flag uint8) uint8 {
	if olc.status&flag > 0 {
		return 1
	} else {
		return 0
	}
}

func (olc *Olc6502) Fetch() uint8 {
	// can't compare functions
	if lookup[olc.opcode].AddrModeName != "IMP" {
		olc.fetched = olc.Read(olc.addrAbs)
	}

	return olc.fetched
}
