package main

import "testing"

func TestCorruptionChecksum(t *testing.T) {
	input := `5 1 9 5
7 5 3
2 4 6 8`
	expected := 18

	result := CorruptionChecksum(input)
	if result != expected {
		t.Errorf("Expected %d to be %d", result, expected)
	}
}
