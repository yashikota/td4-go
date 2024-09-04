package main

// Opcode
const (
	MOV_A  = 0b0011
	MOV_B  = 0b0111
	MOV_AB = 0b0001
	MOV_BA = 0b0100
	ADD_A  = 0b0000
	ADD_B  = 0b0101
	IN_A   = 0b0010
	IN_B   = 0b0110
	OUT    = 0b1011
	OUT_B  = 0b1001
	JMP    = 0b1111
	JNC    = 0b1110
)

// mnemonics
var mnemonics = map[uint8]string{
	MOV_A:  "MOV A",
	MOV_B:  "MOV B",
	MOV_AB: "MOV A, B",
	MOV_BA: "MOV B, A",
	ADD_A:  "ADD A",
	ADD_B:  "ADD B",
	IN_A:   "IN A",
	IN_B:   "IN B",
	OUT:    "OUT",
	OUT_B:  "OUT B",
	JMP:    "JMP",
	JNC:    "JNC",
}
