package mapper

import "fmt"

func CreateMapper(nMapperID, prgBanks, chrBanks uint8) Mapper {
	switch nMapperID {
	case 0:
		return &mapper000{
			mapperInternals: newInternals(prgBanks, chrBanks),
		}
	default:
        panic(fmt.Sprintf("Unsupported mapper: %d", nMapperID))
	}
}
