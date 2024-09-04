package main

import (
	"flag"
	"fmt"
	"time"
)

// Flags
var (
	file  string
	input string
	clk   int
	beep  bool
)

func main() {
	parseFlags()
}

func parseFlags() {
	url := "https://github.com/yashikota/td4-go#supported-file-formats"
	flag.StringVar(&file, "file", "", "File to load\n Supported file formats: "+url)
	flag.StringVar(&input, "input", "", "Default is 0000. Any binary can be specified")
	flag.IntVar(&clk, "clk", -1, "Default is 10Hz. Positive integer is accepted and Negative integer is manual clock")
	flag.BoolVar(&beep, "beep", false, "Default is false. Beep when the program ends")
	flag.Parse()
}

func clock(clk int) {
	if clk != -1 {
		interval := 1 / clk
		time.Sleep(time.Duration(interval) * time.Second)
	} else {
		// manual clock
		fmt.Scan()
	}
}
