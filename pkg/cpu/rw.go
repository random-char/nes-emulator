package cpu

func (cpu *MOSTechnology6502) read(addr uint16) uint8 {
	return cpu.bus.CpuRead(addr)
}

func (cpu *MOSTechnology6502) write(addr uint16, data uint8) {
	cpu.bus.CpuWrite(addr, data)
}
