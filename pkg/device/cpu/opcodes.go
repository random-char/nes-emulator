package cpu

const (
	OPCODE_ADC byte = iota
	OPCODE_AND
	OPCODE_ASL
	OPCODE_BCC
	OPCODE_BCS
	OPCODE_BEQ
	OPCODE_BIT
	OPCODE_BMI
	OPCODE_BNE
	OPCODE_BPL
	OPCODE_BRK
	OPCODE_BVC
	OPCODE_BVS
	OPCODE_CLC
	OPCODE_CLD
	OPCODE_CLI
	OPCODE_CLV
	OPCODE_CMP
	OPCODE_CPX
	OPCODE_CPY
	OPCODE_DEC
	OPCODE_DEX
	OPCODE_DEY
	OPCODE_EOR
	OPCODE_INC
	OPCODE_INX
	OPCODE_INY
	OPCODE_JMP
	OPCODE_JSR
	OPCODE_LDA
	OPCODE_LDX
	OPCODE_LDY
	OPCODE_LSR
	OPCODE_NOP
	OPCODE_ORA
	OPCODE_PHA
	OPCODE_PHP
	OPCODE_PLA
	OPCODE_PLP
	OPCODE_ROL
	OPCODE_ROR
	OPCODE_RTI
	OPCODE_RTS
	OPCODE_SBC
	OPCODE_SEC
	OPCODE_SED
	OPCODE_SEI
	OPCODE_STA
	OPCODE_STX
	OPCODE_STY
	OPCODE_TAX
	OPCODE_TAY
	OPCODE_TSX
	OPCODE_TXA
	OPCODE_TXS
	OPCODE_TYA

	OPCODE_ILL // illegal
	OPCODE_UNK // unknown
)

func (olc *Olc6502) ADC() uint8 { // add with carry
	olc.fetch()

	tmp := uint16(olc.a) + uint16(olc.fetched) + uint16(olc.getFlag(FLAG_C))

	olc.setFlag(FLAG_C, tmp > 255)
	olc.setFlag(FLAG_Z, tmp&0x00FF == 0)
	olc.setFlag(FLAG_N, tmp&0x80 != 0)
	olc.setFlag(FLAG_V, (^(uint16(olc.a)^uint16(olc.fetched))&(uint16(olc.a)^tmp))&0x0080 != 0)

	olc.a = uint8(tmp & 0x00FF)

	return 1
}

func (olc *Olc6502) AND() uint8 { // and (with accumulator)
	olc.fetch()

	olc.a &= olc.fetched

	olc.setFlag(FLAG_Z, olc.a == 0x00)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 1
}

func (olc *Olc6502) ASL() uint8 { // arithmetic shift left
	olc.fetch()

	tmp := uint16(olc.fetched) << 1

	olc.setFlag(FLAG_C, (tmp&0xFF00) > 0)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, (tmp&0x80) != 0)

	if lookup[olc.opcode].AddrModeName == ADDR_MODE_IMP {
		olc.a = uint8(tmp & 0x00FF)
	} else {
		olc.write(olc.addrAbs, uint8(tmp&0x00FF))

	}

	return 0
}

