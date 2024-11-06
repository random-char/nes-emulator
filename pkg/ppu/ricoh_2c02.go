package ppu

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/debugger"
	"nes-emulator/pkg/ppu/register"
	"nes-emulator/pkg/visuals"
)

type Ricoh2c02 struct {
	Nmi           bool
	FrameRendered bool

	cycle    int16
	scanline int16

	fineX        uint8
	addressLatch uint8
	dataBuffer   uint8

	//registers
	vramAddr *register.LoopyReg
	tramAddr *register.LoopyReg
	status   *register.StatusReg
	mask     *register.MaskReg
	control  *register.ControlReg

	//object attribute memory
	oamAddr                    uint8
	OAM                        *OAM
	SpriteScanline             *spriteScanline
	spriteCount                uint8
	spriteZeroHitPossible      bool
	spriteZeroHitBeingRendered bool

	tblName    [2][1024]uint8
	tblPattern [2][4096]uint8 //javidx9 Extension idea
	tblPallete [32]uint8

	sprNameTable    [2]*visuals.Sprite
	sprPatternTable [2]*visuals.Sprite

	//bg rendering
	bgNextTileId   uint8
	bgNextTileAttr uint8
	bgNextTileLsb  uint8
	bgNextTileMsb  uint8
	//shifters
	bgShifterPatternLo     uint16
	bgShifterPatternHi     uint16
	bgShifterAttrLo        uint16
	bgShifterAttrHi        uint16
	spriteShifterPatternLo [8]uint8
	spriteShifterPatternHi [8]uint8

	receiver      VideoReceiver
	debugReceiver DebugReceiver

	cartridge *cartridge.Cartridge
	debugger  *debugger.Debugger
}

func New() *Ricoh2c02 {
	return &Ricoh2c02{
		Nmi: false,

		cycle:    0,
		scanline: 0,

		oamAddr:        0x00,
		OAM:            NewOAM(),
		SpriteScanline: newSpriteScanline(),

		fineX:        0x00,
		addressLatch: 0x00,
		dataBuffer:   0x00,
		vramAddr:     register.NewLoopyReg(),
		tramAddr:     register.NewLoopyReg(),
		status:       register.NewStatusReg(),
		mask:         register.NewMaskReg(),
		control:      register.NewControlReg(),

		sprNameTable: visuals.NewSpriteTable(
			256, 240,
			256, 240,
		),
		sprPatternTable: visuals.NewSpriteTable(
			128, 128,
			128, 128,
		),
	}
}
