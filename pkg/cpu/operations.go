package cpu

const (
	op_ADC = "ADC"
	op_AND = "AND"
	op_ASL = "ASL"
	op_BCC = "BCC"
	op_BCS = "BCS"
	op_BEQ = "BEQ"
	op_BIT = "BIT"
	op_BMI = "BMI"
	op_BNE = "BNE"
	op_BPL = "BPL"
	op_BRK = "BRK"
	op_BVC = "BVC"
	op_BVS = "BVS"
	op_CLC = "CLC"
	op_CLD = "CLD"
	op_CLI = "CLI"
	op_CLV = "CLV"
	op_CMP = "CMP"
	op_CPX = "CPX"
	op_CPY = "CPY"
	op_DEC = "DEC"
	op_DEX = "DEX"
	op_DEY = "DEY"
	op_EOR = "EOR"
	op_INC = "INC"
	op_INX = "INX"
	op_INY = "INY"
	op_JMP = "JMP"
	op_JSR = "JSR"
	op_LDA = "LDA"
	op_LDX = "LDX"
	op_LDY = "LDY"
	op_LSR = "LSR"
	op_NOP = "NOP"
	op_ORA = "ORA"
	op_PHA = "PHA"
	op_PHP = "PHP"
	op_PLA = "PLA"
	op_PLP = "PLP"
	op_ROL = "ROL"
	op_ROR = "ROR"
	op_RTI = "RTI"
	op_RTS = "RTS"
	op_SBC = "SBC"
	op_SEC = "SEC"
	op_SED = "SED"
	op_SEI = "SEI"
	op_STA = "STA"
	op_STX = "STX"
	op_STY = "STY"
	op_TAX = "TAX"
	op_TAY = "TAY"
	op_TSX = "TSX"
	op_TXA = "TXA"
	op_TXS = "TXS"
	op_TYA = "TYA"
	op_ILL = "*NOP"
)

// add with carry
func (cpu *MOSTechnology6502) adc() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.a) + uint16(cpu.fetched) + uint16(cpu.getFlag(flag_C))

	cpu.setFlag(flag_C, tmp > 255)
	cpu.setFlag(flag_Z, tmp&0x00FF == 0)
	cpu.setFlag(flag_N, tmp&0x80 != 0)
	cpu.setFlag(flag_V, ((^(uint16(cpu.a)^uint16(cpu.fetched))&(uint16(cpu.a)^tmp))&0x0080) != 0)

	cpu.a = uint8(tmp & 0x00FF)

	if cpu.debugger != nil && cpu.debugger.CpuNumOperands == 0 {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 1
}

// and (with accumulator)
func (cpu *MOSTechnology6502) and() uint8 {
	cpu.fetch()

	cpu.a &= cpu.fetched

	cpu.setFlag(flag_Z, cpu.a == 0x00)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		if cpu.debugger.CpuNumOperands == 0 {
			cpu.debugger.CpuNumOperands = 1
			cpu.debugger.CpuOperand1 = cpu.fetched
		}
	}

	return 1
}

