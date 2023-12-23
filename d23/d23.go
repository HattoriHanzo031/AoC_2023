package main

import (
	"fmt"
	"maps"
	"time"
	"utils"
)

type coord struct{ x, y int }

type direction struct{ x, y int }

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

func print(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction) {
	c := coord{}
	for c.y = 0; c.y < lenY; c.y++ {
		for c.x = 0; c.x < lenX; c.x++ {
			switch _, ok := slopes[c]; true {
			case stones[c]:
				fmt.Print("#")
			case ok:
				fmt.Print(">")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func moverFn(maxX, maxY int, stones map[coord]bool, slopes map[coord]direction) func(cur coord) []coord {
	visited := map[coord]bool{}
	return func(cur coord) []coord {
		if visited[cur] {
			return []coord{}
		}
		visited[cur] = true

		if d, ok := slopes[cur]; ok {
			return []coord{{cur.x + d.x, cur.y + d.y}}
		}
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

func moverLongestFn(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction) func(cur coord) []coord {
	return func(cur coord) []coord {
		if d, ok := slopes[cur]; ok {
			return []coord{{cur.x + d.x, cur.y + d.y}}
		}
		out := []coord{}
		if cur.x != lenX-1 {
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
		if cur.y != lenY-1 {
			c := coord{cur.x + down.x, cur.y + down.y}
			if !stones[c] {
				out = append(out, c)
			}
		}
		return out
	}
}

func moverLongest2Fn(stones map[coord]bool, slopes map[coord]direction) func(cur coord) []coord {
	return func(cur coord) []coord {
		if d, ok := slopes[cur]; ok {
			return []coord{{cur.x + d.x, cur.y + d.y}}
		}
		out := []coord{}

		c := coord{cur.x + right.x, cur.y + right.y}
		if !stones[c] {
			out = append(out, c)
		}

		c = coord{cur.x + left.x, cur.y + left.y}
		if !stones[c] {
			out = append(out, c)
		}

		c = coord{cur.x + up.x, cur.y + up.y}
		if !stones[c] {
			out = append(out, c)
		}

		c = coord{cur.x + down.x, cur.y + down.y}
		if !stones[c] {
			out = append(out, c)
		}

		return out
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d23/input.txt")
	defer cleanup()

	stones := map[coord]bool{}
	slopes := map[coord]direction{}
	lenX, lenY := 0, 0
	y := 0
	start := coord{}
	exit := coord{}
	for scanner.Scan() {
		lenX = len(scanner.Bytes())

		for x, c := range scanner.Bytes() {
			switch c {
			case '>':
				slopes[coord{x, y}] = right
			case 'v':
				slopes[coord{x, y}] = down
			case '#':
				stones[coord{x, y}] = true
			case '.':
				if start.x == 0 && start.y == 0 {
					start = coord{x, y}
					slopes[coord{x, y}] = down
				}
				exit = coord{x, y}
			}
		}
		y++
	}
	lenY = y

	slopes[start] = down
	slopes[exit] = up

	fmt.Println(lenX, lenY, start)
	print(lenX, lenY, stones, slopes)

	shortest(lenX, lenY, stones, slopes, start, exit)
	longest(0, stones, slopes, start, exit, make(map[coord]bool), moverLongestFn(lenX, lenY, stones, slopes))
	fmt.Println("LONGEST:", longestPath)
}

var longestPath int

func longest(steps int, stones map[coord]bool, slopes map[coord]direction, cur, exit coord, visited map[coord]bool, mover func(cur coord) []coord) {
	if cur.x == exit.x && cur.y == exit.y {
		fmt.Println("EXIT", steps, len(visited))
		// for k := range visited {
		// 	fmt.Println(k)
		// }
		longestPath = max(longestPath, steps)
		return
	}

	if visited[cur] {
		return
	}
	visited[cur] = true

	for _, next := range mover(cur) {
		longest(steps+1, stones, slopes, next, exit, maps.Clone(visited), mover)
	}
}

func shortest(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction, start, exit coord) {
	mover := moverFn(lenX, lenY, stones, slopes)
	q := map[coord]struct{}{start: {}}
	step := 0
	for len(q) > 0 {
		newQ := make(map[coord]struct{})
		for c := range q {
			for _, coord := range mover(c) {
				newQ[coord] = struct{}{}
			}
			if c.x == exit.x && c.y == exit.y {
				fmt.Println("EXIT:", step)
				return
			}
		}

		q = newQ
		step++
	}
	fmt.Println(step)
}
