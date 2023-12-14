package main

import (
	"fmt"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day13")

	// test data
	//dataByRow = []string{
	//	"#.##..##.",
	//	"..#.##.#.",
	//	"##......#",
	//	"##......#",
	//	"..#.##.#.",
	//	"..##..##.",
	//	"#.#.##.#.",
	//	"",
	//	"#...##..#",
	//	"#....#..#",
	//	"..##..###",
	//	"#####.##.",
	//	"#####.##.",
	//	"..##..###",
	//	"#....#..#",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
}

func part01(dataByRow []string) int {

	var result int = 0

	patterns := getPatterns(dataByRow)

	for _, pattern := range patterns {
		result += getReflectionValue(pattern, false, -1, -1)
	}

	return result
}

func part02(dataByRow []string) int {

	result := 0

	patterns := getPatterns(dataByRow)

	for _, pattern := range patterns {
		maxPatternValue := 0
		for r := 0; r < len(pattern); r++ {
			for c := 0; c < len(pattern[0]); c++ {
				smudgedPattern := smudgePattern(pattern, r, c)
				if value := getReflectionValue(smudgedPattern, true, r, c); value > maxPatternValue {
					maxPatternValue = value
				}
			}
		}
		result += maxPatternValue
	}

	return result
}

func smudgePattern(pattern []string, r int, c int) []string {
	char := pattern[r][c]
	newChar := ""
	if char == '#' {
		newChar = "."
	} else {
		newChar = "#"
	}

	clonedPattern := make([]string, len(pattern))
	copy(clonedPattern, pattern)

	clonedPattern[r] = clonedPattern[r][:c] + newChar + clonedPattern[r][c+1:]
	return clonedPattern
}

func getReflectionValue(pattern []string, getMax bool, rowToInclude int, colToInclude int) int {
	result := 0
	verticalValue := 0
	horizontalValue := 0
	// vertical
	for i := 0; i < len(pattern[0]); i++ {
		indexToInclude := -1
		if getMax {
			indexToInclude = colToInclude
		}
		if isRowSymmetrical(pattern[0], i, indexToInclude) {
			// check if the rest of the pattern is symmetrical
			isSymmetrical := true
			for j := 1; j < len(pattern); j++ {
				if !isRowSymmetrical(pattern[j], i, indexToInclude) {
					isSymmetrical = false
					break
				}
			}
			if isSymmetrical {
				verticalValue += i + 1
				break
			}
		}
	}

	// horizontal
	for i := 0; i < len(pattern); i++ {
		indexToInclude := -1
		if getMax {
			indexToInclude = rowToInclude
		}
		if isColumnSymmetrical(pattern, 0, i, indexToInclude) {
			// check if the rest of the pattern is symmetrical
			isSymmetrical := true
			for j := 1; j < len(pattern[0]); j++ {
				if !isColumnSymmetrical(pattern, j, i, indexToInclude) {
					isSymmetrical = false
					break
				}
			}
			if isSymmetrical {
				horizontalValue += (i + 1) * 100
				break
			}
		}
	}

	if getMax {
		if verticalValue > horizontalValue {
			result = verticalValue
		} else {
			result = horizontalValue
		}
	} else {
		result = verticalValue + horizontalValue
	}

	return result
}

func getPatterns(dataByRow []string) [][]string {
	var result [][]string

	var temp []string
	for _, row := range dataByRow {
		if row == "" {
			result = append(result, temp)
			temp = make([]string, 0)
			continue
		}
		temp = append(temp, row)
	}
	if len(temp) > 0 {
		result = append(result, temp)
	}
	return result
}

// the mirror is placed between index and index+1
func isRowSymmetrical(row string, index int, indexToInclude int) bool {
	if index == len(row)-1 {
		return false
	}
	maxStep := len(row) - index - 2
	if index < maxStep {
		maxStep = index
	}

	if indexToInclude != -1 {
		if indexToInclude > index+1+maxStep || indexToInclude < index-maxStep {
			return false
		}
	}

	for i := 0; i <= maxStep; i++ {
		left := row[index-i]
		right := row[index+1+i]
		if left != right {
			return false
		}
	}
	return true
}

// the mirror is placed between index and index+1
func isColumnSymmetrical(pattern []string, column int, index int, indexToInclude int) bool {
	if index == len(pattern)-1 || column >= len(pattern[0]) {
		return false
	}

	// transpose to a row
	row := ""
	for _, line := range pattern {
		row += string(line[column])
	}

	// check if the row is symmetrical
	return isRowSymmetrical(row, index, indexToInclude)
}
