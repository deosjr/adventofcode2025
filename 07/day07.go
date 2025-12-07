package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func main() {
	var fringe map[int]int
	split := 0
	lib.ReadFileByLine(7, func(line string) {
		if fringe == nil {
			start := strings.IndexRune(line, 'S')
			fringe = map[int]int{start:1}
			return
		}
		newfringe := map[int]int{}
		for i, c := range line {
			v, ok := fringe[i]
			if !ok {
				continue
			}
			if c == '.' {
				newfringe[i] += v
				continue
			}
			newfringe[i-1] += v
			newfringe[i+1] += v
			split++
		}
		fringe = newfringe
	})
	lib.WritePart1("%d", split)
	
	sum := 0
	for _, v := range fringe {
		sum += v
	}
	lib.WritePart2("%d", sum)
}
