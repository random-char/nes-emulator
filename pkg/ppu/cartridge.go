package ppu

import "nes-emulator/pkg/cartridge"

func (ppu *Ricoh2c02) ConnectCartridge(cartridge *cartridge.Cartridge) {
	ppu.cartridge = cartridge
}
