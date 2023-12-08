package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {
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

	nodes := make(map[string][]string)
	for scanner.Scan() {
		ss := strings.Split(strings.TrimSuffix(scanner.Text(), ")"), " = (")
		nodes[ss[0]] = strings.Split(ss[1], ", ")
	}

	cur := "AAA"
	turn := 0
	for ; cur != "ZZZ"; turn++ {
		cur = nodes[cur][instructions[turn%len(instructions)]]
	}
	fmt.Println(turn)
}
