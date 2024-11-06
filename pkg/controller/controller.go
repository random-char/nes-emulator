package controller

const (
	right uint8 = 1 << 0
	left  uint8 = 1 << 1
	down  uint8 = 1 << 2
	up    uint8 = 1 << 3
	start uint8 = 1 << 4
	sel   uint8 = 1 << 5
	b     uint8 = 1 << 6
	a     uint8 = 1 << 7
)

type Controller struct {
	state uint8
}

func (c *Controller) GetState() uint8 {
	return c.state
}

func (c *Controller) PressedA() {
	c.state |= a
}

func (c *Controller) PressedB() {
	c.state |= b
}

func (c *Controller) PressedUp() {
	c.state |= up
}

func (c *Controller) PressedDown() {
	c.state |= down
}

func (c *Controller) PressedLeft() {
	c.state |= left
}

func (c *Controller) PressedRight() {
	c.state |= right
}

func (c *Controller) PressedSelect() {
	c.state |= sel
}

func (c *Controller) PressedStart() {
	c.state |= start
}

func (c *Controller) ReleasedA() {
	c.state &= ^a
}

func (c *Controller) ReleasedB() {
	c.state &= ^b
}

func (c *Controller) ReleasedUp() {
	c.state &= ^up
}

func (c *Controller) ReleasedDown() {
	c.state &= ^down
}

func (c *Controller) ReleasedLeft() {
	c.state &= ^left
}

func (c *Controller) ReleasedRight() {
	c.state &= ^right
}

func (c *Controller) ReleasedSelect() {
	c.state &= ^sel
}

func (c *Controller) ReleasedStart() {
	c.state &= ^start
}
