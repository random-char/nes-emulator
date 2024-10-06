package cpu

const (
	AM_IMP byte = iota
	AM_IMM
	AM_ZP0
	AM_ZPX
	AM_ZPY
	AM_REL
	AM_ABS
	AM_ABX
	AM_ABY
	AM_IND
	AM_IZX
	AM_IZY
)

func (cpu *Olc6502) IMP() uint8 {
	cpu.fetched = cpu.a

	return 0
}

func (cpu *Olc6502) IMM() uint8 {
	cpu.pc++
	cpu.addrAbs = cpu.pc

	return 0
}

func (cpu *Olc6502) ZP0() uint8 {
	cpu.addrAbs = uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs &= 0x00FF

	return 0
}

func (cpu *Olc6502) ZPX() uint8 {
	cpu.addrAbs = uint16(cpu.read(cpu.pc) + cpu.x)
	cpu.pc++

	cpu.addrAbs &= 0x00FF

	return 0
}

func (cpu *Olc6502) ZPY() uint8 {
	cpu.addrAbs = uint16(cpu.read(cpu.pc) + cpu.y)
	cpu.pc++

	cpu.addrAbs &= 0x00FF

	return 0
}

func (cpu *Olc6502) REL() uint8 {
	cpu.addrRel = uint16(cpu.read(cpu.pc))
	cpu.pc++

	if cpu.addrRel&0x80 != 0 {
		cpu.addrRel |= 0xFF00
	}

	return 0
}

func (cpu *Olc6502) ABS() uint8 {
	lo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	hi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs = (hi << 8) | lo

	return 0
}

func (cpu *Olc6502) ABX() uint8 {
	lo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	hi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(cpu.x)

	if (cpu.addrAbs & 0x00FF) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

func (cpu *Olc6502) ABY() uint8 {
	lo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	hi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(cpu.y)

	if (cpu.addrAbs & 0x00FF) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

func (cpu *Olc6502) IND() uint8 {
	ptrLo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	ptrHi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	ptr := (ptrHi << 8) | ptrLo

	if ptrLo == 0x00FF {
		//hardware bug simulation
		cpu.addrAbs = (uint16(cpu.read(ptr&0xFF00)) << 8) | uint16(cpu.read(ptr))
	} else {
		cpu.addrAbs = (uint16(cpu.read(ptr+1)) << 8) | uint16(cpu.read(ptr))
	}

	return 0
}

func (cpu *Olc6502) IZX() uint8 {
	t := uint16(cpu.read(cpu.pc))
	cpu.pc++

	lo := uint16(cpu.read(t+uint16(cpu.x)) & 0x00FF)
	hi := uint16(cpu.read(t+uint16(cpu.x)+1) & 0x00FF)

	cpu.addrAbs = (hi << 8) | lo

	return 0
}

func (cpu *Olc6502) IZY() uint8 {
	t := uint16(cpu.read(cpu.pc))
	cpu.pc++

	lo := uint16(cpu.read(t & 0x00FF))
	hi := uint16(cpu.read(t + 1&0x00FF))

	cpu.addrAbs = ((hi << 8) | lo) + uint16(cpu.y)

	if cpu.addrAbs&0xFF00 != (hi << 8) {
		return 1
	} else {
		return 0
	}
}
