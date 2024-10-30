package cpu

type bus interface {
	CpuRead(uint16) uint8
	CpuWrite(uint16, uint8)
}

