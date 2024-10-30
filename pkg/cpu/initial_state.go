package cpu

var InitialState *initialState = newInitialState()

type initialState struct {
	a      uint8
	x      uint8
	y      uint8
	stkp   uint8
	pc     uint16
	status uint8

	fromSpecificPc bool
}

func newInitialState() *initialState {
	return &initialState{
		a: 0,
		x: 0,
		y: 0,

		stkp: 0xFD,

		fromSpecificPc: false,
		pc:             0x0000,
		status:         0x00 | flag_U | flag_I,

	}
}

func (is *initialState) WithPc(pc uint16) *initialState {
    is.fromSpecificPc = true
	is.pc = pc

	return is
}

func (is *initialState) WithDefaultPc() *initialState {
    is.fromSpecificPc = false
	is.pc = 0x0000

	return is
}
