package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {
	sum, err := calculateSum("../input/1_2.test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if sum != 281 {
		t.Errorf("Expected sum to be %d, but got %d", 281, sum)
	}
}
