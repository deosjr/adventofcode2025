package main

import (
	"cmp"
	"maps"
	"math"
	"slices"
	"sort"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type jbox struct {
	x, y, z int64
}

func euclidianDistance(p, q jbox) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	dz := p.z - q.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

type connection struct {
	from, to jbox
	length float64
}

func main() {
	//lib.Test()
	var boxes []jbox
	lib.ReadFileByLine(8, func(line string) {
		split := strings.Split(line, ",")
		x := lib.MustParseInt(split[0])
		y := lib.MustParseInt(split[1])
		z := lib.MustParseInt(split[2])
		boxes = append(boxes, jbox{x, y, z})
	})
	var connections []connection
	for i:=0; i<len(boxes); i++ {
		for j:=i+1; j<len(boxes); j++ {
			from, to := boxes[i], boxes[j]
			dist := euclidianDistance(from, to)
			connections = append(connections, connection{from, to, dist})
		}
	}
	sort.Slice(connections, func(i, j int) bool { return connections[i].length < connections[j].length })

	circuits := map[jbox]int{}
	for _, c := range connections[:1000] {
		idFrom, okFrom := circuits[c.from]
		idTo, okTo := circuits[c.to]
		if okFrom && okTo {
			// connect the two circuits
			for k, v := range circuits {
				if v == idTo {
					circuits[k] = idFrom
				}
			}
			continue
		}
		if okFrom {
			circuits[c.to] = idFrom
			continue
		}
		if okTo {
			circuits[c.from] = idTo
			continue
		}
		newid := len(circuits)
		circuits[c.from] = newid
		circuits[c.to] = newid
	}

	// this could be done in the loop above, I suppose
	sizes := map[int]int{}
	for _, v := range circuits {
		sizes[v] += 1
	}
	sortFunc := func(a, b int) int {
		return cmp.Compare(b, a)
	}
	values := slices.SortedFunc(maps.Values(sizes), sortFunc)[:3]

	lib.WritePart1("%v", values[0] * values[1] * values[2])
}
