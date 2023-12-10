package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day10")

	// test data
	//dataByRow = []string{
	//	"FF7FSF7F7F7F7F7F---7",
	//	"L|LJ||||||||||||F--J",
	//	"FL-7LJLJ||||||LJL-77",
	//	"F--JF--7||LJLJ7F7FJ-",
	//	"L---JF-JLJ.||-FJLJJ7",
	//	"|F|F-JF---7F7-L7L|7|",
	//	"|FFJF7L7F-JF7|JL---7",
	//	"7-L-JL7||F7|L7F-7F7|",
	//	"L.L7LFJ|||||FJL7||LJ",
	//	"L7JLJL-JLJLJL--JLJ.L",
	//}

	startingRow, startingCol := findStartingPoint(dataByRow)

	// 01 part
	part01Result := part01(dataByRow, startingRow, startingCol)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow, startingRow, startingCol)
	fmt.Printf("Second part: %d\n", part02Result)
}

func part01(dataByRow []string, startingRow int, startingCol int) int {

	var result int = 0

	// replace S with the correct pipe
	pipe := findConnectingPipeUnderS(startingRow, startingCol, dataByRow)
	dataByRow[startingRow] = dataByRow[startingRow][:startingCol] + pipe + dataByRow[startingRow][startingCol+1:]

	// run down the pipe
	// priority directions: north, east, south, west

	numberOfSteps := 0
	currentRow := startingRow
	currentCol := startingCol
	comingFrom := "S"
	for numberOfSteps == 0 || currentRow != startingRow || currentCol != startingCol {
		currentPipe := dataByRow[currentRow][currentCol]
		if currentPipe == '|' {
			if comingFrom == "north" {
				currentRow++
				numberOfSteps++
				continue
			}
			if comingFrom == "south" {
				currentRow--
				numberOfSteps++
				continue
			}
			// if S or other (should not happen) go north
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			currentRow--
			numberOfSteps++
			continue
		}
		if currentPipe == '-' {
			if comingFrom == "east" {
				currentCol--
				numberOfSteps++
				continue
			}
			if comingFrom == "west" {
				currentCol++
				numberOfSteps++
				continue
			}
			// if S or other (should not happen) go east
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			currentCol++
			numberOfSteps++
			continue
		}

		if currentPipe == 'L' {
			if comingFrom == "east" {
				currentRow--
				numberOfSteps++
				comingFrom = "south"
				continue
			}
			if comingFrom == "north" {
				currentCol++
				numberOfSteps++
				comingFrom = "west"
				continue
			}
			// if S or other (should not happen) go north
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			currentRow--
			numberOfSteps++
			comingFrom = "south"
			continue
		}

		if currentPipe == 'J' {
			if comingFrom == "west" {
				currentRow--
				numberOfSteps++
				comingFrom = "south"
				continue
			}
			if comingFrom == "north" {
				currentCol--
				numberOfSteps++
				comingFrom = "east"
				continue
			}
			// if S or other (should not happen) go north
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			currentRow--
			numberOfSteps++
			comingFrom = "south"
			continue
		}

		if currentPipe == 'F' {
			if comingFrom == "east" {
				currentRow++
				numberOfSteps++
				comingFrom = "north"
				continue
			}
			if comingFrom == "south" {
				currentCol++
				numberOfSteps++
				comingFrom = "west"
				continue
			}
			// if S or other (should not happen) go east
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			currentCol++
			numberOfSteps++
			comingFrom = "west"
			continue
		}

		if currentPipe == '7' {
			if comingFrom == "west" {
				currentRow++
				numberOfSteps++
				comingFrom = "north"
				continue
			}
			if comingFrom == "south" {
				currentCol--
				numberOfSteps++
				comingFrom = "east"
				continue
			}
			// if S or other (should not happen) go west
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			currentCol--
			numberOfSteps++
			comingFrom = "east"
			continue
		}
		panic(fmt.Sprintf("Unknown pipe: %s, row: %d, col: %d", string(currentPipe), currentRow, currentCol))
	}

	result = numberOfSteps / 2

	return result
}

