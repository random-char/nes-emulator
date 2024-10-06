package cpu

type bus interface {
	CpuRead(uint16, bool) uint8
	CpuWrite(uint16, uint8)
}

