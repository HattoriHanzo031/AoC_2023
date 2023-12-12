package main

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"time"
	"utils"

	"golang.org/x/exp/constraints"
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

		nums := utils.ToUints(bytes.Split(in[1], []byte{','}))
		slices.Reverse(nums)
		record := maybeMask | mustMask
		//fmt.Printf("%32b %v\n", record, nums)

		fit2(record, mustMask, nums, 0, 0)
	}
	fmt.Println(count)
}

func sum[T constraints.Integer](ss []T) T {
	var sum T
	for _, s := range ss {
		sum += s
	}
	return sum
}

var count int

func fit(mustRecord, record uint, nums []uint, mapped uint) {
	fmt.Println("---", record, nums)

	if len(nums) == 0 {
		fmt.Printf("end %b %b\n", mapped, mustRecord)
		count++
		return
	}

	recordBits := uint(math.Log2(float64(record)) + 1)
	fmt.Println("record bits:", recordBits)
	n := nums[0]
	binN := (uint(1) << nums[0]) - 1
	nums = nums[1:]
	remainingBits := sum(nums) + uint(len(nums))
	fmt.Println("remaining:", remainingBits)
	for shift := uint(0); shift < recordBits-remainingBits; shift++ {
		fits := (record & binN) == binN
		if fits {
			fmt.Printf("%32b %v\n", record, fits)
			fit(mustRecord, record>>(n+1), nums, mapped<<(n+1)|binN)
		}
		record >>= 1
		mapped <<= 1
	}
}

func fit2(record, must uint, nums []uint, totalShift uint, mapped uint) {
	if len(nums) == 0 {
		//fmt.Printf("%32b %32b end\n", mapped, must)
		if (mapped | must) == mapped {
			//fmt.Println("REAL END")
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
			//fmt.Printf("%32b %v\n", record, fits)
			fit2(record>>(n+1), must, nums, totalShift+n+1, mapped|(binN<<totalShift))
		}
		record >>= 1
		totalShift++
	}
}
