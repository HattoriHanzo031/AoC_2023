package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func FileScaner(path string) (*bufio.Scanner, func()) {
	file := Must(os.Open(path))
	scanner := bufio.NewScanner(file)
	return scanner, func() {
		Must(struct{}{}, scanner.Err())
		file.Close()
	}
}

func ToInts(ss []string) []int {
	ints := make([]int, 0, len(ss))
	for _, s := range ss {
		ints = append(ints, Must(strconv.Atoi(s)))
	}
	return ints
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Profile(start time.Time) {
	fmt.Println(time.Since(start))
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
