package cpu

type Olc6502 struct {
	//registers
	a uint8 // accumulator
	x uint8 // x
	y uint8 // y

	stkp   uint8  // stack pointer
	pc     uint16 // program counter
	status uint8  // status

	fetched uint8

	addrAbs uint16
	addrRel uint16

	opcode uint8
	cycles uint8

	bus bus
}

func New(bus bus) *Olc6502 {
	cpu := &Olc6502{
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

	initLookupFor(cpu)

	return cpu
}

func (cpu *Olc6502) fetch() uint8 {
	// can't compare functions
	if lookup[cpu.opcode].AMName != ADDR_MODE_IMP {
		cpu.fetched = cpu.read(cpu.addrAbs)
	}

	return cpu.fetched
}
