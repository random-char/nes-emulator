package main

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/debugger"
	"nes-emulator/pkg/nes"
	"nes-emulator/pkg/video"
)

func main() {
	receiver := &video.DummyVideoReceiver{}
	debugger := &debugger.Debugger{PrintData: true}
	cpu.InitialState.WithDefaultPc()

	nes := nes.New(
		receiver,
		nil,
		debugger,
	)

	//todo introduce variable
	cart, err := cartridge.LoadFromFile("../../test/testdata/roms/nestest.nes")
	if err != nil {
		panic(err)
	}

	nes.InsertCartridge(cart)
	nes.Reset()
	nes.Start()

	for {
	}
}
