package nes

func (nes *NES) Reset() {
	nes.cpu.Reset()
	nes.ppu.Reset()

	if nes.cartridge != nil {
		nes.cartridge.Reset()
	}

	nes.nSystemClockCounter = 0

    nes.dmaPage = 0x00
    nes.dmaAddr = 0x00
    nes.dmaData = 0x00
    nes.dmaTransfer = false
    nes.dmaDummy = true
}
