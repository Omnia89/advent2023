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
	M := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pattern := regexp.MustCompile(`\D`)
	var sum int

	for scanner.Scan() {
		s := scanner.Text()

		for key, value := range M {
			s = fmt.Sprintln(strings.ReplaceAll(s, key, key[:1]+value+key[len(key)-1:]))
		}

		s = pattern.ReplaceAllString(s, "")
		switch len(s) {
		case 0:
			s = "0"
		case 1:
			s += s
		default:
			s = s[:1] + s[len(s)-1:]
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		sum += num
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
