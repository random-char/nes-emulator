//go:build wasm
// +build wasm

package main

import (
	"fmt"
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/nes"
	"nes-emulator/pkg/video"
	"syscall/js"
)

func main() {
	doc := js.Global().Get("document")

	canvas := doc.Call("getElementById", "canvas")
	debugCanvas := doc.Call("getElementById", "debug-canvas")

	allInOneCanvasReceiver := video.NewCanvasVideoReceiver(canvas, debugCanvas)
	nes := nes.New().
		WithVideoReceiver(allInOneCanvasReceiver).
		WithDebugReceiver(allInOneCanvasReceiver)

	js.Global().Set("NesInsertCartridgeB64", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			fmt.Println("Cartridge argument not provided")
			return nil
		}

		cart, err := cartridge.LoadFromBase64(args[0].String())
		if err != nil {
			fmt.Println("Cartridge wasn't loaded")
			return nil
		}

		nes.InsertCartridge(cart)
		nes.Reset()

		return nil
	}))

	js.Global().Set("NesFrame", js.FuncOf(func(this js.Value, args []js.Value) any {
		nes.Frame()
		return nil
	}))

	js.Global().Set("NesStart", js.FuncOf(func(this js.Value, args []js.Value) any {
		nes.Start()
		return nil
	}))

	js.Global().Set("NesStop", js.FuncOf(func(this js.Value, args []js.Value) any {
		nes.Stop()
		return nil
	}))

	js.Global().Set("NesReset", js.FuncOf(func(this js.Value, args []js.Value) any {
		nes.Reset()
		return nil
	}))

	js.Global().Call(
		"addEventListener",
		"keyup",
		js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) < 1 {
				return nil
			}

			key := args[0].Get("key").String()
			switch key {
			case "w":
				nes.Controller[0].ReleasedUp()
			case "s":
				nes.Controller[0].ReleasedDown()
			case "a":
				nes.Controller[0].ReleasedLeft()
			case "d":
				nes.Controller[0].ReleasedRight()
			case "j":
				nes.Controller[0].ReleasedB()
			case "k":
				nes.Controller[0].ReleasedA()
			case "c":
				nes.Controller[0].ReleasedSelect()
			case "v":
				nes.Controller[0].ReleasedStart()
			}
			return nil
		}),
	)

	js.Global().Call(
		"addEventListener",
		"keydown",
		js.FuncOf(func(this js.Value, args []js.Value) any {
			if len(args) < 1 {
				return nil
			}

			key := args[0].Get("key").String()
			switch key {
			case "w":
				nes.Controller[0].PressedUp()
			case "s":
				nes.Controller[0].PressedDown()
			case "a":
				nes.Controller[0].PressedLeft()
			case "d":
				nes.Controller[0].PressedRight()
			case "j":
				nes.Controller[0].PressedB()
			case "k":
				nes.Controller[0].PressedA()
			case "c":
				nes.Controller[0].PressedSelect()
			case "v":
				nes.Controller[0].PressedStart()
			}
			return nil
		}),
	)

	exitChan := make(chan struct{})
	<-exitChan
}
