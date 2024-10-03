package cpu

func (cpu *Olc6502) Clock() {
	if cpu.cycles == 0 {
		opcode := cpu.read(cpu.pc)
		cpu.pc++

		cpu.cycles = lookup[opcode].Cycles
		additionalCycle1 := lookup[opcode].AddrMode()
		additionalCycle2 := lookup[opcode].Op()

		cpu.cycles += (additionalCycle1 & additionalCycle2)
	}

	cpu.cycles--
}

func (cpu *Olc6502) Reset() {
	cpu.a = 0
	cpu.x = 0
	cpu.y = 0

	cpu.stkp = 0xFD
	cpu.status = 0x00 | FLAG_U

	cpu.addrAbs = 0xFFFC

	lo := uint16(cpu.read(cpu.addrAbs))
	hi := uint16(cpu.read(cpu.addrAbs + 1))

	cpu.pc = (hi << 8) | lo

	cpu.addrRel = 0
	cpu.addrAbs = 0
	cpu.fetched = 0

	cpu.cycles = 8
}

func (cpu *Olc6502) Irq() {
	if cpu.getFlag(FLAG_I) == 1 {
		return
	}

	cpu.write(
		0x0100+uint16(cpu.stkp),
		uint8((cpu.pc>>8)&0x00FF),
	)
    cpu.stkp--

	cpu.write(
		0x0100+uint16(cpu.stkp),
		uint8(cpu.pc&0x00FF),
	)
    cpu.stkp--

    cpu.setFlag(FLAG_B, false)
    cpu.setFlag(FLAG_U, true)
    cpu.setFlag(FLAG_I, true)
    cpu.write(0x0100 + uint16(cpu.stkp), cpu.status)
    cpu.stkp--

    cpu.addrAbs = 0xFFFE
    lo := uint16(cpu.read(cpu.addrAbs))
    hi := uint16(cpu.read(cpu.addrAbs+1))

    cpu.pc = (hi << 8) | lo

    cpu.cycles = 7
}

func (cpu *Olc6502) Nmi() {
	cpu.write(
		0x0100+uint16(cpu.stkp),
		uint8((cpu.pc>>8)&0x00FF),
	)
    cpu.stkp--

	cpu.write(
		0x0100+uint16(cpu.stkp),
		uint8(cpu.pc&0x00FF),
	)
    cpu.stkp--

    cpu.setFlag(FLAG_B, false)
    cpu.setFlag(FLAG_U, true)
    cpu.setFlag(FLAG_I, true)
    cpu.write(0x0100 + uint16(cpu.stkp), cpu.status)
    cpu.stkp--

    cpu.addrAbs = 0xFFFA
    lo := uint16(cpu.read(cpu.addrAbs))
    hi := uint16(cpu.read(cpu.addrAbs+1))

    cpu.pc = (hi << 8) | lo

    cpu.cycles = 8
}
