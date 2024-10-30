package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	//todo pass filename from args
	var filename string

	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Fprintln(os.Stdout, sEnc)
}
