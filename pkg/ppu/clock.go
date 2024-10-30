package ppu

import (
	"nes-emulator/pkg/ppu/visuals"
)

const (
	width  = 256
	height = 240
)

var pixels [width * height]*visuals.Pixel
var currIndex int

func (ppu *Ricoh2c02) Clock() bool {
	if ppu.scanline >= -1 && ppu.scanline < 240 {
		if ppu.scanline == 0 && ppu.cycle == 0 {
			ppu.cycle = 1
		}

		if ppu.scanline == -1 && ppu.cycle == 1 {
			ppu.status.SetVerticalBlank(0)
		}

		if (ppu.cycle >= 2 && ppu.cycle < 258) || (ppu.cycle >= 321 && ppu.cycle < 338) {
			ppu.updateShifters()

			switch (ppu.cycle - 1) % 8 {
			case 0:
				ppu.loadBackgroundShifters()
				ppu.bgNextTileId = ppu.PpuRead(0x2000 | (ppu.vramAddr.GetReg() & 0x0FFF))
			case 2:
				ppu.bgNextTileAttr = ppu.PpuRead(
					0x23C0 |
						(ppu.vramAddr.GetNametableY() << 11) |
						(ppu.vramAddr.GetNametableX() << 10) |
						((ppu.vramAddr.GetCorseY() >> 2) << 3) |
						((ppu.vramAddr.GetCorseX()) >> 2),
				)

				if ppu.vramAddr.GetCorseY()&0x02 != 0 {
					ppu.bgNextTileAttr >>= 4
				}
				if ppu.vramAddr.GetCorseX()&0x02 != 0 {
					ppu.bgNextTileAttr >>= 2
				}
				ppu.bgNextTileAttr &= 0x03
			case 4:
				ppu.bgNextTileLsb = ppu.PpuRead(
					(uint16(ppu.control.GetPatternBackgroud()) << 12) +
						(uint16(ppu.bgNextTileId) << 4) +
						ppu.vramAddr.GetFineY(),
				)
			case 6:
				ppu.bgNextTileMsb = ppu.PpuRead(
					(uint16(ppu.control.GetPatternBackgroud()) << 12) +
						(uint16(ppu.bgNextTileId) << 4) + // investigate
						ppu.vramAddr.GetFineY() + 8,
				)
			case 7:
				ppu.incrementScrollX()
			}
		}

		if ppu.cycle == 256 {
			ppu.incrementScrollY()
		}

		if ppu.cycle == 257 {
			ppu.loadBackgroundShifters()
			ppu.transferAddressX()
		}

		if ppu.cycle == 338 || ppu.cycle == 340 {
			ppu.bgNextTileId = ppu.PpuRead(0x2000 | (ppu.vramAddr.GetReg() & 0x0FFF))
		}

		if ppu.scanline == -1 && ppu.cycle >= 280 && ppu.cycle < 305 {
			ppu.transferAddressY()
		}
	}

	if ppu.scanline == 240 {
		//do nothing
	}

	if ppu.scanline >= 241 && ppu.scanline < 261 {
		if ppu.scanline == 241 && ppu.cycle == 1 {
			ppu.status.SetVerticalBlank(1)
			if ppu.control.GetEnableNmi() != 0 {
				ppu.Nmi = true
			}
		}
	}

	var bgPixel uint8 = 0x00
	var bgPalette uint8 = 0x00
	if ppu.mask.GetRenderBackground() != 0 {
		var bitMux uint16 = 0x8000 >> ppu.fineX

		var p0Pixel, p1Pixel uint8 = 0x00, 0x00
		var bgPalette0, bgPalette1 uint8 = 0x00, 0x00

		if (ppu.bgShifterPatternLo & bitMux) > 0 {
			p0Pixel = 0x01
		}
		if (ppu.bgShifterPatternHi & bitMux) > 0 {
			p1Pixel = 0x01
		}
		bgPixel = (p1Pixel << 1) | p0Pixel

		if (ppu.bgShifterAttrLo & bitMux) > 0 {
			bgPalette0 = 0x01
		}
		if (ppu.bgShifterAttrHi & bitMux) > 0 {
			bgPalette1 = 0x01
		}
		bgPalette = (bgPalette1 << 1) | bgPalette0
	}

	setPixel(
		int(ppu.cycle-1),
		int(ppu.scanline),
		ppu.getColorFromPaletteRam(bgPalette, bgPixel),
	)

	ppu.cycle++
	if ppu.cycle >= 341 {
		ppu.cycle = 0
		ppu.scanline++

		if ppu.scanline >= 261 {
			ppu.scanline = -1

			if ppu.receiver != nil {
				ppu.receiver.RenderFrame(pixels[:])
			}

			if ppu.debugReceiver != nil {
				ppu.debugReceiver.RenderPatternTables(
					0,
					ppu.getPatternTable(0, 0),
				)
				ppu.debugReceiver.RenderPatternTables(
					1,
					ppu.getPatternTable(1, 1),
				)
			}
			ppu.FrameRendered = true
		}
	}

	return ppu.FrameRendered
}

