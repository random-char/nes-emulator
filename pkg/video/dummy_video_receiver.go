package video

import "nes-emulator/pkg/ppu/visuals"

type DummyVideoReceiver struct{}

func (*DummyVideoReceiver) RenderFrame([]*visuals.Pixel) {}
