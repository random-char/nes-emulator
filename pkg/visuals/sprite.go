package visuals

type Sprite struct {
	width  uint16
	height uint16
	pixels [][]*Pixel
}

func NewSprite(width, height uint16) *Sprite {
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
	height2 uint16,
) [2]*Sprite {
	return [2]*Sprite{
		NewSprite(width1, height1),
		NewSprite(width2, height2),
	}
}

func (s *Sprite) GetWidth() uint16 {
	return s.width
}

func (s *Sprite) GetHeight() uint16 {
	return s.height
}

func (s *Sprite) GetPixel(x, y uint16) (*Pixel, error) {
	if x >= s.width || y >= s.height {
		return nil, OutOfSpriteBoundsErr
	}

	return s.pixels[x][y], nil
}

func (s *Sprite) SetPixel(x, y uint16, pixel *Pixel) error {
	if x < 0 || y < 0 || x >= s.width || y >= s.height {
		return OutOfSpriteBoundsErr
	}

	s.pixels[x][y] = pixel
	return nil
}

func (s *Sprite) GetPixelsData() []uint8 {
	data := make(
		[]uint8,
		int(s.width)*int(s.height)*4,
		int(s.width)*int(s.height)*4,
	)

	var x, y, i uint16 = 0, 0, 0
	for x = 0; x < s.width; x++ {
		for y = 0; y < s.height; y++ {
			data[i] = s.pixels[x][y].R
			data[i+1] = s.pixels[x][y].G
			data[i+2] = s.pixels[x][y].B
			data[i+3] = 255
			i += 4
		}
	}

	return data
}
