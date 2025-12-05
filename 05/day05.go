package main

import (
	"sort"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type IDrange struct {
	from, to int64
}

func (r IDrange) overlap(o IDrange) (IDrange, bool) {
	if r.from > o.to || r.to < o.from {
		return IDrange{}, false
	}
	from := r.from
	if o.from < from {
		from = o.from
	}
	to := r.to
	if o.to > to {
		to = o.to
	}
	return IDrange{from:from, to:to}, true
}

func collapseRanges(list []IDrange) []IDrange {
	sort.SliceStable(list, func(i, j int) bool { return list[i].from < list[j].from })
	newlist := []IDrange{}
	r := list[0]
	for _, o := range list[1:] {
		nr, ok := r.overlap(o)
		if ok {
			r = nr
			continue
		}
		newlist = append(newlist, r)
		r = o
	}
	newlist = append(newlist, r)
	return newlist
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
	ranges = collapseRanges(ranges)
	
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

	var p2 int64
	for _, r := range ranges {
		p2 += r.to - r.from + 1
	}
	lib.WritePart2("%d", p2)
}
