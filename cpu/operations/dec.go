package operations

func DEC(r1 *byte) {
	*r1--
}

func DEC16(r1 *byte, r2 *byte) {
	if *r2 == 0x00 {
		*r2 = 0xFF
		*r1--
		return
	}
	*r2--
}
