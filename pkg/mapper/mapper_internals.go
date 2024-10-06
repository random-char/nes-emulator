package mapper

type mapperInternals struct {
	nPRGBanks uint8
	nCHRBanks uint8
}

func newInternals(prgBanks, chrBanks uint8) *mapperInternals {
	return &mapperInternals{
		nPRGBanks: prgBanks,
		nCHRBanks: chrBanks,
	}
}

