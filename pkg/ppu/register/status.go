package register

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

func (sr *StatusReg) getData(dataLen, shift uint8) uint8 {
	return getRegData(&sr.reg, dataLen, shift)
}

func (sr *StatusReg) setData(data, dataLen, shift uint8) {
	setRegData(&sr.reg, data, dataLen, shift)
}

func (sr *StatusReg) GetSpriteOverflow() uint8 {
	return sr.getData(1, 5)
}

func (sr *StatusReg) GetSpriteZeroHit() uint8 {
	return sr.getData(1, 6)
}

func (sr *StatusReg) GetVerticalBlank() uint8 {
	return sr.getData(1, 7)
}

func (sr *StatusReg) SetVerticalBlank(vb uint8) {
	sr.setData(vb, 1, 7)
}
