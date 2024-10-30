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
// checked
func (cpu *MOSTechnology6502) ADC() uint8 {
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
// checked
func (cpu *MOSTechnology6502) AND() uint8 {
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
// checked
func (cpu *MOSTechnology6502) ASL() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BCC() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BCS() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BEQ() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BIT() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.a & cpu.fetched)

	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, cpu.fetched&(1<<7) != 0)
	cpu.setFlag(flag_V, cpu.fetched&(1<<6) != 0)

	return 0
}

// branch on minus (negative set)
// checked
func (cpu *MOSTechnology6502) BMI() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BNE() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BPL() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BRK() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BVC() uint8 {
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
// checked
func (cpu *MOSTechnology6502) BVS() uint8 {
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
// checked
func (cpu *MOSTechnology6502) CLC() uint8 {
	cpu.setFlag(flag_C, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// clear decimal
// checked
func (cpu *MOSTechnology6502) CLD() uint8 {
	cpu.setFlag(flag_D, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// clear interrupt disable
// checked
func (cpu *MOSTechnology6502) CLI() uint8 {
	cpu.setFlag(flag_I, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// clear overflow
// checked
func (cpu *MOSTechnology6502) CLV() uint8 {
	cpu.setFlag(flag_V, false)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// compare (with accumulator)
// checked
func (cpu *MOSTechnology6502) CMP() uint8 {
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
// checked
func (cpu *MOSTechnology6502) CPX() uint8 {
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
// checked
func (cpu *MOSTechnology6502) CPY() uint8 {
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
// checked
func (cpu *MOSTechnology6502) DEC() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.fetched - 1)

	cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))

	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	return 0
}

// decrement X
// checked
func (cpu *MOSTechnology6502) DEX() uint8 {
	cpu.x--
	cpu.setFlag(flag_Z, cpu.x == 0x00)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// decrement Y
// checked
func (cpu *MOSTechnology6502) DEY() uint8 {
	cpu.y--
	cpu.setFlag(flag_Z, cpu.y == 0x00)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// exclusive or (with accumulator)
// checked
func (cpu *MOSTechnology6502) EOR() uint8 {
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
// checked
func (cpu *MOSTechnology6502) INC() uint8 {
	cpu.fetch()

	tmp := uint16(cpu.fetched) + 1

	cpu.write(cpu.addrAbs, uint8(tmp&0x00FF))
	cpu.setFlag(flag_Z, (tmp&0x00FF) == 0)
	cpu.setFlag(flag_N, (tmp&0x0080) != 0)

	return 0
}

// increment X
// checked
func (cpu *MOSTechnology6502) INX() uint8 {
	cpu.x++

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// increment Y
// checked
func (cpu *MOSTechnology6502) INY() uint8 {
	cpu.y++

	cpu.setFlag(flag_Z, cpu.y == 0)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// jump
// checked
func (cpu *MOSTechnology6502) JMP() uint8 {
	cpu.pc = cpu.addrAbs

	return 0
}

// jump subroutine
// checked
func (cpu *MOSTechnology6502) JSR() uint8 {
	cpu.pc--

	cpu.write(0x0100+uint16(cpu.stkp), uint8((cpu.pc>>8)&0x00FF))
	cpu.stkp--
	cpu.write(0x0100+uint16(cpu.stkp), uint8(cpu.pc&0x00FF))
	cpu.stkp--

	cpu.pc = cpu.addrAbs

	return 0
}

// load accumulator
// checked
func (cpu *MOSTechnology6502) LDA() uint8 {
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
// checked
func (cpu *MOSTechnology6502) LDX() uint8 {
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
// checked
func (cpu *MOSTechnology6502) LDY() uint8 {
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
// checked
func (cpu *MOSTechnology6502) LSR() uint8 {
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
// checked
func (cpu *MOSTechnology6502) NOP() uint8 {
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
// checked
func (cpu *MOSTechnology6502) ORA() uint8 {
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
// checked
func (cpu *MOSTechnology6502) PHA() uint8 {
	cpu.write(0x0100+uint16(cpu.stkp), cpu.a)
	cpu.stkp--

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// push processor status (SR)
// checked
func (cpu *MOSTechnology6502) PHP() uint8 {
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
// checked
func (cpu *MOSTechnology6502) PLA() uint8 {
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
// checked
func (cpu *MOSTechnology6502) PLP() uint8 {
	cpu.stkp++

	cpu.status = cpu.read(0x0100 + uint16(cpu.stkp))
	cpu.setFlag(flag_U, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// rotate left
// checked
func (cpu *MOSTechnology6502) ROL() uint8 {
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
// checked
func (cpu *MOSTechnology6502) ROR() uint8 {
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
// checked
func (cpu *MOSTechnology6502) RTI() uint8 {
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
// checked
func (cpu *MOSTechnology6502) RTS() uint8 {
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
// checked
func (cpu *MOSTechnology6502) SBC() uint8 {
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
// checked
func (cpu *MOSTechnology6502) SEC() uint8 {
	cpu.setFlag(flag_C, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// set decimal
// checked
func (cpu *MOSTechnology6502) SED() uint8 {
	cpu.setFlag(flag_D, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// set interrupt disable
// checked
func (cpu *MOSTechnology6502) SEI() uint8 {
	cpu.setFlag(flag_I, true)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// store accumulator
// checked
func (cpu *MOSTechnology6502) STA() uint8 {
	cpu.write(cpu.addrAbs, cpu.a)

	return 0
}

// store X
// checked
func (cpu *MOSTechnology6502) STX() uint8 {
	cpu.write(cpu.addrAbs, cpu.x)

	return 0
}

// store Y
// checked
func (cpu *MOSTechnology6502) STY() uint8 {
	cpu.write(cpu.addrAbs, cpu.y)

	return 0
}

// transfer accumulator to X
// checked
func (cpu *MOSTechnology6502) TAX() uint8 {
	cpu.x = cpu.a

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer accumulator to Y
// checked
func (cpu *MOSTechnology6502) TAY() uint8 {
	cpu.y = cpu.a

	cpu.setFlag(flag_Z, cpu.y == 0)
	cpu.setFlag(flag_N, (cpu.y&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer stack pointer to X
// checked
func (cpu *MOSTechnology6502) TSX() uint8 {
	cpu.x = cpu.stkp

	cpu.setFlag(flag_Z, cpu.x == 0)
	cpu.setFlag(flag_N, (cpu.x&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer X to accumulator
// checked
func (cpu *MOSTechnology6502) TXA() uint8 {
	cpu.a = cpu.x

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer X to stack pointer
// checked
func (cpu *MOSTechnology6502) TXS() uint8 {
	cpu.stkp = cpu.x

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// transfer Y to accumulator
// checked
func (cpu *MOSTechnology6502) TYA() uint8 {
	cpu.a = cpu.y

	cpu.setFlag(flag_Z, cpu.a == 0)
	cpu.setFlag(flag_N, (cpu.a&0x80) != 0)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 0
	}

	return 0
}

// illegal opcode
func (cpu *MOSTechnology6502) ILL() uint8 {
	return 0
}

func (cpu *MOSTechnology6502) OperationCompleted() bool {
	return cpu.cycles == 0
}
