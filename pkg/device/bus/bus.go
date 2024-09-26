package bus

import "nes-emulator/pkg/device/cpu"

type Bus struct {
	cpu *cpu.Olc6502
	ram []uint8
}

func New(cpu *cpu.Olc6502) *Bus {
	return &Bus{
		cpu: cpu,
		ram: make([]uint8, 64*1024),
	}
}

func (b *Bus) Read(addr uint16, readOnly bool) uint8 {
	if addr >= 0 && addr <= 0xFFFF {
        return b.ram[addr]
    }

    return 0
}

func (b *Bus) Write(addr uint16, data uint8) {
	if addr >= 0 && addr <= 0xFFFF {
		b.ram[addr] = data
	}
}
