package cpu

const (
	am_IMP = "IMP"
	am_IMM = "IMM"
	am_ZP0 = "ZP0"
	am_ZPX = "ZPX"
	am_ZPY = "ZPY"
	am_REL = "REL"
	am_ABS = "ABS"
	am_ABX = "ABX"
	am_ABY = "ABY"
	am_IND = "IND"
	am_IZX = "IZX"
	am_IZY = "IZY"
)

// checked
func (cpu *MOSTechnology6502) IMP() uint8 {
	cpu.fetched = cpu.a

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) IMM() uint8 {
	cpu.addrAbs = cpu.pc
	cpu.pc++

	return 0
}

// checked
func (cpu *MOSTechnology6502) ZP0() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrAbs = uint16(valueRead) & 0x00FF
	cpu.pc++

	//todo recheck
	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) ZPX() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrAbs = uint16(valueRead+cpu.x) & 0x00FF
	cpu.pc++

	//todo recheck
	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) ZPY() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrAbs = uint16(valueRead+cpu.y) & 0x00FF
	cpu.pc++

	//todo recheck
	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) REL() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrRel = uint16(valueRead)
	cpu.pc++

	if cpu.addrRel&0x80 != 0 {
		cpu.addrRel |= 0xFF00
	}

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) ABS() uint8 {
	lo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	hi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs = (hi << 8) | lo

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 2
		cpu.debugger.CpuOperand1 = uint8(lo & 0x00FF)
		cpu.debugger.CpuOperand2 = uint8(hi & 0x00FF)
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) ABX() uint8 {
	lo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	hi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(cpu.x)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 2
		cpu.debugger.CpuOperand1 = uint8(lo & 0x00FF)
		cpu.debugger.CpuOperand2 = uint8(hi & 0x00FF)
	}

	if (cpu.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

// checked
func (cpu *MOSTechnology6502) ABY() uint8 {
	lo := uint16(cpu.read(cpu.pc))
	cpu.pc++

	hi := uint16(cpu.read(cpu.pc))
	cpu.pc++

	cpu.addrAbs = (hi << 8) | lo
	cpu.addrAbs += uint16(cpu.y)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 2
		cpu.debugger.CpuOperand1 = uint8(lo & 0x00FF)
		cpu.debugger.CpuOperand2 = uint8(hi & 0x00FF)
	}

	if (cpu.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

// checked
func (cpu *MOSTechnology6502) IND() uint8 {
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

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 2
		cpu.debugger.CpuOperand1 = uint8(ptrLo & 0x00FF)
		cpu.debugger.CpuOperand2 = uint8(ptrHi & 0x00FF)
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) IZX() uint8 {
	t := cpu.read(cpu.pc)
	cpu.pc++

	lo := uint16(cpu.read((uint16(t) + uint16(cpu.x) + 0) & 0x00FF))
	hi := uint16(cpu.read((uint16(t) + uint16(cpu.x) + 1) & 0x00FF))

	cpu.addrAbs = (hi << 8) | lo

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = uint8(t)
	}

	return 0
}

// checked
func (cpu *MOSTechnology6502) IZY() uint8 {
	t := uint16(cpu.read(cpu.pc))
	cpu.pc++

	lo := uint16(cpu.read((t + 0) & 0x00FF))
	hi := uint16(cpu.read((t + 1) & 0x00FF))

	cpu.addrAbs = ((hi << 8) | lo) + uint16(cpu.y)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = uint8(t & 0x00FF)
	}

	if cpu.addrAbs&0xFF00 != (hi << 8) {
		return 1
	} else {
		return 0
	}
}
