package ppu

func (ppu *Ricoh2c02) Reset() {
	ppu.fineX = 0
	ppu.addressLatch = 0
	ppu.dataBuffer = 0x00
	ppu.scanline = 0
	ppu.cycle = 0
	ppu.bgNextTileId = 0x00
	ppu.bgNextTileAttr = 0x00
	ppu.bgNextTileLsb = 0x00
	ppu.bgNextTileMsb = 0x00
	ppu.bgShifterPatternLo = 0x0000
	ppu.bgShifterPatternLo = 0x0000
	ppu.bgShifterAttrLo = 0x0000
	ppu.bgShifterAttrHi = 0x0000
	ppu.status.SetReg(0x00)
	ppu.mask.SetReg(0x00)
	ppu.control.SetReg(0x00)
	ppu.vramAddr.SetReg(0x0000)
	ppu.tramAddr.SetReg(0x0000)
}
