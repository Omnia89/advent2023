package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {
	sum, err := calculateSum("../input/1.test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if sum != 8 {
		t.Errorf("Expected sum to be %d, but got %d", 8, sum)
	}
}
