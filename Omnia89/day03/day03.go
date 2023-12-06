package main

import (
	"fmt"
	"strconv"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day03")

	// testData
	//dataByRow := []string{
	//	"467..114..",
	//	"...*......",
	//	"..35..633.",
	//	"......#...",
	//	"617*......",
	//	".....+.58.",
	//	"..592.....",
	//	"......755.",
	//	"...$.*....",
	//	".664.598..",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)

}

func part01(dataByRow []string) int {

	sum := 0

	for rowIndex, row := range dataByRow {
		accumulatedNumber := ""
		initIndex := -1
		for charIndex, char := range row {
			if isNumber(uint8(char)) {
				accumulatedNumber += string(char)
				if initIndex == -1 {
					initIndex = charIndex
				}
				continue
			}

			// if not a number, search symbols
			if accumulatedNumber != "" {
				if isNextToSymbol(dataByRow, rowIndex, initIndex, charIndex-1) {
					number, _ := strconv.Atoi(accumulatedNumber)
					sum += number
				}
				accumulatedNumber = ""
				initIndex = -1
			}
		}

		// to catch last number
		// if not a number, search symbols
		if accumulatedNumber != "" {
			if isNextToSymbol(dataByRow, rowIndex, initIndex, len(row)-1) {
				number, _ := strconv.Atoi(accumulatedNumber)
				sum += number
			}
		}
	}

	return sum
}

func part02(dataByRow []string) int {

	type rowColCoord = string

	gearMap := map[rowColCoord]int{}

	sum := 0

	for rowIndex, row := range dataByRow {
		accumulatedNumber := ""
		initIndex := -1
		for charIndex, char := range row {
			if isNumber(uint8(char)) {
				accumulatedNumber += string(char)
				if initIndex == -1 {
					initIndex = charIndex
				}
				continue
			}

			// if not a number, search gear
			if accumulatedNumber != "" {
				if c, gearRow, gearCol := isNextToGear(dataByRow, rowIndex, initIndex, charIndex-1); c {
					number, _ := strconv.Atoi(accumulatedNumber)

					index := fmt.Sprintf("%d-%d", gearRow, gearCol)
					if _, ok := gearMap[index]; !ok {
						gearMap[index] = number
					} else {
						toAdd := gearMap[index] * number
						sum += toAdd
					}
				}
				accumulatedNumber = ""
				initIndex = -1
			}
		}

		// to catch last number
		// if not a number, search gear
		if accumulatedNumber != "" {
			if c, gearRow, gearCol := isNextToGear(dataByRow, rowIndex, initIndex, len(row)-1); c {
				number, _ := strconv.Atoi(accumulatedNumber)

				index := fmt.Sprintf("%d-%d", gearRow, gearCol)
				if _, ok := gearMap[index]; !ok {
					gearMap[index] = number
				} else {
					toAdd := gearMap[index] * number
					sum += toAdd
				}
			}
		}
	}

	return sum

}

func isNumber(intChar uint8) bool {
	char := string(intChar)
	return char == "0" ||
		char == "1" ||
		char == "2" ||
		char == "3" ||
		char == "4" ||
		char == "5" ||
		char == "6" ||
		char == "7" ||
		char == "8" ||
		char == "9"
}

func isSymbol(char uint8) bool {
	strChar := string(char)
	return !isNumber(char) && strChar != "."
}

func isGear(char uint8) bool {
	strChar := string(char)
	return strChar == "*"
}

func isNextToSymbol(dataMap []string, rowIndex int, initColumn int, endColumn int) (res bool) {

	diagonalInitColumn := initColumn
	diagonalEndColumn := endColumn

	if initColumn != 0 {
		diagonalInitColumn = initColumn - 1
		if isSymbol(dataMap[rowIndex][initColumn-1]) {
			return true
		}
	}

	if endColumn < len(dataMap[rowIndex])-1 {
		diagonalEndColumn = endColumn + 1
		if isSymbol(dataMap[rowIndex][endColumn+1]) {
			return true
		}
	}

	if rowIndex != 0 {
		for i := diagonalInitColumn; i <= diagonalEndColumn; i++ {
			if isSymbol(dataMap[rowIndex-1][i]) {
				return true
			}
		}
	}

	if rowIndex < len(dataMap)-1 {
		for i := diagonalInitColumn; i <= diagonalEndColumn; i++ {
			if isSymbol(dataMap[rowIndex+1][i]) {
				return true
			}
		}
	}

	return false
}

func isNextToGear(dataMap []string, rowIndex int, initColumn int, endColumn int) (res bool, row int, col int) {

	diagonalInitColumn := initColumn
	diagonalEndColumn := endColumn

	if initColumn != 0 {
		diagonalInitColumn = initColumn - 1
		if isGear(dataMap[rowIndex][initColumn-1]) {
			return true, rowIndex, initColumn - 1
		}
	}

	if endColumn < len(dataMap[rowIndex])-1 {
		diagonalEndColumn = endColumn + 1
		if isGear(dataMap[rowIndex][endColumn+1]) {
			return true, rowIndex, endColumn + 1
		}
	}

	if rowIndex != 0 {
		for i := diagonalInitColumn; i <= diagonalEndColumn; i++ {
			if isGear(dataMap[rowIndex-1][i]) {
				return true, rowIndex - 1, i
			}
		}
	}

	if rowIndex < len(dataMap)-1 {
		for i := diagonalInitColumn; i <= diagonalEndColumn; i++ {
			if isGear(dataMap[rowIndex+1][i]) {
				return true, rowIndex + 1, i
			}
		}
	}

	return false, -1, -1
}
