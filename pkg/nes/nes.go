package nes

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/debugger"
	"nes-emulator/pkg/ppu"
	"time"
)

type NES struct {
	cpu       *cpu.MOSTechnology6502
	ppu       *ppu.Ricoh2c02
	cartridge *cartridge.Cartridge

	cpuRam [2048]uint8

	running             bool
	nSystemClockCounter int
	stopChan            chan struct{}
}

func New(
	receiver ppu.VideoReceiver,
	debugReceiver ppu.DebugReceiver,
	debugger *debugger.Debugger,
) *NES {
	nes := &NES{
		running:             false,
		nSystemClockCounter: 0,
		stopChan:            make(chan struct{}),
	}

	nes.cpu = cpu.New(nes, debugger)
	nes.ppu = ppu.New(receiver, debugReceiver)

	return nes
}

func (nes *NES) InsertCartridge(cartridge *cartridge.Cartridge) {
	nes.cartridge = cartridge
	nes.ppu.ConnectCartridge(cartridge)
}

func (nes *NES) Reset() {
	nes.cpu.Reset()
	nes.ppu.Reset()

	if nes.cartridge != nil {
		nes.cartridge.Reset()
	}

	nes.nSystemClockCounter = 0
}

func (nes *NES) Clock() {
	nes.ppu.Clock()

	if nes.nSystemClockCounter%3 == 0 {
		nes.cpu.Clock()
	}

	if nes.ppu.Nmi {
		nes.ppu.Nmi = false
		nes.cpu.Nmi()
	}

	nes.nSystemClockCounter++
}

func (nes *NES) Frame() {
	for !nes.ppu.FrameRendered {
		nes.Clock()
	}
	nes.ppu.FrameRendered = false
}

func (nes *NES) Start() {
	if nes.running {
		return
	}

	nes.running = true

	ticker := time.NewTicker(100 * time.Nanosecond)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				nes.Frame()
			case <-nes.stopChan:
				return
			}
		}
	}(ticker)
}

func (nes *NES) Stop() {
	if nes.running {
		nes.running = false
		nes.stopChan <- struct{}{}
	}
}
