package nes

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/ppu"
	"time"
)

type NES struct {
	cpu       *cpu.Olc6502
	ppu       *ppu.Olc2c02
	cartridge *cartridge.Cartridge

	cpuRam []uint8

	running             bool
	nSystemClockCounter int
	stopChan            chan struct{}
}

func New(
	receiver ppu.VideoReceiver,
) *NES {
	nes := &NES{
		cpuRam: make([]uint8, 2048),

		running:             false,
		nSystemClockCounter: 0,
		stopChan:            make(chan struct{}),
	}

	nes.cpu = cpu.New(nes)
	nes.ppu = ppu.New(receiver)

	return nes
}

func (nes *NES) CpuRead(addr uint16, readOnly bool) uint8 {
	var data uint8 = 0

	_, fromCartridge := nes.cartridge.CpuRead(addr, readOnly)
	if fromCartridge {
		// allow extension via cartridge
	} else if addr >= 0 && addr <= 0x1FFF {
		data = nes.cpuRam[addr&0x07FF]
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		nes.ppu.CpuRead(addr&0x0007, readOnly)
	}

	return data
}

func (nes *NES) CpuWrite(addr uint16, data uint8) {
	if nes.cartridge.CpuWrite(addr, data) {
		// allow extension via cartridge
	} else if addr >= 0 && addr <= 0x1FFF {
		nes.cpuRam[addr&0x07FF] = data
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		nes.ppu.CpuWrite(addr&0x0007, data)
	}
}

func (nes *NES) InsertCartridge(cartridge *cartridge.Cartridge) {
	nes.cartridge = cartridge
	nes.ppu.ConnectCartridge(cartridge)
}

func (nes *NES) Reset() {
	nes.cpu.Reset()
	nes.nSystemClockCounter = 0
}

func (nes *NES) Clock() {
	nes.ppu.Clock()

	if nes.nSystemClockCounter%3 == 0 {
		nes.cpu.Clock()
	}

	nes.nSystemClockCounter++
}

func (nes *NES) Start() {
	if nes.running {
		return
	}

	nes.running = true

	ticker := time.NewTicker(46 * time.Nanosecond)
	for {
		select {
		case <-ticker.C:
			nes.Clock()
			break
		case <-nes.stopChan:
			return
		default:
			time.Sleep(5 * time.Nanosecond)
		}
	}
}

func (nes *NES) Stop() {
	if nes.running {
		nes.running = false
		nes.stopChan <- struct{}{}
	}
}
