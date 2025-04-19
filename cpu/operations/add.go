package operations

import "radu.macocian.me/goboy/memory"

func ADD(r1 *byte, r2 *byte) {
	*r1 += *r2
}

func ADDFromMem(r1 *byte, r2 uint) {
	*r1 += memory.Read8(r2)
}
