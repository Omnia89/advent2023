package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {
	sum, err := calculateSum("../input/1.test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if sum != 4361 {
		t.Errorf("Expected sum to be %d, but got %d", 4361, sum)
	}
}
