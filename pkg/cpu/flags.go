package cpu

const (
	flag_C = uint8(1 << 0) // carry bit
	flag_Z = uint8(1 << 1) // zero
	flag_I = uint8(1 << 2) // disable interrupts
	flag_D = uint8(1 << 3) // decimal mode
	flag_B = uint8(1 << 4) // break
	flag_U = uint8(1 << 5) // unused
	flag_V = uint8(1 << 6) // overflow
	flag_N = uint8(1 << 7) // negative
)

func (cpu *MOSTechnology6502) setFlag(flag uint8, value bool) {
	if value {
		cpu.status |= flag
	} else {
		cpu.status &= ^flag
	}
}

func (cpu *MOSTechnology6502) getFlag(flag uint8) uint8 {
	if cpu.status&flag != 0 {
		return 1
	} else {
		return 0
	}
}
