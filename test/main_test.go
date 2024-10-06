package test

import (
	"nes-emulator/pkg/cartridge/loader/file"
	"nes-emulator/pkg/nes"
	"nes-emulator/pkg/video"
	"testing"
)

func TestFromNestestRom(t *testing.T) {
    receiver := video.DummyVideoReceiver{}
	nes := nes.New(&receiver)

	cartridge, err := file.LoadFromFile("./testdata/roms/nestest.nes")
	if err != nil {
		t.Error(err)
	}

	nes.InsertCartridge(cartridge)
    nes.Reset()
}
