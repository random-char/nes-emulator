package nes

func (nes *NES) Clock() {
	nes.ppu.Clock()

	if nes.nSystemClockCounter%3 == 0 {
		if nes.dmaTransfer {
			if nes.dmaDummy {
				if nes.nSystemClockCounter%2 == 1 {
					nes.dmaDummy = false
				}
			} else {
				if nes.nSystemClockCounter%2 == 0 {
					//read
					nes.dmaData = nes.CpuRead(uint16(nes.dmaPage)<<8 | uint16(nes.dmaAddr))
				} else {
					//write
					nes.ppu.OAM.Set(nes.dmaAddr, nes.dmaData)
					nes.dmaAddr++

					if nes.dmaAddr == 0x00 {
						nes.dmaTransfer = false
						nes.dmaDummy = true
					}
				}
			}
		} else {
			nes.cpu.Clock()
		}
	}

	if nes.ppu.Nmi {
		nes.ppu.Nmi = false
		nes.cpu.Nmi()
	}

	nes.nSystemClockCounter++
}
