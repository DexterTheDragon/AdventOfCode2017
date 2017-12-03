package main

import "testing"

func TestInverseCaptcha(t *testing.T) {
	inputs := map[string]int{"1122": 3, "1111": 4, "1234": 0, "91212129": 9}

	for input, expected := range inputs {
		result := InverseCaptcha(input)
		if result != expected {
			t.Errorf("%s: Expected %d to be %d", input, result, expected)
		}
	}
}
