package operations

func INC(r1 *byte) {
	*r1++
}

func INC16(r1 *byte, r2 *byte) {
	if *r2 == 0xFF {
		*r2 = 0
		*r1++
		return
	}
	*r2++
}
