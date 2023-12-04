package utils

import (
	"bufio"
	"os"
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