func (olc *Olc6502) BCC() uint8 { // branch on carry clear
	if olc.getFlag(FLAG_C) == 0 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BCS() uint8 { // branch on carry set
	if olc.getFlag(FLAG_C) == 1 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BEQ() uint8 { // branch on equal (zero set)
	if olc.getFlag(FLAG_Z) == 1 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BIT() uint8 { // bit test
	olc.fetch()

	tmp := olc.a & olc.fetched

	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, olc.fetched&(1<<7) != 0)
	olc.setFlag(FLAG_V, olc.fetched&(1<<6) != 0)

	return 0
}

func (olc *Olc6502) BMI() uint8 { // branch on minus (negative set)
	if olc.getFlag(FLAG_N) == 1 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BNE() uint8 { // branch on not equal (zero clear)
	if olc.getFlag(FLAG_Z) == 0 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BPL() uint8 { // branch on plus (negative clear)
	if olc.getFlag(FLAG_N) == 0 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BRK() uint8 { // break / interrupt
	olc.pc++

	olc.setFlag(FLAG_I, true)
	olc.write(0x0100+uint16(olc.stkp), uint8((olc.pc>>8)&0x00FF))
	olc.stkp--
	olc.write(0x0100+uint16(olc.stkp), uint8(olc.pc&0x00FF))
	olc.stkp--

	olc.setFlag(FLAG_B, true)
	olc.write(0x0100+uint16(olc.stkp), olc.status)
	olc.stkp--
	olc.setFlag(FLAG_B, false)

	olc.pc = uint16(olc.read(0xFFFE)) | (uint16(olc.read(0xFFFF)) << 8)

	return 0
}

func (olc *Olc6502) BVC() uint8 { // branch on overflow clear
	if olc.getFlag(FLAG_V) == 0 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) BVS() uint8 { // branch on overflow set
	if olc.getFlag(FLAG_V) == 1 {
		olc.cycles++

		olc.addrAbs = olc.pc + olc.addrRel
		if olc.addrAbs&0xFF00 != olc.pc&0xFF00 {
			olc.cycles++
		}

		olc.pc = olc.addrAbs
	}

	return 0
}

func (olc *Olc6502) CLC() uint8 { // clear carry
	olc.setFlag(FLAG_C, false)

	return 0
}

func (olc *Olc6502) CLD() uint8 { // clear decimal
	olc.setFlag(FLAG_D, false)

	return 0
}

func (olc *Olc6502) CLI() uint8 { // clear interrupt disable
	olc.setFlag(FLAG_I, false)

	return 0
}

func (olc *Olc6502) CLV() uint8 { // clear overflow
	olc.setFlag(FLAG_V, false)

	return 0
}

func (olc *Olc6502) CMP() uint8 { // compare (with accumulator)
	olc.fetch()

	tmp := olc.a - olc.fetched

	olc.setFlag(FLAG_C, olc.a >= olc.fetched)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0x0000)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 1
}

func (olc *Olc6502) CPX() uint8 { // compare with X
	olc.fetch()

	tmp := olc.x - olc.fetched

	olc.setFlag(FLAG_C, olc.x >= olc.fetched)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0x0000)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

func (olc *Olc6502) CPY() uint8 { // compare with Y
	olc.fetch()

	tmp := olc.y - olc.fetched

	olc.setFlag(FLAG_C, olc.y >= olc.fetched)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0x0000)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

func (olc *Olc6502) DEC() uint8 { // decrement
	olc.fetch()

	tmp := olc.fetched - 1

	olc.write(olc.addrAbs, tmp&0x00FF)

	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

func (olc *Olc6502) DEX() uint8 { // decrement X
	olc.x--
	olc.setFlag(FLAG_Z, olc.x == 0x00)
	olc.setFlag(FLAG_N, (olc.x&0x80) != 0)

	return 0
}

func (olc *Olc6502) DEY() uint8 { // decrement Y
	olc.y--
	olc.setFlag(FLAG_Z, olc.y == 0x00)
	olc.setFlag(FLAG_N, (olc.y&0x80) != 0)

	return 0
}

func (olc *Olc6502) EOR() uint8 { // exclusive or (with accumulator)
	olc.fetch()

	olc.a = olc.a ^ olc.fetched

	olc.setFlag(FLAG_Z, olc.a == 0)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 1
}

func (olc *Olc6502) INC() uint8 { // increment
	olc.fetch()

	tmp := olc.fetched + 1

	olc.write(olc.addrAbs, tmp&0x00FF)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

func (olc *Olc6502) INX() uint8 { // increment X
	olc.x++

	olc.setFlag(FLAG_Z, olc.x == 0)
	olc.setFlag(FLAG_N, (olc.x&0x80) != 0)

	return 0
}

func (olc *Olc6502) INY() uint8 { // increment Y
	olc.y++

	olc.setFlag(FLAG_Z, olc.y == 0)
	olc.setFlag(FLAG_N, (olc.y&0x80) != 0)

	return 0
}

func (olc *Olc6502) JMP() uint8 { // jump
	olc.pc = olc.addrAbs

	return 0
}

func (olc *Olc6502) JSR() uint8 { // jump subroutine
	olc.pc--

	olc.write(0x0100+uint16(olc.stkp), uint8((olc.pc>>8)&0x00FF))
	olc.stkp--
	olc.write(0x0100+uint16(olc.stkp), uint8(olc.pc&0x00FF))
	olc.stkp--

	olc.pc = olc.addrAbs

	return 0
}

func (olc *Olc6502) LDA() uint8 { // load accumulator
	olc.fetch()

	olc.a = olc.fetched

	olc.setFlag(FLAG_Z, olc.a == 0)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 1
}

func (olc *Olc6502) LDX() uint8 { // load X
	olc.fetch()

	olc.x = olc.fetched

	olc.setFlag(FLAG_Z, olc.x == 0)
	olc.setFlag(FLAG_N, (olc.x&0x80) != 0)

	return 1
}

func (olc *Olc6502) LDY() uint8 { // load Y
	olc.fetch()

	olc.y = olc.fetched

	olc.setFlag(FLAG_Z, olc.y == 0)
	olc.setFlag(FLAG_N, (olc.y&0x80) != 0)

	return 1
}

func (olc *Olc6502) LSR() uint8 { // logical shift right
	olc.fetch()

	olc.setFlag(FLAG_C, (olc.fetched&0x0001) != 0)
	tmp := uint16(olc.fetched) >> 1

	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	if lookup[olc.opcode].AddrModeName == ADDR_MODE_IMP {
		olc.a = uint8(tmp & 0x00FF)
	} else {
		olc.write(olc.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

func (olc *Olc6502) NOP() uint8 { // no operation
	// based on https://wiki.nesdev.com/w/index.php/CPU_unofficial_opcodes
	switch olc.opcode {
	case 0x1C:
	case 0x3C:
	case 0x5C:
	case 0x7C:
	case 0xDC:
	case 0xFC:
		return 1
	}

	return 0
}

func (olc *Olc6502) ORA() uint8 { // or with accumulator
	olc.fetch()

	olc.a = olc.a | olc.fetched

	olc.setFlag(FLAG_Z, olc.a == 0)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 1
}

func (olc *Olc6502) PHA() uint8 { // push accumulator
	olc.write(0x0100+uint16(olc.stkp), olc.a)
	olc.stkp--

	return 0
}

func (olc *Olc6502) PHP() uint8 { // push processor status (SR)
	olc.write(0x0100+uint16(olc.stkp), olc.status|FLAG_B|FLAG_U)

	olc.setFlag(FLAG_B, false)
	olc.setFlag(FLAG_U, false)

	olc.stkp--

	return 0
}

// pull accumulator
func (olc *Olc6502) PLA() uint8 {
	olc.stkp++
	olc.a = olc.read(0x0100 + uint16(olc.stkp))

	olc.setFlag(FLAG_Z, olc.a == 0)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 0
}

// pull processor status (SR)
func (olc *Olc6502) PLP() uint8 {
	olc.stkp++

	olc.status = olc.read(0x0100 + uint16(olc.stkp))
	olc.setFlag(FLAG_U, true)

	return 0
}

func (olc *Olc6502) ROL() uint8 { // rotate left
	olc.fetch()

	tmp := uint16((olc.fetched << 1) | (olc.getFlag(FLAG_C)))

	olc.setFlag(FLAG_C, (tmp&0xFF00) != 0)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	if lookup[olc.opcode].AddrModeName == ADDR_MODE_IMP {
		olc.a = uint8(tmp & 0x00FF)
	} else {
		olc.write(olc.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

func (olc *Olc6502) ROR() uint8 { // rotate right
	olc.fetch()

	tmp := uint16(olc.getFlag(FLAG_C))<<7 | uint16(olc.fetched>>1)

	olc.setFlag(FLAG_C, (olc.fetched&0x01) != 0)
	olc.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	olc.setFlag(FLAG_N, (tmp&0x0080) != 0)

	if lookup[olc.opcode].AddrModeName == ADDR_MODE_IMP {
		olc.a = uint8(tmp & 0x00FF)
	} else {
		olc.write(olc.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

func (olc *Olc6502) RTI() uint8 { // return from interrupt
	olc.stkp++

	olc.status = olc.read(0x0100 + uint16(olc.stkp))
	olc.status &= ^FLAG_B
	olc.status &= ^FLAG_U

	olc.stkp++
	olc.pc = uint16(olc.read(0x0100 + uint16(olc.stkp)))
	olc.stkp++
	olc.pc |= uint16(olc.read(0x0100+uint16(olc.stkp))) << 8

	return 0
}

func (olc *Olc6502) RTS() uint8 { // return from subroutine
	olc.stkp++
	olc.pc = uint16(olc.read(0x0100 + uint16(olc.stkp)))
	olc.stkp++
	olc.pc |= uint16(olc.read(0x0100+uint16(olc.stkp))) << 8

	olc.pc++

	return 0
}

func (olc *Olc6502) SBC() uint8 { // subtract with carry
	olc.fetch()

	value := uint16(olc.fetched) ^ 0x00FF
	tmp := uint16(olc.a) + value + uint16(olc.getFlag(FLAG_C))

	olc.setFlag(FLAG_C, tmp&0xFF00 != 0)
	olc.setFlag(FLAG_Z, tmp&0x00FF == 0)
	olc.setFlag(FLAG_V, (tmp^uint16(olc.a)&(tmp^value))&0x0080 != 0)
	olc.setFlag(FLAG_N, tmp&0x0080 != 0)

	olc.a = uint8(tmp & 0x00FF)

	return 0
}

func (olc *Olc6502) SEC() uint8 { // set carry
	olc.setFlag(FLAG_C, true)
	return 0
}

func (olc *Olc6502) SED() uint8 { // set decimal
	olc.setFlag(FLAG_D, true)
	return 0
}

func (olc *Olc6502) SEI() uint8 { // set interrupt disable
	olc.setFlag(FLAG_I, true)
	return 0
}

func (olc *Olc6502) STA() uint8 { // store accumulator
	olc.write(olc.addrAbs, olc.a)
	return 0
}

func (olc *Olc6502) STX() uint8 { // store X
	olc.write(olc.addrAbs, olc.x)
	return 0
}

func (olc *Olc6502) STY() uint8 { // store Y
	olc.write(olc.addrAbs, olc.y)
	return 0
}

func (olc *Olc6502) TAX() uint8 { // transfer accumulator to X
	olc.x = olc.a

	olc.setFlag(FLAG_Z, olc.x == 0)
	olc.setFlag(FLAG_N, (olc.x&0x80) != 0)

	return 0
}

func (olc *Olc6502) TAY() uint8 { // transfer accumulator to Y
	olc.y = olc.a

	olc.setFlag(FLAG_Z, olc.y == 0)
	olc.setFlag(FLAG_N, (olc.y&0x80) != 0)

	return 0
}

func (olc *Olc6502) TSX() uint8 { // transfer stack pointer to X
	olc.x = olc.a

	olc.setFlag(FLAG_Z, olc.x == 0)
	olc.setFlag(FLAG_N, (olc.x&0x80) != 0)

	return 0
}

func (olc *Olc6502) TXA() uint8 { // transfer X to accumulator
	olc.a = olc.x

	olc.setFlag(FLAG_Z, olc.a == 0)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 0
}

func (olc *Olc6502) TXS() uint8 { // transfer X to stack pointer
    olc.stkp = olc.x

	return 0
}

func (olc *Olc6502) TYA() uint8 { // transfer Y to accumulator
	olc.a = olc.y

	olc.setFlag(FLAG_Z, olc.a == 0)
	olc.setFlag(FLAG_N, (olc.a&0x80) != 0)

	return 0
}

func (olc *Olc6502) ILL() uint8 { // illegal opcode
	return 0
}
