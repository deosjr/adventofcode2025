package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func add(nums ...int64) int64 {
	var sum int64 = 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func mul(nums ...int64) int64 {
	var prod int64 = 1
	for _, n := range nums {
		prod *= n
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

	var p1, p2 int64
	for index, op := range ops {
		list := []int64{}
		for i:=0; i<len(matrix); i++ {
			list = append(list, matrix[i][index])
		}
		if op == "*" {
			p1 += mul(list...)
			continue
		}
		p1 += add(list...)
	}
	lib.WritePart1("%d", p1)

	s := strings.Split(strings.Trim(lib.ReadFile(6), "\n"), "\n")
	rawops := s[len(s)-1]
	spaces := strings.Split(strings.ReplaceAll(rawops, "*", "+"), "+")[1:]
	lens := []int{}
	for _, space := range spaces {
		lens = append(lens, len(space))
	}
	lens[len(lens)-1] = lens[len(lens)-1]+1

	offset := 0
	for i, op := range ops {
		list := []int64{}
		for j:=0; j<lens[i]; j++ {
			var n int64
			for k:=0; k<len(s)-1; k++ {
				c := s[k][offset]
				if c == ' ' {
					continue
				}
				n *= 10
				n += int64(c) - 48
			}
			offset++
			list = append(list, n)
		}
		offset++
		if op == "*" {
			p2 += mul(list...)
			continue
		}
		p2 += add(list...)
	}
	
	lib.WritePart2("%d", p2)
}
