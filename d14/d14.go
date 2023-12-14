package main

import (
	"bytes"
	"fmt"
	"time"
	"utils"
)

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d14/input.txt")
	defer cleanup()

	var input [][]byte
	for scanner.Scan() {
		if input == nil {
			input = make([][]byte, len(scanner.Bytes()))
		}
		for i, b := range scanner.Bytes() {
			input[i] = append(input[i], b)
		}
	}

	for i, line := range input {
		out := [][]byte{}
		for _, s := range bytes.Split(line, []byte{'#'}) {
			count := bytes.Count(s, []byte{'.'})
			out = append(out, append(bytes.ReplaceAll(s, []byte{'.'}, []byte{}), bytes.Repeat([]byte{'.'}, count)...))
		}
		input[i] = []byte(bytes.Join(out, []byte{'#'}))
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
