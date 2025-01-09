package ppu

type VideoReceiver interface {
	Render([]uint8)
}
