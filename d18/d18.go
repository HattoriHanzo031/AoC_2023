package main

import (
	"fmt"
	"time"
	"utils"
)

type coord struct{ x, y int }

type direction coord

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	left  = direction{-1, 0}
	right = direction{1, 0}
)

var dirMapP1 = map[byte]direction{
	'U': up,
	'D': down,
	'L': left,
	'R': right,
}

var dirMapP2 = map[int]direction{
	3: up,
	1: down,
	2: left,
	0: right,
}

type instruction struct {
	dir direction
	num int
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d18/input.txt")
	defer cleanup()

	var instructionsP1 []instruction
	var instructionsP2 []instruction
	for scanner.Scan() {
		var dir byte
		var color, num int
		fmt.Sscanf(scanner.Text(), "%c %d (#%X)", &dir, &num, &color)
		instructionsP1 = append(instructionsP1, instruction{dirMapP1[dir], num})
		instructionsP2 = append(instructionsP2, instruction{dirMapP2[color&0x0F], color >> 4})
	}

	fmt.Println("P1:", dig(instructionsP1))
	fmt.Println("P2:", dig(instructionsP2))
}

func dig(instructions []instruction) int {
	points := []coord{{0, 0}}
	circumference := 0
	cur := coord{0, 0}
	for _, ins := range instructions {
		cur.x += ins.dir.x * ins.num
		cur.y += ins.dir.y * ins.num
		points = append(points, coord{cur.x, cur.y})
		circumference += ins.num
	}
	return area(points) + circumference/2 + 1
}

func area(points []coord) int {
	points = append(points, points[0])
	a := 0
	for i := 0; i < len(points)-1; i++ {
		a += (points[i].x * points[i+1].y) - (points[i].y * points[i+1].x)
	}
	return a / 2
}
