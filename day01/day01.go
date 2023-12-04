package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent2023/util"
)

func main() {
	// get data by row
	dataByRow := util.GetDataByRow("day01")

	// 01 part

	sum := 0

	re := regexp.MustCompile(`\d`)

	for _, row := range dataByRow {
		firstDigit := ""
		lastDigit := ""

		allDigits := re.FindAllString(row, -1)

		if len(allDigits) > 0 {
			firstDigit = allDigits[0]
			lastDigit = allDigits[len(allDigits)-1]
			number, _ := strconv.Atoi(firstDigit + lastDigit)
			sum += number
		}
	}

	println(fmt.Sprintf("First part: %d", sum))

	//testData := []string{
	//	"two1nine",
	//	"eightwothree",
	//	"abcone2threexyz",
	//	"xtwone3four",
	//	"4nineeightseven2",
	//	"zoneight234",
	//	"7pqrstsixteen",
	//}

	// 02 part

	sum = 0

	for _, row := range dataByRow {
		digits := getDigits(row)
		//fmt.Printf("%s -> %d\n", row, firstDigit+lastDigit)
		sum += digits[0]*10 + digits[len(digits)-1]
	}

	println(fmt.Sprintf("Second part: %d", sum))
}

func getDigits(row string) []int {
	numbers := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var values []int

	accumulated := ""
	for _, char := range row {
		accumulated += string(char)

		// Check if a characters is a number
		if char >= 48 && char <= 57 {
			values = append(values, int(char)-'0')
		}

		for value, number := range numbers {
			if strings.HasSuffix(accumulated, value) {
				values = append(values, number)
			}
		}
	}

	return values
}
