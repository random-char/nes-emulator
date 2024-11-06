package nes

func (nes *NES) CpuRead(addr uint16) uint8 {
	var data uint8 = 0
	if cartData, fromCartridge := nes.cartridge.CpuRead(addr); fromCartridge {
		data = cartData
	} else if addr >= 0x0000 && addr <= 0x1FFF {
		//cpu range
		data = nes.cpuRam[addr&0x07FF]
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		//ppu range
		data = nes.ppu.CpuRead(addr & 0x0007)
	} else if addr >= 0x4016 && addr <= 0x4017 {
		//controller range
		if (nes.controllerState[addr&0x0001] & 0x80) > 0 {
			data = 0x01
		} else {
			data = 0x00
		}

		nes.controllerState[addr&0x0001] <<= 1
	}

	return data
}

func (nes *NES) CpuWrite(addr uint16, data uint8) {
	if nes.cartridge.CpuWrite(addr, data) {
		// allow extension via cartridge
	} else if addr >= 0x0000 && addr <= 0x1FFF {
		//cpu range
		nes.cpuRam[addr&0x07FF] = data
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		//ppu range
		nes.ppu.CpuWrite(addr&0x0007, data)
	} else if addr == 0x4014 {
		nes.dmaPage = data
		nes.dmaAddr = 0x00
		nes.dmaTransfer = true
	} else if addr >= 0x4016 && addr <= 0x4017 {
		//controller range
		nes.controllerState[addr&0x0001] = nes.Controller[addr&0x0001].GetState()
	}
}
