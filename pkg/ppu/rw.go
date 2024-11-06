package ppu

func (ppu *Ricoh2c02) CpuRead(addr uint16) uint8 {
	var data uint8 = 0x00

	switch addr {
	// Control
	case 0x0000:
	// Mask
	case 0x0001:
	// Status
	case 0x0002:
		data = (ppu.status.GetReg() & 0xE0) | (ppu.dataBuffer & 0x1F)
		ppu.status.SetVerticalBlank(false)
		ppu.addressLatch = 0
	// OAM Address
	case 0x0003:
	// OAM Data
	case 0x0004:
		//todo recheck
		data = ppu.OAM.Get(ppu.oamAddr)
	// Scroll
	case 0x0005:
	// PPU Address
	case 0x0006:
	// PPU Data
	case 0x0007:
		data = ppu.dataBuffer
		ppu.dataBuffer = ppu.PpuRead(ppu.vramAddr.GetReg())

		if ppu.vramAddr.GetReg() >= 0x3F00 {
			data = ppu.dataBuffer
		}

		if ppu.control.GetIncrementMode() != 0 {
			ppu.vramAddr.IncrementReg(32)
		} else {
			ppu.vramAddr.IncrementReg(1)
		}
	}

	return data
}

func (ppu *Ricoh2c02) CpuWrite(addr uint16, data uint8) {
	switch addr {
	// Control
	case 0x0000:
		ppu.control.SetReg(data)
		ppu.tramAddr.SetNametableX(
			uint16(ppu.control.GetNamtableX()),
		)
		ppu.tramAddr.SetNametableY(
			uint16(ppu.control.GetNamtableY()),
		)
	// Mask
	case 0x0001:
		ppu.mask.SetReg(data)
	// Status
	case 0x0002:
	// OAM Address
	case 0x0003:
		ppu.oamAddr = data
	// OAM Data
	case 0x0004:
		ppu.OAM.Set(ppu.oamAddr, data)
	// Scroll
	case 0x0005:
		if ppu.addressLatch == 0 {
			ppu.fineX = data & 0x07
			ppu.tramAddr.SetCoarseX(uint16(data >> 3))
			ppu.addressLatch = 1
		} else {
			ppu.tramAddr.SetFineY(uint16(data & 0x07))
			ppu.tramAddr.SetCoarseY(uint16(data >> 3))
			ppu.addressLatch = 0
		}
	// PPU Address
	case 0x0006:
		if ppu.addressLatch == 0 {
			ppu.tramAddr.SetReg((ppu.tramAddr.GetReg() & 0x00FF) | (uint16(data&0x3F) << 8))
			ppu.addressLatch = 1
		} else {
			ppu.tramAddr.SetReg((ppu.tramAddr.GetReg() & 0xFF00) | uint16(data))
			ppu.vramAddr.SetReg(ppu.tramAddr.GetReg())
			ppu.addressLatch = 0
		}
	// PPU Data
	case 0x0007:
		ppu.PpuWrite(ppu.vramAddr.GetReg(), data)

		if ppu.control.GetIncrementMode() == 1 {
			ppu.vramAddr.IncrementReg(32)
		} else {
			ppu.vramAddr.IncrementReg(1)
		}
	}
}

func (ppu *Ricoh2c02) PpuRead(addr uint16) uint8 {
	var data uint8 = 0
	addr &= 0x3FFF

	if cartData, fromCartridge := ppu.cartridge.PpuRead(addr); fromCartridge {
		data = cartData
	} else if addr >= 0x0000 && addr <= 0x1FFF {
		data = ppu.tblPattern[(addr&0x1000)>>12][addr&0x0FFF]
	} else if addr >= 0x2000 && addr <= 0x3EFF {
		addr &= 0x0FFF

		if ppu.cartridge.IsMirrorVertical() {
			if addr >= 0x0000 && addr <= 0x03FF {
				data = ppu.tblName[0][addr&0x03FF]
			} else if addr >= 0x0400 && addr <= 0x07FF {
				data = ppu.tblName[1][addr&0x03FF]
			} else if addr >= 0x0800 && addr <= 0x0BFF {
				data = ppu.tblName[0][addr&0x03FF]
			} else if addr >= 0x0C00 && addr <= 0x0FFF {
				data = ppu.tblName[1][addr&0x03FF]
			}
		} else if ppu.cartridge.IsMirrorHorizontal() {
			if addr >= 0x0000 && addr <= 0x03FF {
				data = ppu.tblName[0][addr&0x03FF]
			} else if addr >= 0x0400 && addr <= 0x07FF {
				data = ppu.tblName[0][addr&0x03FF]
			} else if addr >= 0x0800 && addr <= 0x0BFF {
				data = ppu.tblName[1][addr&0x03FF]
			} else if addr >= 0x0C00 && addr <= 0x0FFF {
				data = ppu.tblName[1][addr&0x03FF]
			}
		}

	} else if addr >= 0x3F00 && addr <= 0x3FFF {
		addr &= 0x001F

		if addr == 0x0010 {
			addr = 0x0000
		} else if addr == 0x0014 {
			addr = 0x0004
		} else if addr == 0x0018 {
			addr = 0x0008
		} else if addr == 0x001C {
			addr = 0x000C
		}

		var mask uint8 = 0x3F
		if ppu.mask.GetGrayscale() == 1 {
			mask = 0x30
		}

		data = ppu.tblPallete[addr] & mask
	}

	return data
}

func (ppu *Ricoh2c02) PpuWrite(addr uint16, data uint8) {
	addr &= 0x3FFF

	if ppu.cartridge.PpuWrite(addr, data) {
		// allow extension via cartridge
	} else if addr >= 0x0000 && addr <= 0x1FFF {
		// usually is ROM, but in some cases is used as RAM
		ppu.tblPattern[(addr&0x1000)>>12][addr&0x0FFF] = data
	} else if addr >= 0x2000 && addr <= 0x3EFF {
		addr &= 0x0FFF

		if ppu.cartridge.IsMirrorVertical() {
			// Vertical
			if addr >= 0x0000 && addr <= 0x03FF {
				ppu.tblName[0][addr&0x03FF] = data
			} else if addr >= 0x0400 && addr <= 0x07FF {
				ppu.tblName[1][addr&0x03FF] = data
			} else if addr >= 0x0800 && addr <= 0x0BFF {
				ppu.tblName[0][addr&0x03FF] = data
			} else if addr >= 0x0C00 && addr <= 0x0FFF {
				ppu.tblName[1][addr&0x03FF] = data
			}
		} else if ppu.cartridge.IsMirrorHorizontal() {
			// Horizontal
			if addr >= 0x0000 && addr <= 0x03FF {
				ppu.tblName[0][addr&0x03FF] = data
			} else if addr >= 0x0400 && addr <= 0x07FF {
				ppu.tblName[0][addr&0x03FF] = data
			} else if addr >= 0x0800 && addr <= 0x0BFF {
				ppu.tblName[1][addr&0x03FF] = data
			} else if addr >= 0x0C00 && addr <= 0x0FFF {
				ppu.tblName[1][addr&0x03FF] = data
			}
		}

	} else if addr >= 0x3F00 && addr <= 0x3FFF {
		addr &= 0x001F
		if addr == 0x0010 {
			addr = 0x0000
		} else if addr == 0x0014 {
			addr = 0x0004
		} else if addr == 0x0018 {
			addr = 0x0008
		} else if addr == 0x001C {
			addr = 0x000C
		}

		ppu.tblPallete[addr] = data
	}
}
