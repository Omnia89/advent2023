package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isValidSymbol(symbol rune) bool {
	return !unicode.IsNumber(symbol) && symbol != '.' && symbol != '\n'
}

func calculateSum(filename string) (int, error) {

	content, err := os.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(content), "\n")
	sum := 0

	for lineIndex, line := range lines {
		lineLength := len(line)
		var lineBefore string
		var lineAfter string
		if lineIndex == 0 {
			lineBefore = strings.Repeat(".", lineLength)
		} else {
			lineBefore = lines[lineIndex-1]
		}
		if lineIndex == len(lines)-1 {
			lineAfter = strings.Repeat(".", lineLength)
		} else {
			lineAfter = lines[lineIndex+1]
		}
		var numberString string
		var numberShouldCount bool
		numberEndedAtIndex := -1
		var indexesToCheck []int
		for charIndex, char := range line {

			lineRunes := []rune(line)
			lineBeforeRunes := []rune(lineBefore)
			lineAfterRunes := []rune(lineAfter)

			if unicode.IsNumber(char) {
				numberString += string(char)
				indexesToCheck = append(indexesToCheck, charIndex)
				if charIndex != 0 && !unicode.IsNumber(lineRunes[charIndex-1]) {
					if isValidSymbol(lineRunes[charIndex-1]) {
						numberShouldCount = true
					} else {
						indexesToCheck = append(indexesToCheck, charIndex-1)
					}
				}
				if !(charIndex == len(lineRunes)-1) && !unicode.IsNumber(lineRunes[charIndex+1]) {
					numberEndedAtIndex = charIndex + 1
					if isValidSymbol(lineRunes[charIndex+1]) {
						numberShouldCount = true
					} else {
						indexesToCheck = append(indexesToCheck, charIndex+1)
					}
				}
				if charIndex == len(lineRunes)-1 {
					numberEndedAtIndex = charIndex
				}
			}

			if numberEndedAtIndex != -1 {
				if numberShouldCount {
					number, err := strconv.Atoi(numberString)
					if err != nil {
						return 0, err
					}
					sum += number
				} else {
					for _, index := range indexesToCheck {
						if isValidSymbol(lineAfterRunes[index]) || isValidSymbol(lineBeforeRunes[index]) {
							number, err := strconv.Atoi(numberString)
							if err != nil {
								return 0, err
							}
							sum += number
							break
						}
					}
				}
				numberString = ""
				numberShouldCount = false
				numberEndedAtIndex = -1
				indexesToCheck = nil
			}
		}
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
