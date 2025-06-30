package operations

func ShiftLeft(r1 *byte) {
	*r1 = *r1 << 1
}

func ShiftRight(r1 *byte) {
	*r1 = *r1 >> 1
}
