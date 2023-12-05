package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculateSum(filename string) (int, error) {
	R := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// match digits between a space and a semicolon
	gameNumberPattern := regexp.MustCompile(`\s(\d+):`)
	gameDataPattern := regexp.MustCompile(`:(.+)`)
	// gameDataArrayPattern := regexp.MustCompile(`(.+);(.+);(.+)`)
	var sum int

	for scanner.Scan() {
		s := scanner.Text()
		gameNumber := gameNumberPattern.FindStringSubmatch(s)[1]
		gameCanBeAdded := true
		gameData := gameDataPattern.FindStringSubmatch(s)[1]
		gameDataArray := strings.Split(gameData, "; ")
		for _, v := range gameDataArray {
			singleGameDataArray := strings.Split(v, ", ")
			for _, v := range singleGameDataArray {
				v = strings.TrimSpace(v)
				ballsArray := strings.Split(v, " ")
				ballsNumberString := ballsArray[0]
				ballsColor := ballsArray[1]
				ballsNumber, err := strconv.Atoi(ballsNumberString)
				if err != nil {
					return 0, err
				}
				if ballsNumber > R[ballsColor] {
					gameCanBeAdded = false
				}
			}
		}
		if gameCanBeAdded {
			gameNumberInt, err := strconv.Atoi(gameNumber)
			if err != nil {
				return 0, err
			}
			sum += gameNumberInt
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func main() {
	sum, err := calculateSum("../input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
