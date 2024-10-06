package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("./test/testdata/roms/nestest.nes")
	if err != nil {
		panic(err)
	}

	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Println(sEnc)
}
