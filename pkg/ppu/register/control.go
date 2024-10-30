package register

type ControlReg struct {
	//0_0_0_0_0_0_0_0
	//| | | | | | | |
	//| | | | | | | +- [1, 0] nametable_x
	//| | | | | | +--- [1, 1] nametable_y
	//| | | | | +----- [1, 2] increment_mode
	//| | | | +------- [1, 3] pattern_sprite
	//| | | +--------- [1, 4] pattern_background
	//| | +----------- [1, 5] sprite_mode
	//| +------------- [1, 6] slave_mode
	//+--------------- [1, 7] enable_nmi
	reg uint8
}

func NewControlReg() *ControlReg {
	return &ControlReg{
		reg: 0x00,
	}
}

func (cr *ControlReg) GetReg() uint8 {
	return cr.reg
}

func (cr *ControlReg) SetReg(reg uint8) {
	cr.reg = reg
}

func (cr *ControlReg) IncrementReg(i uint8) {
	cr.reg += i
}

func (cr *ControlReg) getData(dataLen, shift uint8) uint8 {
	return getRegData(&cr.reg, dataLen, shift)
}

func (cr *ControlReg) setData(data, dataLen, shift uint8) {
	setRegData(&cr.reg, data, dataLen, shift)
}

func (cr *ControlReg) GetNamtableX() uint8 {
	return cr.getData(1, 0)
}

func (cr *ControlReg) GetNamtableY() uint8 {
	return cr.getData(1, 1)
}

func (cr *ControlReg) GetIncrementMode() uint8 {
	return cr.getData(1, 2)
}

func (cr *ControlReg) GetPatternSprite() uint8 {
	return cr.getData(1, 3)
}

func (cr *ControlReg) GetPatternBackgroud() uint8 {
	return cr.getData(1, 4)
}

func (cr *ControlReg) GetSpriteSize() uint8 {
	return cr.getData(1, 5)
}

func (cr *ControlReg) GetEnableNmi() uint8 {
	return cr.getData(1, 7)
}
