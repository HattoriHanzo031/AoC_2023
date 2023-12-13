package main

import (
	"fmt"
	"slices"
	"time"
	"utils"
)

func isReflection(first, second [][]byte, smudge int) bool {
	for i, f := range first {
		if slices.CompareFunc(f, second[len(second)-1-i], func(b1, b2 byte) int {
			if smudge == 1 && b1 != b2 {
				smudge = 0
				return 0
			}
			return int(b1) - int(b2)
		}) != 0 {
			return false
		}
	}
	return smudge == 0 // we must use up all the smudges
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d13/input.txt")
	defer cleanup()

	var patternsNormal [][][]byte
	var pattern [][]byte
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			patternsNormal = append(patternsNormal, slices.Clone(pattern))
			pattern = pattern[:0]
			continue
		}
		pattern = append(pattern, slices.Clone(scanner.Bytes()))
	}
	patternsNormal = append(patternsNormal, pattern)

	// Transposed for vertical search
	patternsTransposed := make([][][]byte, 0, len(patternsNormal))
	for _, pp := range patternsNormal {
		pattern := make([][]byte, len(pp[1]))
		for _, p := range pp {
			for i, c := range p {
				pattern[i] = append(pattern[i], c)
			}
		}
		patternsTransposed = append(patternsTransposed, pattern)
	}

	solution := func(smudges int) int {
		total := 0
		for hv, patterns := range [][][][]byte{patternsTransposed, patternsNormal} {
			for _, pattern := range patterns {
				for i := 1; i < len(pattern); i++ {
					size := min(i, len(pattern)-i)
					if isReflection(pattern[i-size:i], pattern[i:i+size], smudges) {
						total += ((hv * 99) + 1) * i // 1 for vertical and 100 for horizontal
						break
					}
				}
			}
		}
		return total
	}

	fmt.Println("P1:", solution(0))
	fmt.Println("P2", solution(1))
}
