package cartridge

import (
	"encoding/binary"
	"io"
	"nes-emulator/pkg/mapper"
)

func Load(rsc io.ReadSeeker) (*Cartridge, error) {
	header := ReadHeader(rsc)

	if (header.Mapper1 & 0x04) != 0 {
		// reserved for trainers, ignore for now
		rsc.Seek(512, 1)
	}

	nMapperID := ((header.Mapper2 >> 4) << 4) | (header.Mapper1 >> 4)

	var nFileType uint8 = 1
	var nPRGBanks uint8 = 0
	var nCHRBanks uint8 = 0

	vPRGMemory := make([]uint8, 0)
	vCHRMemory := make([]uint8, 0)

	switch nFileType {
	case 0:
		break
	case 1:
		nPRGBanks = header.PRGRomChunks
		vPRGMemory = make([]uint8, uint(nPRGBanks)*16384)
		binary.Read(rsc, binary.LittleEndian, &vPRGMemory)

		nCHRBanks = header.CHRRomChunks
		vCHRMemory = make([]uint8, uint(nCHRBanks)*8192)
		binary.Read(rsc, binary.LittleEndian, &vCHRMemory)

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

		mapper: mapper.CreateMapper(nMapperID, nPRGBanks, nCHRBanks),
	}, nil
}
