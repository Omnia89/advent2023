package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day06")

	// test data
	//dataByRow = []string{
	//	"Time:      7  15   30",
	//	"Distance:  9  40  200",
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

	result := 1

	timesPart := strings.Split(dataByRow[0], ":")
	distancesPart := strings.Split(dataByRow[1], ":")

	times := util.StringToIntSlice(timesPart[1], " ")
	distances := util.StringToIntSlice(distancesPart[1], " ")

	recordBeats := []int{}
	for i := 0; i < len(times); i++ {

		beatCounter := 0

		// for each possible "hold time"
		for j := 0; j <= times[i]; j++ {
			if getDistance(times[i], j) > distances[i] {
				beatCounter++
			}
		}

		recordBeats = append(recordBeats, beatCounter)
	}

	for _, v := range recordBeats {
		result *= v
	}

	return result
}

func part02(dataByRow []string) int {

	result := 0

	timePart := strings.Split(dataByRow[0], ":")
	distancePart := strings.Split(dataByRow[1], ":")

	stringTime := strings.ReplaceAll(timePart[1], " ", "")
	stringDistance := strings.ReplaceAll(distancePart[1], " ", "")

	time := util.ToInt(stringTime)
	distance := util.ToInt(stringDistance)

	for j := 1; j < time; j++ {
		if getDistance(time, j) > distance {
			result++
		}
	}

	return result
}

func getDistance(totTime int, holdTime int) int {
	if holdTime >= totTime || holdTime == 0 {
		return 0
	}

	travelTime := totTime - holdTime

	speed := holdTime

	return travelTime * speed
}
