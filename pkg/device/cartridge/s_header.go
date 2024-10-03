package cartridge

type sHeader struct {
	name         []byte
	prgRomChunks uint8
	chrRomChunks uint8
	mapper1      uint8
	mapper2      uint8
	prgRamSize   uint8
	tvSystem1    uint8
	tvSystem2    uint8
	unused       []byte
}

func newHeader() *sHeader {
	return &sHeader{
		name:   make([]byte, 4),
		unused: make([]byte, 5),
	}
}
