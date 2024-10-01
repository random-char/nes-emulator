package cpu

const (
	ADDR_MODE_IMP byte = iota
	ADDR_MODE_IMM
	ADDR_MODE_ZP0
	ADDR_MODE_ZPX
	ADDR_MODE_ZPY
	ADDR_MODE_REL
	ADDR_MODE_ABS
	ADDR_MODE_ABX
	ADDR_MODE_ABY
	ADDR_MODE_IND
	ADDR_MODE_IZX
	ADDR_MODE_IZY
)

func (olc *Olc6502) IMP() uint8 {
	olc.fetched = olc.a

	return 0
}

func (olc *Olc6502) IMM() uint8 {
	olc.pc++
	olc.addrAbs = olc.pc

	return 0
}

func (olc *Olc6502) ZP0() uint8 {
	olc.addrAbs = uint16(olc.read(olc.pc))
	olc.pc++

	olc.addrAbs &= 0x00FF

	return 0
}

func (olc *Olc6502) ZPX() uint8 {
	olc.addrAbs = uint16(olc.read(olc.pc) + olc.x)
	olc.pc++

	olc.addrAbs &= 0x00FF

	return 0
}

func (olc *Olc6502) ZPY() uint8 {
	olc.addrAbs = uint16(olc.read(olc.pc) + olc.y)
	olc.pc++

	olc.addrAbs &= 0x00FF

	return 0
}

func (olc *Olc6502) REL() uint8 {
	olc.addrRel = uint16(olc.read(olc.pc))
	olc.pc++

	if olc.addrRel&0x80 != 0 {
		olc.addrRel |= 0xFF00
	}

	return 0
}

func (olc *Olc6502) ABS() uint8 {
	lo := uint16(olc.read(olc.pc))
	olc.pc++

	hi := uint16(olc.read(olc.pc))
	olc.pc++

	olc.addrAbs = (hi << 8) | lo

	return 0
}

func (olc *Olc6502) ABX() uint8 {
	lo := uint16(olc.read(olc.pc))
	olc.pc++

	hi := uint16(olc.read(olc.pc))
	olc.pc++

	olc.addrAbs = (hi << 8) | lo
	olc.addrAbs += uint16(olc.x)

	if (olc.addrAbs & 0x00FF) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

func (olc *Olc6502) ABY() uint8 {
	lo := uint16(olc.read(olc.pc))
	olc.pc++

	hi := uint16(olc.read(olc.pc))
	olc.pc++

	olc.addrAbs = (hi << 8) | lo
	olc.addrAbs += uint16(olc.y)

	if (olc.addrAbs & 0x00FF) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

func (olc *Olc6502) IND() uint8 {
	ptrLo := uint16(olc.read(olc.pc))
	olc.pc++

	ptrHi := uint16(olc.read(olc.pc))
	olc.pc++

	ptr := (ptrHi << 8) | ptrLo

	if ptrLo == 0x00FF {
		//hardware bug simulation
		olc.addrAbs = (uint16(olc.read(ptr&0xFF00)) << 8) | uint16(olc.read(ptr))
	} else {
		olc.addrAbs = (uint16(olc.read(ptr+1)) << 8) | uint16(olc.read(ptr))
	}

	return 0
}

func (olc *Olc6502) IZX() uint8 {
	t := uint16(olc.read(olc.pc))
	olc.pc++

	lo := uint16(olc.read(t+uint16(olc.x)) & 0x00FF)
	hi := uint16(olc.read(t+uint16(olc.x)+1) & 0x00FF)

	olc.addrAbs = (hi << 8) | lo

	return 0
}

func (olc *Olc6502) IZY() uint8 {
	t := uint16(olc.read(olc.pc))
	olc.pc++

	lo := uint16(olc.read(t & 0x00FF))
	hi := uint16(olc.read(t + 1&0x00FF))

	olc.addrAbs = ((hi << 8) | lo) + uint16(olc.y)

	if olc.addrAbs&0xFF00 != (hi << 8) {
		return 1
	} else {
		return 0
	}
}
