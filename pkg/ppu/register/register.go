package register

func setRegData[T uint8 | uint16](reg *T, data T, dataLen, shift uint8) {
	var mask T = 0x01
	for dataLen > 1 {
		mask <<= 1
		mask |= 0x01
		dataLen--
	}

	*reg &= ^T(mask << shift)         //clear
	*reg |= (T(data) & mask) << shift //set
}

func getRegData[T uint8|uint16](reg *T, dataLen, shift uint8) T {
	var mask T = 0x01
	for dataLen > 1 {
		mask <<= 1
		mask |= 0x01
		dataLen--
	}

    return (*reg >> shift) & mask
}
