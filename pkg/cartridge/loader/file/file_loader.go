package file

import (
	"nes-emulator/pkg/cartridge"
	"os"
)

func LoadFromFile(filename string) (*cartridge.Cartridge, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return cartridge.Load(f)
}
