package cartridge

import (
	"encoding/base64"
	"strings"
)

func LoadFromBase64(input string) (*Cartridge, error) {
	b, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(string(b))
	return load(r)
}
