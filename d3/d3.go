package main

import (
	"fmt"
	"utils"
)

type coord struct {
	x, y int
}

type partNum struct {
	num  int
	len  int
	coor coord
}

type adjacent struct {
	sym  byte
	coor coord
}

func (pn partNum) findSymbols(symbols map[coord]byte) []adjacent {
	adj := []adjacent{}

	if s := symbols[coord{pn.coor.x + 1, pn.coor.y}]; s != 0 {
		adj = append(adj, adjacent{s, coord{pn.coor.x + 1, pn.coor.y}})
	}

	if s := symbols[coord{pn.coor.x - pn.len, pn.coor.y}]; s != 0 {
		adj = append(adj, adjacent{s, coord{pn.coor.x - pn.len, pn.coor.y}})
	}

	for _, y := range []int{-1, 1} {
		for x := pn.coor.x - pn.len; x <= pn.coor.x+1; x++ {
			if s := symbols[coord{x, pn.coor.y + y}]; s != 0 {
				adj = append(adj, adjacent{s, coord{x, pn.coor.y + y}})
			}
		}
	}

	if len(adj) > 0 {
		return adj
	}

	return nil
}

func main() {
	scanner, cleanup := utils.FileScaner("d3/input.txt")
	defer cleanup()

	//numbers := make(map[coord]byte)
	partNumbers := []partNum{}
	symbols := make(map[coord]byte)
	var y, x int
	for scanner.Scan() {
		var c byte
		pn := partNum{}
		for x, c = range scanner.Bytes() {
			switch {
			case c == '.':
			case isNum(c):
				//numbers[coord{x, y}] = c
				pn.num = pn.num*10 + int(c-'0')
				pn.coor = coord{x, y}
				pn.len++
				continue
			default:
				symbols[coord{x, y}] = c
			}
			if pn.len > 0 {
				partNumbers = append(partNumbers, pn)
				pn = partNum{}
			}
		}
		if pn.len > 0 {
			partNumbers = append(partNumbers, pn)
			pn = partNum{}
		}
		y++
	}
	//printMap(x+1, y, numbers, symbols)

	sum := 0
	type gear struct {
		num, ratio int
	}
	gears := make(map[coord]gear)
	for _, pn := range partNumbers {
		if ss := pn.findSymbols(symbols); ss != nil {
			for _, s := range ss {
				if s.sym == '*' {
					g := gears[s.coor]
					if g.ratio == 0 {
						g.ratio = 1
					}
					gears[s.coor] = gear{g.num + 1, g.ratio * pn.num}
				}
			}
			sum += pn.num
		}
	}
	fmt.Println("SUM:", sum)

	sum = 0
	for _, gear := range gears {
		if gear.num == 2 {
			sum += gear.ratio
		}
	}
	fmt.Println("ratio", sum)
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func printMap(width, height int, numbers, symbols map[coord]byte) {
	c := coord{}
	for c.y = 0; c.y < height; c.y++ {
		for c.x = 0; c.x < width; c.x++ {
			ch := numbers[c] + symbols[c]
			if ch == 0 {
				ch = '.'
			}
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}
