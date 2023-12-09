package main

import (
	"fmt"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day09")

	// test data
	//dataByRow = []string{
	//	"0 3 6 9 12 15",
	//	"1 3 6 10 15 21",
	//	"10 13 16 21 30 45",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
}

func part01(dataByRow []string) int {

	result := 0

	for _, row := range dataByRow {
		numbers := util.StringToIntSlice(row, " ")

		firstEndingNumber := numbers[len(numbers)-1]

		endingDifferences := make([]int, 0)

		for !allZero(numbers) {
			numbers = getDifferenceArray(numbers)
			endingDifferences = append(endingDifferences, numbers[len(numbers)-1])
		}

		finalDelta := sumArray(endingDifferences)
		result += firstEndingNumber + finalDelta

	}

	return result
}

func part02(dataByRow []string) int {

	result := 0

	for _, row := range dataByRow {
		numbers := util.StringToIntSlice(row, " ")

		firstStartingNumber := numbers[0]
		startingDifferences := make([]int, 0)

		for !allZero(numbers) {
			numbers = getDifferenceArray(numbers)
			startingDifferences = append(startingDifferences, numbers[0])
		}

		delta := 0
		// skip last zero
		for i := len(startingDifferences) - 2; i >= 0; i-- {
			delta = startingDifferences[i] - delta
		}

		result += firstStartingNumber - delta

	}

	return result
}

func sumArray(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func allZero(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}
	return true
}

func getDifferenceArray(numbers []int) []int {
	result := make([]int, len(numbers)-1)

	for i := 0; i < len(numbers)-1; i++ {
		result[i] = numbers[i+1] - numbers[i]
	}
	return result
}
