package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"time"
	"utils"
)

func hash(s []byte) int {
	cv := 0
	for _, b := range s {
		cv = ((cv + int(b)) * 17) % 256
	}
	return cv
}

type box struct {
	label string
	focal int
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d15/input.txt")
	defer cleanup()

	scanner.Scan()
	input := bytes.Split(slices.Clone(scanner.Bytes()), []byte{','})

	boxes := make(map[int][]box)
	for _, in := range input {
		if bytes.HasSuffix(in, []byte{'-'}) {
			in = bytes.TrimSuffix(in, []byte{'-'})
		}

		step := bytes.Split(in, []byte{'='})
		h := hash(step[0])
		index := slices.IndexFunc(boxes[h], func(b box) bool {
			return b.label == string(step[0])
		})

		if len(step) == 1 {
			if index != -1 {
				boxes[h] = append(boxes[h][:index], boxes[h][index+1:]...)
			}
			continue
		}

		focal := utils.Must(strconv.Atoi(string(step[1])))
		if index == -1 {
			boxes[h] = append(boxes[h], box{string(step[0]), focal})
		} else {
			boxes[h][index].focal = focal
		}
	}

	fmt.Println("P1:", calcP1(input))
	fmt.Println("P2:", calcP2(boxes))
}

func print(boxes map[int][]box) {
	for k, v := range boxes {
		fmt.Print("Box:", k, " ")
		for _, box := range v {
			fmt.Print(box.label, " ", box.focal, ",")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func calcP1(input [][]byte) int {
	total := 0
	for _, in := range input {
		h := hash(in)
		total += h
	}
	return total
}

func calcP2(boxes map[int][]box) int {
	total := 0
	for k, v := range boxes {
		for i, box := range v {
			total += (k + 1) * (i + 1) * (box.focal)
		}
	}
	return total
}
