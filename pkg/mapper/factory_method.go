package mapper

func CreateMapper(nMapperID, prgBanks, chrBanks uint8) (Mapper, error) {
	switch nMapperID {
	case 0:
		return &mapper000{
			mapperInternals: newInternals(prgBanks, chrBanks),
		}, nil
	default:
		return nil, &UnsupportedMapperErr{mapperId: nMapperID}
	}
}
