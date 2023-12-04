package main

import (
	"fmt"
	"strings"
	"unicode"

	"utils"
)

func main() {
	scanner, cleanup := utils.FileScaner("d1/input.txt")
	defer cleanup()

	sumOne, sumTwo := 0, 0
	for scanner.Scan() {
		sumOne += lineNumberOne(scanner.Text())
		sumTwo += lineNumberTwo(scanner.Text())
	}

	fmt.Println(sumOne)
	fmt.Println(sumTwo)
}

func lineNumberOne(line string) int {
	var num byte
	for i := 0; i < len(line); i++ {
		if unicode.IsNumber(rune(line[i])) {
			num = line[i] - '0'
			break
		}
	}
	num *= 10
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(line[i])) {
			num += line[i] - '0'
			break
		}
	}
	return int(num)
}

func lineNumberTwo(s string) int {
	firstIndex, lastIndex := len(s), -1
	first, last := 0, 0
	for i, num := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		index := strings.Index(s, num)
		if index != -1 && index < firstIndex {
			firstIndex = index
			first = (i % 9) + 1
		}
		index = strings.LastIndex(s, num)
		if index != -1 && index > lastIndex {
			lastIndex = index
			last = (i % 9) + 1
		}
	}

	return first*10 + last
}
