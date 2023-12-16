package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day16")

	// test data
	//dataByRow = []string{
	//	`.|...\....`,
	//	`|.-.\.....`,
	//	`.....|-...`,
	//	`........|.`,
	//	`..........`,
	//	`.........\`,
	//	`..../.\\..`,
	//	`.-.-/..|..`,
	//	`.|....-|.\`,
	//	`..//.|....`,
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

	// first beam, -1 one to get the first advance
	firstBeam := beam{0, -1, Right}
	energizedRows := travel(dataByRow, firstBeam)
	result = countEnergy(energizedRows)

	return result
}

func part02(dataByRow []string) int {
	result := 0

	// vertical
	for i := 0; i < len(dataByRow[0]); i++ {
		// downward
		firstBeam := beam{-1, i, Down}
		energizedRows := countEnergy(travel(dataByRow, firstBeam))
		if energizedRows > result {
			result = energizedRows
		}

		// upward
		firstBeam = beam{len(dataByRow), i, Up}
		energizedRows = countEnergy(travel(dataByRow, firstBeam))
		if energizedRows > result {
			result = energizedRows
		}
	}

	// horizontal
	for i := 0; i < len(dataByRow); i++ {
		// rightward
		firstBeam := beam{i, -1, Right}
		energizedRows := countEnergy(travel(dataByRow, firstBeam))
		if energizedRows > result {
			result = energizedRows
		}

		// leftward
		firstBeam = beam{i, len(dataByRow[0]), Left}
		energizedRows = countEnergy(travel(dataByRow, firstBeam))
		if energizedRows > result {
			result = energizedRows
		}
	}

	return result
}

// directions:
const (
	Up    = 8
	Down  = 4
	Left  = 2
	Right = 1
)

func isGoingTo(energizedCell int, direction int) bool {
	switch energizedCell {
	case Up:
		return direction&Up == Up
	case Down:
		return direction&Down == Down
	case Left:
		return direction&Left == Left
	case Right:
		return direction&Right == Right
	}
	return false
}

type beam struct {
	row       int
	col       int
	direction int
}

func (b *beam) advance() {
	switch b.direction {
	case Up:
		b.row--
	case Down:
		b.row++
	case Left:
		b.col--
	case Right:
		b.col++
	}
}

type beamStack []beam

func (s *beamStack) Push(v beam) {
	*s = append(*s, v)
}

func (s *beamStack) Pop() *beam {
	beam := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &beam
}

func travel(dataByRow []string, firstBeam beam) []string {
	result := make([]string, len(dataByRow))
	directionHistory := make([][]int, len(dataByRow))
	for i := range result {
		result[i] = strings.Repeat(".", len(dataByRow[i]))
		directionHistory[i] = make([]int, len(dataByRow[i]))
	}

	beams := make(beamStack, 0)

	beams.Push(firstBeam)

	var currentBeam *beam
	currentBeam = beams.Pop()

	for len(beams) > 0 || currentBeam != nil {
		if currentBeam == nil {
			currentBeam = beams.Pop()
		}
		currentBeam.advance()

		// check if out of bounds
		if currentBeam.row < 0 || currentBeam.row >= len(dataByRow) || currentBeam.col < 0 || currentBeam.col >= len(dataByRow[currentBeam.row]) {
			currentBeam = nil
			continue
		}

		// check if cell already energized in the same direction
		cell := directionHistory[currentBeam.row][currentBeam.col]
		if isGoingTo(cell, currentBeam.direction) {
			//path already energized
			currentBeam = nil
			continue
		}

		// energize cell
		directionHistory[currentBeam.row][currentBeam.col] = cell | currentBeam.direction
		result[currentBeam.row] = result[currentBeam.row][:currentBeam.col] + "#" + result[currentBeam.row][currentBeam.col+1:]

		// check mirrors
		currBeam, newBeam := newDirection(currentBeam, dataByRow[currentBeam.row][currentBeam.col])
		currentBeam = currBeam
		if newBeam != nil {
			beams.Push(*newBeam)
		}
	}

	return result
}

func newDirection(currentBeam *beam, mirror byte) (modifiedBeam *beam, newBeam *beam) {
	newBeam = nil
	if mirror == '/' {
		if currentBeam.direction == Right {
			currentBeam.direction = Up
		} else if currentBeam.direction == Left {
			currentBeam.direction = Down
		} else if currentBeam.direction == Up {
			currentBeam.direction = Right
		} else if currentBeam.direction == Down {
			currentBeam.direction = Left
		}
	} else if mirror == '\\' {
		if currentBeam.direction == Right {
			currentBeam.direction = Down
		} else if currentBeam.direction == Left {
			currentBeam.direction = Up
		} else if currentBeam.direction == Up {
			currentBeam.direction = Left
		} else if currentBeam.direction == Down {
			currentBeam.direction = Right
		}
	} else if mirror == '-' {
		if currentBeam.direction == Up || currentBeam.direction == Down {
			currentBeam.direction = Left
			newBeam = &beam{row: currentBeam.row, col: currentBeam.col, direction: Right}
		}
	} else if mirror == '|' {
		if currentBeam.direction == Left || currentBeam.direction == Right {
			currentBeam.direction = Up
			newBeam = &beam{row: currentBeam.row, col: currentBeam.col, direction: Down}
		}
	}

	return currentBeam, newBeam
}

func countEnergy(dataByRow []string) int {
	result := 0
	for _, row := range dataByRow {
		for _, cell := range row {
			if cell == '#' {
				result++
			}
		}
	}
	return result
}
