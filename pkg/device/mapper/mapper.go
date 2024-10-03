package mapper

type Mapper struct {
	nPRGBanks uint8
	nCHRBanks uint8
}

func New(prgBanks, chrBanks uint8) *Mapper {
	return &Mapper{
		nPRGBanks: prgBanks,
		nCHRBanks: chrBanks,
	}
}
