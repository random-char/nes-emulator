package register

import "nes-emulator/pkg/binary"

type StatusReg struct {
	//vertical_blank[1]
	//sprite_zero_hit[1]
	//sprite_overflow[1]
	//unused[5]
	//0_0_0_00000
	//| | | |||||
	//| | | +++++- [5, 0] unused
	//| | |
	//| | |
	//| | |
	//| | |
	//| | ------------ [1, 5] sprite_overflow
	//| -------------- [1, 6] sprite_zero_hit
	//---------------- [1, 7] vertical_blank
	reg uint8
}

func NewStatusReg() *StatusReg {
	return &StatusReg{
		reg: 0x00,
	}
}

func (sr *StatusReg) GetReg() uint8 {
	return sr.reg
}

func (sr *StatusReg) SetReg(reg uint8) {
	sr.reg = reg
}

func (sr *StatusReg) GetSpriteOverflow() uint8 {
	return sr.getData(1, 5)
}

func (sr *StatusReg) SetSpriteOverflow(v bool) {
	if v {
		sr.setData(1, 1, 5)
	} else {
		sr.setData(0, 1, 5)
	}
}

func (sr *StatusReg) GetSpriteZeroHit() uint8 {
	return sr.getData(1, 6)
}

func (sr *StatusReg) SetSpriteZeroHit(v bool) {
	if v {
		sr.setData(1, 1, 6)
	} else {
		sr.setData(0, 1, 6)
	}
}

func (sr *StatusReg) GetVerticalBlank() uint8 {
	return sr.getData(1, 7)
}

func (sr *StatusReg) SetVerticalBlank(v bool) {
	if v {
		sr.setData(1, 1, 7)
	} else {
		sr.setData(0, 1, 7)
	}
}

func (sr *StatusReg) getData(dataLen, shift uint8) uint8 {
	return binary.GetData(&sr.reg, dataLen, shift)
}

func (sr *StatusReg) setData(data, dataLen, shift uint8) {
	binary.SetData(&sr.reg, data, dataLen, shift)
}
