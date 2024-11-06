package register

import "nes-emulator/pkg/binary"

type MaskReg struct {
	//0_0_0_0_0_0_0_0
	//| | | | | | | |
	//| | | | | | | +- [1, 0] grayscale
	//| | | | | | +--- [1, 1] render_background_left
	//| | | | | +----- [1, 2] render_sprites_left
	//| | | | +------- [1, 3] render_backgroud
	//| | | +--------- [1, 4] render_sprites
	//| | +----------- [1, 5] enhance_red
	//| +------------- [1, 6] enhance_green
	//+--------------- [1, 7] enhance_blue
	reg uint8
}

func NewMaskReg() *MaskReg {
	return &MaskReg{
		reg: 0x00,
	}
}

func (mr *MaskReg) GetReg() uint8 {
	return mr.reg
}

func (mr *MaskReg) SetReg(reg uint8) {
	mr.reg = reg
}

func (mr *MaskReg) GetGrayscale() uint8 {
	return mr.getData(1, 0)
}

func (mr *MaskReg) GetRenderBackgroundLeft() uint8 {
	return mr.getData(1, 1)
}

func (mr *MaskReg) GetRenderSpritesLeft() uint8 {
	return mr.getData(1, 2)
}

func (mr *MaskReg) GetRenderBackground() uint8 {
	return mr.getData(1, 3)
}

func (mr *MaskReg) GetRenderSprites() uint8 {
	return mr.getData(1, 4)
}

func (mr *MaskReg) GetEnhanceRed() uint8 {
	return mr.getData(1, 5)
}

func (mr *MaskReg) GetEnhanceGreen() uint8 {
	return mr.getData(1, 6)
}

func (mr *MaskReg) GetEnhanceBlue() uint8 {
	return mr.getData(1, 7)
}

func (mr *MaskReg) getData(dataLen, shift uint8) uint8 {
	return binary.GetData(&mr.reg, dataLen, shift)
}

func (mr *MaskReg) setData(data, dataLen, shift uint8) {
	binary.SetData(&mr.reg, data, dataLen, shift)
}
