package cpu

var additionalCycle1, additionalCycle2 uint8
var currentOp *instruction

func (cpu *MOSTechnology6502) Clock() {
	if cpu.cycles == 0 {
		cpu.opcode = cpu.read(cpu.pc)
		currentOp = lookup[cpu.opcode]

		if cpu.debugger != nil {
			cpu.debugger.CpuPc = cpu.pc
			cpu.debugger.CpuCycles = cpu.clockCount - 1 //investigate
			cpu.debugger.CpuOpCode = cpu.opcode
			cpu.debugger.CpuOpName = currentOp.OpName
			cpu.debugger.CpuRegA = cpu.a
			cpu.debugger.CpuRegX = cpu.x
			cpu.debugger.CpuRegY = cpu.y
			cpu.debugger.CpuSP = cpu.stkp
			cpu.debugger.CpuStatus = cpu.status
		}

		cpu.setFlag(flag_U, true)
		cpu.pc++

		cpu.cycles = currentOp.Cycles
		additionalCycle1 = currentOp.AddrMode()
		additionalCycle2 = currentOp.Op()

		cpu.cycles += (additionalCycle1 & additionalCycle2)

		cpu.setFlag(flag_U, true)

		if cpu.debugger != nil {
			if cpu.debugger.PrintData {
				cpu.debugger.Print()
				cpu.debugger.Reset()
			}
		}
	}

	cpu.cycles--
	cpu.clockCount++
}
