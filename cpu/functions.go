package cpu

import (
	"radu.macocian.me/goboy/cpu/operations"
	"radu.macocian.me/goboy/memory"
)

type op_context struct {
	cpu  CPU_struct
	addr uint
}

func Noop(context op_context) {
	return
}

// LDBCd16 Load the 2 bytes of immediate data into register pair BC.
func LDBCd16(context op_context) {
	operations.LD(&context.cpu.C, memory.Read8(context.addr))
	operations.LD(&context.cpu.B, memory.Read8(context.addr+8))
}

// LDBCa Store the contents of register A in the memory location specified by register pair BC.
func LDBCa(context op_context) {
	operations.LDInMemory8(context.cpu.BC(), context.cpu.A)
	return
}

// INCBC Increment the contents of register pair BC by 1.
func INCBC(context op_context) {
	operations.INC16(&context.cpu.B, &context.cpu.C)
	return
}

// INCB Increment the contents of register B by 1.
func INCB(context op_context) {
	operations.INC(&context.cpu.B)
}

// DECB Decrement the contents of register B by 1.
func DECB(context op_context) {
	operations.DEC(&context.cpu.B)
}

// LD8B Load the 8-bit immediate operand d8 into register B.
func LD8B(context op_context) {
	operations.LD(&context.cpu.B, memory.Read8(context.addr))
}
