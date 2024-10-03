package cartridge

func (c *Cartridge) CpuRead(addr uint16, readOnly bool) bool {
	return false
}

func (c *Cartridge) CpuWrite(addr uint16, data uint8) bool {
	return false
}

func (c *Cartridge) PpuRead(addr uint16) bool {
	return false
}

func (c *Cartridge) PpuWrite(addr uint16, data uint8) bool {
	return false
}
