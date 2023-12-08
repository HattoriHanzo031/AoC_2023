package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

func main() {
	defer utils.Profile(time.Now())

	scanner, cleanup := utils.FileScaner("d8/input.txt")
	defer cleanup()

	scanner.Scan()
	instructions := make([]byte, len(scanner.Bytes()))
	for i, ins := range scanner.Bytes() {
		instructions[i] = 0
		if ins == 'R' {
			instructions[i] = 1
		}
	}
	scanner.Scan()

	lanes := []string{}
	endingPoints := make(map[string]bool)
	nodes := make(map[string][]string)
	for scanner.Scan() {
		ss := strings.Split(strings.TrimSuffix(scanner.Text(), ")"), " = (")
		nodes[ss[0]] = strings.Split(ss[1], ", ")
		if ss[0][2] == 'A' {
			lanes = append(lanes, ss[0])
		} else if ss[0][2] == 'Z' {
			endingPoints[ss[0]] = true
		}
	}

	cur := "AAA"
	turn := 0
	for ; cur != "ZZZ"; turn++ {
		cur = nodes[cur][instructions[turn%len(instructions)]]
	}
	fmt.Println("P1:", turn)

	periods := []int{}
	for turn = 0; len(periods) < len(lanes); turn++ {
		for i := 0; i < len(lanes); i++ {
			lanes[i] = nodes[lanes[i]][instructions[turn%len(instructions)]]
			if endingPoints[lanes[i]] {
				periods = append(periods, turn+1)
			}
		}
	}

	fmt.Println("P2:", utils.LCM(periods[0], periods[1], periods[2:]...))
}
