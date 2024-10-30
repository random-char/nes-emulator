package nes

func (nes *NES) CpuRead(addr uint16) uint8 {
	var data uint8 = 0
	if cartData, fromCartridge := nes.cartridge.CpuRead(addr); fromCartridge {
		data = cartData
	} else if addr >= 0x0000 && addr <= 0x1FFF {
		data = nes.cpuRam[addr&0x07FF]
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		data = nes.ppu.CpuRead(addr & 0x0007)
	}

	return data
}

func (nes *NES) CpuWrite(addr uint16, data uint8) {
	if nes.cartridge.CpuWrite(addr, data) {
		// allow extension via cartridge
	} else if addr >= 0x0000 && addr <= 0x1FFF {
		nes.cpuRam[addr&0x07FF] = data
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		nes.ppu.CpuWrite(addr&0x0007, data)
	}
}
