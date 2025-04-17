package memory

import (
	"radu.macocian.me/goboy/errorHandling"
)

var memory = new([64000]byte)

func Read(addr int) []byte {
	checkInside(addr)
	checkInside(addr + 8)

	return memory[addr : addr+9]
}

func Write(addr int, val [8]byte) {
	checkInside(addr)
	for i := 0; i < 8; i++ {
		memory[addr+i] = val[i]
	}
}

func WriteAll(addr int, val []byte) {
	checkInside(addr)
	checkInside(addr + len(val))
	for b := 0; b < len(val); b++ {
		memory[addr+b] = val[b]
	}
}

func checkInside(addr int) {
	if addr < 0 || addr >= len(memory) {
		panic(errorHandling.OutOfMemoryBoundsError)
	}
}
