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
	IncAndSetFlags(context, &context.cpu.B)
}

// 0x05 DECB Decrement the contents of register B by 1.
func DECB(context op_context) {
	DecAndSetFlags(context, &context.cpu.B)
}

// 0x06 LDBd8 Load the 8-bit immediate operand d8 into register B.
func LDBd8(context op_context) {
	operations.LD(&context.cpu.B, byte(context.immediate))
}

// 0x07 RLCA Rotate the contents of register A to the left. The contents of bit 7 are placed in both the CY flag and bit 0 of register A.
func RLCA(context op_context) {
	a7 := context.cpu.A >> 7
	operations.ShiftLeft(&context.cpu.A)
	context.cpu.A = context.cpu.A | byte(a7)

	context.cpu.SetZF(false)
	context.cpu.SetNF(false)
	context.cpu.SetHF(false)
	context.cpu.SetCF(a7 == 1)
}

// 0x08 LDa16SP Store the lower byte of stack pointer SP at the address specified by the 16-bit immediate operand a16, and store the upper byte of SP at address a16 + 1.
func LDa16SP(context op_context) {
	operations.LDInMemory8(context.immediate, operations.GetLowerByte(context.cpu.SP))
	operations.LDInMemory8(context.immediate+1, operations.GetHigherByte(context.cpu.SP))
}

// 0x09 ADDHLBC Add the contents of register pair BC to the contents of register pair HL, and store the results in register pair HL.
func ADDHLBC(context op_context) {
	result := Add16BitsAndSetFlags(context, context.cpu.HL(), context.cpu.BC())
	context.cpu.SetHL(result)
}

// 0x0A LDABC Load the 8-bit contents of memory specified by register pair BC into register A.
func LDABC(context op_context) {
	operations.LDFromMem(&context.cpu.A, uint(context.cpu.BC()))
}

// 0x0B DECBC Decrement the contents of register pair BC by 1.
func DECBC(context op_context) {
	operations.DEC16(&context.cpu.B, &context.cpu.C)
}

// 0x0C INCC Increment the contents of register C by 1.
func INCC(context op_context) {
	IncAndSetFlags(context, &context.cpu.C)
}

// 0x0D DECC Decrement the contents of register C by 1.
func DECC(context op_context) {
	DecAndSetFlags(context, &context.cpu.C)
}

// 0x0E LDCd8 Load the 8-bit immediate operand d8 into register C.
func LDCd8(context op_context) {
	context.cpu.C = byte(context.immediate)
}

// 0x0F RRRCA Rotate the contents of register A to the right. That is, the contents of bit 7 are copied to bit 6, and the previous contents of bit 6 (before the copy) are copied to bit 5. The same operation is repeated in sequence for the rest of the register. The contents of bit 0 are placed in both the CY flag and bit 7 of register A.
func RRRCA(context op_context) {
	a0 := context.cpu.A & 0x01
	operations.ShiftRight(&context.cpu.A)
	context.cpu.A = context.cpu.A | byte(a0)<<7

	context.cpu.SetZF(false)
	context.cpu.SetNF(false)
	context.cpu.SetHF(false)
	context.cpu.SetCF(a0 == 1)

}
