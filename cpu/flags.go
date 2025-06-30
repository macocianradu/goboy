package cpu

func Set8BitAddFlags(context op_context, byte1 byte, byte2 byte) byte {
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
