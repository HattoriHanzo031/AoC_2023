package main

import (
	"fmt"
	"slices"
	"time"
	"utils"
)

func print(input [][]byte) {
	for _, line := range input {
		fmt.Println(string(line))
	}
	fmt.Println("")
}

func rotate(input [][]byte) {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input[i])-1-i; j++ {
			input[i][j], input[len(input[i])-1-j][len(input)-1-i] = input[len(input[i])-1-j][len(input)-1-i], input[i][j]
		}
	}
	slices.Reverse(input)
}

func slide(input [][]byte) {
	for i, line := range input {
		storeRoundTo := len(line) - 1
		for j := len(line) - 1; j >= 0; j-- {
			switch line[j] {
			case 'O':
				input[i][storeRoundTo] = 'O'
				if j != storeRoundTo {
					input[i][j] = '.'
				}
				storeRoundTo--
			case '#':
				storeRoundTo = j - 1
			}
		}
	}
}

func cycle(input [][]byte) {
	for i := 0; i < 4; i++ {
		rotate(input)
		slide(input)
	}
}

func hash(input [][]byte) string {
	key := []byte{}
	for i, line := range input {
		for j, c := range line {
			if c == 'O' {
				key = append(key, byte(i), byte(j))
			}
		}
	}
	return string(key)
}

func totalLoad(input [][]byte) int {
	total := 0
	for j := 0; j < len(input[0]); j++ {
		for i := 0; i < len(input); i++ {
			if input[i][j] == 'O' {
				total += len(input) - i
			}
		}
	}
	return total
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d14/input.txt")
	defer cleanup()

	var input [][]byte
	for scanner.Scan() {
		input = append(input, slices.Clone(scanner.Bytes()))
	}

	cycleDetection := make(map[string]int)
	for i := 0; i < 1000_000_000; i++ {
		cycle(input)

		key := hash(input)
		if prev, found := cycleDetection[key]; found {
			cycleLen := i - prev
			// skip all the cycles
			i = 1000_000_000 - ((1000_000_000 - i) % cycleLen)
		}
		cycleDetection[key] = i
	}

	fmt.Println("P2:", totalLoad(input))
}
