package binary

func SetBoolFlag[T uint8 | uint16](bitfield *T, flagMask T, value bool) {
	if value {
		*bitfield |= flagMask
	} else {
		*bitfield &= ^flagMask
	}
}

func GetBoolFlag[T uint8 | uint16](bitfield, flagMask T) bool {
	return bitfield&flagMask != 0
}

func SetData[T uint8 | uint16](bitfield *T, data T, dataLen, shift uint8) {
	var mask T = 0x01
	for dataLen > 1 {
		mask <<= 1
		mask |= 0x01
		dataLen--
	}

	*bitfield &= ^T(mask << shift)         //clear
	*bitfield |= (T(data) & mask) << shift //set
}

func GetData[T uint8 | uint16](bitfield *T, dataLen, shift uint8) T {
	var mask T = 0x01
	for dataLen > 1 {
		mask <<= 1
		mask |= 0x01
		dataLen--
	}

	return (*bitfield >> shift) & mask
}
