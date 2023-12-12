package main

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"time"
	"utils"
)

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d12/input.txt")
	defer cleanup()

	for scanner.Scan() {
		in := bytes.Fields(scanner.Bytes())
		maybeMask := uint(0)
		mustMask := uint(0)
		for _, b := range in[0] {
			maybeMask <<= 1
			mustMask <<= 1
			switch b {
			case '?':
				maybeMask++
			case '#':
				mustMask++
			}
		}

		nums := utils.ToIntegers[[]byte, uint](bytes.Split(in[1], []byte{','}))
		slices.Reverse(nums)
		record := maybeMask | mustMask

		fit(record, mustMask, nums, 0, 0)
	}
	fmt.Println(count)
}

var count int

func fit(record, must uint, nums []uint, totalShift uint, mapped uint) {
	if len(nums) == 0 {
		if (mapped | must) == mapped {
			count++
		}
		return
	}

	recordBits := uint(math.Log2(float64(record)) + 1)
	n := nums[0]
	if recordBits < n {
		return
	}

	binN := (uint(1) << nums[0]) - 1
	nums = nums[1:]
	for shift := uint(0); shift <= recordBits-n; shift++ {
		fits := (record & binN) == binN
		if fits {
			fit(record>>(n+1), must, nums, totalShift+n+1, mapped|(binN<<totalShift))
		}
		record >>= 1
		totalShift++
	}
}
