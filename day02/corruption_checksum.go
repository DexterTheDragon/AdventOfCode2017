package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func CorruptionChecksum(input string) int {
	result := 0
	// Split the input into lines
	lines := strings.Split(input, "\n")
	// Loop through all of the lines
	for _, line := range lines {
		var numbers = []int{} // Array that will hold our numbers
		// Loop through each number
		for _, item := range strings.Split(line, " ") {
			// Convert it from a string to an int
			num, err := strconv.Atoi(item)
			// Skip errors because yolo
			if err != nil {
				continue
			}
			// Append it to our numbers array
			numbers = append(numbers, num)
		}

		result += (maxIntSlice(numbers) - minIntSlice(numbers))
	}
	return result
}

func minIntSlice(list []int) int {
	// We should probably throw an error if the list is empty but returning
	// 0 will work for our case
	if len(list) == 0 {
		return 0
	}
	result := list[0]
	for _, item := range list {
		if item < result {
			result = item
		}
	}
	return result
}

func maxIntSlice(list []int) int {
	// We should probably throw an error if the list is empty but returning
	// 0 will work for our case
	if len(list) == 0 {
		return 0
	}
	result := list[0]
	for _, item := range list {
		if item > result {
			result = item
		}
	}
	return result
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input filename is missing.")
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", CorruptionChecksum(string(dat)))
}
