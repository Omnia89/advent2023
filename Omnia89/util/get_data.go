package util

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetRawData(day string) string {
	// open "data.txt" file in current directory
	file, err := os.Open(fmt.Sprintf(`%s/data.txt`, day))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file
	res, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(res)
}

func GetDataByRow(day string) []string {
	// get raw data
	rawData := GetRawData(day)

	// split raw data by row
	dataByRow := strings.Split(rawData, "\n")

	return dataByRow
}

func ToInt(s string) int {
	t := strings.TrimSpace(s)
	r, _ := strconv.Atoi(t)
	return r
}

func StringToIntSlice(s string, separator string) []int {
	parts := strings.Split(s, separator)
	var numbers []int
	for _, s := range parts {
		if s != "" {
			numbers = append(numbers, ToInt(s))
		}
	}
	return numbers
}

func IntContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