func (ppu *Ricoh2c02) incrementScrollX() {
	if ppu.mask.GetRenderBackground() != 0 || ppu.mask.GetRenderSprites() != 0 {
		if ppu.vramAddr.GetCorseX() == 31 {
			ppu.vramAddr.SetCoarseX(0)
			ppu.vramAddr.SetNametableX(
				^ppu.vramAddr.GetNametableX(),
			)
		} else {
			ppu.vramAddr.SetCoarseX(
				ppu.vramAddr.GetCorseX() + 1,
			)
		}
	}
}

func (ppu *Ricoh2c02) incrementScrollY() {
	if ppu.mask.GetRenderBackground() != 0 || ppu.mask.GetRenderSprites() != 0 {
		if ppu.vramAddr.GetFineY() < 7 {
			ppu.vramAddr.SetFineY(
				ppu.vramAddr.GetFineY() + 1,
			)
		} else {
			ppu.vramAddr.SetFineY(0)
			if ppu.vramAddr.GetCorseY() == 29 {
				ppu.vramAddr.SetCoarseY(0)
				ppu.vramAddr.SetNametableY(
					^ppu.vramAddr.GetNametableY(),
				)
			} else if ppu.vramAddr.GetCorseY() == 31 {
				ppu.vramAddr.SetCoarseY(0)
			} else {

				ppu.vramAddr.SetCoarseY(
					ppu.vramAddr.GetCorseY() + 1,
				)
			}
		}
	}
}

func (ppu *Ricoh2c02) transferAddressX() {
	if ppu.mask.GetRenderBackground() != 0 || ppu.mask.GetRenderSprites() != 0 {
		ppu.vramAddr.SetNametableX(ppu.tramAddr.GetNametableX())
		ppu.vramAddr.SetCoarseX(ppu.tramAddr.GetCorseX())
	}
}

func (ppu *Ricoh2c02) transferAddressY() {
	if ppu.mask.GetRenderBackground() != 0 || ppu.mask.GetRenderSprites() != 0 {
		ppu.vramAddr.SetFineY(ppu.tramAddr.GetFineY())
		ppu.vramAddr.SetNametableY(ppu.tramAddr.GetNametableY())
		ppu.vramAddr.SetCoarseY(ppu.tramAddr.GetCorseY())
	}
}

func (ppu *Ricoh2c02) loadBackgroundShifters() {
	ppu.bgShifterPatternLo = (ppu.bgShifterPatternLo & 0xFF00) | uint16(ppu.bgNextTileLsb)
	ppu.bgShifterPatternHi = (ppu.bgShifterPatternHi & 0xFF00) | uint16(ppu.bgNextTileMsb)

	var loShift uint16 = 0x0000
	if (ppu.bgNextTileAttr & 0b01) != 0 {
		loShift = 0x00FF
	}
	ppu.bgShifterAttrLo = (ppu.bgShifterAttrLo & 0xFF00) | loShift

	var hiShift uint16 = 0x0000
	if (ppu.bgNextTileAttr & 0b10) != 0 {
		hiShift = 0x00FF
	}
	ppu.bgShifterAttrHi = (ppu.bgShifterAttrHi & 0xFF00) | hiShift
}

func (ppu *Ricoh2c02) updateShifters() {
	if ppu.mask.GetRenderBackground() != 0 {
		ppu.bgShifterPatternLo <<= 1
		ppu.bgShifterPatternHi <<= 1
		ppu.bgShifterAttrLo <<= 1
		ppu.bgShifterAttrHi <<= 1
	}
}

func setPixel(x, y int, pixel *visuals.Pixel) {
	if x < 0 || y < 0 || x >= width || y >= height {
		return
	}

	pixels[(y*width + x)] = pixel
}
