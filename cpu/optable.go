package cpu

type operation func(byte)

var OpTable = [256]func(context op_context){
	Noop,    //0x00
	LDBCd16, //0x01
	LDBCa,   //0x02
	INCBC,   //0x03
	INCB,    //0x04
	DECB,    //0x05
	LD8B,    //0x06
}

func (cpu *CPU_struct) execute(op byte, addr uint) {
}
