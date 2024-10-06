package b64string

import (
	"encoding/base64"
	"nes-emulator/pkg/cartridge"
	"strings"
)

func LoadFromBase64(input string) (*cartridge.Cartridge, error) {
	b, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(string(b))
	return cartridge.Load(r)
}
