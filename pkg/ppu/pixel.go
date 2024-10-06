package ppu

type Pixel struct {
	R uint8
	G uint8
	B uint8
}

func newPixel(r, g, b uint8) *Pixel {
	return &Pixel{
		R: r,
		G: g,
		B: b,
	}
}
