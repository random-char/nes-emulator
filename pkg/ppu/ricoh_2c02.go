package ppu

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/debugger"
	"nes-emulator/pkg/ppu/register"
	"nes-emulator/pkg/ppu/visuals"
)

type Ricoh2c02 struct {
	cycle         int16
	scanline      int16
	FrameRendered bool

	//registers
	vramAddr *register.LoopyReg
	tramAddr *register.LoopyReg
	fineX    uint8

	status  *register.StatusReg
	mask    *register.MaskReg
	control *register.ControlReg

	addressLatch uint8
	dataBuffer   uint8

	tblName    [2][1024]uint8 //[2][1024]
	tblPattern [2][4096]uint8 //[2][4096] javidx9 Extension idea
	tblPallete [32]uint8      //[32]

	sprNameTable    [2]*visuals.Sprite
	sprPatternTable [2]*visuals.Sprite

	//bg rendering
	bgNextTileId   uint8
	bgNextTileAttr uint8
	bgNextTileLsb  uint8
	bgNextTileMsb  uint8
	//shifters
	bgShifterPatternLo uint16
	bgShifterPatternHi uint16
	bgShifterAttrLo    uint16
	bgShifterAttrHi    uint16

	Nmi bool

	receiver      VideoReceiver
	debugReceiver DebugReceiver

	cartridge *cartridge.Cartridge
	debugger  *debugger.Debugger
}

func New(
	receiver VideoReceiver,
	debugReceiver DebugReceiver,
) *Ricoh2c02 {
	return &Ricoh2c02{
		Nmi: false,

		cycle:    0,
		scanline: 0,

		status:  register.NewStatusReg(),
		mask:    register.NewMaskReg(),
		control: register.NewControlReg(),

		vramAddr:     register.NewLoopyReg(),
		tramAddr:     register.NewLoopyReg(),
		fineX:        0x00,
		addressLatch: 0x00,
		dataBuffer:   0x00,

		sprNameTable: visuals.NewSpriteTable(
			256, 240,
			256, 240,
		),
		sprPatternTable: visuals.NewSpriteTable(
			128, 128,
			128, 128,
		),

		receiver:      receiver,
		debugReceiver: debugReceiver,
		debugger:      &debugger.Debugger{},
	}
}

