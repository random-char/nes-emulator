package video

import "nes-emulator/pkg/ppu"

type DummyVideoReceiver struct{}

func (*DummyVideoReceiver) SetPixel(int16, int16, *ppu.Pixel) {}

func (*DummyVideoReceiver) DrawFrame() {}
