package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"utils"
)

func toInts(ss []string) []int {
	ints := make([]int, 0, len(ss))
	for _, s := range ss {
		ints = append(ints, utils.Must(strconv.Atoi(s)))
	}
	return ints
}

type mapping func(in int) (int, bool)

func mappingFn(from, to, offset int) mapping {
	//fmt.Println(from, to, offset)
	return func(in int) (int, bool) {
		if in < from || in > to {
			return in, false
		}
		return in + offset, true
	}
}

func stageFn(mm []mapping) mapping {
	return func(in int) (int, bool) {
		for _, m := range mm {
			if out, found := m(in); found {
				//fmt.Println("stage:", in, "->", out)
				return out, true
			}
		}
		//fmt.Println("stage:", in)
		return in, true
	}
}

func pipelineFn(mm []mapping) mapping {
	return func(in int) (int, bool) {
		for _, m := range mm {
			in, _ = m(in)
		}
		//fmt.Println("pipeline:", in)
		return in, true
	}
}

func main() {
	scanner, cleanup := utils.FileScaner("d5/test.txt")
	defer cleanup()
	scanner.Scan()
	seeds := toInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))
	fmt.Println("seeds:", seeds)

	scanner.Scan() // skip empty line
	scanner.Scan() // skip title

	stages := []mapping{}
	mappings := []mapping{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			stages = append(stages, stageFn(slices.Clone(mappings)))
			mappings = mappings[:0]
			if !scanner.Scan() { // skip empty line
				break
			}
			scanner.Scan() // skip title
		}
		m := toInts(strings.Fields(scanner.Text()))
		fmt.Println(m)
		mappings = append(mappings, mappingFn(m[1], m[1]+m[2]-1, m[0]-m[1]))
	}

	pipeline := pipelineFn(stages)

	closest := math.MaxInt
	for _, seed := range seeds {
		location, _ := pipeline(seed)
		closest = min(closest, location)
	}
	fmt.Println(closest)

	closest = math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println(i)
		for count := 0; count < seeds[i+1]; count++ {
			location, _ := pipeline(seeds[i] + count)
			closest = min(closest, location)
		}
	}
	fmt.Println(closest)
}
