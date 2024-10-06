package mapper

type mapper000 struct {
	*mapperInternals
}

func (m *mapper000) CpuMapRead(addr uint16) (uint32, bool) {
	var mappedAddr uint32 = 0

	if addr >= 0x8000 && addr <= 0xFFFF {
		if m.nPRGBanks > 1 {
			mappedAddr = uint32(addr & 0x7FFF)
		} else {
			mappedAddr = uint32(addr & 0x3FFF)
		}

		return mappedAddr, true
	}

	return mappedAddr, false
}

func (m *mapper000) CpuMapWrite(addr uint16) (uint32, bool) {
	var mappedAddr uint32 = 0

	if addr >= 0x8000 && addr <= 0xFFFF {
		return mappedAddr, true
	}

	return mappedAddr, false
}

func (m *mapper000) PpuMapRead(addr uint16) (uint32, bool) {
	if addr >= 0x0000 && addr <= 0x1FFF {
		return uint32(addr), true
	}

	return uint32(addr), false
}

func (m *mapper000) PpuMapWrite(addr uint16) (uint32, bool) {
	var mappedAddr uint32 = 0

	return mappedAddr, false
}