func part02(dataByRow []string, startingRow int, startingCol int) int {

	newMap := blankMap(dataByRow)

	// replace S with the correct pipe
	startingPipe := findConnectingPipeUnderS(startingRow, startingCol, dataByRow)

	setPipe := func(row int, col int, pipe string) {
		newMap[row] = newMap[row][:col] + pipe + newMap[row][col+1:]
	}

	setPipe(startingRow, startingCol, startingPipe)

	// run down the pipe and set walls
	currentRow := startingRow
	currentCol := startingCol
	comingFrom := "S"
	for comingFrom == "S" || currentRow != startingRow || currentCol != startingCol {
		currentPipe := dataByRow[currentRow][currentCol]
		if currentPipe == '|' {
			if comingFrom == "north" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentRow++
				continue
			}
			if comingFrom == "south" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentRow--
				continue
			}
			// if S or other (should not happen) go north
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			setPipe(currentRow, currentCol, string(currentPipe))
			currentRow--
			continue
		}
		if currentPipe == '-' {
			if comingFrom == "east" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentCol--
				continue
			}
			if comingFrom == "west" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentCol++
				continue
			}
			// if S or other (should not happen) go east
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			setPipe(currentRow, currentCol, string(currentPipe))
			currentCol++
			continue
		}

		if currentPipe == 'L' {
			if comingFrom == "east" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentRow--
				comingFrom = "south"
				continue
			}
			if comingFrom == "north" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentCol++
				comingFrom = "west"
				continue
			}
			// if S or other (should not happen) go north
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			setPipe(currentRow, currentCol, string(currentPipe))
			currentRow--
			comingFrom = "south"
			continue
		}

		if currentPipe == 'J' {
			if comingFrom == "west" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentRow--
				comingFrom = "south"
				continue
			}
			if comingFrom == "north" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentCol--
				comingFrom = "east"
				continue
			}
			// if S or other (should not happen) go north
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			setPipe(currentRow, currentCol, string(currentPipe))
			currentRow--
			comingFrom = "south"
			continue
		}

		if currentPipe == 'F' {
			if comingFrom == "east" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentRow++
				comingFrom = "north"
				continue
			}
			if comingFrom == "south" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentCol++
				comingFrom = "west"
				continue
			}
			// if S or other (should not happen) go east
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			setPipe(currentRow, currentCol, string(currentPipe))
			currentCol++
			comingFrom = "west"
			continue
		}

		if currentPipe == '7' {
			if comingFrom == "west" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentRow++
				comingFrom = "north"
				continue
			}
			if comingFrom == "south" {
				setPipe(currentRow, currentCol, string(currentPipe))
				currentCol--
				comingFrom = "east"
				continue
			}
			// if S or other (should not happen) go west
			fmt.Printf("WARNING: only one timel row: %d, col: %d\n", currentRow, currentCol)
			setPipe(currentRow, currentCol, string(currentPipe))
			currentCol--
			comingFrom = "east"
			continue
		}
		panic(fmt.Sprintf("Unknown pipe: %s, row: %d, col: %d", string(currentPipe), currentRow, currentCol))
	}

	// find area between walls, using "point in polygon" algorithm (https://en.wikipedia.org/wiki/Point_in_polygon)

	area := 0
	for row := 0; row < len(newMap); row++ {
		insideWall := false
		//tempArea := 0
		lastVerticalDirection := ""
		for col := 0; col < len(newMap[row]); col++ {

			// se terreno e dentro muro, incremento area
			if newMap[row][col] == '.' {
				if insideWall {
					//tempArea++
					area++
				}
				continue
			}

			// se incontro un muro
			if newMap[row][col] == '|' {
				//if insideWall {
				//	area += tempArea
				//	tempArea = 0
				//}

				insideWall = !insideWall
			}

			// se orizontale, skippo
			if newMap[row][col] == '-' {
				continue
			}

			verticalDirection := ""
			if newMap[row][col] == 'L' || newMap[row][col] == 'J' {
				verticalDirection = "n"
			} else if newMap[row][col] == 'F' || newMap[row][col] == '7' {
				verticalDirection = "s"
			}

			if lastVerticalDirection == "" {
				lastVerticalDirection = verticalDirection
				continue
			}

			if verticalDirection != lastVerticalDirection {
				insideWall = !insideWall
			}
			lastVerticalDirection = ""
		}
	}

	return area
}

func findStartingPoint(dataByRow []string) (row int, col int) {
	for i, line := range dataByRow {
		if charIndex := strings.Index(line, "S"); charIndex != -1 {
			return i, charIndex
		}
	}
	return -1, -1
}

func findConnectingPipeUnderS(row int, col int, dataByRow []string) (pipe string) {
	connectNorth := false
	connectSouth := false
	connectWest := false
	connectEast := false

	if row-1 >= 0 {
		if dataByRow[row-1][col] == '|' ||
			dataByRow[row-1][col] == 'F' ||
			dataByRow[row-1][col] == '7' {
			connectNorth = true
		}
	}

	if row+1 < len(dataByRow) {
		if dataByRow[row+1][col] == '|' ||
			dataByRow[row+1][col] == 'L' ||
			dataByRow[row+1][col] == 'J' {
			connectSouth = true
		}
	}

	if col-1 >= 0 {
		if dataByRow[row][col-1] == '-' ||
			dataByRow[row][col-1] == 'F' ||
			dataByRow[row][col-1] == 'L' {
			connectWest = true
		}
	}

	if col+1 < len(dataByRow[row]) {
		if dataByRow[row][col+1] == '-' ||
			dataByRow[row][col+1] == '7' ||
			dataByRow[row][col+1] == 'J' {
			connectEast = true
		}
	}

	if connectNorth && connectSouth {
		return "|"
	}
	if connectWest && connectEast {
		return "-"
	}

	if connectNorth && connectEast {
		return "L"
	}
	if connectNorth && connectWest {
		return "J"
	}

	if connectSouth && connectEast {
		return "F"
	}
	if connectSouth && connectWest {
		return "7"
	}

	return "."
}

func blankMap(dataByRow []string) []string {
	result := make([]string, len(dataByRow))
	for i := range result {
		result[i] = strings.Repeat(".", len(dataByRow[i]))
	}
	return result
}
