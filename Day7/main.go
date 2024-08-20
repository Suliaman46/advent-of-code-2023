package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cardStrength = map[string]int{
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

type Hand struct {
	hand      string
	bet       int
	cardCount []CardCount
	score     int
}

type CardCount struct {
	card  string
	count int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Could not read file!")
		return
	}
	defer file.Close()

	partOne(file)
}

func partOne(file *os.File) {
	hands := parseHands(file)

	for i := 0; i < len(hands); i++ {
		score := resolvePrimaryType(hands[i])
		hands[i].score = score
	}

	var handByTypes [][]Hand

	for i := 1; i <= 7; i++ {
		temp := make([]Hand, 0)
		appendToAppSlice(hands, &temp, i)
		handByTypes = append(handByTypes, temp)
	}

	for i := 0; i < len(handByTypes); i++ {
		if len(handByTypes[i]) > 1 {
			performSecondarySort(&handByTypes[i])
		}
	}

	total := 0
	rank := 1
	for i := 0; i < len(handByTypes); i++ {
		for j := 0; j < len(handByTypes[i]); j++ {
			total += handByTypes[i][j].bet * rank
			rank++
		}
	}

	fmt.Println("Total Part One:", total)

}

func performSecondarySort(hands *[]Hand) {
	sort.Slice(*hands, func(i, j int) bool {
		return sortCardByStrength(hands, 0, i, j)
	})
}

func sortCardByStrength(hands *[]Hand, cardIdx int, i int, j int) bool {
	iStr := cardStrength[string((*hands)[i].hand[cardIdx])]
	jStr := cardStrength[string((*hands)[j].hand[cardIdx])]

	if jStr == iStr && cardIdx+1 < len(*hands) {
		return sortCardByStrength(hands, cardIdx+1, i, j)
	} else {
		return iStr < jStr
	}
}

func appendToAppSlice(hands []Hand, arr *[]Hand, score int) {
	for i := 0; i < len(hands); i++ {
		if (hands)[i].score == score {
			*arr = append(*arr, (hands)[i])
		}
	}
}

func resolvePrimaryType(hand Hand) int {
	switch {
	case isFiveOfaKind(hand):
		return 7
	case isFourOfaKind(hand):
		return 6
	case isFullHouse(hand):
		return 5
	case isThreeOfaKind(hand):
		return 4
	case isTwoPair(hand):
		return 3
	case isOnePair(hand):
		return 2
	default:
		return 1
	}
}

func isFiveOfaKind(hand Hand) bool {
	return len(hand.cardCount) == 1
}

func isFourOfaKind(hand Hand) bool {
	return len(hand.cardCount) == 2 && hand.cardCount[0].count == 4
}

func isFullHouse(hand Hand) bool {
	return len(hand.cardCount) == 2 && hand.cardCount[0].count == 3
}

func isThreeOfaKind(hand Hand) bool {
	return len(hand.cardCount) == 3 && hand.cardCount[0].count == 3
}

func isTwoPair(hand Hand) bool {
	return len(hand.cardCount) == 3 && hand.cardCount[0].count == 2
}

func isOnePair(hand Hand) bool {
	return len(hand.cardCount) == 4 && hand.cardCount[0].count == 2
}

func resolveUniqueCharCount(hand string) []CardCount {
	unsorted := map[string]int{}
	for _, char := range hand {
		val, present := unsorted[string(char)]
		if present {
			unsorted[string(char)] = val + 1
			continue
		}

		unsorted[string(char)] = 1
	}

	//fmt.Println("Map", unsorted)

	keys := make([]string, 0, len(unsorted))
	for key := range unsorted {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return unsorted[keys[i]] > unsorted[keys[j]]
	})

	cardCounts := make([]CardCount, 0, len(keys))

	for _, key := range keys {
		cardCounts = append(cardCounts, CardCount{key, unsorted[key]})
	}

	return cardCounts
}

func parseHands(file *os.File) []Hand {
	scanner := bufio.NewScanner(file)
	var hands []Hand

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		hand := line[0]
		bet, _ := strconv.Atoi(line[1])

		charCounts := resolveUniqueCharCount(hand)

		hands = append(hands, Hand{hand, bet, charCounts, 0})
	}
	return hands
}
