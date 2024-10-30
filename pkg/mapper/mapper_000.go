package mapper

type mapper000 struct {
	*mapperInternals
}

func (m *mapper000) Reset() {}

func (m *mapper000) CpuMapRead(addr uint16) (uint32, bool) {
	var mappedAddr uint32 = 0

	if addr >= 0x8000 && addr <= 0xFFFF {
		var mask uint16 = 0x3FFF
		if m.nPRGBanks > 1 {
			mask = 0x7FFF
		}

		mappedAddr = uint32(addr & mask)

		return mappedAddr, true
	}

	return mappedAddr, false
}

func (m *mapper000) CpuMapWrite(addr uint16) (uint32, bool) {
	var mappedAddr uint32 = 0

	if addr >= 0x8000 && addr <= 0xFFFF {
		var mask uint16 = 0x3FFF
		if m.nPRGBanks > 1 {
			mask = 0x7FFF
		}

		mappedAddr = uint32(addr & mask)

		return mappedAddr, true
	}

	return mappedAddr, false
}

func (m *mapper000) PpuMapRead(addr uint16) (uint32, bool) {
	if addr >= 0x0000 && addr <= 0x1FFF {
		return uint32(addr), true
	}

	return 0, false
}

func (m *mapper000) PpuMapWrite(addr uint16) (uint32, bool) {
	if addr >= 0x0000 && addr <= 0x1FFF {
		if m.nCHRBanks == 0 {
			return uint32(addr), true
		}
	}

	return 0, false
}
