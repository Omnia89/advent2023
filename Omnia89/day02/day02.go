package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent2023/util"
)

func main() {
	// get data by row
	dataByRow := util.GetDataByRow("day02")

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
}

func part01(dataByRow []string) int {
	sum := 0

	checks := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

outerLoop:
	for _, row := range dataByRow {
		gameParts := strings.Split(row, ":")
		gameNumber, _ := strconv.ParseInt(strings.Split(gameParts[0], " ")[1], 10, 64)

		rounds := strings.Split(gameParts[1], ";")
		allExtractions := []string{}

		for _, round := range rounds {
			smallParts := strings.Split(round, ",")
			for _, smallPart := range smallParts {
				allExtractions = append(allExtractions, strings.TrimSpace(smallPart))
			}
		}

		for _, extraction := range allExtractions {
			extractionParts := strings.Split(extraction, " ")
			extractionNumber, _ := strconv.ParseInt(extractionParts[0], 10, 64)
			if extractionNumber > checks[extractionParts[1]] {
				continue outerLoop
			}
		}
		sum += int(gameNumber)
	}
	return sum
}

func part02(dataByRow []string) int64 {
	var sum int64 = 0
	for _, row := range dataByRow {
		gameParts := strings.Split(row, ":")

		rounds := strings.Split(gameParts[1], ";")
		allExtractions := []string{}

		minimunValues := map[string]int64{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range rounds {
			smallParts := strings.Split(round, ",")
			for _, smallPart := range smallParts {
				allExtractions = append(allExtractions, strings.TrimSpace(smallPart))
			}
		}

		for _, extraction := range allExtractions {
			extractionParts := strings.Split(extraction, " ")
			extractionNumber, _ := strconv.ParseInt(extractionParts[0], 10, 64)
			if extractionNumber > minimunValues[extractionParts[1]] {
				minimunValues[extractionParts[1]] = extractionNumber
			}
		}
		power := minimunValues["red"] * minimunValues["green"] * minimunValues["blue"]
		sum += power
	}
	return sum
}
