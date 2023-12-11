package main

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"time"
	"utils"
)

type coord struct{ x, y int }

func (c *coord) distance(other coord) int {
	return int(math.Abs(float64(c.x-other.x)) + math.Abs(float64(c.y-other.y)))
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d11/input.txt")
	defer cleanup()

	emptyRows := map[int]int{}
	universe := [][]byte{}
	for row := 0; scanner.Scan(); row++ {
		universe = append(universe, slices.Clone(scanner.Bytes()))
		if bytes.Count(scanner.Bytes(), []byte{'#'}) == 0 {
			emptyRows[row] = 1
		}
	}

	emptyColumns := map[int]int{}
	for col := 0; col < len(universe[0]); col++ {
		count := 0
		for row := 0; row < len(universe); row++ {
			if universe[row][col] == '#' {
				count++
			}
		}
		if count == 0 {
			emptyColumns[col] = 1
		}
	}

	galaxies := map[coord]bool{}
	yExpanded := 0
	for y, row := range universe {
		xExpanded := 0
		for x, c := range row {
			if c == '#' {
				galaxies[coord{xExpanded, yExpanded}] = true
			}
			xExpanded += emptyColumns[x] + 1
		}
		yExpanded += emptyRows[y] + 1
	}

	fmt.Println(galaxies)

	total := 0
	for galaxy1 := range galaxies {
		for galaxy2 := range galaxies {
			total += galaxy1.distance(galaxy2)
		}
	}
	fmt.Println(total / 2)
}
