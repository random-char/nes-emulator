package visuals

type Sprite struct {
	width  int16
	height int16
	pixels [][]*Pixel
}

func NewSprite(width, height int16) *Sprite {
	pixels := make([][]*Pixel, width)
	for i := 0; i < int(width); i++ {
		pixels[i] = make([]*Pixel, height)
	}

	return &Sprite{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

func NewSpriteTable(
	width1,
	height1,
	width2,
	height2 int16,
) [2]*Sprite {
	return [2]*Sprite{
		NewSprite(width1, height1),
		NewSprite(width2, height2),
	}
}

func (s *Sprite) GetWidth() int16 {
	return s.width
}

func (s *Sprite) GetHeight() int16 {
	return s.height
}

func (s *Sprite) GetPixel(x, y int16) *Pixel {
	//todo add error if out of bounds
	return s.pixels[x][y]
}

func (s *Sprite) SetPixel(x, y uint16, pixel *Pixel) {
	//todo add error if out of bounds
	s.pixels[x][y] = pixel
}
