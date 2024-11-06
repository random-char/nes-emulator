package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

const usage string = "usage: encoder -f rom_filename"

func main() {
	filename := flag.String("f", "", "Rom file")
	flag.Parse()

	if *filename == "" {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	b, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file: %s\n", err)
		os.Exit(1)
	}

	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Fprintln(os.Stdout, sEnc)
}
