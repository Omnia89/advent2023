package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func calculateSum(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pattern := regexp.MustCompile(`\D`)
	var sum int

	for scanner.Scan() {
		s := pattern.ReplaceAllString(scanner.Text(), "")
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
	sum, err := calculateSum("./1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
