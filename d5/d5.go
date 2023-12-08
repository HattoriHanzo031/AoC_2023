package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
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

	partTwo()
}

// PART TWO
type singleRange struct {
	start, end int
}

func (sr singleRange) isIn(a int) bool {
	return a >= sr.start && a <= sr.end
}

type mappingRanges func(in singleRange) (mapped []singleRange, unmapped []singleRange)

func mappingRangesFn(from, to, offset int) mappingRanges {
	return func(inR singleRange) (mapped []singleRange, unmapped []singleRange) {
		if from <= inR.start && to >= inR.end {
			return []singleRange{{inR.start + offset, inR.end + offset}}, nil
		}

		if !inR.isIn(from) && !inR.isIn(to) {
			return nil, []singleRange{inR}
		}

		if inR.isIn(from - 1) {
			unmapped = append(unmapped, singleRange{inR.start, from - 1})
			inR.start = from
		}
		if inR.isIn(to + 1) {
			unmapped = append(unmapped, singleRange{to + 1, inR.end})
			inR.end = to
		}
		if inR.start <= inR.end {
			mapped = append(mapped, singleRange{inR.start + offset, inR.end + offset})
		}
		return mapped, unmapped
	}
}

type stageRanges func(inRs []singleRange) (out []singleRange)

func stageRangesFn(mappers []mappingRanges) stageRanges {
	return func(inRs []singleRange) (out []singleRange) {
		for _, inR := range inRs {
			unmapped := []singleRange{inR}
			for _, mapper := range mappers {
				accUm := []singleRange{}
				for _, in := range unmapped {
					mm, um := mapper(in)
					out = append(out, mm...)
					accUm = append(accUm, um...)
				}
				unmapped = accUm
			}
			out = append(out, unmapped...)
		}
		return out
	}
}

func partTwo() {
	scanner, cleanup := utils.FileScaner("d5/input.txt")
	defer cleanup()
	scanner.Scan()
	seeds := utils.ToInts(strings.Fields(strings.Split(scanner.Text(), ":")[1]))

	scanner.Scan() // skip empty line
	scanner.Scan() // skip title

	stage := []mappingRanges{}
	pipeline := []stageRanges{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			pipeline = append(pipeline, stageRangesFn(slices.Clone(stage)))
			stage = stage[:0]
			scanner.Scan() // skip title
			continue
		}
		m := utils.ToInts(strings.Fields(scanner.Text()))
		stage = append(stage, mappingRangesFn(m[1], m[1]+m[2]-1, m[0]-m[1]))
	}

	ranges := []singleRange{}
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, singleRange{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	for _, stage := range pipeline {
		ranges = stage(ranges)
	}

	slices.SortFunc(ranges, func(a, b singleRange) int {
		return a.start - b.start
	})

	fmt.Println("closest:", ranges[0].start)
}
