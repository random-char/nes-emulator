package ppu

func (ppu *Olc2c02) CpuRead(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0

	switch addr {
	// Control
	case 0x0000:
		break
	// Mask
	case 0x0001:
		break
	// Status
	case 0x0002:
		break
	// OAM Address
	case 0x0003:
		break
	// OAM Data
	case 0x0004:
		break
	// Scroll
	case 0x0005:
		break
	// PPU Address
	case 0x0006:
		break
	// PPU Data
	case 0x0007:
		break
	}

	return data
}

func (ppu *Olc2c02) CpuWrite(addr uint16, data uint8) {
	switch addr {
	// Control
	case 0x0000:
		break
	// Mask
	case 0x0001:
		break
	// Status
	case 0x0002:
		break
	// OAM Address
	case 0x0003:
		break
	// OAM Data
	case 0x0004:
		break
	// Scroll
	case 0x0005:
		break
	// PPU Address
	case 0x0006:
		break
	// PPU Data
	case 0x0007:
		break
	}
}

func (ppu *Olc2c02) PpuRead(addr uint16) uint8 {
	var data uint8 = 0
	addr &= 0x3FFF

	_, toCartridge := ppu.cartridge.PpuRead(addr)
	if toCartridge {
		// allow extension via cartridge
	}

	return data
}

func (ppu *Olc2c02) PpuWrite(addr uint16, data uint8) {
	addr &= 0x3FFF

	if ppu.cartridge.PpuWrite(addr, data) {
		// allow extension via cartridge
	}
}
