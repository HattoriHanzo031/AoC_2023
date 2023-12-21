package main

import (
	"fmt"
	"slices"
	"time"
	"utils"
)

type coord struct{ x, y int }

type direction struct{ x, y int }

func moverFn(maxX, maxY int, stones map[coord]bool) func(cur coord) []coord {
	var (
		up    = direction{0, -1}
		down  = direction{0, 1}
		left  = direction{-1, 0}
		right = direction{1, 0}
	)

	return func(cur coord) []coord {
		out := []coord{}
		if cur.x != maxX {
			c := coord{cur.x + right.x, cur.y + right.y}
			if !stones[c] {
				out = append(out, c)
			}
		}
		if cur.x != 0 {
			c := coord{cur.x + left.x, cur.y + left.y}
			if !stones[c] {
				out = append(out, c)
			}
		}
		if cur.y != 0 {
			c := coord{cur.x + up.x, cur.y + up.y}
			if !stones[c] {
				out = append(out, c)
			}
		}
		if cur.y != maxY {
			c := coord{cur.x + down.x, cur.y + down.y}
			if !stones[c] {
				out = append(out, c)
			}
		}
		return out
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d21/input.txt")
	defer cleanup()

	stones := map[coord]bool{}
	maxX, maxY := 0, 0
	y := 0
	start := coord{}
	for scanner.Scan() {
		maxX = len(scanner.Bytes()) - 1

		for x, c := range scanner.Bytes() {
			switch c {
			case 'S':
				start = coord{x, y}
			case '#':
				stones[coord{x, y}] = true
			}
		}
		y++
	}
	maxY = y - 1
	fmt.Println(maxX, maxY, start)
	//print(maxX, maxY, stones, nil)

	mover := moverFn(maxX, maxY, stones)
	q := []coord{start}
	for step := 0; step < 64; step++ {
		newQ := []coord{}
		for _, c := range q {
			newQ = append(newQ, mover(c)...)
		}
		slices.SortFunc(newQ, func(a, b coord) int {
			if difY := a.y - b.y; difY != 0 {
				return difY
			} else {
				return a.x - b.x
			}
		})
		q = slices.Compact(newQ)
		//print(maxX, maxY, stones, q)
		//fmt.Println("")
	}
	fmt.Println(len(q))
}

func print(maxX, maxY int, stones map[coord]bool, destinations []coord) {
	ds := make(map[coord]bool, len(destinations))
	for _, d := range destinations {
		ds[d] = true
	}

	c := coord{}
	for c.y = 0; c.y < maxY+1; c.y++ {
		for c.x = 0; c.x < maxX+1; c.x++ {
			switch {
			case stones[c]:
				fmt.Print("#")
			case ds[c]:
				fmt.Print("O")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
