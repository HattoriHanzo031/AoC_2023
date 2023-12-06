package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

func numWins(t, d int) int {
	sqr := math.Sqrt(float64((t * t) - 4*d))
	return int((float64(t)+sqr)/2) - int((float64(t)-sqr)/2)
}

func main() {
	scanner, cleanup := utils.FileScaner("d6/input.txt")
	defer cleanup()

	scanner.Scan()
	times := utils.ToInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))
	combinedTime := utils.Must(strconv.Atoi(strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", "")))
	scanner.Scan()
	distances := utils.ToInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))
	combinedDistance := utils.Must(strconv.Atoi(strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", "")))

	wins := 1
	for i, t := range times {
		wins *= numWins(t, distances[i])
	}
	fmt.Println("P1:", wins)

	fmt.Println("P2", numWins(combinedTime, combinedDistance))
}
