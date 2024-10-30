package cartridge

import (
	"os"
)

func LoadFromFile(filename string) (*Cartridge, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return load(f)
}
