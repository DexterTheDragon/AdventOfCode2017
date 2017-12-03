package main

import "testing"

func TestInverseCaptcha(t *testing.T) {
	inputs := map[string]int{"1122": 3, "1111": 4, "1234": 0, "91212129": 9}

	for input, expected := range inputs {
		result := InverseCaptcha(input, 1)
		if result != expected {
			t.Errorf("%s: Expected %d to be %d", input, result, expected)
		}
	}
}

func TestInverseCaptchaPartTwo(t *testing.T) {
	inputs := map[string]int{
		"1212":     6,
		"1221":     0,
		"123425":   4,
		"123123":   12,
		"12131415": 4,
	}

	for input, expected := range inputs {
		step := len(input) / 2
		result := InverseCaptcha(input, step)
		if result != expected {
			t.Errorf("%s: Expected %d to be %d", input, result, expected)
		}
	}
}
