package main

import "fmt"

// CPU registers
var (
	pc     int   = 0      // Program Counter
	regA   uint8 = 0b0000 // Register A
	regB   uint8 = 0b0000 // Register B
	cFlag  int   = 0      // Carry Flag
	output uint8 = 0b0000 // Output
)

func fetch(rom []string) string {
	r := rom[pc]
	pc++
	return r
}

func decode(inst string) (string, string) {
	return inst[:4], inst[4:]
}

func execute(op string, im uint8, input uint8) {
	cFlag = 0
	switch op {
	case fmt.Sprint(MOV_A): // Move A, im
		regA = im
	case fmt.Sprint(MOV_B): // Move B, im
		regB = im
	case fmt.Sprint(MOV_AB): // Move A, B
		regA = regB
	case fmt.Sprint(MOV_BA): // Move B, A
		regB = regA
	case fmt.Sprint(ADD_A): // Add A, im
		regA, cFlag = add(regA, im)
	case fmt.Sprint(ADD_B): // Add B, im
		regB, cFlag = add(regB, im)
	case fmt.Sprint(IN_A): // Input A
		regA = input
	case fmt.Sprint(IN_B): // Input B
		regB = input
	case fmt.Sprint(OUT): // Output im
		output = im
	case fmt.Sprint(OUT_B): // Output B
		output = regB
	case fmt.Sprint(JMP): // Jump im
		pc = int(im)
	case fmt.Sprint(JNC): // Jump if no carry
		if cFlag == 0 {
			pc = int(im)
		}
	}
}

func add(a uint8, b uint8) (uint8, int) {
	sum := a + b
	carry := 0
	if sum > 0b1111 {
		carry = 1
		sum = sum & 0b1111
	}
	return uint8(sum), carry
}
