package operations

func GetLowerByte(value uint16) byte {
	return (byte)(value)
}

func GetHigherByte(value uint16) byte {
	return (byte)(value >> 8)
}

