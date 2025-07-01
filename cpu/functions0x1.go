package cpu

import (
	"radu.macocian.me/goboy/cpu/operations"
)

/*
	0x10 STOP Execution of a STOP instruction stops both the system clock and oscillator circuit. STOP mode is entered and the LCD controller also stops. However, the status of the internal RAM register ports remains unchanged.

STOP mode can be cancelled by a reset signal.
If the RESET terminal goes LOW in STOP mode, it becomes that of a normal reset status.
The following conditions should be met before a STOP instruction is executed and stop mode is entered:

	All interrupt-enable (IE) flags are reset.
	Input to P10-P13 is LOW for all.
*/
func STOP(context op_context) {
	//TODO
	return
}

// 0x11 LDDEd16 Load the 2 bytes of immediate data into register pair DE.
func LDDEd16(context op_context) {
	context.cpu.SetDE(context.immediate)
}

// 0x12 LDDEA Store the contents of register A in the memory speicifed by the register pair DE.
func LDDEA(context op_context) {
	operations.LDInMemory8(context.cpu.DE(), context.cpu.A)
	return
}

// 0x13 INCDE Increment the contents of register pair DE by 1.
func INCDE(context op_context) {
	operations.INC16(&context.cpu.D, &context.cpu.E)
}

// 0x14 INCD Increment the contents of register D by 1.
func INCD(context op_context) {
	IncAndSetFlags(context, &context.cpu.D)
}

// 0x15 DECD Decrement the contents of register D by 1.
func DECD(context op_context) {
	DecAndSetFlags(context, &context.cpu.D)
}

// 0x16 LDDd8 Load the 8-bit immediate operand d8 into register D.
func LDDd8(context op_context) {
	operations.LD(&context.cpu.D, byte(context.immediate))
}

// 0x17 RLA Rotate the contents of register A to the left, through the carry (CY) flag.
func RLA(context op_context) {
	a7 := context.cpu.A >> 7
	operations.ShiftLeft(&context.cpu.A)
	carry := 0
	if context.cpu.CF() {
		carry = 1
	}
	context.cpu.A = context.cpu.A | byte(carry)

	context.cpu.SetZF(false)
	context.cpu.SetNF(false)
	context.cpu.SetHF(false)
	context.cpu.SetCF(a7 == 1)
}

// 0x18 JRs8 Jump s8 steps from the current address in the program counter (PC). (Jump relative.)
func JRs8(context op_context) {
	signedValue := int8(context.immediate & 0x00FF)
	context.cpu.PC += uint16(int16(signedValue))
}

// 0x19 ADDHLDE Add the contents of register pair DE to the contents of register pair HL, and store the results in register pair HL.
func ADDHLDE(context op_context) {
	result := Add16BitsAndSetFlags(context, context.cpu.DE(), context.cpu.HL())
	context.cpu.SetHL(result)
}

// 0x1A LDADE Load the 8-bit contents of memory specified by register pair DE into register A.
func LDADE(context op_context) {
	operations.LDFromMem(&context.cpu.A, uint(context.cpu.DE()))
}

// 0x1B DECDE Decrement the contents of register pair DE by 1.
func DECDE(context op_context) {
	operations.DEC16(&context.cpu.D, &context.cpu.E)
}

// 0x1C INCE Increment the contents of register E by 1.
func INCE(context op_context) {
	IncAndSetFlags(context, &context.cpu.E)
}

// 0x1D DECE Decrement the contents of register E by 1.
func DECE(context op_context) {
	DecAndSetFlags(context, &context.cpu.E)
}

// 0x1E LDEd8 Load the 8-bit immediate operand d8 into register E.
func LDEd8(context op_context) {
	context.cpu.E = byte(context.immediate)
}

// 0x1F RRA Rotate the contents of register A to the right, through the carry (CY) flag.
func RRA(context op_context) {
	a0 := context.cpu.A & 0x01
	operations.ShiftRight(&context.cpu.A)
	carry := 0
	if context.cpu.CF() {
		carry = 1
	}
	context.cpu.A = context.cpu.A | byte(carry)<<7

	context.cpu.SetZF(false)
	context.cpu.SetNF(false)
	context.cpu.SetHF(false)
	context.cpu.SetCF(a0 == 1)
}
