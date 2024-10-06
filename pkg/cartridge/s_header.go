package cartridge

import (
	"encoding/binary"
	"io"
)

type sHeader struct {
	Name         []byte
	PRGRomChunks uint8
	CHRRomChunks uint8
	Mapper1      uint8
	Mapper2      uint8
	PRGRamSize   uint8
	TvSystem1    uint8
	TvSystem2    uint8
	Unused       []byte
}

func ReadHeader(r io.Reader) *sHeader {
    header := &sHeader{
		Name:   make([]byte, 4),
		Unused: make([]byte, 5),
	}

	binary.Read(r, binary.LittleEndian, &header.Name)
	binary.Read(r, binary.LittleEndian, &header.PRGRomChunks)
	binary.Read(r, binary.LittleEndian, &header.CHRRomChunks)
	binary.Read(r, binary.LittleEndian, &header.Mapper1)
	binary.Read(r, binary.LittleEndian, &header.Mapper2)
	binary.Read(r, binary.LittleEndian, &header.PRGRamSize)
	binary.Read(r, binary.LittleEndian, &header.TvSystem1)
	binary.Read(r, binary.LittleEndian, &header.TvSystem2)
	binary.Read(r, binary.LittleEndian, &header.Unused)

	return header
}

