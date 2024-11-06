package ppu

import (
	"nes-emulator/pkg/visuals"
)

var nTileX, nTileY, nOffset uint16
var tileLsb, tileMsb, tilePixel uint8
var row, col uint16

func (ppu *Ricoh2c02) getPatternTable(i, palette uint8) *visuals.Sprite {
	for nTileX = 0; nTileX < 16; nTileX++ {
		for nTileY = 0; nTileY < 16; nTileY++ {
			nOffset = nTileY*256 + nTileX*16

			for row = 0; row < 8; row++ {
				tileLsb = ppu.PpuRead(uint16(i)*0x1000 + nOffset + row)
				tileMsb = ppu.PpuRead(uint16(i)*0x1000 + nOffset + row + 8)

				for col = 0; col < 8; col++ {
					tilePixel = ((tileLsb & 0x01) << 1) | (tileMsb & 0x01)
					tileLsb >>= 1
					tileMsb >>= 1
					ppu.sprPatternTable[i].SetPixel(
						nTileX*8+(7-col),
						nTileY*8+row,
						ppu.getColorFromPaletteRam(palette, tilePixel),
					)
				}
			}
		}
	}

	return ppu.sprPatternTable[i]
}

func (ppu *Ricoh2c02) getColorFromPaletteRam(pallete, pixel uint8) *visuals.Pixel {
	return palScreen[ppu.PpuRead(0x3F00+uint16(pallete<<2+pixel))&0x3F]
}
