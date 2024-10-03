package mapper

type MapperInternals struct {
	nPRGBanks uint8
	nCHRBanks uint8
}

func New(prgBanks, chrBanks uint8) *MapperInternals {
	return &MapperInternals{
		nPRGBanks: prgBanks,
		nCHRBanks: chrBanks,
	}
}

