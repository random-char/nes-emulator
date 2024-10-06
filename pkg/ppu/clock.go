package ppu

func (ppu *Olc2c02) Clock() {
	go ppu.receiver.SetPixel(ppu.cycle, ppu.scanline, PalScreen[0x11])

	ppu.cycle++
	if ppu.cycle >= 341 {
		ppu.cycle = 0
		ppu.scanline++

		if ppu.scanline >= 261 {
			ppu.scanline = -1
			ppu.frameComplete = true

			go ppu.receiver.DrawFrame()
		}
	}
}
