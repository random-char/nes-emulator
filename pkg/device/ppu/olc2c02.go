package ppu

import (
	"nes-emulator/pkg/device/cartridge"
)

type Olc2c02 struct {
	tblName    [][]uint8
	tblPallete []uint8
	tblPattern [][]uint8 //javidx9 Extension idea

	bus       bus
	cartridge *cartridge.Cartridge
}

func New() *Olc2c02 {
	tblName := make([][]uint8, 2)
	tblName[0] = make([]uint8, 1024)
	tblName[1] = make([]uint8, 1024)

	tblPattern := make([][]uint8, 2)
	tblPattern[0] = make([]uint8, 4096)
	tblPattern[1] = make([]uint8, 4096)

	return &Olc2c02{
		tblName:    tblName,
		tblPallete: make([]uint8, 32),
		tblPattern: tblPattern,
	}
}
