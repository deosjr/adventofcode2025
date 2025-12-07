package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func main() {
	var fringe map[int]struct{}
	split := 0
	lib.ReadFileByLine(7, func(line string) {
		if fringe == nil {
			start := strings.IndexRune(line, 'S')
			fringe = map[int]struct{}{start:struct{}{}}
			return
		}
		for i, c := range line {
			if c == '.' {
				continue
			}
			if _, ok := fringe[i]; !ok {
				continue
			}
			delete(fringe, i)
			fringe[i-1] = struct{}{}
			fringe[i+1] = struct{}{}
			split++
		}
	})
	lib.WritePart1("%d", split)
}
