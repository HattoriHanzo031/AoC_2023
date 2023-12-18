package main

import (
	"fmt"
	"time"
	"utils"
)

type coord struct{ x, y int }

type direction coord

func (d direction) cw() direction {
	return direction{-d.y, d.x}
}

func (d direction) ccw() direction {
	return direction{d.y, -d.x}
}

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

type state struct {
	dir         direction
	numStraight int
}

type block struct {
	val        int
	prevStates map[state]int
}

func (b *block) interactP1(s state, heatLoss int) []state {
	if prevHl, ok := b.prevStates[s]; ok && heatLoss >= prevHl {
		return nil
	}
	b.prevStates[s] = heatLoss

	out := []state{{s.dir.cw(), 0}, {s.dir.ccw(), 0}}
	if s.numStraight < 2 {
		out = append(out, state{s.dir, s.numStraight + 1})
	}

	return out
}

func (b *block) interact(s state, heatLoss int) []state {
	if prevHl, ok := b.prevStates[s]; ok && heatLoss >= prevHl {
		return nil
	}
	b.prevStates[s] = heatLoss

	if s.numStraight < 4 {
		return []state{{s.dir, s.numStraight + 1}}
	}

	out := []state{{s.dir.cw(), 1}, {s.dir.ccw(), 1}}
	if s.numStraight < 10 {
		out = append(out, state{s.dir, s.numStraight + 1})
	}

	return out
}

var xMax, yMax int

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d17/input.txt")
	defer cleanup()

	input := make(map[coord]*block)
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Bytes() {
			input[coord{x, y}] = &block{int(c - '0'), make(map[state]int)}
		}
		y++
		xMax = len(scanner.Bytes())
	}
	yMax = y
	fmt.Println(xMax, yMax)

	p1(input, coord{1, 0}, state{right, 1}, 0)
	p1(input, coord{0, 1}, state{down, 1}, 0)
	fmt.Println(minHeatLoss)
}

var minHeatLoss = 2000

func p1(blocks map[coord]*block, cur coord, s state, heatLoss int) {
	var b *block
	var ok bool
	if b, ok = blocks[cur]; !ok {
		return
	}
	heatLoss += b.val

	if heatLoss >= minHeatLoss-(xMax-1-cur.x+yMax-1-cur.y) {
		return
	}

	if cur.x == xMax-1 && cur.y == yMax-1 {
		if s.numStraight < 4 {
			return
		}
		fmt.Println("Path found:", heatLoss)
		minHeatLoss = min(minHeatLoss, heatLoss)
		return
	}

	for _, next := range blocks[cur].interact(s, heatLoss) {
		p1(blocks, coord{next.dir.x + cur.x, next.dir.y + cur.y}, next, heatLoss)
	}
}
