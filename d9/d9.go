package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d9/test.txt")
	defer cleanup()

	histories := [][]int{}
	for scanner.Scan() {
		histories = append(histories, utils.ToInts(strings.Fields(scanner.Text())))
	}
	//fmt.Println(histories)

	total := 0
	for _, history := range histories {
		prediction := history[len(history)-1]
		fmt.Println(history)
		for sum := -1; sum != 0; {
			sum = 0
			newHistory := []int{}
			for i := 0; i < len(history)-1; i++ {
				newHistory = append(newHistory, history[i+1]-history[i])
				//history[i] = history[i+1] - history[i]
				sum += history[i]
			}
			//history = history[:len(history)-1]
			//fmt.Println(history)
			prediction += newHistory[len(newHistory)-1]
			history = newHistory
		}
		fmt.Println(prediction)
		total += prediction
	}
	fmt.Println(total)
}
