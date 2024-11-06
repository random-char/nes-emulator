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

func (cpu *MOSTechnology6502) imp() uint8 {
	cpu.fetched = cpu.a

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = cpu.fetched
	}

	return 0
}

func (cpu *MOSTechnology6502) imm() uint8 {
	cpu.addrAbs = cpu.pc
	cpu.pc++

	return 0
}

func (cpu *MOSTechnology6502) zp0() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrAbs = uint16(valueRead) & 0x00FF
	cpu.pc++

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

func (cpu *MOSTechnology6502) zpx() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrAbs = uint16(valueRead+cpu.x) & 0x00FF
	cpu.pc++

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

func (cpu *MOSTechnology6502) zpy() uint8 {
	valueRead := cpu.read(cpu.pc)
	cpu.addrAbs = uint16(valueRead+cpu.y) & 0x00FF
	cpu.pc++

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = valueRead
	}

	return 0
}

func (cpu *MOSTechnology6502) rel() uint8 {
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

func (cpu *MOSTechnology6502) abs() uint8 {
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

func (cpu *MOSTechnology6502) abx() uint8 {
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

func (cpu *MOSTechnology6502) aby() uint8 {
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

func (cpu *MOSTechnology6502) ind() uint8 {
	ptrLo := cpu.read(cpu.pc)
	cpu.pc++

	ptrHi := cpu.read(cpu.pc)
	cpu.pc++

	ptr := (uint16(ptrHi) << 8) | uint16(ptrLo)

	if ptrLo == 0x00FF {
		//hardware bug simulation
		cpu.addrAbs = (uint16(cpu.read(ptr&0xFF00)) << 8) | uint16(cpu.read(ptr))
	} else {
		cpu.addrAbs = (uint16(cpu.read(ptr+1)) << 8) | uint16(cpu.read(ptr))
	}

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 2
		cpu.debugger.CpuOperand1 = ptrLo
		cpu.debugger.CpuOperand2 = ptrHi
	}

	return 0
}

func (cpu *MOSTechnology6502) izx() uint8 {
	t := cpu.read(cpu.pc)
	cpu.pc++

	lo := uint16(cpu.read((uint16(t) + uint16(cpu.x) + 0) & 0x00FF))
	hi := uint16(cpu.read((uint16(t) + uint16(cpu.x) + 1) & 0x00FF))

	cpu.addrAbs = (hi << 8) | lo

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = t
	}

	return 0
}

func (cpu *MOSTechnology6502) izy() uint8 {
	t := cpu.read(cpu.pc)
	cpu.pc++

	lo := uint16(cpu.read((uint16(t) + 0) & 0x00FF))
	hi := uint16(cpu.read((uint16(t) + 1) & 0x00FF))

	cpu.addrAbs = ((hi << 8) | lo) + uint16(cpu.y)

	if cpu.debugger != nil {
		cpu.debugger.CpuNumOperands = 1
		cpu.debugger.CpuOperand1 = t
	}

	if cpu.addrAbs&0xFF00 != (hi << 8) {
		return 1
	} else {
		return 0
	}
}
