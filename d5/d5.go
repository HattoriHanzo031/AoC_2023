package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"
	"utils"
)

type mapping func(in int) (int, bool)

func mappingFn(from, to, offset int) mapping {
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
				return out, true
			}
		}
		return in, true
	}
}

func pipelineFn(mm []mapping) mapping {
	return func(in int) (int, bool) {
		for _, m := range mm {
			in, _ = m(in)
		}
		return in, true
	}
}

func main() {
	scanner, cleanup := utils.FileScaner("d5/input.txt")
	defer cleanup()
	scanner.Scan()
	seeds := utils.ToInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))

	scanner.Scan() // skip empty line
	scanner.Scan() // skip title

	stages := []mapping{}
	mappings := []mapping{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			stages = append(stages, stageFn(slices.Clone(mappings)))
			mappings = mappings[:0]
			scanner.Scan() // skip title
			continue
		}
		m := utils.ToInts(strings.Fields(scanner.Text()))
		mappings = append(mappings, mappingFn(m[1], m[1]+m[2]-1, m[0]-m[1]))
	}

	pipeline := pipelineFn(stages)

	closest := math.MaxInt
	for _, seed := range seeds {
		location, _ := pipeline(seed)
		closest = min(closest, location)
	}
	fmt.Println(closest)

	var wg sync.WaitGroup
	ch := make(chan int, 10)
	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			closest = math.MaxInt
			for count := 0; count < seeds[i+1]; count++ {
				location, _ := pipeline(seeds[i] + count)
				closest = min(closest, location)
			}
			ch <- closest
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	closest = math.MaxInt
	for location := range ch {
		closest = min(closest, location)
	}
	fmt.Println(closest)
}
