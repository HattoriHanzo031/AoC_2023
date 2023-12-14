package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d14/input.txt")
	defer cleanup()

	var input []string
	for scanner.Scan() {
		if input == nil {
			input = make([]string, len(scanner.Text()))
		}
		for i, b := range scanner.Text() {
			input[i] += string(b)
		}
	}

	for i, line := range input {
		out := []string{}
		for _, s := range strings.Split(line, "#") {
			count := strings.Count(s, ".")
			out = append(out, strings.ReplaceAll(s, ".", "")+strings.Repeat(".", count))
		}
		input[i] = strings.Join(out, "#")
	}

	total := 0
	for j := 0; j < len(input[0]); j++ {
		for i := 0; i < len(input); i++ {
			fmt.Print(string(input[i][j]))
			if input[i][j] == 'O' {
				total += len(input[0]) - j
			}
		}
		fmt.Println("")
	}
	fmt.Println(total)
}
