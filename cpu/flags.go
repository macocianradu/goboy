package cpu

import (
	"radu.macocian.me/goboy/cpu/operations"
)

func Add8BitsAndSetFlags(context op_context, byte1 byte, byte2 byte) byte {
	result := uint16(byte1) + uint16(byte2)
	halfcarry := (byte1&0x0F)+(byte2&0x0F) > 0x0F
	context.cpu.SetNF(false)
	context.cpu.SetHF(halfcarry)
	context.cpu.SetCF(result > 0xFFFF)
	return byte(result)
}

func Add16BitsAndSetFlags(context op_context, op1 uint16, op2 uint16) uint16 {
	result := uint32(op1) + uint32(op2)
	halfcarry := (op1&0x0FFF)+(op2&0x0FFF) > 0x0FFF
	context.cpu.SetNF(false)
	context.cpu.SetHF(halfcarry)
	context.cpu.SetCF(result > 0xFFFF)
	return uint16(result)
}

func IncAndSetFlags(context op_context, op1 *byte) {
	halfcarry := *op1&0x0F == 0x0F
	context.cpu.SetNF(false)
	context.cpu.SetHF(halfcarry)
	context.cpu.SetZF(*op1 == byte(0))
	operations.INC(op1)
}

func DecAndSetFlags(context op_context, op1 *byte) {
	halfcarry := *op1&0x0F == 0x00
	context.cpu.SetZF(*op1 == byte(0))
	context.cpu.SetNF(true)
	context.cpu.SetHF(halfcarry)
	operations.DEC(op1)
}
