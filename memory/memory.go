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

func Read8(addr uint) byte {
	checkInside(addr)

	return memory[addr]
}

func Read16(addr uint) uint16 {
	checkInside(addr)
	checkInside(addr + 8)

	return uint16(memory[addr])<<8 | uint16(memory[addr+1])
}

func Write(addr uint, val [8]byte) {
	checkInside(addr)
	for i := 0; i < 8; i++ {
		memory[int(addr)+i] = val[i]
	}
}

func Write8(addr uint, val byte) {
	checkInside(addr)
	memory[int(addr)] = val
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
