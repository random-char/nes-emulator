package cartridge

import (
	"nes-emulator/pkg/mapper"
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
