package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {
	sum, err := calculateSum("../input/2.test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if sum != 2286 {
		t.Errorf("Expected sum to be %d, but got %d", 2286, sum)
	}
}
