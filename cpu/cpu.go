package cpu

import (
	"radu.macocian.me/goboy/memory"
)

type cpu_struct struct {
	A  int8
	F  int8
	B  int8
	C  int8
	D  int8
	E  int8
	H  int8
	L  int8
	IR int8
	IE int8

	PC int16
	SP int16
}

var cpu = new(cpu_struct)

func Execute(addr int) {
	instr := memory.Read(addr + int(cpu.PC))
	cpu.PC++
	print(string(instr))
}
