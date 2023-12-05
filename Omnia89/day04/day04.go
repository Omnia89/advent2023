package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day04")

	// test data
	//dataByRow = []string{
	//	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	//	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	//	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	//	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	//	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	//	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
	// 5554894

}

func part01(dataByRow []string) int {

	sum := 0

	for _, row := range dataByRow {

		parts := strings.Split(row, ":")
		numbersString := strings.Split(parts[1], "|")
		winningNumbers := util.StringToIntSlice(numbersString[0], " ")
		playedNumbers := util.StringToIntSlice(numbersString[1], " ")

		var matches []int
		for _, n := range winningNumbers {
			if util.IntContains(playedNumbers, n) {
				matches = append(matches, n)
			}
		}

		if len(matches) > 0 {
			points := 1
			for i := 0; i < len(matches)-1; i++ {
				points *= 2
			}
			sum += points
		}
	}
	return sum
}

func part02(dataByRow []string) int {

	winningScratchCards := map[int]int{}
	addScore := func(n int) {
		if _, ok := winningScratchCards[n]; !ok {
			winningScratchCards[n] = 1
		} else {
			winningScratchCards[n]++
		}
	}
	removeScore := func(n int) {
		if _, ok := winningScratchCards[n]; ok {
			winningScratchCards[n]--
		}
	}

	numberScratched := 0

	for n := 0; n < len(dataByRow); {

		row := dataByRow[n]
		numberScratched++

		gameNumber := n + 1
		parts := strings.Split(row, ":")
		numbersString := strings.Split(parts[1], "|")
		winningNumbers := util.StringToIntSlice(numbersString[0], " ")
		playedNumbers := util.StringToIntSlice(numbersString[1], " ")

		var matches int
		for _, num := range winningNumbers {
			if util.IntContains(playedNumbers, num) {
				matches++
			}
		}

		if matches > 0 {
			for i := gameNumber + 1; i <= gameNumber+matches; i++ {
				addScore(i + winningScratchCards[gameNumber])
			}
		}

		if winningScratchCards[gameNumber] > 0 {
			// se è presente un'altra schedina con lo stesso numero, scalo dalle "vinte" e la rigioco, non
			// incrementando il contatore "n"
			removeScore(gameNumber)
		} else {
			n++
		}
	}

	return numberScratched
}
