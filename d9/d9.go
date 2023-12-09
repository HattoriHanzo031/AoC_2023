package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

func alternateAddSubFn() func(a, b int) int {
	invocation := 0
	return func(a, b int) int {
		invocation++
		return []func(a, b int) int{
			func(a, b int) int { return a - b },
			func(a, b int) int { return a + b },
		}[invocation%2](a, b)
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d9/input.txt")
	defer cleanup()

	histories := [][]int{}
	for scanner.Scan() {
		histories = append(histories, utils.ToInts(strings.Fields(scanner.Text())))
	}

	totalP1 := 0
	totalP2 := 0
	for _, history := range histories {
		fn := alternateAddSubFn()
		totalP1 += history[len(history)-1]
		totalP2 = fn(totalP2, history[0])
		for allZeroes := false; !allZeroes; {
			allZeroes = true
			for i := 0; i < len(history)-1; i++ {
				allZeroes = allZeroes && history[i] == 0
				history[i] = history[i+1] - history[i]
			}
			history = history[:len(history)-1]

			totalP1 += history[len(history)-1]
			totalP2 = fn(totalP2, history[0])
		}
	}

	fmt.Println("P1:", totalP1)
	fmt.Println("P2:", totalP2)
}
