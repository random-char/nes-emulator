package binary

// https://stackoverflow.com/a/2602885
func ReverseUint8(b *uint8) {
	*b = (*b&0xF0)>>4 | (*b&0x0F)<<4
	*b = (*b&0xCC)>>2 | (*b&0x33)<<2
	*b = (*b&0xAA)>>1 | (*b&0x55)<<1
}
