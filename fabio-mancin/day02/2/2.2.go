package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	//"strconv"
	"strings"
)

func calculateSum(filename string) (int, error) {

	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int

	for scanner.Scan() {
		s := scanner.Text()
		// gameNumber := gameNumberPattern.FindStringSubmatch(s)[1]
		counter := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		gameDataPattern := regexp.MustCompile(`:(.+)`)
		gameDataString := gameDataPattern.FindStringSubmatch(s)[1]
		gameDataString = strings.ReplaceAll(gameDataString, "; ", ", ")
		gameDataArray := strings.Split(gameDataString, ", ")
		for _, v := range gameDataArray {
			v = strings.TrimSpace(v)
			ballsArray := strings.Split(v, " ")
			ballsNumberString := ballsArray[0]
			ballsNumberString = strings.TrimSpace(ballsNumberString)
			ballsNumberInt, err := strconv.Atoi(ballsNumberString)
			if err != nil {
				return 0, err
			}
			ballsColor := ballsArray[1]
			for k, v := range counter {
				if k == ballsColor {
					if ballsNumberInt > v {
						counter[k] = ballsNumberInt
					}
				}
			}
		}
		power := 1
		for _, v := range counter {
			power *= v
		}
		sum += power
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
