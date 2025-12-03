package main

import (
	"fmt"
	"github.com/deosjr/adventofcode2025/lib"
)

func part1(line string) int64 {
	length := len(line)
	dec, uni := line[length-2], line[length-1]
	for i:=length-3; i >=0; i-- {
		c := line[i]
		if c < dec {
			continue
		}
		dec, c = c, dec
		if c > uni {
			uni = c
		}
	}
	return lib.MustParseInt(fmt.Sprintf("%c%c", dec, uni))
}

func main() {
	var p1 int64
	lib.ReadFileByLine(3, func(line string) {
		p1 += part1(line)
	})
	lib.WritePart1("%d", p1)
}
