package cpu

const (
	OP_ADC byte = iota
	OP_AND
	OP_ASL
	OP_BCC
	OP_BCS
	OP_BEQ
	OP_BIT
	OP_BMI
	OP_BNE
	OP_BPL
	OP_BRK
	OP_BVC
	OP_BVS
	OP_CLC
	OP_CLD
	OP_CLI
	OP_CLV
	OP_CMP
	OP_CPX
	OP_CPY
	OP_DEC
	OP_DEX
	OP_DEY
	OP_EOR
	OP_INC
	OP_INX
	OP_INY
	OP_JMP
	OP_JSR
	OP_LDA
	OP_LDX
	OP_LDY
	OP_LSR
	OP_NOP
	OP_ORA
	OP_PHA
	OP_PHP
	OP_PLA
	OP_PLP
	OP_ROL
	OP_ROR
	OP_RTI
	OP_RTS
	OP_SBC
	OP_SEC
	OP_SED
	OP_SEI
	OP_STA
	OP_STX
	OP_STY
	OP_TAX
	OP_TAY
	OP_TSX
	OP_TXA
	OP_TXS
	OP_TYA

	OP_ILL // illegal
	OP_UNK // unknown
)

// add with carry
func (cpu *Olc6502) ADC() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.a) + uint16(cpu.fetched) + uint16(cpu.getFlag(FLAG_C))

	cpu.setFlag(FLAG_C, tmp > 255)
	cpu.setFlag(FLAG_Z, tmp&0x00FF == 0)
	cpu.setFlag(FLAG_N, tmp&0x80 != 0)
	cpu.setFlag(FLAG_V, (^(uint16(cpu.a)^uint16(cpu.fetched))&(uint16(cpu.a)^tmp))&0x0080 != 0)

	cpu.a = uint8(tmp & 0x00FF)

	return 1
}

// and (with accumulator)
func (cpu *Olc6502) AND() uint8 {
	cpu.fetch()

	cpu.a &= cpu.fetched

	cpu.setFlag(FLAG_Z, cpu.a == 0x00)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 1
}

// arithmetic shift left
func (cpu *Olc6502) ASL() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.fetched) << 1

	cpu.setFlag(FLAG_C, (tmp&0xFF00) > 0)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, (tmp&0x80) != 0)

	if lookup[cpu.opcode].AMName == AM_IMP {
		cpu.a = uint8(tmp & 0x00FF)
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))

	}

	return 0
}

