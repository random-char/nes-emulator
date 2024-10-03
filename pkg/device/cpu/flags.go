package cpu

const (
	FLAG_C = uint8(1 << 0) // carry bit
	FLAG_Z = uint8(1 << 1) // zero
	FLAG_I = uint8(1 << 2) // disable interrupts
	FLAG_D = uint8(1 << 3) // decimal mode
	FLAG_B = uint8(1 << 4) // break
	FLAG_U = uint8(1 << 5) // unused
	FLAG_V = uint8(1 << 6) // overflow
	FLAG_N = uint8(1 << 7) // negative
)

func (cpu *Olc6502) setFlag(flag uint8, value bool) {
	if value {
		cpu.status |= flag
	} else {
		cpu.status &= ^flag
	}
}

func (cpu *Olc6502) getFlag(flag uint8) uint8 {
	if cpu.status&flag > 0 {
		return 1
	} else {
		return 0
	}
}
