package cartridge

import (
	"nes-emulator/pkg/mapper"
)

const (
	horizontal byte = iota
	vertical
	onescreen_lo
	onescreen_hi
)

type Cartridge struct {
	vPRGMemory []uint8
	vCHRMemory []uint8

	nMapperID uint8
	nPRGBanks uint8
	nCHRBanks uint8
	mirror    byte

	mapper mapper.Mapper
}

func (c *Cartridge) Reset() {
	if c.mapper != nil {
		c.mapper.Reset()
	}
}

func (c *Cartridge) IsMirrorVertical() bool {
	return c.mirror == vertical
}

func (c *Cartridge) IsMirrorHorizontal() bool {
	return c.mirror == horizontal
}