// branch on carry clear
func (cpu *Olc6502) BCC() uint8 {
	if cpu.getFlag(FLAG_C) == 0 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// branch on carry set
func (cpu *Olc6502) BCS() uint8 {
	if cpu.getFlag(FLAG_C) == 1 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// branch on equal (zero set)
func (cpu *Olc6502) BEQ() uint8 {
	if cpu.getFlag(FLAG_Z) == 1 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// bit test
func (cpu *Olc6502) BIT() uint8 {
	cpu.fetch()

	tmp := cpu.a & cpu.fetched

	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, cpu.fetched&(1<<7) != 0)
	cpu.setFlag(FLAG_V, cpu.fetched&(1<<6) != 0)

	return 0
}

// branch on minus (negative set)
func (cpu *Olc6502) BMI() uint8 {
	if cpu.getFlag(FLAG_N) == 1 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// branch on not equal (zero clear)
func (cpu *Olc6502) BNE() uint8 {
	if cpu.getFlag(FLAG_Z) == 0 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// branch on plus (negative clear)
func (cpu *Olc6502) BPL() uint8 {
	if cpu.getFlag(FLAG_N) == 0 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// break / interrupt
func (cpu *Olc6502) BRK() uint8 {
	cpu.pc++

	cpu.setFlag(FLAG_I, true)
	cpu.write(0x0100+uint16(cpu.stkp), uint8((cpu.pc>>8)&0x00FF))
	cpu.stkp--
	cpu.write(0x0100+uint16(cpu.stkp), uint8(cpu.pc&0x00FF))
	cpu.stkp--

	cpu.setFlag(FLAG_B, true)
	cpu.write(0x0100+uint16(cpu.stkp), cpu.status)
	cpu.stkp--
	cpu.setFlag(FLAG_B, false)

	cpu.pc = uint16(cpu.read(0xFFFE)) | (uint16(cpu.read(0xFFFF)) << 8)

	return 0
}

// branch on overflow clear
func (cpu *Olc6502) BVC() uint8 {
	if cpu.getFlag(FLAG_V) == 0 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// branch on overflow set
func (cpu *Olc6502) BVS() uint8 {
	if cpu.getFlag(FLAG_V) == 1 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if cpu.addrAbs&0xFF00 != cpu.pc&0xFF00 {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// clear carry
func (cpu *Olc6502) CLC() uint8 {
	cpu.setFlag(FLAG_C, false)

	return 0
}

// clear decimal
func (cpu *Olc6502) CLD() uint8 {
	cpu.setFlag(FLAG_D, false)

	return 0
}

// clear interrupt disable
func (cpu *Olc6502) CLI() uint8 {
	cpu.setFlag(FLAG_I, false)

	return 0
}

// clear overflow
func (cpu *Olc6502) CLV() uint8 {
	cpu.setFlag(FLAG_V, false)

	return 0
}

// compare (with accumulator)
func (cpu *Olc6502) CMP() uint8 {
	cpu.fetch()

	tmp := cpu.a - cpu.fetched

	cpu.setFlag(FLAG_C, cpu.a >= cpu.fetched)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0x0000)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 1
}

// compare with X
func (cpu *Olc6502) CPX() uint8 {
	cpu.fetch()

	tmp := cpu.x - cpu.fetched

	cpu.setFlag(FLAG_C, cpu.x >= cpu.fetched)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0x0000)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

// compare with Y
func (cpu *Olc6502) CPY() uint8 {
	cpu.fetch()

	tmp := cpu.y - cpu.fetched

	cpu.setFlag(FLAG_C, cpu.y >= cpu.fetched)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0x0000)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

// decrement
func (cpu *Olc6502) DEC() uint8 {
	cpu.fetch()

	tmp := cpu.fetched - 1

	cpu.write(cpu.addrAbs, tmp&0x00FF)

	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

// decrement X
func (cpu *Olc6502) DEX() uint8 {
	cpu.x--
	cpu.setFlag(FLAG_Z, cpu.x == 0x00)
	cpu.setFlag(FLAG_N, (cpu.x&0x80) != 0)

	return 0
}

// decrement Y
func (cpu *Olc6502) DEY() uint8 {
	cpu.y--
	cpu.setFlag(FLAG_Z, cpu.y == 0x00)
	cpu.setFlag(FLAG_N, (cpu.y&0x80) != 0)

	return 0
}

// exclusive or (with accumulator)
func (cpu *Olc6502) EOR() uint8 {
	cpu.fetch()

	cpu.a = cpu.a ^ cpu.fetched

	cpu.setFlag(FLAG_Z, cpu.a == 0)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 1
}

// increment
func (cpu *Olc6502) INC() uint8 {
	cpu.fetch()

	tmp := cpu.fetched + 1

	cpu.write(cpu.addrAbs, tmp&0x00FF)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	return 0
}

// increment X
func (cpu *Olc6502) INX() uint8 {
	cpu.x++

	cpu.setFlag(FLAG_Z, cpu.x == 0)
	cpu.setFlag(FLAG_N, (cpu.x&0x80) != 0)

	return 0
}

// increment Y
func (cpu *Olc6502) INY() uint8 {
	cpu.y++

	cpu.setFlag(FLAG_Z, cpu.y == 0)
	cpu.setFlag(FLAG_N, (cpu.y&0x80) != 0)

	return 0
}

// jump
func (cpu *Olc6502) JMP() uint8 {
	cpu.pc = cpu.addrAbs

	return 0
}

// jump subroutine
func (cpu *Olc6502) JSR() uint8 {
	cpu.pc--

	cpu.write(0x0100+uint16(cpu.stkp), uint8((cpu.pc>>8)&0x00FF))
	cpu.stkp--
	cpu.write(0x0100+uint16(cpu.stkp), uint8(cpu.pc&0x00FF))
	cpu.stkp--

	cpu.pc = cpu.addrAbs

	return 0
}

// load accumulator
func (cpu *Olc6502) LDA() uint8 {
	cpu.fetch()

	cpu.a = cpu.fetched

	cpu.setFlag(FLAG_Z, cpu.a == 0)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 1
}

// load X
func (cpu *Olc6502) LDX() uint8 {
	cpu.fetch()

	cpu.x = cpu.fetched

	cpu.setFlag(FLAG_Z, cpu.x == 0)
	cpu.setFlag(FLAG_N, (cpu.x&0x80) != 0)

	return 1
}

// load Y
func (cpu *Olc6502) LDY() uint8 {
	cpu.fetch()

	cpu.y = cpu.fetched

	cpu.setFlag(FLAG_Z, cpu.y == 0)
	cpu.setFlag(FLAG_N, (cpu.y&0x80) != 0)

	return 1
}

// logical shift right
func (cpu *Olc6502) LSR() uint8 {
	cpu.fetch()

	cpu.setFlag(FLAG_C, (cpu.fetched&0x0001) != 0)
	tmp := uint16(cpu.fetched) >> 1

	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	if lookup[cpu.opcode].AMName == AM_IMP {
		cpu.a = uint8(tmp & 0x00FF)
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// no operation
func (cpu *Olc6502) NOP() uint8 {
	// based on https://wiki.nesdev.com/w/index.php/CPU_unofficial_opcodes
	switch cpu.opcode {
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

// or with accumulator
func (cpu *Olc6502) ORA() uint8 {
	cpu.fetch()

	cpu.a = cpu.a | cpu.fetched

	cpu.setFlag(FLAG_Z, cpu.a == 0)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 1
}

// push accumulator
func (cpu *Olc6502) PHA() uint8 {
	cpu.write(0x0100+uint16(cpu.stkp), cpu.a)
	cpu.stkp--

	return 0
}

// push processor status (SR)
func (cpu *Olc6502) PHP() uint8 {
	cpu.write(0x0100+uint16(cpu.stkp), cpu.status|FLAG_B|FLAG_U)

	cpu.setFlag(FLAG_B, false)
	cpu.setFlag(FLAG_U, false)

	cpu.stkp--

	return 0
}

// pull accumulator
func (cpu *Olc6502) PLA() uint8 {
	cpu.stkp++
	cpu.a = cpu.read(0x0100 + uint16(cpu.stkp))

	cpu.setFlag(FLAG_Z, cpu.a == 0)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 0
}

// pull processor status (SR)
func (cpu *Olc6502) PLP() uint8 {
	cpu.stkp++

	cpu.status = cpu.read(0x0100 + uint16(cpu.stkp))
	cpu.setFlag(FLAG_U, true)

	return 0
}

// rotate left
func (cpu *Olc6502) ROL() uint8 {
	cpu.fetch()

	tmp := uint16((cpu.fetched << 1) | (cpu.getFlag(FLAG_C)))

	cpu.setFlag(FLAG_C, (tmp&0xFF00) != 0)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	if lookup[cpu.opcode].AMName == AM_IMP {
		cpu.a = uint8(tmp & 0x00FF)
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// rotate right
func (cpu *Olc6502) ROR() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.getFlag(FLAG_C))<<7 | uint16(cpu.fetched>>1)

	cpu.setFlag(FLAG_C, (cpu.fetched&0x01) != 0)
	cpu.setFlag(FLAG_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(FLAG_N, (tmp&0x0080) != 0)

	if lookup[cpu.opcode].AMName == AM_IMP {
		cpu.a = uint8(tmp & 0x00FF)
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// return from interrupt
func (cpu *Olc6502) RTI() uint8 {
	cpu.stkp++

	cpu.status = cpu.read(0x0100 + uint16(cpu.stkp))
	cpu.status &= ^FLAG_B
	cpu.status &= ^FLAG_U

	cpu.stkp++
	cpu.pc = uint16(cpu.read(0x0100 + uint16(cpu.stkp)))
	cpu.stkp++
	cpu.pc |= uint16(cpu.read(0x0100+uint16(cpu.stkp))) << 8

	return 0
}

// return from subroutine
func (cpu *Olc6502) RTS() uint8 {
	cpu.stkp++
	cpu.pc = uint16(cpu.read(0x0100 + uint16(cpu.stkp)))
	cpu.stkp++
	cpu.pc |= uint16(cpu.read(0x0100+uint16(cpu.stkp))) << 8

	cpu.pc++

	return 0
}

// subtract with carry
func (cpu *Olc6502) SBC() uint8 {
	cpu.fetch()

	value := uint16(cpu.fetched) ^ 0x00FF
	tmp := uint16(cpu.a) + value + uint16(cpu.getFlag(FLAG_C))

	cpu.setFlag(FLAG_C, tmp&0xFF00 != 0)
	cpu.setFlag(FLAG_Z, tmp&0x00FF == 0)
	cpu.setFlag(FLAG_V, (tmp^uint16(cpu.a)&(tmp^value))&0x0080 != 0)
	cpu.setFlag(FLAG_N, tmp&0x0080 != 0)

	cpu.a = uint8(tmp & 0x00FF)

	return 0
}

// set carry
func (cpu *Olc6502) SEC() uint8 {
	cpu.setFlag(FLAG_C, true)
	return 0
}

// set decimal
func (cpu *Olc6502) SED() uint8 {
	cpu.setFlag(FLAG_D, true)
	return 0
}

// set interrupt disable
func (cpu *Olc6502) SEI() uint8 {
	cpu.setFlag(FLAG_I, true)
	return 0
}

// store accumulator
func (cpu *Olc6502) STA() uint8 {
	cpu.write(cpu.addrAbs, cpu.a)
	return 0
}

// store X
func (cpu *Olc6502) STX() uint8 {
	cpu.write(cpu.addrAbs, cpu.x)
	return 0
}

// store Y
func (cpu *Olc6502) STY() uint8 {
	cpu.write(cpu.addrAbs, cpu.y)
	return 0
}

// transfer accumulator to X
func (cpu *Olc6502) TAX() uint8 {
	cpu.x = cpu.a

	cpu.setFlag(FLAG_Z, cpu.x == 0)
	cpu.setFlag(FLAG_N, (cpu.x&0x80) != 0)

	return 0
}

// transfer accumulator to Y
func (cpu *Olc6502) TAY() uint8 {
	cpu.y = cpu.a

	cpu.setFlag(FLAG_Z, cpu.y == 0)
	cpu.setFlag(FLAG_N, (cpu.y&0x80) != 0)

	return 0
}

// transfer stack pointer to X
func (cpu *Olc6502) TSX() uint8 {
	cpu.x = cpu.a

	cpu.setFlag(FLAG_Z, cpu.x == 0)
	cpu.setFlag(FLAG_N, (cpu.x&0x80) != 0)

	return 0
}

// transfer X to accumulator
func (cpu *Olc6502) TXA() uint8 {
	cpu.a = cpu.x

	cpu.setFlag(FLAG_Z, cpu.a == 0)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 0
}

// transfer X to stack pointer
func (cpu *Olc6502) TXS() uint8 {
	cpu.stkp = cpu.x

	return 0
}

// transfer Y to accumulator
func (cpu *Olc6502) TYA() uint8 {
	cpu.a = cpu.y

	cpu.setFlag(FLAG_Z, cpu.a == 0)
	cpu.setFlag(FLAG_N, (cpu.a&0x80) != 0)

	return 0
}

// illegal opcode
func (cpu *Olc6502) ILL() uint8 {
	return 0
}
