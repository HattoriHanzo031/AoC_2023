package main

import (
	"bytes"
	"fmt"
	"slices"
	"time"
	"utils"
)

type pipeMap [][]byte

func (pm *pipeMap) printWithPath(path map[int][]node) {
	pipeMappingPath := map[byte]rune{
		'S': '┃',
		'|': '┃',
		'-': '━',
		'L': '┗',
		'J': '┛',
		'F': '┏',
		'7': '┓',
		'.': '.',
	}
	pipeMapping := map[byte]rune{
		'S': '│',
		'|': '│',
		'-': '─',
		'L': '└',
		'J': '┘',
		'F': '┌',
		'7': '┐',
		'.': '.',
	}

	for y := 0; y < 140; y++ {
		for x := 0; x < 140; x++ {
			p := pm.get(x, y)
			c := pipeMapping[p]
			if slices.Contains(path[x], node{x, y, p}) {
				c = pipeMappingPath[p]
			}
			fmt.Print(string(c))
		}
		fmt.Println("")
	}
}

func (pm *pipeMap) get(x, y int) byte {
	if y < 0 || y > len((*pm))-1 {
		return '.'
	}

	if x < 0 || x > len((*pm)[0])-1 {
		return '.'
	}
	return (*pm)[y][x]
}

type direction struct{ x, y int }

func (d direction) opposite() direction {
	d.x *= -1
	d.y *= -1
	return d
}

type node struct {
	x, y int
	t    byte
}

func pipeIteratorFn(start node, pm *pipeMap) func() node {
	var (
		up    = direction{0, -1}
		down  = direction{0, 1}
		left  = direction{-1, 0}
		right = direction{1, 0}
	)

	directions := map[byte][]direction{
		'S': {up, down, left, right},
		'|': {up, down},
		'-': {left, right},
		'F': {down, right},
		'7': {down, left},
		'J': {up, left},
		'L': {up, right},
	}

	cur := start
	prev := start
	return func() node {
		for _, dir := range directions[cur.t] {
			next := node{
				x: cur.x + dir.x,
				y: cur.y + dir.y,
				t: pm.get(cur.x+dir.x, cur.y+dir.y),
			}
			if next.t != '.' && next != prev && slices.Contains(directions[next.t], dir.opposite()) {
				prev = cur
				cur = next
				return next
			}
		}
		panic("END")
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d10/input.txt")
	defer cleanup()

	start := node{-1, -1, 'S'}
	pipes := pipeMap{}
	for scanner.Scan() {
		pipes = append(pipes, slices.Clone(scanner.Bytes()))
		if start.x == -1 { // find staring point
			start.y++
			start.x = bytes.IndexByte(pipes[start.y], 'S')
		}
	}

	fmt.Println("start:", start.x, start.y)

	next := pipeIteratorFn(start, &pipes)
	path := make(map[int][]node)
	countP1 := 0
	var cur node
	for cur.t != 'S' {
		cur = next()
		path[cur.x] = append(path[cur.x], cur)
		countP1++
	}
	fmt.Println(countP1 / 2)

	pipes.printWithPath(path)

	countP2 := 0
	for _, p := range path {
		slices.SortFunc(p, func(a, b node) int { return a.y - b.y })
		crossing := 0
		inside := false
		for i, n := range p[:len(p)-1] {
			switch n.t {
			case '-':
				crossing += 2
			case 'S', '|': // in my case S is '|' pipe
				crossing += 0
			case '7', 'L':
				crossing += -1
			case 'J', 'F':
				crossing += 1
			}

			if crossing == 2 || crossing == -2 {
				crossing = 0
				inside = !inside
			}
			if inside {
				countP2 += p[i+1].y - n.y - 1
			}
		}
	}
	fmt.Println(countP2)
}
