package cpu

func (olc *Olc6502) Clock() {
	if olc.cycles == 0 {
		opcode := olc.Read(olc.pc)
		olc.pc++

		olc.cycles = lookup[opcode].Cycles
		additionalCycle1 := lookup[opcode].AddressMode()
		additionalCycle2 := lookup[opcode].Operation()

		olc.cycles += (additionalCycle1 & additionalCycle2)
	}

	olc.cycles--
}

func (olc *Olc6502) Reset() {}

func (olc *Olc6502) Irq() {}

func (olc *Olc6502) Nmi() {}
