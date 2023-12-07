package main

import (
	"fmt"
	"strings"

	"advent2023/util"
)

func main() {
	dataByRow := util.GetDataByRow("day07")

	// test data
	//dataByRow = []string{
	//	"32T3K 765",
	//	"T55J5 684",
	//	"KK677 28",
	//	"KTJJT 220",
	//	"QQQJA 483",
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

	hands, bids := parseHands(dataByRow)

	bubbleSort(hands, bids, compareHands)

	for i := 0; i < len(bids); i++ {
		result += bids[i] * (i + 1)
	}

	return result
}

func part02(dataByRow []string) int {

	result := 0

	hands, bids := parseHands(dataByRow)

	bubbleSort(hands, bids, compareHandsP2)

	for i := 0; i < len(bids); i++ {
		result += bids[i] * (i + 1)
	}

	return result
}

func parseHands(dataByRow []string) (hands []string, bids []int) {

	hands = make([]string, 0, len(dataByRow))
	bids = make([]int, 0, len(dataByRow))

	for _, row := range dataByRow {
		parts := strings.Split(row, " ")
		hands = append(hands, parts[0])
		bids = append(bids, util.ToInt(parts[1]))
	}

	return
}

// map[card]numberOfCards
// map[groupCount]numberOfGroups
func groupCards(hand string) (map[string]int, map[int]int) {
	groups := make(map[string]int)
	for _, card := range hand {
		groups[string(card)]++
	}

	final := make(map[int]int)
	for _, count := range groups {
		final[count]++
	}

	return groups, final
}

func bubbleSort(hands []string, bids []int, compareFunc func(string, string) int) {
	for i := 0; i < len(hands); i++ {
		for j := 0; j < len(hands)-1; j++ {
			if compareFunc(hands[j], hands[j+1]) == 1 {
				hands[j], hands[j+1] = hands[j+1], hands[j]
				bids[j], bids[j+1] = bids[j+1], bids[j]
			}
		}
	}
}

// -1: hand1 < hand2
// 0: hand1 == hand2
// 1: hand1 > hand2
func compareHands(hand1 string, hand2 string) (result int) {
	_, hand1Group := groupCards(hand1)
	_, hand2Group := groupCards(hand2)

	hand1Value := handValue(hand1Group)
	hand2Value := handValue(hand2Group)

	if hand1Value > hand2Value {
		return 1
	}
	if hand1Value < hand2Value {
		return -1
	}

	for i := 0; i < len(hand1); i++ {
		hand1Card := string(hand1[i])
		hand2Card := string(hand2[i])
		if cardValue(hand1Card) > cardValue(hand2Card) {
			return 1
		}
		if cardValue(hand1Card) < cardValue(hand2Card) {
			return -1
		}
	}
	fmt.Printf("pareggio = hand1: %s, hand2: %s\n", hand1, hand2)
	return 0
}

// -1: hand1 < hand2
// 0: hand1 == hand2
// 1: hand1 > hand2
func compareHandsP2(hand1 string, hand2 string) (result int) {
	hand1Group, hand1Couples := groupCards(hand1)
	hand2Group, hand2Couples := groupCards(hand2)

	hand1Value := handValueP2(hand1Couples, hand1Group["J"])
	hand2Value := handValueP2(hand2Couples, hand2Group["J"])

	if hand1Value > hand2Value {
		return 1
	}
	if hand1Value < hand2Value {
		return -1
	}

	for i := 0; i < len(hand1); i++ {
		hand1Card := string(hand1[i])
		hand2Card := string(hand2[i])
		if cardValueP2(hand1Card) > cardValueP2(hand2Card) {
			return 1
		}
		if cardValueP2(hand1Card) < cardValueP2(hand2Card) {
			return -1
		}
	}
	fmt.Printf("pareggio = hand1: %s, hand2: %s\n", hand1, hand2)
	return 0
}

func cardValue(card string) int {
	points := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}
	return points[card]
}

func cardValueP2(card string) int {
	points := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}
	return points[card]
}

func handValue(groupedHand map[int]int) int {
	if _, ok := groupedHand[5]; ok {
		return 7
	}
	if _, ok := groupedHand[4]; ok {
		return 6
	}
	if groupedHand[3] == 1 && groupedHand[2] == 1 {
		return 5
	}
	if _, ok := groupedHand[3]; ok {
		return 4
	}
	if groupedHand[2] == 2 {
		return 3
	}
	if groupedHand[2] == 1 {
		return 2
	}
	if groupedHand[1] == 5 {
		return 1
	}
	return 0
}

func handValueP2(groupedHand map[int]int, jokerNumber int) int {
	if groupedHand[5] == 1 ||
		(groupedHand[4] == 1 && jokerNumber == 1) ||
		(groupedHand[3] == 1 && jokerNumber == 2) ||
		(groupedHand[2] == 1 && jokerNumber == 3) ||
		(groupedHand[1] == 1 && jokerNumber == 4) ||
		(jokerNumber == 5) {
		return 7
	}

	if groupedHand[4] == 1 ||
		(groupedHand[3] == 1 && jokerNumber == 1) ||
		(groupedHand[2] >= 1 && jokerNumber == 2) ||
		(groupedHand[1] >= 1 && jokerNumber == 3) {
		return 6
	}

	if (groupedHand[3] == 1 && groupedHand[2] == 1) ||
		(groupedHand[2] >= 2 && jokerNumber == 1) {
		return 5
	}

	if groupedHand[3] == 1 ||
		(groupedHand[2] >= 1 && jokerNumber == 1) ||
		(groupedHand[1] >= 1 && jokerNumber == 2) {
		return 4
	}

	if groupedHand[2] == 2 ||
		(groupedHand[2] >= 1 && jokerNumber == 1) {
		return 3
	}
	if groupedHand[2] == 1 {
		return 2
	}
	if groupedHand[1] == 5 {
		return 1
	}
	return 0
}
