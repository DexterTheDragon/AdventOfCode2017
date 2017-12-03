package main

import (
	"flag"
	"fmt"
	"os"
)

func InverseCaptcha(input string) int {
	result := 0
	prev := 0
	var first int
	for pos, char := range input {
		i := int(char - '0')
		// Save the first number so we can refer back to it
		if pos == 0 {
			first = i
		}
		// Add the number if it matches the previous
		if prev == i {
			result += i
		}
		// When we reach the end, check against the first number
		if pos == len(input)-1 && first == i {
			result += i
		}
		prev = i
	}
	return result
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input is missing.")
		os.Exit(1)
	}
	fmt.Printf("Result: %d\n", InverseCaptcha(args[0]))
}
