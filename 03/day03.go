package main

import (
	"github.com/deosjr/adventofcode2025/lib"
)

func joltage(line string, n int) int64 {
	length := len(line)
	list := []byte(line[length-n:])
Loop:
	for i:=length-n-1; i >=0; i-- {
		c := line[i]
		for j:=0; j<n; j++ {
			if c < list[j] {
				continue Loop
			}
			list[j], c = c, list[j]
		}
	}
	return lib.MustParseInt(string(list))
}

func main() {
	var p1, p2 int64
	lib.ReadFileByLine(3, func(line string) {
		p1 += joltage(line, 2)
		p2 += joltage(line, 12)
	})
	lib.WritePart1("%d", p1)
	lib.WritePart2("%d", p2)
}
