package ppu

type bus interface {
	CpuRead(uint16, bool) uint8
	CpuWrite(uint16, uint8)
	PpuRead(uint16, bool) uint8
	PpuWrite(uint16, uint8)
}

