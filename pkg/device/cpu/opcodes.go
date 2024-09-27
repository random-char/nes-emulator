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
	olc.Fetch()

	tmp := uint16(olc.a) + uint16(olc.fetched) + uint16(olc.GetFlag(FLAG_C))

	olc.SetFlag(FLAG_C, tmp > 255)
	olc.SetFlag(FLAG_Z, tmp&0x00FF == 0)
	olc.SetFlag(FLAG_N, tmp&0x80 != 0)
	olc.SetFlag(FLAG_V, (^(uint16(olc.a)^uint16(olc.fetched))&(uint16(olc.a)^tmp))&0x0080 != 0)

	olc.a = uint8(tmp & 0x00FF)

	return 1
}

func (olc *Olc6502) AND() uint8 { // and (with accumulator)
	olc.Fetch()

	olc.a &= olc.fetched

	olc.SetFlag(FLAG_Z, olc.a == 0x00)
	olc.SetFlag(FLAG_N, (olc.a&0x80) != 0)

	return 1
}

func (olc *Olc6502) ASL() uint8 { // arithmetic shift left
	return 0
}

func (olc *Olc6502) BCC() uint8 { // branch on carry clear
	if olc.GetFlag(FLAG_C) == 0 {
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
	if olc.GetFlag(FLAG_C) == 1 {
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
	if olc.GetFlag(FLAG_Z) == 1 {
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
	return 0
}

func (olc *Olc6502) BMI() uint8 { // branch on minus (negative set)
	if olc.GetFlag(FLAG_N) == 1 {
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
	if olc.GetFlag(FLAG_Z) == 0 {
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
	if olc.GetFlag(FLAG_N) == 0 {
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
	return 0
}

func (olc *Olc6502) BVC() uint8 { // branch on overflow clear
	if olc.GetFlag(FLAG_V) == 0 {
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
	if olc.GetFlag(FLAG_V) == 1 {
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
	olc.SetFlag(FLAG_C, false)

	return 0
}

func (olc *Olc6502) CLD() uint8 { // clear decimal
	olc.SetFlag(FLAG_D, false)

	return 0
}

func (olc *Olc6502) CLI() uint8 { // clear interrupt disable
	olc.SetFlag(FLAG_I, false)

	return 0
}

func (olc *Olc6502) CLV() uint8 { // clear overflow
	olc.SetFlag(FLAG_V, false)

	return 0
}

func (olc *Olc6502) CMP() uint8 { // compare (with accumulator)
	return 0
}

func (olc *Olc6502) CPX() uint8 { // compare with X
	return 0
}

func (olc *Olc6502) CPY() uint8 { // compare with Y
	return 0
}

func (olc *Olc6502) DEC() uint8 { // decrement
	return 0
}

func (olc *Olc6502) DEX() uint8 { // decrement X
	return 0
}

func (olc *Olc6502) DEY() uint8 { // decrement Y
	return 0
}

func (olc *Olc6502) EOR() uint8 { // exclusive or (with accumulator)
	return 0
}

func (olc *Olc6502) INC() uint8 { // increment
	return 0
}

func (olc *Olc6502) INX() uint8 { // increment X
	return 0
}

func (olc *Olc6502) INY() uint8 { // increment Y
	return 0
}

func (olc *Olc6502) JMP() uint8 { // jump
	return 0
}

func (olc *Olc6502) JSR() uint8 { // jump subroutine
	return 0
}

func (olc *Olc6502) LDA() uint8 { // load accumulator
	return 0
}

func (olc *Olc6502) LDX() uint8 { // load X
	return 0
}

func (olc *Olc6502) LDY() uint8 { // load Y
	return 0
}

func (olc *Olc6502) LSR() uint8 { // logical shift right
	return 0
}

func (olc *Olc6502) NOP() uint8 { // no operation
	return 0
}

func (olc *Olc6502) ORA() uint8 { // or with accumulator
	return 0
}

func (olc *Olc6502) PHA() uint8 { // push accumulator
    olc.Write(0x0100 + uint16(olc.stkp), olc.a)
    olc.stkp--

	return 0
}

func (olc *Olc6502) PHP() uint8 { // push processor status (SR)
	return 0
}

func (olc *Olc6502) PLA() uint8 { // pull accumulator
    olc.stkp++
    olc.a = olc.Read(0x0100 + uint16(olc.stkp))

    olc.SetFlag(FLAG_Z, olc.a == 0)
    olc.SetFlag(FLAG_N, olc.a & 0x80 != 0)

	return 0
}

func (olc *Olc6502) PLP() uint8 { // pull processor status (SR)
	return 0
}

func (olc *Olc6502) ROL() uint8 { // rotate left
	return 0
}

func (olc *Olc6502) ROR() uint8 { // rotate right
	return 0
}

func (olc *Olc6502) RTI() uint8 { // return from interrupt
	return 0
}

func (olc *Olc6502) RTS() uint8 { // return from subroutine
	return 0
}

func (olc *Olc6502) SBC() uint8 { // subtract with carry
	olc.Fetch()

	value := uint16(olc.fetched) ^ 0x00FF
	tmp := uint16(olc.a) + value + uint16(olc.GetFlag(FLAG_C))

	olc.SetFlag(FLAG_C, tmp&0xFF00 != 0)
	olc.SetFlag(FLAG_Z, tmp&0x00FF == 0)
	olc.SetFlag(FLAG_V, (tmp^uint16(olc.a)&(tmp^value))&0x0080 != 0)
	olc.SetFlag(FLAG_N, tmp&0x0080 != 0)

	olc.a = uint8(tmp & 0x00FF)

	return 0
}

func (olc *Olc6502) SEC() uint8 { // set carry
	return 0
}

func (olc *Olc6502) SED() uint8 { // set decimal
	return 0
}

func (olc *Olc6502) SEI() uint8 { // set interrupt disable
	return 0
}

func (olc *Olc6502) STA() uint8 { // store accumulator
	return 0
}

func (olc *Olc6502) STX() uint8 { // store X
	return 0
}

func (olc *Olc6502) STY() uint8 { // store Y
	return 0
}

func (olc *Olc6502) TAX() uint8 { // transfer accumulator to X
	return 0
}

func (olc *Olc6502) TAY() uint8 { // transfer accumulator to Y
	return 0
}

func (olc *Olc6502) TSX() uint8 { // transfer stack pointer to X
	return 0
}

func (olc *Olc6502) TXA() uint8 { // transfer X to accumulator
	return 0
}

func (olc *Olc6502) TXS() uint8 { // transfer X to stack pointer
	return 0
}

func (olc *Olc6502) TYA() uint8 { // transfer Y to accumulator
	return 0
}

func (olc *Olc6502) ILL() uint8 { // illegal opcode
	return 0
}
