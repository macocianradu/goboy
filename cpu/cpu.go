package cpu

import (
	"radu.macocian.me/goboy/memory"
)

type CPU_struct struct {
	A  byte
	F  byte
	B  byte
	C  byte
	D  byte
	E  byte
	H  byte
	L  byte
	IR byte
	IE byte

	PC uint16
	SP uint16
}

var cpu = new(CPU_struct)

func Execute(addr uint) {
	instr := memory.Read(addr + uint(cpu.PC))
	cpu.PC++
	print(string(instr))
}

func readDoubleRegis(r1 byte, r2 byte) uint16 {
	return uint16(r1)<<8 | uint16(r2)
}

func setDoubleRegis(r1 *byte, r2 *byte, value uint16) {
	*r1 = byte((value & 0xFF00) >> 8)
	*r2 = byte(value & 0x00FF)
}

func (cpu *CPU_struct) AF() uint16 {
	return readDoubleRegis(cpu.A, cpu.F)
}

func (cpu *CPU_struct) SetAF(value uint16) {
	setDoubleRegis(&cpu.A, &cpu.F, value)
}

func (cpu *CPU_struct) BC() uint16 {
	return readDoubleRegis(cpu.B, cpu.C)
}

func (cpu *CPU_struct) SetBC(value uint16) {
	setDoubleRegis(&cpu.B, &cpu.C, value)
}

func (cpu *CPU_struct) DE() uint16 {
	return readDoubleRegis(cpu.D, cpu.E)
}

func (cpu *CPU_struct) SetDE(value uint16) {
	setDoubleRegis(&cpu.D, &cpu.E, value)
}

func (cpu *CPU_struct) HL() uint16 {
	return readDoubleRegis(cpu.H, cpu.L)
}

func (cpu *CPU_struct) SetHL(value uint16) {
	setDoubleRegis(&cpu.H, &cpu.L, value)
}
