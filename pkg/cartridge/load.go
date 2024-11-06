package cartridge

import (
	"io"
	"nes-emulator/pkg/mapper"
)

func load(rsc io.ReadSeeker) (*Cartridge, error) {
	header ,err := loadHeader(rsc)
    if err != nil {
        return nil, ReadingHeaderError
    }

	if (header.Mapper1 & 0x04) != 0 {
		// reserved for trainers, ignore for now
		rsc.Seek(512, io.SeekCurrent)
	}

	nMapperID := ((header.Mapper2 >> 4) << 4) | (header.Mapper1 >> 4)
	mirror := horizontal
	if (header.Mapper1 & 0x01) != 0 {
		mirror = vertical
	}

	var nFileType uint8 = 0x01
	var nPRGBanks uint8 = 0x00
	var nCHRBanks uint8 = 0x00

	var vPRGMemory, vCHRMemory []uint8

	switch nFileType {
	case 1:
		nPRGBanks = header.PRGRomChunks
		vPRGMemory = make([]uint8, uint(nPRGBanks)*0x4000)

		if _, err := io.ReadFull(rsc, vPRGMemory); err != nil {
			return nil, err
		}

		nCHRBanks = header.CHRRomChunks
		length := uint(nCHRBanks) * 0x2000
		if length == 0 {
			length = 0x2000
		}
		vCHRMemory = make([]uint8, length)

		if _, err := io.ReadFull(rsc, vCHRMemory); err != nil {
			return nil, err
		}
	default:
		return nil, UnsupportedFileTypeErr
	}

	m, err := mapper.CreateMapper(nMapperID, nPRGBanks, nCHRBanks)
	if err != nil {
		return nil, err
	}

	return &Cartridge{
		vPRGMemory: vPRGMemory,
		vCHRMemory: vCHRMemory,

		nMapperID: nMapperID,
		nPRGBanks: nPRGBanks,
		nCHRBanks: nCHRBanks,
		mirror:    mirror,

		mapper: m,
	}, nil
}
