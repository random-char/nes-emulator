package cartridge

func (c *Cartridge) CpuRead(addr uint16, readOnly bool) (uint8, bool) {
	var data uint8 = 0

	mappedAddr, fromCartridge := c.mapper.CpuMapRead(addr)
	if fromCartridge {
		data = c.vPRGMemory[mappedAddr]
	}

	return data, fromCartridge
}

func (c *Cartridge) CpuWrite(addr uint16, data uint8) bool {
	mappedAddr, toCartridge := c.mapper.CpuMapWrite(addr)
	if toCartridge {
		c.vPRGMemory[mappedAddr] = data
	}

	return toCartridge
}

func (c *Cartridge) PpuRead(addr uint16) (uint8, bool) {
	var data uint8 = 0

	mappedAddr, fromCartridge := c.mapper.PpuMapRead(addr)
	if fromCartridge {
		data = c.vPRGMemory[mappedAddr]
	}

	return data, fromCartridge
}

func (c *Cartridge) PpuWrite(addr uint16, data uint8) bool {
	mappedAddr, toCartridge := c.mapper.PpuMapWrite(addr)
	if toCartridge {
		c.vCHRMemory[mappedAddr] = data
	}

	return toCartridge
}
