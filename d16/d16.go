package main

import (
	"fmt"
	"time"
	"utils"
)

type coord struct{ x, y int }

type direction coord

func (d *direction) cw() {
	d.x, d.y = -d.y, d.x
}

func (d *direction) ccw() {
	d.x, d.y = d.y, -d.x
}

type object struct {
	t            byte
	energizedDir direction
}

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

func (o *object) interact(dir direction) []direction {
	if o == nil {
		return nil
	}

	if o.energizedDir == dir {
		return nil
	}
	o.energizedDir = dir

	switch o.t {
	case '.':
		return []direction{dir}
	case '-':
		if dir == left || dir == right {
			return []direction{dir}
		}
		return []direction{left, right}
	case '|':
		if dir == up || dir == down {
			return []direction{dir}
		}
		return []direction{up, down}
	case '/':
		if dir == up || dir == down {
			dir.cw()
		} else {
			dir.ccw()
		}
		return []direction{dir}
	case '\\':
		if dir == up || dir == down {
			dir.ccw()
		} else {
			dir.cw()
		}
		return []direction{dir}
	default:
		return nil
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d16/input.txt")
	defer cleanup()

	objects := make(map[coord]*object)
	y := 0
	xMax := 0
	for scanner.Scan() {
		for x, c := range scanner.Bytes() {
			objects[coord{x, y}] = &object{c, direction{}}
		}
		y++
		xMax = len(scanner.Bytes())
	}
	yMax := y

	maxP2 := 0
	for y := 0; y < yMax; y++ {
		p1(objects, coord{0, y}, right)
		total := countEnergized(objects)
		maxP2 = max(maxP2, total)
	}

	for y := 0; y < yMax; y++ {
		p1(objects, coord{xMax - 1, y}, left)
		total := countEnergized(objects)
		maxP2 = max(maxP2, total)
	}

	for x := 0; x < xMax; x++ {
		p1(objects, coord{x, 0}, down)
		total := countEnergized(objects)
		maxP2 = max(maxP2, total)
	}

	for x := 0; x < xMax; x++ {
		p1(objects, coord{x, yMax - 1}, up)
		total := countEnergized(objects)
		maxP2 = max(maxP2, total)
	}

	fmt.Println("P2", maxP2)
}

func p1(objects map[coord]*object, cur coord, dir direction) {
	for _, next := range objects[cur].interact(dir) {
		p1(objects, coord{next.x + cur.x, next.y + cur.y}, next)
	}
}

func countEnergized(objects map[coord]*object) int {
	total := 0
	for _, v := range objects {
		if v.energizedDir.x != 0 || v.energizedDir.y != 0 {
			total++
			v.energizedDir = direction{}
		}
	}
	return total
}
