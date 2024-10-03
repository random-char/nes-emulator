package mapper

type Mapper000 struct {
	MapperInternals
}

func (m *Mapper000) CpuMapRead(addr uint16, mappedArrd uint32) bool {
    if addr >= 0x8000 && addr <= 0xFFFF {
        return true
    }

    return false
}

func (m *Mapper000) CpuMapWrite(addr uint16, mappedArrd uint32) bool {
    if addr >= 0x8000 && addr <= 0xFFFF {
        return true
    }

    return false
}

func (m *Mapper000) PpuMapRead(addr uint16, mappedArrd uint32) bool {
    if addr >= 0x0000 && addr <= 0x1FFF {
        return true
    }

    return false
}

func (m *Mapper000) PpuMapWrite(addr uint16, mappedArrd uint32) bool {
    if addr >= 0x0000 && addr <= 0x1FFF {
        return true
    }

    return false
}