// arithmetic shift left
func (cpu *MOSTechnology6502) asl() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.fetched) << 1

	cpu.setFlag(flag_C, (tmp&0xFF00) > 0)
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x80) != 0)

	if lookup[cpu.opcode].AMName == am_IMP {
		cpu.a = uint8(tmp & 0x00FF)

		if cpu.debugger != nil {
			cpu.debugger.CpuNumOperands = 0
		}
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// branch on carry clear
func (cpu *MOSTechnology6502) bcc() uint8 {
	if cpu.getFlag(flag_C) == 0 {
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
func (cpu *MOSTechnology6502) bcs() uint8 {
	if cpu.getFlag(flag_C) == 1 {
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
func (cpu *MOSTechnology6502) beq() uint8 {
	if cpu.getFlag(flag_Z) == 1 {
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
func (cpu *MOSTechnology6502) bit() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.a & cpu.fetched)

	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, cpu.fetched&(1<<7) != 0)
	cpu.setFlag(flag_V, cpu.fetched&(1<<6) != 0)

	return 0
}

// branch on minus (negative set)
func (cpu *MOSTechnology6502) bmi() uint8 {
	if cpu.getFlag(flag_N) == 1 {
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
func (cpu *MOSTechnology6502) bne() uint8 {
	if cpu.getFlag(flag_Z) == 0 {
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
func (cpu *MOSTechnology6502) bpl() uint8 {
	if cpu.getFlag(flag_N) == 0 {
		cpu.cycles++

		cpu.addrAbs = cpu.pc + cpu.addrRel
		if (cpu.addrAbs & 0xFF00) != (cpu.pc & 0xFF00) {
			cpu.cycles++
		}

		cpu.pc = cpu.addrAbs
	}

	return 0
}

// break / interrupt
func (cpu *MOSTechnology6502) brk() uint8 {
	cpu.pc++

	cpu.setFlag(flag_I, true)
	cpu.write(0x0100+uint16(cpu.stkp), uint8((cpu.pc>>8)&0x00FF))
	cpu.stkp--
	cpu.write(0x0100+uint16(cpu.stkp), uint8(cpu.pc&0x00FF))
	cpu.stkp--

	cpu.setFlag(flag_B, true)
	cpu.write(0x0100+uint16(cpu.stkp), cpu.status)
	cpu.stkp--
	cpu.setFlag(flag_B, false)

	cpu.pc = uint16(cpu.read(0xFFFE)) | (uint16(cpu.read(0xFFFF)) << 8)

	return 0
}

// branch on overflow clear
func (cpu *MOSTechnology6502) bvc() uint8 {
	if cpu.getFlag(flag_V) == 0 {
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
func (cpu *MOSTechnology6502) bvs() uint8 {
	if cpu.getFlag(flag_V) == 1 {
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
func (cpu *MOSTechnology6502) clc() uint8 {
	cpu.setFlag(flag_C, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// clear decimal
func (cpu *MOSTechnology6502) cld() uint8 {
	cpu.setFlag(flag_D, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// clear interrupt disable
func (cpu *MOSTechnology6502) cli() uint8 {
	cpu.setFlag(flag_I, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// clear overflow
func (cpu *MOSTechnology6502) clv() uint8 {
	cpu.setFlag(flag_V, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// compare (with accumulator)
func (cpu *MOSTechnology6502) cmp() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.a - cpu.fetched)

	cpu.setFlag(flag_C, cpu.a >= cpu.fetched)
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0x0000)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	if cpu.debugger != nil && cpu.debugger.CpuNumOperands == 0 {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 1
}

// compare with X
func (cpu *MOSTechnology6502) cpx() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.x - cpu.fetched)

	cpu.setFlag(flag_C, cpu.x >= cpu.fetched)
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0x0000)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	if cpu.debugger != nil && cpu.debugger.CpuNumOperands == 0 {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 0
}

// compare with Y
func (cpu *MOSTechnology6502) cpy() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.y - cpu.fetched)

	cpu.setFlag(flag_C, cpu.y >= cpu.fetched)
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0x0000)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	if cpu.debugger != nil && cpu.debugger.CpuNumOperands == 0 {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 0
}

// decrement
func (cpu *MOSTechnology6502) dec() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.fetched - 1)

	cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))

	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	return 0
}

// decrement X
func (cpu *MOSTechnology6502) dex() uint8 {
	cpu.x--
	cpu.setFlag(flag_Z, cpu.x == 0x00)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// decrement Y
func (cpu *MOSTechnology6502) dey() uint8 {
	cpu.y--
	cpu.setFlag(flag_Z, cpu.y == 0x00)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// exclusive or (with accumulator)
func (cpu *MOSTechnology6502) eor() uint8 {
	cpu.fetch()

	cpu.a = cpu.a ^ cpu.fetched

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil && cpu.debugger.CpuNumOperands == 0 {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 1
}

// increment
func (cpu *MOSTechnology6502) inc() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.fetched) + 1

	cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	return 0
}

// increment X
func (cpu *MOSTechnology6502) inx() uint8 {
	cpu.x++

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// increment Y
func (cpu *MOSTechnology6502) iny() uint8 {
	cpu.y++

	cpu.setFlag(flag_Z, cpu.y == 0)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// jump
func (cpu *MOSTechnology6502) jmp() uint8 {
	cpu.pc = cpu.addrAbs

	return 0
}

// jump subroutine
func (cpu *MOSTechnology6502) jsr() uint8 {
	cpu.pc--

	cpu.write(0x0100+uint16(cpu.stkp), uint8((cpu.pc>>8)&0x00FF))
	cpu.stkp--
	cpu.write(0x0100+uint16(cpu.stkp), uint8(cpu.pc&0x00FF))
	cpu.stkp--

	cpu.pc = cpu.addrAbs

	return 0
}

// load accumulator
func (cpu *MOSTechnology6502) lds() uint8 {
	cpu.fetch()
	cpu.a = cpu.fetched

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		if cpu.debugger.CpuNumOperands == 0 {
			cpu.debugger.CpuNumOperands = 1
			cpu.debugger.CpuOperand1 = cpu.fetched
		}
	}

	return 1
}

// load X
func (cpu *MOSTechnology6502) ldx() uint8 {
	cpu.fetch()
	cpu.x = cpu.fetched

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		if cpu.debugger.CpuNumOperands == 0 {
			cpu.debugger.CpuNumOperands = 1
			cpu.debugger.CpuOperand1 = cpu.fetched
		}
	}

	return 1
}

// load Y
func (cpu *MOSTechnology6502) ldy() uint8 {
	cpu.fetch()
	cpu.y = cpu.fetched

	cpu.setFlag(flag_Z, cpu.y == 0)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		if cpu.debugger.CpuNumOperands == 0 {
			cpu.debugger.CpuNumOperands = 1
			cpu.debugger.CpuOperand1 = cpu.fetched
		}
	}

	return 1
}

// logical shift right
func (cpu *MOSTechnology6502) lsr() uint8 {
	cpu.fetch()

	cpu.setFlag(flag_C, (cpu.fetched&0x0001) != 0)
	tmp := uint16(cpu.fetched) >> 1

	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	if lookup[cpu.opcode].AMName == am_IMP {
		cpu.a = uint8(tmp & 0x00FF)

		if cpu.debugger != nil {
			cpu.debugger.CpuNumOperands = 0
		}
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// no operation
func (cpu *MOSTechnology6502) nop() uint8 {
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
func (cpu *MOSTechnology6502) ora() uint8 {
	cpu.fetch()

	cpu.a = cpu.a | cpu.fetched

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil && cpu.debugger.CpuNumOperands == 0 {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 1
}

// push accumulator
func (cpu *MOSTechnology6502) pha() uint8 {
	cpu.write(0x0100+uint16(cpu.stkp), cpu.a)
	cpu.stkp--

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// push processor status (SR)
func (cpu *MOSTechnology6502) php() uint8 {
	cpu.write(0x0100+uint16(cpu.stkp), cpu.status|flag_B|flag_U)

	cpu.setFlag(flag_B, false)
	cpu.setFlag(flag_U, false)

	cpu.stkp--

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// pull accumulator
func (cpu *MOSTechnology6502) pla() uint8 {
	cpu.stkp++
	cpu.a = cpu.read(0x0100 + uint16(cpu.stkp))

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// pull processor status (SR)
func (cpu *MOSTechnology6502) plp() uint8 {
	cpu.stkp++

	cpu.status = cpu.read(0x0100 + uint16(cpu.stkp))
	cpu.setFlag(flag_U, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// rotate left
func (cpu *MOSTechnology6502) rol() uint8 {
	cpu.fetch()

	tmp := (uint16(cpu.fetched) << 1) | uint16(cpu.getFlag(flag_C))

	cpu.setFlag(flag_C, (tmp&0xFF00) != 0)
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	if lookup[cpu.opcode].AMName == am_IMP {
		cpu.a = uint8(tmp & 0x00FF)

		if cpu.debugger != nil {
			cpu.debugger.CpuNumOperands = 0
		}
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// rotate right
func (cpu *MOSTechnology6502) ror() uint8 {
	cpu.fetch()

	tmp := (uint16(cpu.getFlag(flag_C)) << 7) | uint16(cpu.fetched>>1)

	cpu.setFlag(flag_C, (cpu.fetched&0x01) != 0)
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	if lookup[cpu.opcode].AMName == am_IMP {
		cpu.a = uint8(tmp & 0x00FF)

		if cpu.debugger != nil {
			cpu.debugger.CpuNumOperands = 0
		}
	} else {
		cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	}

	return 0
}

// return from interrupt
func (cpu *MOSTechnology6502) rti() uint8 {
	cpu.stkp++

	cpu.status = cpu.read(0x0100 + uint16(cpu.stkp))
	cpu.status &= ^flag_B
	cpu.status &= ^flag_U

	cpu.stkp++
	cpu.pc = uint16(cpu.read(0x0100 + uint16(cpu.stkp)))
	cpu.stkp++
	cpu.pc |= uint16(cpu.read(0x0100+uint16(cpu.stkp))) << 8

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// return from subroutine
func (cpu *MOSTechnology6502) rts() uint8 {
	cpu.stkp++
	cpu.pc = uint16(cpu.read(0x0100 + uint16(cpu.stkp)))
	cpu.stkp++
	cpu.pc |= uint16(cpu.read(0x0100+uint16(cpu.stkp))) << 8

	cpu.pc++

	//recheck
	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// subtract with carry
func (cpu *MOSTechnology6502) sbc() uint8 {
	cpu.fetch()

	value := uint16(cpu.fetched) ^ 0x00FF
	tmp := uint16(cpu.a) + value + uint16(cpu.getFlag(flag_C))

	cpu.setFlag(flag_C, tmp&0xFF00 != 0)
	cpu.setFlag(flag_Z, tmp&0x00FF == 0)
	cpu.setFlag(flag_V, ((tmp^uint16(cpu.a))&(tmp^value)&0x0080) != 0)
	cpu.setFlag(flag_N, tmp&0x0080 != 0)

	cpu.a = uint8(tmp & 0x00FF)

	if cpu.debugger != nil {
		if cpu.debugger.CpuNumOperands == 0 {
			cpu.debugger.CpuNumOperands = 1
			cpu.debugger.CpuOperand1 = cpu.fetched
		}
	}

	return 1
}

// set carry
func (cpu *MOSTechnology6502) sec() uint8 {
	cpu.setFlag(flag_C, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// set decimal
func (cpu *MOSTechnology6502) sed() uint8 {
	cpu.setFlag(flag_D, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// set interrupt disable
func (cpu *MOSTechnology6502) sei() uint8 {
	cpu.setFlag(flag_I, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// store accumulator
func (cpu *MOSTechnology6502) sta() uint8 {
	cpu.write(cpu.addrAbs, cpu.a)

	return 0
}

// store X
func (cpu *MOSTechnology6502) stx() uint8 {
	cpu.write(cpu.addrAbs, cpu.x)

	return 0
}

// store Y
func (cpu *MOSTechnology6502) sty() uint8 {
	cpu.write(cpu.addrAbs, cpu.y)

	return 0
}

// transfer accumulator to X
func (cpu *MOSTechnology6502) tax() uint8 {
	cpu.x = cpu.a

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer accumulator to Y
func (cpu *MOSTechnology6502) tay() uint8 {
	cpu.y = cpu.a

	cpu.setFlag(flag_Z, cpu.y == 0)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer stack pointer to X
func (cpu *MOSTechnology6502) tsx() uint8 {
	cpu.x = cpu.stkp

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer X to accumulator
func (cpu *MOSTechnology6502) txa() uint8 {
	cpu.a = cpu.x

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer X to stack pointer
func (cpu *MOSTechnology6502) txs() uint8 {
	cpu.stkp = cpu.x

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer Y to accumulator
func (cpu *MOSTechnology6502) tya() uint8 {
	cpu.a = cpu.y

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// illegal opcode
func (cpu *MOSTechnology6502) ill() uint8 {
	return 0
}

func (cpu *MOSTechnology6502) OperationCompleted() bool {
	return cpu.cycles == 0
}
