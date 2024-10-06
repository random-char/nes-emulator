package ppu

type VideoReceiver interface {
	SetPixel(x, y int16, pixel *Pixel)
	DrawFrame()
}
