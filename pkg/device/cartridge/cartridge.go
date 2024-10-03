package cartridge

import (
	"encoding/binary"
	"os"
)

type Cartridge struct {
	vPRGMemory []uint8
	vCHRMemory []uint8

	nMapperID uint8
	nPRGBanks uint8
	nCHRBanks uint8
}

func FromFile(filename string) (*Cartridge, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	header := newHeader()
	binary.Read(f, binary.LittleEndian, header)

	if (header.mapper1 & 0x04) != 0 {
		// reserved for trainers, ignore for now
		f.Seek(512, 1)
	}

	nMapperID := ((header.mapper2 >> 4) << 4) | (header.mapper1 >> 4)

	var nFileType uint8 = 1
	var nPRGBanks uint8 = 0
	var nCHRBanks uint8 = 0

	vPRGMemory := make([]uint8, 0)
	vCHRMemory := make([]uint8, 0)

	switch nFileType {
	case 0:
		break
	case 1:
		nPRGBanks = header.prgRomChunks
		vPRGMemory = make([]uint8, uint(nPRGBanks)*16384)
		binary.Read(f, binary.LittleEndian, &vPRGMemory)

		nCHRBanks = header.chrRomChunks
		vCHRMemory = make([]uint8, uint(nCHRBanks)*8192)
		binary.Read(f, binary.LittleEndian, &vCHRMemory)

		break
	case 2:
		break
	}

	return &Cartridge{
		vPRGMemory: vPRGMemory,
		vCHRMemory: vCHRMemory,

		nMapperID: nMapperID,
		nPRGBanks: nPRGBanks,
		nCHRBanks: nCHRBanks,
	}, nil
}
