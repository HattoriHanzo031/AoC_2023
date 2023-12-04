package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

type card struct {
	index    int
	score    int
	totalWon *int
}

func parseCard(index int, s string) card {
	ss := strings.Split(strings.Split(s, ":")[1], "|")

	winning := make(map[int]bool)
	for _, num := range strings.Fields(ss[0]) {
		winning[utils.Must(strconv.Atoi(num))] = true
	}

	c := card{index: index}
	for _, num := range strings.Fields(ss[1]) {
		if winning[utils.Must(strconv.Atoi(num))] {
			c.score++
		}
	}
	return c
}

func (c card) cardsWon(cards []card) int {
	if c.totalWon != nil {
		return *c.totalWon
	}

	c.totalWon = new(int)
	*c.totalWon = c.score
	end := c.index + *c.totalWon + 1
	for i := c.index + 1; i < end; i++ {
		*c.totalWon += cards[i].cardsWon(cards)
	}

	return *c.totalWon
}

func main() {
	scanner, cleanup := utils.FileScaner("d4/input.txt")
	defer cleanup()

	total := 0
	cards := []card{}
	index := 0
	for scanner.Scan() {
		cards = append(cards, parseCard(index, scanner.Text()))
		index++
	}

	for _, card := range cards {
		total += int(math.Pow(2, float64(card.score-1)))
	}
	fmt.Println("Part 1", total)

	total = len(cards)
	for _, card := range cards {
		total += card.cardsWon(cards)
	}
	fmt.Println("Part 2", total)
}
