package main

import (
	"fmt"
	"time"
	"utils"
)

type coord struct{ x, y int }

func (c *coord) move(dir direction) coord {
	return coord{c.x + dir.x, c.y + dir.y}
}

type direction coord

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

var dirMap = map[byte]direction{
	'U': up,
	'D': down,
	'L': left,
	'R': right,
}

var dirMapP2 = map[byte]direction{
	3: up,
	1: down,
	2: left,
	0: right,
}

type instruction struct {
	dir   direction
	num   uint32
	color uint32
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d18/input.txt")
	defer cleanup()

	var instructions []instruction
	var instructionsP2 []instruction
	for scanner.Scan() {
		var ins instruction
		var dir byte
		fmt.Sscanf(scanner.Text(), "%c %d (#%X)", &dir, &ins.num, &ins.color)
		ins.dir = dirMap[dir]
		instructions = append(instructions, ins)
		instructionsP2 = append(instructionsP2, instruction{
			dir: dirMapP2[byte(ins.color&0x0F)],
			num: ins.color >> 4,
		})
		//fmt.Println(dirMapP2[byte(ins.color&0x0F)], ins.color>>4)
	}

	maxX, maxY := 0, 0
	minX, minY := 0, 0
	cur := coord{0, 0}
	trench := map[coord]uint32{cur: 0xFFFFFF}
	for _, ins := range instructions {
		for i := uint32(0); i < ins.num; i++ {
			cur.x += ins.dir.x
			cur.y += ins.dir.y
			maxX, maxY = max(maxX, cur.x), max(maxY, cur.y)
			minX, minY = min(minX, cur.x), min(minY, cur.y)
			trench[cur] = ins.color
		}
	}

	inside := map[coord]bool{}
	fill(trench, inside, coord{1, 1})

	fmt.Println("TRENCH:", len(trench))
	fmt.Println("INSIDE:", len(inside))
	fmt.Println("TOTAL:", len(inside)+len(trench))
	//print(trench, inside, maxX, maxY, minX, minY)
}

func fill(trench map[coord]uint32, inside map[coord]bool, cur coord) {
	if trench[cur] != 0 {
		return
	}
	if inside[cur] {
		return
	}
	inside[cur] = true

	fill(trench, inside, cur.move(up))
	fill(trench, inside, cur.move(down))
	fill(trench, inside, cur.move(left))
	fill(trench, inside, cur.move(right))
}

func print(trench map[coord]uint32, inside map[coord]bool, maxX, maxY, minX, minY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if trench[coord{x, y}] == 0 {
				if inside[coord{x, y}] {
					fmt.Print(".")
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
}
