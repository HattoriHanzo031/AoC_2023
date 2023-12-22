package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
	"utils"
)

type coord struct{ x, y int }

type brick struct {
	blocks     []coord
	minZ, maxZ int
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d22/input.txt")
	defer cleanup()

	// 1,0,1~1,2,1
	bricks := []brick{}
	for i := 0; scanner.Scan(); i++ {
		b := brick{}
		s, e, _ := strings.Cut(scanner.Text(), "~")
		start := utils.ToInts(strings.Split(s, ","))
		end := utils.ToInts(strings.Split(e, ","))
		for x := min(start[0], end[0]); x <= max(start[0], end[0]); x++ {
			for y := min(start[1], end[1]); y <= max(start[1], end[1]); y++ {
				b.blocks = append(b.blocks, coord{x, y})
			}
		}
		b.minZ, b.maxZ = min(start[2], end[2]), max(start[2], end[2])+1
		bricks = append(bricks, b)
	}

	drop(bricks) // initial drop

	totalP1 := 0
	totalP2 := 0
	for i := range bricks {
		dropped := drop(slices.Delete(slices.Clone(bricks), i, i+1))
		if dropped == 0 {
			totalP1++
		}
		totalP2 += dropped
	}
	fmt.Println("P1:", totalP1)
	fmt.Println("P2:", totalP2)
}

func drop(bricks []brick) int {
	// no need to sort each time, but sorting sorted slice is very fast
	slices.SortFunc(bricks, func(a, b brick) int {
		return a.minZ - b.minZ
	})

	dropped := 0
	surfaces := map[coord]int{}
	for i, brick := range bricks {
		dropTo := 0
		for _, bl := range brick.blocks {
			dropTo = max(dropTo, surfaces[bl])
		}

		if dropTo != bricks[i].minZ {
			bricks[i].maxZ = bricks[i].maxZ - (bricks[i].minZ - dropTo)
			bricks[i].minZ = dropTo
			dropped++
		}

		for _, bl := range brick.blocks {
			surfaces[bl] = bricks[i].maxZ
		}
	}
	return dropped
}
