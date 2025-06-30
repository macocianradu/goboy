package cpu

var OpTable = [256]func(context op_context){
	Noop,    //0x00
	LDBCd16, //0x01
	LDBCa,   //0x02
	INCBC,   //0x03
	INCB,    //0x04
	DECB,    //0x05
	LDBd8,   //0x06
	RLCA,    //0x07
	LDa16SP, //0x08
	ADDHLBC, //0x09
	LDABC,   //0x0A
	DECBC,   //0x0B
	INCC,    //0x0C
	DECC,    //0x0D
	LDCd8,   //0x0E
	RRRCA,   //0x0F
}
