package main

import (
	"bytes"
	"fmt"
	"slices"
	"time"
	"utils"
)

type coord struct{ x, y int }

func (c *coord) distance(other coord) int {
	return utils.Abs(c.x-other.x) + utils.Abs(c.y-other.y)
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d11/input.txt")
	defer cleanup()

	expandedRows := map[int]int{}
	universe := [][]byte{}
	for row := 0; scanner.Scan(); row++ {
		universe = append(universe, slices.Clone(scanner.Bytes()))
		if bytes.Count(scanner.Bytes(), []byte{'#'}) == 0 {
			expandedRows[row] = 1
		}
	}

	expandedColumns := map[int]int{}
	for col := 0; col < len(universe[0]); col++ {
		count := 0
		for row := 0; row < len(universe); row++ {
			if universe[row][col] == '#' {
				count++
			}
		}
		if count == 0 {
			expandedColumns[col] = 1
		}
	}

	solution := func(expansionRate int) int {
		expansionRate-- // to compensate that we are always adding 1 in calculations
		galaxies := []coord{}
		yExpanded := 0
		for y, row := range universe {
			xExpanded := 0
			for x, c := range row {
				if c == '#' {
					galaxies = append(galaxies, coord{xExpanded, yExpanded})
				}
				xExpanded += expandedColumns[x]*expansionRate + 1
			}
			yExpanded += expandedRows[y]*expansionRate + 1
		}

		total := 0
		for g1 := range galaxies {
			for g2 := g1 + 1; g2 < len(galaxies); g2++ {
				total += galaxies[g1].distance(galaxies[g2])
			}
		}
		return total
	}

	fmt.Println("P1:", solution(2))
	fmt.Println("P2:", solution(1000_000))
}
