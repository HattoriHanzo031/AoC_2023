package main

import (
	"fmt"
	"strconv"
	"strings"

	"utils"
)

func main() {
	scanner, cleanup := utils.FileScaner("d2/input.txt")
	defer cleanup()

	sumOne, sumTwo := 0, 0
	for scanner.Scan() {
		sumOne += lineOne(scanner.Text())
		sumTwo += lineTwo(scanner.Text())
	}

	fmt.Println(sumOne)
	fmt.Println(sumTwo)
}

func lineOne(line string) int {
	// 12 red cubes, 13 green cubes, and 14 blue cubes
	var allowed = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	ss := strings.Split(line, ":")
	id := utils.Must(strconv.Atoi(strings.TrimPrefix(ss[0], "Game ")))

	for _, s := range strings.Split(ss[1], ";") {
		for _, cubes := range strings.Split(s, ",") {
			ss := strings.Fields(cubes)
			if utils.Must(strconv.Atoi(ss[0])) > allowed[ss[1]] {
				return 0
			}
		}
	}
	return id
}

func lineTwo(line string) int {
	var minCubes = map[string]int{}
	ss := strings.Split(line, ":")
	for _, s := range strings.Split(ss[1], ";") {
		for _, cubes := range strings.Split(s, ",") {
			ss := strings.Fields(cubes)
			minCubes[ss[1]] = max(minCubes[ss[1]], utils.Must(strconv.Atoi(ss[0])))
		}
	}

	return minCubes["red"] * minCubes["green"] * minCubes["blue"]
}
