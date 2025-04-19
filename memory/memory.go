package memory

import (
	"radu.macocian.me/goboy/errorHandling"
)

var memory = new([64000]byte)

func Read(addr uint) []byte {
	checkInside(addr)
	checkInside(addr + 8)

	return memory[addr : addr+9]
}

func Write(addr uint, val [8]byte) {
	checkInside(addr)
	for i := 0; i < 8; i++ {
		memory[int(addr)+i] = val[i]
	}
}

func WriteAll(addr uint, val []byte) {
	checkInside(addr)
	checkInside(addr + uint(len(val)))
	for b := 0; b < len(val); b++ {
		memory[int(addr)+b] = val[b]
	}
}

func checkInside(addr uint) {
	if addr < 0 || addr >= uint(len(memory)) {
		panic(errorHandling.OutOfMemoryBoundsError)
	}
}
