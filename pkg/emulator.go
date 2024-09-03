package main

import (
	"flag"
	"fmt"
)

// Flags
var (
	verbose bool
	file   string
	input  string
	clk   int
	beep  bool
)

func main() {
	// pc := 0     // Program Counter
	// reg_a := 0  // Register A
	// reg_b := 0  // Register B
	// c_flag := 0 // Carry Flag
	// output := 0 // Output

	// Parse flags
	flags()

	fmt.Println("verbose:", verbose)
	fmt.Println("file:", file)
	fmt.Println("input:", input)
	fmt.Println("clk:", clk)
	fmt.Println("beep:", beep)
}

func flags() {
	flag.BoolVar(&verbose, "verbose", false, "Prints the output of each instruction")
	flag.StringVar(&file, "file", "", "File to load the program from")
	flag.StringVar(&input, "input", "", "Input to the program")
	flag.IntVar(&clk, "clk", -1, "Clock speed")
	flag.BoolVar(&beep, "beep", false, "Beep on output")
	flag.Parse()
}
