package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func add(matrix [][]int64, index int) int64 {
	var sum int64 = 0
	for i:=0; i<len(matrix); i++ {
		sum += matrix[i][index]
	}
	return sum
}

func mul(matrix [][]int64, index int) int64 {
	var prod int64 = 1
	for i:=0; i<len(matrix); i++ {
		prod *= matrix[i][index]
	}
	return prod
}

func main() {
	matrix := [][]int64{}
	ops := []string{}
	lib.ReadFileByLine(6, func(line string) {
		if line[0] == '*' || line[0] == '+' {
			ops = strings.Fields(line)
			return
		}
		list := []int64{}
		for _, s := range strings.Fields(line) {
			list = append(list, lib.MustParseInt(s))
		}
		matrix = append(matrix, list)
	})

	var p1 int64
	for i, op := range ops {
		if op == "*" {
			p1 += mul(matrix, i)
			continue
		}
		p1 += add(matrix, i)
	}

	lib.WritePart1("%d", p1)
}
