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
	num        int
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
		b := brick{num: i}
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

	slices.SortFunc(bricks, func(a, b brick) int {
		return a.minZ - b.minZ
	})

	type surface struct{ z, brickNum int }

	surfaces := map[coord]surface{}
	doNotDisintegrate := map[int]bool{}
	for i, brick := range bricks {
		dropTo := 0
		for _, bl := range brick.blocks {
			dropTo = max(dropTo, surfaces[bl].z)
		}
		bricks[i].maxZ = bricks[i].maxZ - (bricks[i].minZ - dropTo)
		bricks[i].minZ = dropTo

		supportedBy := []int{}
		for _, bl := range brick.blocks {
			if s, ok := surfaces[bl]; ok && s.z == dropTo {
				supportedBy = append(supportedBy, s.brickNum)
			}
			surfaces[bl] = surface{bricks[i].maxZ, bricks[i].num}
		}
		slices.Sort(supportedBy)
		supportedBy = slices.Compact(supportedBy)

		if len(supportedBy) == 1 {
			doNotDisintegrate[supportedBy[0]] = true
		}
	}
	slices.SortFunc(bricks, func(a, b brick) int {
		return a.minZ - b.minZ
	})

	totalP1 := 0
	totalP2 := 0
	for i, b := range bricks {
		if !doNotDisintegrate[b.num] {
			totalP1++
		} else {
			totalP2 += drop(slices.Delete(slices.Clone(bricks), i, i+1))
		}
	}

	fmt.Println("P1:", totalP1)
	fmt.Println("P2:", totalP2)
}

func drop(bricks []brick) int {
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
