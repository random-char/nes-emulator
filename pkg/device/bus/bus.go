package bus

import (
	"nes-emulator/pkg/device/cartridge"
	"nes-emulator/pkg/device/cpu"
	"nes-emulator/pkg/device/ppu"
)

type Bus struct {
	cpu       *cpu.Olc6502
	ppu       *ppu.Olc2c02
	cartridge *cartridge.Cartridge

	cpuRam []uint8

	nSystemClockCounter int
}

func New(
	cpu *cpu.Olc6502,
	ppu *ppu.Olc2c02,
) *Bus {
	return &Bus{
		cpu:    cpu,
		ppu:    ppu,
		cpuRam: make([]uint8, 2048),

		nSystemClockCounter: 0,
	}
}

func (b *Bus) CpuRead(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0

	if b.cartridge.CpuRead(addr, readOnly) {
		// allow extension via cartridge
	} else if addr >= 0 && addr <= 0x1FFF {
		data = b.cpuRam[addr&0x07FF]
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		b.ppu.CpuRead(addr&0x0007, readOnly)
	}

	return data
}

func (b *Bus) CpuWrite(addr uint16, data uint8) {
	if b.cartridge.CpuWrite(addr, data) {
		// allow extension via cartridge
	} else if addr >= 0 && addr <= 0x1FFF {
		b.cpuRam[addr&0x07FF] = data
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		b.ppu.CpuWrite(addr&0x0007, data)
	}
}

func (b *Bus) InsertCartridge(cartridge *cartridge.Cartridge) {
	b.cartridge = cartridge
	b.ppu.ConnectCartridge(cartridge)
}

func (b *Bus) Reset() {
	b.cpu.Reset()
	b.nSystemClockCounter = 0
}

func (b *Bus) Clock() {}
