package ppu

import "nes-emulator/pkg/cartridge"

func (ppu *Olc2c02) ConnectCartridge(cartridge *cartridge.Cartridge) {
	ppu.cartridge = cartridge
}
