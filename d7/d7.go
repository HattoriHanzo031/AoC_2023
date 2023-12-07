package main

import (
	"fmt"
	"math"
	"slices"
	"utils"
)

var strengths = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

type hand struct {
	cards    []byte
	bid      int
	t        int
	strength int
}

func getType(cards []byte) (t int) {
	acc := map[byte]int{}
	for _, card := range cards {
		acc[card] = acc[card] + 1
	}

	for _, v := range acc {
		t += int(math.Pow10(v))
	}
	return t
}

func getTypeP2(cards []byte) (t int) {
	acc := map[byte]int{}
	numJ := 0
	mostCards := 0
	mostOftenCard := cards[0]
	for _, card := range cards {
		if card == 'J' {
			numJ++
			continue
		}
		acc[card] = acc[card] + 1
		if mostCards < acc[card] {
			mostCards = acc[card]
			mostOftenCard = card
		}
	}

	acc[mostOftenCard] += numJ

	for _, v := range acc {
		t += int(math.Pow10(v))
	}
	return t
}

func getStrength(cards []byte) (strength int) {
	for _, card := range cards {
		strength *= 100
		strength += strengths[card]
	}

	return strength
}

func getStrengthP2(cards []byte) (strength int) {
	for _, card := range cards {
		strength *= 100
		if card == 'J' {
			continue
		}
		strength += strengths[card]
	}

	return strength
}

func newHand(cards []byte, bid int) (h hand) {
	h.bid = bid
	h.cards = cards
	h.t = getTypeP2(cards)
	h.strength = getStrengthP2(cards)
	return h
}

func main() {
	scanner, cleanup := utils.FileScaner("d7/input.txt")
	defer cleanup()

	hands := []hand{}
	for scanner.Scan() {
		cards := []byte{}
		var bid int
		fmt.Sscanf(scanner.Text(), "%s %d", &cards, &bid)
		hands = append(hands, newHand(cards, bid))
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.t == b.t {
			return a.strength - b.strength
		}
		return a.t - b.t
	})

	totalWinnings := 0
	rank := 1
	for _, h := range hands {
		totalWinnings += h.bid * rank
		rank++
	}
	fmt.Println("totalWinnings:", totalWinnings)
}
