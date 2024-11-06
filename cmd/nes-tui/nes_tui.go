package main

import (
	"flag"
	"fmt"
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/nes"
	"nes-emulator/pkg/ppu"
	"os"
)

const usage string = "usage: nes-tui -f rom_filename"

func main() {
	filename := flag.String("f", "", "Rom filename")
	flag.Parse()

	var receiver ppu.VideoReceiver = nil //todo add tui
	cpu.InitialState.SetDefaultStartingPc()

	nes := nes.New().WithVideoReceiver(receiver)

	cart, err := cartridge.LoadFromFile(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading rom: %s\n", err)
		os.Exit(1)
	}

	nes.InsertCartridge(cart)
	nes.Reset()
	nes.Start()

	exitChan := make(chan struct{})
	<-exitChan
}
