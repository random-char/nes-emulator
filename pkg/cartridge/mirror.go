package cartridge

const (
	horizontal byte = iota
	vertical
	onescreen_lo
	onescreen_hi
)

func (c *Cartridge) IsMirrorVertical() bool {
	return c.mirror == vertical
}

func (c *Cartridge) IsMirrorHorizontal() bool {
	return c.mirror == horizontal
}
