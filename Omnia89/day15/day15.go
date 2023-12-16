package main

import (
	"fmt"
	"sort"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day15")

	// test data
	//dataByRow = []string{
	//	"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
	//}

	// 01 part
	part01Result := part01(dataByRow)
	fmt.Printf("First part: %d\n", part01Result)

	// 02 part
	part02Result := part02(dataByRow)
	fmt.Printf("Second part: %d\n", part02Result)
}

func part01(dataByRow []string) int {
	var result int = 0

	values := strings.Split(dataByRow[0], ",")

	for _, v := range values {
		result += hash(v)
	}

	return result
}

func part02(dataByRow []string) int {
	result := 0

	boxes := make([]box, 256)
	for i := 0; i < 256; i++ {
		boxes[i].values = make(map[string]boxedValue)
	}

	values := strings.Split(dataByRow[0], ",")

	for _, v := range values {
		hashValue, label, lens := hashAndBox(v)
		if lens == -1 {
			boxes[hashValue].remove(label)
		} else {
			boxes[hashValue].add(label, lens)
		}
	}

	for k, box := range boxes {
		boxValues := box.getOrderedValues()
		for i, v := range boxValues {
			result += v * (i + 1) * (k + 1)
		}
	}

	return result
}

type boxedValue struct {
	label    string
	value    int
	position int
}

type box struct {
	values       map[string]boxedValue
	lastPosition int
}

func (b *box) add(label string, value int) {
	var position int
	if _, ok := b.values[label]; ok {
		position = b.values[label].position
	} else {
		position = b.lastPosition + 1
		b.lastPosition = position
	}

	b.values[label] = boxedValue{
		label:    label,
		value:    value,
		position: position,
	}
}

func (b *box) remove(label string) {
	delete(b.values, label)
}

func (b *box) getOrderedValues() []int {
	var result []int

	var boxes []boxedValue
	for _, v := range b.values {
		boxes = append(boxes, v)
	}

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].position < boxes[j].position
	})

	for _, v := range boxes {
		result = append(result, v.value)
	}

	return result
}

func hash(s string) int {
	value := 0

	for _, c := range s {
		value += int(c)
		value *= 17
		value = value % 256
	}

	return value
}

func hashAndBox(s string) (hashValue int, label string, lens int) {
	value := 0
	label = ""

	for i, c := range s {
		if c == '-' {
			lens = -1
			break
		}
		if c == '=' {
			lens = util.ToInt(s[i+1:])
			break
		}
		value += int(c)
		value *= 17
		value = value % 256
		label += string(c)
	}

	return value, label, lens
}
