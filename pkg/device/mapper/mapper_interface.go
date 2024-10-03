package mapper

type Mapper interface {
	CpuMapRead(addr uint16, mappedArrd uint32) bool
	CpuMapWrite(addr uint16, mappedArrd uint32) bool
	PpuMapRead(addr uint16, mappedArrd uint32) bool
	PpuMapWrite(addr uint16, mappedArrd uint32) bool
}
