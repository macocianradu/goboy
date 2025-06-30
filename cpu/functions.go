package cpu

import (
	"radu.macocian.me/goboy/cpu/operations"
	"radu.macocian.me/goboy/memory"
)

type op_context struct {
	cpu       CPU_struct
	immediate uint16
}

// 0x00 NOOP function
func Noop(context op_context) {
	return
}

// 0x01 LDBCd16 Load the 2 bytes of immediate data into register pair BC.
func LDBCd16(context op_context) {
	context.cpu.SetBC(context.immediate)
}

// 0x02 LDBCa Store the contents of register A in the memory location specified by register pair BC.
func LDBCa(context op_context) {
	operations.LDInMemory8(context.cpu.BC(), context.cpu.A)
	return
}

// 0x03 INCBC Increment the contents of register pair BC by 1.
func INCBC(context op_context) {
	operations.INC16(&context.cpu.B, &context.cpu.C)
	return
}

// 0x04 INCB Increment the contents of register B by 1.
func INCB(context op_context) {
	if context.cpu.B&0x0F == 0x0F {
		context.cpu.SetHF(true)
	} else {
		context.cpu.SetHF(false)
	}
	operations.INC(&context.cpu.B)

	context.cpu.SetZF(context.cpu.B == byte(0))
	context.cpu.SetNF(false)
}

// 0x05 DECB Decrement the contents of register B by 1.
func DECB(context op_context) {
	if context.cpu.B&0x0F == 0x00 {
		context.cpu.SetHF(true)
	} else {
		context.cpu.SetHF(false)
	}
	operations.DEC(&context.cpu.B)

	context.cpu.SetZF(context.cpu.B == byte(0))
	context.cpu.SetNF(true)
}

// 0x06 LD8B Load the 8-bit immediate operand d8 into register B.
func LD8B(context op_context) {
	operations.LD(&context.cpu.B, memory.Read8(uint(context.immediate)))
}

// 0x07 RLCA Rotate the contents of register A to the left. The contents of bit 7 are placed in both the CY flag and bit 0 of register A.
func RLCA(context op_context) {
	a7 := context.cpu.A>>7 == 1
	operations.Shift(&context.cpu.A)

	context.cpu.SetZF(false)
	context.cpu.SetNF(false)
	context.cpu.SetHF(false)
	context.cpu.SetCF(a7)
}

// 0x08 LDA16 Store the lower byte of stack pointer SP at the address specified by the 16-bit immediate operand a16, and store the upper byte of SP at address a16 + 1.
func LDA16(context op_context) {
	operations.LDInMemory8(context.immediate, operations.GetLowerByte(context.cpu.SP))
	operations.LDInMemory8(context.immediate+1, operations.GetHigherByte(context.cpu.SP))
}

// 0x09 ADDHLBC Add the contents of register pair BC to the contents of register pair HL, and store the results in register pair HL.
func ADDHLBC(context op_context) {
	result := Add16BitsAndSetFlags(context, context.cpu.HL(), context.cpu.BC())
	context.cpu.SetHL(result)
}
