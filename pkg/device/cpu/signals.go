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

func (olc *Olc6502) Reset() {
	olc.a = 0
	olc.x = 0
	olc.y = 0

	olc.stkp = 0xFD
	olc.status = 0x00 | FLAG_U

	olc.addrAbs = 0xFFFC

	lo := uint16(olc.Read(olc.addrAbs))
	hi := uint16(olc.Read(olc.addrAbs + 1))

	olc.pc = (hi << 8) | lo

	olc.addrRel = 0
	olc.addrAbs = 0
	olc.fetched = 0

	olc.cycles = 8
}

func (olc *Olc6502) Irq() {
	if olc.GetFlag(FLAG_I) == 1 {
		return
	}

	olc.Write(
		0x0100+uint16(olc.stkp),
		uint8((olc.pc>>8)&0x00FF),
	)
    olc.stkp--

	olc.Write(
		0x0100+uint16(olc.stkp),
		uint8(olc.pc&0x00FF),
	)
    olc.stkp--

    olc.SetFlag(FLAG_B, false)
    olc.SetFlag(FLAG_U, true)
    olc.SetFlag(FLAG_I, true)
    olc.Write(0x0100 + uint16(olc.stkp), olc.status)
    olc.stkp--

    olc.addrAbs = 0xFFFE
    lo := uint16(olc.Read(olc.addrAbs))
    hi := uint16(olc.Read(olc.addrAbs+1))

    olc.pc = (hi << 8) | lo

    olc.cycles = 7
}

func (olc *Olc6502) Nmi() {
	olc.Write(
		0x0100+uint16(olc.stkp),
		uint8((olc.pc>>8)&0x00FF),
	)
    olc.stkp--

	olc.Write(
		0x0100+uint16(olc.stkp),
		uint8(olc.pc&0x00FF),
	)
    olc.stkp--

    olc.SetFlag(FLAG_B, false)
    olc.SetFlag(FLAG_U, true)
    olc.SetFlag(FLAG_I, true)
    olc.Write(0x0100 + uint16(olc.stkp), olc.status)
    olc.stkp--

    olc.addrAbs = 0xFFFA
    lo := uint16(olc.Read(olc.addrAbs))
    hi := uint16(olc.Read(olc.addrAbs+1))

    olc.pc = (hi << 8) | lo

    olc.cycles = 8
}
