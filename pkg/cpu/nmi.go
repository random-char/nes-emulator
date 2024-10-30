package cpu

func (cpu *MOSTechnology6502) Nmi() {
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

	cpu.setFlag(flag_B, false)
	cpu.setFlag(flag_U, true)
	cpu.setFlag(flag_I, true)
	cpu.write(0x0100+uint16(cpu.stkp), cpu.status)
	cpu.stkp--

	cpu.addrAbs = 0xFFFA
	lo := uint16(cpu.read(cpu.addrAbs))
	hi := uint16(cpu.read(cpu.addrAbs + 1))

	cpu.pc = (hi << 8) | lo

	cpu.cycles = 8
}
