package devices

type Bus interface {
	Read(uint16, bool) uint8
	Write(uint16, uint8)
}

