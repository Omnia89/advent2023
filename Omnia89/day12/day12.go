package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day12")

	// test data
	//dataByRow = []string{
	//	"???.### 1,1,3",
	//	".??..??...?##. 1,1,3",
	//	"?#?#?#?#?#?#?#? 1,3,1,6",
	//	"????.#...#... 4,1,1",
	//	"????.######..#####. 1,6,5",
	//	"?###???????? 3,2,1",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
}

type HotSpringRow struct {
	Row        string
	DamageList []int
}

func part01(dataByRow []string) int {

	var result int = 0

	for _, hotRow := range parseData(dataByRow) {
		rowResult := recursive(hotRow.Row, hotRow.DamageList, 0)
		result += rowResult
	}

	return result
}

func part02(dataByRow []string) int {

	result := 0

	foldedHotSpringRow := parseData(dataByRow)
	unfoldedHotSpringRow := unFoldHotSpringRow(foldedHotSpringRow)

	for _, hotRow := range unfoldedHotSpringRow {
		rowCache := initRowCache(hotRow.Row, hotRow.DamageList)
		rowResult := recursiveP2(hotRow.Row, hotRow.DamageList, 0, 0, rowCache)
		result += rowResult
	}

	return result
}

func unfoldRecord(record string) string {
	var res strings.Builder
	for i := 0; i < len(record)*5; i++ {
		if i != 0 && i%len(record) == 0 {
			res.WriteByte('?')
		}
		res.WriteByte(record[i%len(record)])
	}

	return res.String()
}

func parseData(dataByRow []string) []HotSpringRow {
	var result []HotSpringRow

	for _, row := range dataByRow {
		parts := strings.Split(row, " ")
		damageList := util.StringToIntSlice(parts[1], ",")
		result = append(result, HotSpringRow{Row: parts[0], DamageList: damageList})
	}

	return result
}

func unFoldHotSpringRow(hotRows []HotSpringRow) []HotSpringRow {
	var result []HotSpringRow

	for _, hotRow := range hotRows {
		var newRow string
		var newDamageList []int
		for i := 0; i < 5; i++ {
			newRow += "?" + hotRow.Row
			newDamageList = append(newDamageList, hotRow.DamageList...)
		}
		result = append(result,
			HotSpringRow{
				Row:        newRow[1:],
				DamageList: newDamageList,
			},
		)
	}

	return result
}

func recursive(row string, damageList []int, index int) int {
	if index == len(row) {
		if isValid(row, damageList) {
			return 1
		}
		return 0
	}
	if row[index] == '?' {
		return recursive(row[:index]+"."+row[index+1:], damageList, index+1) +
			recursive(row[:index]+"#"+row[index+1:], damageList, index+1)
	} else {
		return recursive(row, damageList, index+1)
	}
}

func initRowCache(row string, damageList []int) [][]int {
	var cache [][]int
	for i := 0; i < len(row); i++ {
		cache = append(cache, make([]int, len(damageList)+1))
		for j := 0; j < len(damageList)+1; j++ {
			cache[i][j] = -1
		}
	}
	return cache
}

func recursiveP2(row string, damageList []int, indexRow int, indexList int, rowCache [][]int) int {
	if indexRow >= len(row) {
		// if indexList is greater than the damageList, it means that the row is valid
		if indexList >= len(damageList) {
			return 1
		}
		return 0
	}

	// after the check to prevent out of bounds
	if rowCache[indexRow][indexList] != -1 {
		return rowCache[indexRow][indexList]
	}

	tempSum := 0
	if row[indexRow] == '.' {
		// if empty, just go to the next
		tempSum = recursiveP2(row, damageList, indexRow+1, indexList, rowCache)
	} else {
		if row[indexRow] == '?' {
			// if ?, simply go ahead and try to find the sub-solutions in rowCache
			tempSum += recursiveP2(row, damageList, indexRow+1, indexList, rowCache)
		}
		if indexList < len(damageList) {
			// find the current damage
			damageCount := 0
			for k := indexRow; k < len(row); k++ {
				if damageCount > damageList[indexList] || // if damage is greater than the expected, break
					row[k] == '.' || // if empty, break
					damageCount == damageList[indexList] && row[k] == '?' { // if damage is equal to the expected, but the next is '?' (so it can be '.'), break
					break
				}
				damageCount += 1
			}

			// if the damage is equal to the expected, go ahead
			if damageCount == damageList[indexList] {
				// se l'indice e il damage trovato sono contenuti nella lunghezza della stringa, va avanti di uno
				if indexRow+damageCount < len(row) { //&& row[indexRow+damageCount] != '#'
					tempSum += recursiveP2(row, damageList, indexRow+damageCount+1, indexList+1, rowCache)
				} else {
					// altrimenti avanza solo nel gruppo
					tempSum += recursiveP2(row, damageList, indexRow+damageCount, indexList+1, rowCache)
				}
			}

		}
	}

	// salvataggio in rowCache per il x5
	rowCache[indexRow][indexList] = tempSum
	return tempSum
}

func isValid(row string, damageList []int) bool {
	var found []int

	count := 0
	for _, char := range row {
		if char == '#' {
			count++
		} else if char == '.' {
			if count != 0 {
				found = append(found, count)
				count = 0
			}
		}
	}

	// append the last count if present
	if count != 0 {
		found = append(found, count)
	}

	if len(found) != len(damageList) {
		return false
	}

	for i := 0; i < len(found); i++ {
		if found[i] != damageList[i] {
			return false
		}
	}
	return true
}
