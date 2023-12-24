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

func print(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction, nodes map[coord]node) {
	c := coord{}
	for c.y = 0; c.y < lenY; c.y++ {
		for c.x = 0; c.x < lenX; c.x++ {
			node, nodeOk := nodes[c]
			switch _, ok := slopes[c]; true {
			case stones[c]:
				fmt.Print("░")
			case nodeOk:
				fmt.Print(len(node.connected))
			case ok:
				fmt.Print(">")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

type node struct {
	connected map[coord]int
}

func findJunction(prev, cur coord, mover func(cur coord) []coord) (coord, int) {
	count := 0
	for {
		count++
		next := mover(cur)
		if len(next) != 2 {
			return cur, count
		}
		if next[0].x == prev.x && next[0].y == prev.y {
			prev = cur
			cur = next[1]
		} else {
			prev = cur
			cur = next[0]
		}
	}
}

func makeNodes(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction) map[coord]node {
	nodes := make(map[coord]node)
	mover := moverLongestFn(lenX-1, lenY-1, stones, slopes)
	c := coord{}
	for c.y = 0; c.y < lenY; c.y++ {
		for c.x = 0; c.x < lenX; c.x++ {
			if stones[c] {
				continue
			}
			if neighbors := mover(c); len(neighbors) != 2 {
				nodes[c] = node{connected: make(map[coord]int, len(neighbors))}
				for _, ngbh := range neighbors {
					next, count := findJunction(c, ngbh, mover)
					nodes[c].connected[next] = count
				}
			}
		}
	}
	return nodes
}

func moverLongestFn(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction) func(cur coord) []coord {
	return func(cur coord) []coord {
		if d, ok := slopes[cur]; ok {
			return []coord{{cur.x + d.x, cur.y + d.y}}
		}
		out := []coord{}
		if cur.x != lenX {
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
		if cur.y != lenY {
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
	nodes := makeNodes(lenX, lenY, stones, slopes)
	print(lenX, lenY, stones, slopes, nodes)
	findAllPaths(0, start, exit, nodes, make(map[coord]bool))
	fmt.Println("LONGEST P1:", longestPath)

	slopes = map[coord]direction{} // clear all slopes
	slopes[start] = down
	slopes[exit] = up
	longestPath = 0
	nodes = makeNodes(lenX, lenY, stones, slopes)
	//print(lenX, lenY, stones, slopes, nodes)
	findAllPaths(0, start, exit, nodes, make(map[coord]bool))
	fmt.Println("LONGEST P2:", longestPath)
}

var longestPath int

func findAllPaths(count int, cur, exit coord, nodes map[coord]node, visited map[coord]bool) {
	if cur.x == exit.x && cur.y == exit.y {
		//fmt.Println("END", count)
		longestPath = max(longestPath, count)
	}
	if visited[cur] {
		return
	}
	visited[cur] = true

	for node, distance := range nodes[cur].connected {
		findAllPaths(count+distance, node, exit, nodes, maps.Clone(visited))
	}
}

// func moverFn(maxX, maxY int, stones map[coord]bool, slopes map[coord]direction) func(cur coord) []coord {
// 	visited := map[coord]bool{}
// 	return func(cur coord) []coord {
// 		if visited[cur] {
// 			return []coord{}
// 		}
// 		visited[cur] = true

// 		if d, ok := slopes[cur]; ok {
// 			return []coord{{cur.x + d.x, cur.y + d.y}}
// 		}
// 		out := []coord{}
// 		if cur.x != maxX {
// 			c := coord{cur.x + right.x, cur.y + right.y}
// 			if !stones[c] {
// 				out = append(out, c)
// 			}
// 		}
// 		if cur.x != 0 {
// 			c := coord{cur.x + left.x, cur.y + left.y}
// 			if !stones[c] {
// 				out = append(out, c)
// 			}
// 		}
// 		if cur.y != 0 {
// 			c := coord{cur.x + up.x, cur.y + up.y}
// 			if !stones[c] {
// 				out = append(out, c)
// 			}
// 		}
// 		if cur.y != maxY {
// 			c := coord{cur.x + down.x, cur.y + down.y}
// 			if !stones[c] {
// 				out = append(out, c)
// 			}
// 		}
// 		return out
// 	}
// }

// func shortest(lenX, lenY int, stones map[coord]bool, slopes map[coord]direction, start, exit coord) {
// 	mover := moverFn(lenX, lenY, stones, slopes)
// 	q := map[coord]struct{}{start: {}}
// 	step := 0
// 	for len(q) > 0 {
// 		newQ := make(map[coord]struct{})
// 		for c := range q {
// 			for _, coord := range mover(c) {
// 				newQ[coord] = struct{}{}
// 			}
// 			if c.x == exit.x && c.y == exit.y {
// 				fmt.Println("EXIT:", step)
// 				return
// 			}
// 		}

// 		q = newQ
// 		step++
// 	}
// 	fmt.Println(step)
// }
