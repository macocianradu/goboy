package operations

import "radu.macocian.me/goboy/memory"

func LD(r1 *byte, val byte) {
	*r1 = val
}

func LDFromMem(r1 *byte, r2 uint) {
	*r1 = memory.Read8(r2)
}

func LDInMemory8(addr uint16, val byte) {
	memory.Write8(uint(addr), val)
}
