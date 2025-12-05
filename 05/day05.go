package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type IDrange struct {
	from, to int64
}

func main() {
	input := strings.Split(strings.TrimSpace(lib.ReadFile(5)), "\n\n")
	ranges := []IDrange{}
	for _, s := range strings.Split(input[0], "\n") {
		split := strings.Split(s, "-")
		from, to := lib.MustParseInt(split[0]), lib.MustParseInt(split[1])
		ranges = append(ranges, IDrange{from:from, to:to})
	}
	ids := []int64{}
	for _, s := range strings.Split(input[1], "\n") {
		ids = append(ids, lib.MustParseInt(s))
	}
	
	var p1 int64
	for _, id := range ids {
		for _, r := range ranges {
			if id < r.from || id > r.to {
				continue
			}
			p1++
			break
		}
	}
	
	lib.WritePart1("%d", p1)
}
