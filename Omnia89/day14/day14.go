package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day14")

	// test data
	//dataByRow = []string{
	//	"O....#....",
	//	"O.OO#....#",
	//	".....##...",
	//	"OO.#O....O",
	//	".O.....O#.",
	//	"O.#..O.#.#",
	//	"..O..#O..O",
	//	".......O..",
	//	"#....###..",
	//	"#OO..#....",
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

	tiltedRows := tiltNorth(dataByRow)
	result = getLoadValue(tiltedRows)

	return result
}

func part02(dataByRow []string) int {

	result := 0

	hashes := map[string]int{}

	cycles := 1_000_000_000
	tiltedRows := dataByRow
	for i := 0; i < cycles; i++ {
		tiltedRows = tiltCycle(tiltedRows)
		hash := strings.Join(tiltedRows, "")
		if _, ok := hashes[hash]; ok {
			i = cycles - (cycles-i)%(i-hashes[hash])
		}
		hashes[hash] = i
	}
	result = getLoadValue(tiltedRows)

	return result
}

func tiltCycle(rows []string) []string {
	tiltedRows := tiltNorth(rows)
	tiltedRows = tiltWest(tiltedRows)
	tiltedRows = tiltSouth(tiltedRows)
	tiltedRows = tiltEast(tiltedRows)
	return tiltedRows
}

func tiltNorth(rows []string) []string {
	result := make([]string, len(rows))
	for i := 0; i < len(rows); i++ {
		result[i] = strings.Repeat(".", len(rows))
	}

	placeChar := func(row, col int, char string) {
		result[row] = result[row][:col] + char + result[row][col+1:]
	}

	for col := 0; col < len(rows[0]); col++ {
		emptySpaceIndex := -1
		rocksCount := 0
		for row := 0; row < len(rows); row++ {
			currentChar := rows[row][col]

			if emptySpaceIndex == -1 {
				if currentChar != '#' {
					emptySpaceIndex = row
				} else {
					// if square rock and no empty space, write it (consecutives rocks)
					placeChar(row, col, "#")
					continue
				}
			}

			if currentChar == 'O' {
				rocksCount++
			}
			if currentChar == '#' {
				// shift rocks
				for i := emptySpaceIndex; i < row; i++ {
					if rocksCount > 0 {
						placeChar(i, col, "O")
						rocksCount--
					} else {
						placeChar(i, col, ".")
					}
				}
				// add current char
				placeChar(row, col, string(currentChar))
				emptySpaceIndex = -1
			}
		}
		// shift rocks
		if rocksCount > 0 {
			for i := emptySpaceIndex; i < len(rows); i++ {
				if rocksCount > 0 {
					placeChar(i, col, "O")
					rocksCount--
				} else {
					placeChar(i, col, ".")
				}
			}
		}
	}

	return result
}

func tiltSouth(rows []string) []string {
	result := make([]string, len(rows))
	for i := 0; i < len(rows); i++ {
		result[i] = strings.Repeat(".", len(rows))
	}

	placeChar := func(row, col int, char string) {
		result[row] = result[row][:col] + char + result[row][col+1:]
	}

	for col := 0; col < len(rows[0]); col++ {
		emptySpaceIndex := -1
		rocksCount := 0
		for row := len(rows) - 1; row >= 0; row-- {
			currentChar := rows[row][col]

			if emptySpaceIndex == -1 {
				if currentChar != '#' {
					emptySpaceIndex = row
				} else {
					// if square rock and no empty space, write it (consecutives rocks)
					placeChar(row, col, "#")
					continue
				}
			}

			if currentChar == 'O' {
				rocksCount++
			}
			if currentChar == '#' {
				// shift rocks
				for i := emptySpaceIndex; i > row; i-- {
					if rocksCount > 0 {
						placeChar(i, col, "O")
						rocksCount--
					} else {
						placeChar(i, col, ".")
					}
				}
				// add current char
				placeChar(row, col, string(currentChar))
				emptySpaceIndex = -1
			}
		}
		// shift rocks
		if rocksCount > 0 {
			for i := emptySpaceIndex; i >= 0; i-- {
				if rocksCount > 0 {
					placeChar(i, col, "O")
					rocksCount--
				} else {
					placeChar(i, col, ".")
				}
			}
		}
	}

	return result
}

func tiltEast(rows []string) []string {
	result := make([]string, len(rows))
	for i := 0; i < len(rows); i++ {
		result[i] = strings.Repeat(".", len(rows))
	}

	placeChar := func(row, col int, char string) {
		result[row] = result[row][:col] + char + result[row][col+1:]
	}

	for row := 0; row < len(rows); row++ {
		emptySpaceIndex := -1
		rocksCount := 0
		for col := len(rows[0]) - 1; col >= 0; col-- {
			currentChar := rows[row][col]

			if emptySpaceIndex == -1 {
				if currentChar != '#' {
					emptySpaceIndex = col
				} else {
					// if square rock and no empty space, write it (consecutives rocks)
					placeChar(row, col, "#")
					continue
				}
			}

			if currentChar == 'O' {
				rocksCount++
			}
			if currentChar == '#' {
				// shift rocks
				for i := emptySpaceIndex; i > col; i-- {
					if rocksCount > 0 {
						placeChar(row, i, "O")
						rocksCount--
					} else {
						placeChar(row, i, ".")
					}
				}
				// add current char
				placeChar(row, col, string(currentChar))
				emptySpaceIndex = -1
			}
		}
		// shift rocks
		if rocksCount > 0 {
			for i := emptySpaceIndex; i >= 0; i-- {
				if rocksCount > 0 {
					placeChar(row, i, "O")
					rocksCount--
				} else {
					placeChar(row, i, ".")
				}
			}
		}
	}

	return result
}

func tiltWest(rows []string) []string {
	result := make([]string, len(rows))
	for i := 0; i < len(rows); i++ {
		result[i] = strings.Repeat(".", len(rows))
	}

	placeChar := func(row, col int, char string) {
		result[row] = result[row][:col] + char + result[row][col+1:]
	}

	for row := 0; row < len(rows); row++ {
		emptySpaceIndex := -1
		rocksCount := 0
		for col := 0; col < len(rows[0]); col++ {
			currentChar := rows[row][col]

			if emptySpaceIndex == -1 {
				if currentChar != '#' {
					emptySpaceIndex = col
				} else {
					// if square rock and no empty space, write it (consecutives rocks)
					placeChar(row, col, "#")
					continue
				}
			}

			if currentChar == 'O' {
				rocksCount++
			}
			if currentChar == '#' {
				// shift rocks
				for i := emptySpaceIndex; i < col; i++ {
					if rocksCount > 0 {
						placeChar(row, i, "O")
						rocksCount--
					} else {
						placeChar(row, i, ".")
					}
				}
				// add current char
				placeChar(row, col, string(currentChar))
				emptySpaceIndex = -1
			}
		}
		// shift rocks
		if rocksCount > 0 {
			for i := emptySpaceIndex; i < len(rows[0]); i++ {
				if rocksCount > 0 {
					placeChar(row, i, "O")
					rocksCount--
				} else {
					placeChar(row, i, ".")
				}
			}
		}
	}

	return result
}

func getLoadValue(rows []string) int {
	result := 0
	for i := 0; i < len(rows); i++ {
		weight := len(rows) - i
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == 'O' {
				result += weight
			}
		}
	}
	return result
}
