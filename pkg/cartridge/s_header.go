package cartridge

import (
	"encoding/binary"
	"io"
)

type sHeader struct {
	Name         [4]byte
	PRGRomChunks uint8
	CHRRomChunks uint8
	Mapper1      uint8
	Mapper2      uint8
	PRGRamSize   uint8
	TvSystem1    uint8
	TvSystem2    uint8
	Unused       [5]byte
}

func loadHeader(r io.Reader) (*sHeader, error) {
	header := &sHeader{}

	if err := binary.Read(r, binary.LittleEndian, header); err != nil {
		return nil, err
	}

	return header, nil
}
