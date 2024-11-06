package ppu

//todo refactor into smaller bits

import (
	"nes-emulator/pkg/binary"
	"nes-emulator/pkg/visuals"
)

const (
	imageWidth  = 256
	imageHeight = 240
)

var pixels [imageWidth * imageHeight]*visuals.Pixel
var nOAMEntry uint8
var spriteSize uint8

var spritePatternBitsLo, spritePatternBitsHi uint8
var spritePatternAddrLo, spritePatternAddrHi uint16

var bgPixel uint8 = 0x00
var bgPalette uint8 = 0x00
var fgPixel uint8 = 0x00
var fgPalette uint8 = 0x00
var fgPriority uint8 = 0x00

func (ppu *Ricoh2c02) Clock() bool {
	if ppu.scanline >= -1 && ppu.scanline < 240 {
		if ppu.scanline == 0 && ppu.cycle == 0 {
			ppu.cycle = 1
		}

		if ppu.scanline == -1 && ppu.cycle == 1 {
			ppu.status.SetVerticalBlank(false)
			ppu.status.SetSpriteOverflow(false)
			ppu.status.SetSpriteZeroHit(false)

			for i := 0; i < 8; i++ {
				ppu.spriteShifterPatternLo[i] = 0x00
				ppu.spriteShifterPatternHi[i] = 0x00
			}
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

		//foreground
		//approximation for now, might improve later
		if ppu.cycle == 257 && ppu.scanline >= 0 {
			//todo find a better way for clearing
			ppu.SpriteScanline.Reset()
			ppu.spriteCount = 0

			for i := 0; i < 8; i++ {
				ppu.spriteShifterPatternLo[i] = 0
				ppu.spriteShifterPatternHi[i] = 0
			}

			ppu.spriteZeroHitPossible = false
			for nOAMEntry = 0; nOAMEntry < 64 && ppu.spriteCount < 9; nOAMEntry++ {
				diff := uint8(ppu.scanline) - ppu.OAM[nOAMEntry].y

				if ppu.control.GetSpriteSize() == 0 {
					spriteSize = 8
				} else {
					spriteSize = 16
				}

				if diff >= 0 && diff < spriteSize {
					if ppu.spriteCount < 8 {

						if nOAMEntry == 0 {
							ppu.spriteZeroHitPossible = true
						}
						ppu.SpriteScanline[ppu.spriteCount] = ppu.OAM[nOAMEntry]
						ppu.spriteCount++
					}
				}
			}

			ppu.status.SetSpriteOverflow(ppu.spriteCount > 8)
		}

		if ppu.cycle == 340 {
			var i uint8
			for i = 0; i < ppu.spriteCount; i++ {
				if ppu.control.GetSpriteSize() == 0 {
					//8x8
					if ppu.SpriteScanline[i].attr&0x80 == 0 {
						//sprite not flipped
						spritePatternAddrLo = (uint16(ppu.control.GetPatternSprite()) << 12) |
							(uint16(ppu.SpriteScanline[i].id) << 4) |
							(uint16(ppu.scanline) - uint16(ppu.SpriteScanline[i].y))
					} else {
						//sprite is flipped vertically
						spritePatternAddrLo = (uint16(ppu.control.GetPatternSprite()) << 12) |
							(uint16(ppu.SpriteScanline[i].id) << 4) |
							(7 - (uint16(ppu.scanline) - uint16(ppu.SpriteScanline[i].y)))
					}
				} else {
					//8x16
					if ppu.SpriteScanline[i].attr&0x80 == 0 {
						//sprite not flipped
						if ppu.scanline-int16(ppu.SpriteScanline[i].y) < 8 {
							//read top half
							spritePatternAddrLo = (uint16(ppu.SpriteScanline[i].id&0x01) << 12) |
								(uint16(ppu.SpriteScanline[i].id&0xFE) << 4) |
								((uint16(ppu.scanline) - uint16(ppu.SpriteScanline[i].y)) & 0x07)
						} else {
							//read bottom half
							spritePatternAddrLo = (uint16(ppu.SpriteScanline[i].id&0x01) << 12) |
								((uint16(ppu.SpriteScanline[i].id&0xFE) + 1) << 4) |
								((uint16(ppu.scanline) - uint16(ppu.SpriteScanline[i].y)) & 0x07)
						}
					} else {
						//sprite flipped vertically
						if ppu.scanline-int16(ppu.SpriteScanline[i].y) < 8 {
							//read top half
							spritePatternAddrLo = (uint16(ppu.SpriteScanline[i].id&0x01) << 12) |
								(((uint16(ppu.SpriteScanline[i].id) & 0xFE) + 1) << 4) |
								(7 - (uint16(ppu.scanline)-uint16(ppu.SpriteScanline[i].y))&0x07)
						} else {
							//read bottom half
							spritePatternAddrLo = (uint16(ppu.SpriteScanline[i].id&0x01) << 12) |
								((uint16(ppu.SpriteScanline[i].id & 0xFE)) << 4) |
								(7 - (uint16(ppu.scanline)-uint16(ppu.SpriteScanline[i].y))&0x07)
						}
					}
				}

				spritePatternAddrHi = spritePatternAddrLo + 8

				spritePatternBitsLo = ppu.PpuRead(spritePatternAddrLo)
				spritePatternBitsHi = ppu.PpuRead(spritePatternAddrHi)

				if ppu.SpriteScanline[i].attr&0x40 != 0 {
					//flip horizontally
					binary.ReverseUint8(&spritePatternBitsLo)
					binary.ReverseUint8(&spritePatternBitsHi)
				}

				ppu.spriteShifterPatternLo[i] = spritePatternBitsLo
				ppu.spriteShifterPatternHi[i] = spritePatternBitsHi
			}
		}
	}

	if ppu.scanline == 240 {
		//do nothing
	}

	if ppu.scanline >= 241 && ppu.scanline < 261 {
		if ppu.scanline == 241 && ppu.cycle == 1 {
			ppu.status.SetVerticalBlank(true)
			if ppu.control.GetEnableNmi() != 0 {
				ppu.Nmi = true
			}
		}
	}

	//background rendering prep
	bgPixel = 0x00
	bgPalette = 0x00
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

	//sprite rendering prep
	fgPixel = 0x00
	fgPalette = 0x00
	fgPriority = 0x00
	if ppu.mask.GetRenderSprites() != 0 {
		ppu.spriteZeroHitBeingRendered = false

		for i := 0; i < int(ppu.spriteCount); i++ {
			if ppu.SpriteScanline[i].x == 0 {
				var fgPixelLo, fgPixelHi uint8 = 0x00, 0x00
				if ppu.spriteShifterPatternLo[i]&0x80 > 0 {
					fgPixelLo = 0x01
				}

				if ppu.spriteShifterPatternHi[i]&0x80 > 0 {
					fgPixelHi = 0x01
				}

				fgPixel = (fgPixelHi << 1) | fgPixelLo

				fgPalette = ppu.SpriteScanline[i].attr&0x03 + 0x04

				if ppu.SpriteScanline[i].attr&0x20 == 0 {
					fgPriority = 0x01
				} else {
					fgPriority = 0x00
				}

				if fgPixel != 0 {
					if i == 0 {
						ppu.spriteZeroHitBeingRendered = true
					}
					break
				}
			}
		}
	}

	//selecting pixel for rendering
	var pixel, palette uint8 = 0x00, 0x00
	if bgPixel != 0x00 && (fgPixel == 0x00 || fgPriority == 0x00) {
		pixel = bgPixel
		palette = bgPalette
	} else if fgPixel != 0x00 && (bgPixel == 0x00 || fgPriority != 0x00) {
		pixel = fgPixel
		palette = fgPalette
	}

	if bgPixel != 0x00 && fgPixel != 0x00 {
		if ppu.spriteZeroHitPossible && ppu.spriteZeroHitBeingRendered {
			if ppu.mask.GetRenderBackground() != 0x00 && ppu.mask.GetRenderSprites() != 0x00 {
				if ^(ppu.mask.GetRenderBackgroundLeft() | ppu.mask.GetRenderSpritesLeft()) != 0x00 {
					if ppu.cycle >= 9 && ppu.cycle < 258 {
						ppu.status.SetSpriteZeroHit(true)
					}
				} else {
					if ppu.cycle >= 1 && ppu.cycle < 258 {
						ppu.status.SetSpriteZeroHit(true)
					}
				}
			}
		}
	}

	setPixel(
		int(ppu.cycle-1),
		int(ppu.scanline),
		ppu.getColorFromPaletteRam(palette, pixel),
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
					ppu.getPatternTable(1, 0),
				)
			}
			ppu.FrameRendered = true
		}
	}

	if ppu.debugger != nil {
		ppu.debugger.PpuCycle = ppu.cycle
		ppu.debugger.PpuScanline = ppu.scanline
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

	if ppu.mask.GetRenderSprites() != 0 && ppu.cycle >= 1 && ppu.cycle < 258 {
		var i uint8
		for i = 0; i < ppu.spriteCount; i++ {
			if ppu.SpriteScanline[i].x > 0 {
				ppu.SpriteScanline[i].x--
			} else {
				ppu.spriteShifterPatternLo[i] <<= 1
				ppu.spriteShifterPatternHi[i] <<= 1
			}
		}
	}
}

func setPixel(x, y int, pixel *visuals.Pixel) {
	if x < 0 || y < 0 || x >= imageWidth || y >= imageHeight {
		return
	}

	pixels[(y*imageWidth + x)] = pixel
}
