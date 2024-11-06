package nes

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/controller"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/debugger"
	"nes-emulator/pkg/ppu"
	"time"
)

type NES struct {
	cpu       *cpu.MOSTechnology6502
	ppu       *ppu.Ricoh2c02
	cartridge *cartridge.Cartridge

	Controller      [2]*controller.Controller
	controllerState [2]uint8

	cpuRam [2048]uint8

	dmaPage     uint8
	dmaAddr     uint8
	dmaData     uint8
	dmaTransfer bool
	dmaDummy    bool

	running             bool
	nSystemClockCounter int
	stopChan            chan struct{}
}

func New() *NES {
	nes := &NES{
		dmaPage:     0x00,
		dmaAddr:     0x00,
		dmaData:     0x00,
		dmaTransfer: false,
		dmaDummy:    true,

		running:             false,
		nSystemClockCounter: 0,
		stopChan:            make(chan struct{}),
	}

	nes.Controller[0] = &controller.Controller{}
	nes.Controller[1] = &controller.Controller{}

	nes.cpu = cpu.New(nes)
	nes.ppu = ppu.New()

	return nes
}

func (nes *NES) WithVideoReceiver(
	receiver ppu.VideoReceiver,
) *NES {
	nes.ppu.SetVideoReceiver(receiver)

	return nes
}

func (nes *NES) WithDebugReceiver(
	debugReceiver ppu.DebugReceiver,
) *NES {
	nes.ppu.SetDebugReceiver(debugReceiver)

	return nes
}

func (nes *NES) WithDebugger(
	debugger *debugger.Debugger,
) *NES {
	nes.cpu.SetDebugger(debugger)
	nes.ppu.SetDebugger(debugger)

	return nes
}

func (nes *NES) Frame() {
	for !nes.ppu.FrameRendered {
		nes.Clock()
	}
	nes.ppu.FrameRendered = false
}

func (nes *NES) Start() error {
	if nes.running {
		return AlreadyRunningErr
	}

	if nes.cartridge == nil {
		return NoCartridgeErr
	}

	nes.running = true

	ticker := time.NewTicker(75 * time.Millisecond)
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

	return nil
}

func (nes *NES) Stop() {
	if nes.running {
		nes.running = false
		nes.stopChan <- struct{}{}
	}
}
