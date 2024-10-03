package cpu

func (cpu *Olc6502) read(addr uint16) uint8 {
	return cpu.bus.CpuRead(addr, true)
}

func (cpu *Olc6502) write(addr uint16, data uint8) {
	cpu.bus.CpuWrite(addr, data)
}
