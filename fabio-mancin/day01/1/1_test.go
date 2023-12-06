package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {
	sum, err := calculateSum("../input/1_1.test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if sum != 142 {
		t.Errorf("Expected sum to be %d, but got %d", 123, sum)
	}
}
