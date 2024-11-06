package ppu

type spriteObjectAttributeEntry struct {
	y    uint8
	id   uint8
	attr uint8
	x    uint8
}

type OAM [64]spriteObjectAttributeEntry

func NewOAM() *OAM {
	oam := &OAM{}

	for i := 0; i < 64; i++ {
		oam[i].y = 0x00
		oam[i].id = 0x00
		oam[i].attr = 0x00
		oam[i].x = 0x00
	}

	return oam
}

func (oam *OAM) Get(addr uint8) uint8 {
	index := addr / 4

	switch addr % 4 {
	case 0:
		return oam[index].y
	case 1:
		return oam[index].id
	case 2:
		return oam[index].attr
	case 3:
		return oam[index].x
	default:
		//shouldn't happen
		panic("Can't get OAM at this address")
	}
}

func (oam *OAM) Set(addr, data uint8) {
	index := addr / 4

	switch addr % 4 {
	case 0:
		oam[index].y = data
	case 1:
		oam[index].id = data
	case 2:
		oam[index].attr = data
	case 3:
		oam[index].x = data
	default:
		//shouldn't happen
		panic("Can't set OAM at this address")
	}
}

type spriteScanline [8]spriteObjectAttributeEntry

func newSpriteScanline() *spriteScanline {
	ss := &spriteScanline{}

	for i := 0; i < 8; i++ {
		ss[i].y = 0xFF
		ss[i].id = 0xFF
		ss[i].attr = 0xFF
		ss[i].x = 0xFF
	}

	return ss
}

func (ss *spriteScanline) Reset() {
	for i := 0; i < 8; i++ {
		ss[i].y = 0xFF
		ss[i].id = 0xFF
		ss[i].attr = 0xFF
		ss[i].x = 0xFF
	}
}

