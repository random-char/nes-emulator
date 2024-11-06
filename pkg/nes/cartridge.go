package nes

import "nes-emulator/pkg/cartridge"

func (nes *NES) InsertCartridge(cartridge *cartridge.Cartridge) {
	nes.cartridge = cartridge
	nes.ppu.ConnectCartridge(cartridge)
}
