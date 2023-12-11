package main

import (
	"fmt"
	"slices"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day11")

	// test data
	//dataByRow = []string{
	//	"...#......",
	//	".......#..",
	//	"#.........",
	//	"..........",
	//	"......#...",
	//	".#........",
	//	".........#",
	//	"..........",
	//	".......#..",
	//	"#...#.....",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
}

type galaxy struct {
	row int
	col int
}

func part01(dataByRow []string) int {

	var result int = 0

	newMap := expandMap(dataByRow)

	galaxies := findGalaxies(newMap)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += getDistance(galaxies[i], galaxies[j])
		}
	}

	return result
}

func part02(dataByRow []string) int {

	result := 0

	newMap := superExpandMap(dataByRow)

	galaxies := findGalaxies(newMap)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			result += getSuperExpandedDistance(newMap, galaxies[i], galaxies[j])
		}
	}

	return result
}

func expandMap(dataByRow []string) []string {

	var result []string

	columnsToExpand := []int{}

	for i := 0; i < len(dataByRow[0]); i++ {
		allEmpty := true
		for _, row := range dataByRow {
			if row[i] != '.' {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			columnsToExpand = append(columnsToExpand, i)
		}
	}

	for row := 0; row < len(dataByRow); row++ {
		newRow := ""
		allEmpty := true
		for column := 0; column < len(dataByRow[row]); column++ {
			if dataByRow[row][column] != '.' {
				allEmpty = false
			}
			newRow += string(dataByRow[row][column])
			if slices.Contains(columnsToExpand, column) {
				newRow += "."
			}
		}
		result = append(result, newRow)
		if allEmpty {
			result = append(result, newRow)
		}
	}

	return result
}

func superExpandMap(dataByRow []string) []string {

	var result []string

	columnsToExpand := []int{}

	for i := 0; i < len(dataByRow[0]); i++ {
		allEmpty := true
		for _, row := range dataByRow {
			if row[i] != '.' {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			columnsToExpand = append(columnsToExpand, i)
		}
	}

	for row := 0; row < len(dataByRow); row++ {
		newRow := ""
		allEmpty := true
		for column := 0; column < len(dataByRow[row]); column++ {
			if dataByRow[row][column] != '.' {
				allEmpty = false
			}

			if slices.Contains(columnsToExpand, column) {
				newRow += "X"
			} else {
				newRow += string(dataByRow[row][column])
			}
		}

		if allEmpty {
			result = append(result, strings.Repeat("X", len(newRow)))
		} else {
			result = append(result, newRow)
		}
	}

	return result
}

func findGalaxies(dataByRow []string) []galaxy {

	var result []galaxy

	for row := 0; row < len(dataByRow); row++ {
		for column := 0; column < len(dataByRow[row]); column++ {
			if dataByRow[row][column] == '#' {
				result = append(result, galaxy{row, column})
			}
		}
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getDistance(galaxy1 galaxy, galaxy2 galaxy) int {
	return abs(galaxy1.row-galaxy2.row) + abs(galaxy1.col-galaxy2.col)
}

func getSuperExpandedDistance(dataByRow []string, galaxy1 galaxy, galaxy2 galaxy) int {

	distance := 0

	times := 1000000

	lowerRow := galaxy1.row
	if galaxy2.row < galaxy1.row {
		lowerRow = galaxy2.row
	}
	higherRow := galaxy1.row
	if galaxy2.row > galaxy1.row {
		higherRow = galaxy2.row
	}
	lowerCol := galaxy1.col
	if galaxy2.col < galaxy1.col {
		lowerCol = galaxy2.col
	}
	higherCol := galaxy1.col
	if galaxy2.col > galaxy1.col {
		higherCol = galaxy2.col
	}

	// cols
	for c := lowerCol; c <= higherCol; c++ {
		if dataByRow[lowerRow][c] == 'X' {
			distance += times - 1 // minus one because we count the same step in the "delta" operation
		}
	}

	// rows
	for r := lowerRow; r <= higherRow; r++ {
		if dataByRow[r][lowerCol] == 'X' {
			distance += times - 1 // minus one because we count the same step in the "delta" operation
		}
	}

	distance += abs(galaxy1.row-galaxy2.row) + abs(galaxy1.col-galaxy2.col)

	return distance
}
