//go:build wasm
// +build wasm

package main

import (
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/nes"
	"nes-emulator/pkg/video"
	"syscall/js"
)

var nesEmulator *nes.NES

func main() {
	doc := js.Global().Get("document")

	canvas := doc.Call("getElementById", "canvas")
	debugCanvas := doc.Call("getElementById", "debug-canvas")

	allInOneCanvasReceiver := video.NewCanvasVideoReceiver(canvas, debugCanvas)
	nesEmulator = nes.New().
		WithVideoReceiver(allInOneCanvasReceiver).
		WithDebugReceiver(allInOneCanvasReceiver)

	exitChan := make(chan struct{})
	<-exitChan
}

//export start
func start() {
	js.FuncOf(func(this js.Value, args []js.Value) any {
		nesEmulator.Start()
		return nil
	}).Invoke()
}

//export frame
func frame() {
	nesEmulator.Frame()
}

//export stop
func stop() {
	nesEmulator.Stop()
}

//export reset
func reset() {
	nesEmulator.Reset()
}

//export keyDown
func keyDown(key byte) {
	switch key {
	case 'w':
		nesEmulator.Controller[0].PressedUp()
	case 's':
		nesEmulator.Controller[0].PressedDown()
	case 'a':
		nesEmulator.Controller[0].PressedLeft()
	case 'd':
		nesEmulator.Controller[0].PressedRight()
	case 'j':
		nesEmulator.Controller[0].PressedB()
	case 'k':
		nesEmulator.Controller[0].PressedA()
	case 'c':
		nesEmulator.Controller[0].PressedSelect()
	case 'v':
		nesEmulator.Controller[0].PressedStart()
	}
}

//export keyUp
func keyUp(key byte) {
	switch key {
	case 'w':
		nesEmulator.Controller[0].ReleasedUp()
	case 's':
		nesEmulator.Controller[0].ReleasedDown()
	case 'a':
		nesEmulator.Controller[0].ReleasedLeft()
	case 'd':
		nesEmulator.Controller[0].ReleasedRight()
	case 'j':
		nesEmulator.Controller[0].ReleasedB()
	case 'k':
		nesEmulator.Controller[0].ReleasedA()
	case 'c':
		nesEmulator.Controller[0].ReleasedSelect()
	case 'v':
		nesEmulator.Controller[0].ReleasedStart()
	}
}

//export insertCartridge
func insertCartridgeB64(b64Cartridge string) {
	cart, err := cartridge.LoadFromBase64(b64Cartridge)
	if err != nil {
		println("Cartridge wasn't loaded")
		return
	}

	nesEmulator.InsertCartridge(cart)
	nesEmulator.Reset()
}
