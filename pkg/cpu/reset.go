package cpu

func (cpu *MOSTechnology6502) Reset() {
	if InitialState.fromSpecificPc {
		cpu.pc = InitialState.pc
	} else {
		cpu.addrAbs = 0xFFFC

		lo := uint16(cpu.read(cpu.addrAbs))
		hi := uint16(cpu.read(cpu.addrAbs + 1))

		cpu.pc = (hi << 8) | lo
	}

	cpu.a = InitialState.a
	cpu.x = InitialState.x
	cpu.y = InitialState.y
	cpu.stkp = InitialState.stkp
	cpu.status = InitialState.status

	cpu.addrRel = 0
	cpu.addrAbs = 0
	cpu.fetched = 0

	cpu.cycles = 8
}
