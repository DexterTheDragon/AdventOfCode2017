package main

import (
	"flag"
	"fmt"
	"os"
)

func InverseCaptcha(input string, step int) int {
	input_length := len(input)
	result := 0
	for pos, char := range input {
		i := int(char - '0')

		check_pos := pos + step
		// If our check_pos is out of bounds of the array wrap it
		// around to the start of the array
		for check_pos >= input_length {
			check_pos -= input_length
		}
		check := int(input[check_pos] - '0')
		if i == check {
			result += i
		}
	}
	return result
}

func main() {
	isHalfway := flag.Bool("halfway", false, "Use halfway step for Part Two")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input is missing.")
		os.Exit(1)
	}

	input := args[0]
	step := 1
	if *isHalfway {
		step = len(input) / 2
	}

	fmt.Printf("Result: %d\n", InverseCaptcha(input, step))
}
