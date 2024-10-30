package register

type LoopyReg struct {
	//0_000_0_0_00000_00000
    //| ||| | | ||||| |||||
    //| ||| | | ||||| +++++- [5,  0] coarse_x
    //| ||| | | +++++------- [5,  5] coarse_y
    //| ||| | +------------- [1, 10] nametable_x
    //| ||| +--------------- [1, 11] nametable_y
    //| +++----------------- [3, 12] fine_y
    //+--------------------- [1, 15] unused
	reg uint16
}

func NewLoopyReg() *LoopyReg {
	return &LoopyReg{
		reg: 0x0000,
	}
}

func (lr *LoopyReg) GetReg() uint16 {
	return lr.reg
}

func (lr *LoopyReg) SetReg(reg uint16) {
	lr.reg = reg
}

func (lr *LoopyReg) IncrementReg(i uint16) {
    lr.reg += i
}

func (lr *LoopyReg) getData(dataLen, shift uint8) uint16 {
	return getRegData(&lr.reg, dataLen, shift)
}

func (lr *LoopyReg) setData(data uint16, dataLen, shift uint8) {
	setRegData(&lr.reg, data, dataLen, shift)
}

func (lr *LoopyReg) GetCorseX() uint16 {
	return lr.getData(5, 0)
}

func (lr *LoopyReg) GetCorseY() uint16 {
	return lr.getData(5, 5)
}

func (lr *LoopyReg) GetNametableX() uint16 {
	return lr.getData(1, 10)
}

func (lr *LoopyReg) GetNametableY() uint16 {
	return lr.getData(1, 11)
}

func (lr *LoopyReg) GetFineY() uint16 {
	return lr.getData(3, 12)
}

func (lr *LoopyReg) SetCoarseX(coarseX uint16) {
	lr.setData(coarseX, 5, 0)
}

func (lr *LoopyReg) SetCoarseY(coarseY uint16) {
	lr.setData(coarseY, 5, 5)
}

func (lr *LoopyReg) SetNametableX(nametableX uint16) {
	lr.setData(nametableX, 1, 10)
}

func (lr *LoopyReg) SetNametableY(nametableY uint16) {
	lr.setData(nametableY, 1, 11)
}

func (lr *LoopyReg) SetFineY(fineY uint16) {
	lr.setData(fineY, 3, 12)
}
