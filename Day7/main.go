package main

import (
	"aoc23/utils"
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
	"J": 1, // Needs to be changed to 11 for Part 1
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

	utils.SeekToFileStart(file)

	partTwo(file)
}

func partOne(file *os.File) {
	hands := parseHands(file)
	sort.Slice(hands, func(i, j int) bool {
		return sortHand(&hands, i, j)
	})

	total := 0
	rank := 1
	for i := 0; i < len(hands); i++ {
		total += hands[i].bet * rank
		rank++
	}

	fmt.Println("Total Part One:", total)

}

func partTwo(file *os.File) {
	hands := parseHands(file)

	var updatedHands []*Hand
	newHandToOldHand := make(map[*Hand]*Hand, len(hands))
	for i := 0; i < len(hands); i++ {
		if strings.Contains(hands[i].hand, "J") {
			jCount := 0
			for _, cardCount := range hands[i].cardCount {
				if cardCount.card == "J" {
					jCount = cardCount.count
				}
			}
			newHand := resolveBestJokerReplacement(&hands[i], jCount)
			updatedHands = append(updatedHands, newHand)
			newHandToOldHand[newHand] = &hands[i]
		} else {
			updatedHands = append(updatedHands, &Hand{hands[i].hand, hands[i].bet, hands[i].cardCount})
		}
	}

	sort.Slice(updatedHands, func(i, j int) bool {
		return sortHandPartTwo(&updatedHands, newHandToOldHand, i, j)
	})

	total := 0
	rank := 1
	for i := 0; i < len(updatedHands); i++ {
		total += updatedHands[i].bet * rank
		rank++
	}

	fmt.Println("Total Part Two:", total)
}

func sortHand(hands *[]Hand, i int, j int) bool {
	iType := resolvePrimaryType((*hands)[i].cardCount)
	jType := resolvePrimaryType((*hands)[j].cardCount)

	if iType == jType {
		return sortCardByStrength((*hands)[i].hand, (*hands)[j].hand, 0)
	}

	return iType < jType
}

func sortHandPartTwo(hands *[]*Hand, newHandToOldHand map[*Hand]*Hand, i int, j int) bool {
	iType := resolvePrimaryType((*hands)[i].cardCount)
	jType := resolvePrimaryType((*hands)[j].cardCount)

	if iType == jType {
		first := (*hands)[i].hand
		second := (*hands)[j].hand

		oldFirst, oldFirstExists := newHandToOldHand[(*hands)[i]]
		oldSecond, oldSecondExists := newHandToOldHand[(*hands)[j]]

		if oldFirstExists {
			first = oldFirst.hand
		}
		if oldSecondExists {
			second = oldSecond.hand
		}
		return sortCardByStrength(
			first, second,
			0)
	}

	return iType < jType
}

func sortCardByStrength(handOne string, handTwo string, cardIdx int) bool {
	iStr := cardStrength[string(handOne[cardIdx])]
	jStr := cardStrength[string(handTwo[cardIdx])]

	if jStr == iStr && cardIdx+1 < len(handOne) {
		return sortCardByStrength(handOne, handTwo, cardIdx+1)
	} else {
		return iStr < jStr
	}
}

func resolvePrimaryType(cardCount []CardCount) int {
	switch {
	case isFiveOfaKind(cardCount):
		return 7
	case isFourOfaKind(cardCount):
		return 6
	case isFullHouse(cardCount):
		return 5
	case isThreeOfaKind(cardCount):
		return 4
	case isTwoPair(cardCount):
		return 3
	case isOnePair(cardCount):
		return 2
	default:
		return 1
	}
}

func isFiveOfaKind(cardCount []CardCount) bool {
	return len(cardCount) == 1
}

func isFourOfaKind(cardCount []CardCount) bool {
	return len(cardCount) == 2 && cardCount[0].count == 4
}

func isFullHouse(cardCount []CardCount) bool {
	return len(cardCount) == 2 && cardCount[0].count == 3
}

func isThreeOfaKind(cardCount []CardCount) bool {
	return len(cardCount) == 3 && cardCount[0].count == 3
}

func isTwoPair(cardCount []CardCount) bool {
	return len(cardCount) == 3 && cardCount[0].count == 2
}

func isOnePair(cardCount []CardCount) bool {
	return len(cardCount) == 4 && cardCount[0].count == 2
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

		hands = append(hands, Hand{hand, bet, charCounts})
	}
	return hands
}

func resolveBestJokerReplacement(hand *Hand, jokerOccur int) *Hand {
	maxCount := 0
	var maxString string
	var maxCardCount []CardCount
	for i := 0; i < jokerOccur; i++ {
		for card, _ := range cardStrength {
			tempString := strings.Replace(hand.hand, "J", card, i+1)
			cardCount := resolveUniqueCharCount(tempString)
			val := resolvePrimaryType(cardCount)
			if val == maxCount {
				if sortCardByStrength(maxString, tempString, 0) {
					maxCount = val
					maxString = tempString
					maxCardCount = cardCount
				}
			}
			if val > maxCount {
				maxCount = val
				maxString = tempString
				maxCardCount = cardCount
			}
		}
	}

	return &Hand{maxString, hand.bet, maxCardCount}
}
