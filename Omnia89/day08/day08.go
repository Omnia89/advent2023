package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"advent2023/util"
)

type Directions struct {
	index         int
	directionList []string
}

func (d *Directions) getNextDirection() string {
	if d.index >= len(d.directionList) {
		d.index = 0
	}
	direction := d.directionList[d.index]
	d.index++
	return direction
}

type Node struct {
	leftIndex  string
	rightIndex string
}

func (n *Node) getNextKey(direction string) string {
	if direction == "L" {
		return n.leftIndex
	}
	return n.rightIndex
}

func main() {
	dataByRow := util.GetDataByRow("day08")

	// test data
	//dataByRow = []string{
	//"RL",
	//"",
	//"AAA = (BBB, CCC)",
	//"BBB = (DDD, EEE)",
	//"CCC = (ZZZ, GGG)",
	//"DDD = (DDD, DDD)",
	//"EEE = (EEE, EEE)",
	//"GGG = (GGG, GGG)",
	//"ZZZ = (ZZZ, ZZZ)",
	//"LLR",
	//"",
	//"AAA = (BBB, BBB)",
	//"BBB = (AAA, ZZZ)",
	//"ZZZ = (ZZZ, ZZZ)",
	//	"LR",
	//	"",
	//	"11A = (11B, XXX)",
	//	"11B = (XXX, 11Z)",
	//	"11Z = (11B, XXX)",
	//	"22A = (22B, XXX)",
	//	"22B = (22C, 22C)",
	//	"22C = (22Z, 22Z)",
	//	"22Z = (22B, 22B)",
	//	"XXX = (XXX, XXX)",
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

	result := 0

	directions, nodes := parseInput(dataByRow)

	currentKey := "AAA"
	for {
		node := nodes[currentKey]
		direction := directions.getNextDirection()

		currentKey = node.getNextKey(direction)
		result++
		if currentKey == "ZZZ" {
			break
		}
	}

	return result
}

func part02(dataByRow []string) int {

	result := 1

	directions, nodes := parseInput(dataByRow)

	currentKeys := []string{}
	for key, _ := range nodes {
		if strings.HasSuffix(key, "A") {
			currentKeys = append(currentKeys, key)
		}
	}
	fmt.Printf("currentKeys: %v\n", currentKeys)

	steps := make([]int, len(currentKeys))

keysLoop:
	for i, currentKey := range currentKeys {
		for {
			node := nodes[currentKey]
			direction := directions.getNextDirection()

			currentKey = node.getNextKey(direction)
			steps[i]++
			if strings.HasSuffix(currentKey, "Z") {
				continue keysLoop
			}
		}
	}

	// retrieve the prime factors of each step, and than found the LCM of all the factors (minimo comune multiplo)
	var factors []int
	for _, step := range steps {
		primeNumberFactors := primeNumbers(step)
		for _, primeNumberFactor := range primeNumberFactors {
			if !slices.Contains(factors, primeNumberFactor) {
				factors = append(factors, primeNumberFactor)
			}
		}
	}

	for _, factor := range factors {
		result *= factor
	}

	return result
}

func parseInput(dataByRow []string) (directions Directions, nodes map[string]Node) {

	rawDirections := strings.Split(dataByRow[0], "")
	directions = Directions{
		index:         0,
		directionList: rawDirections,
	}
	nodes = make(map[string]Node)

	for i, node := range dataByRow {
		if i == 0 || node == "" {
			continue
		}
		reg := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)
		match := reg.FindStringSubmatch(node)
		keyNode := match[1]
		leftIndex := match[2]
		rightIndex := match[3]

		nodes[keyNode] = Node{
			leftIndex:  leftIndex,
			rightIndex: rightIndex,
		}
	}
	return
}

func primeNumbers(n int) []int {
	var result []int
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
			n = n / i
			i = 1
		}
	}
	return result
}
