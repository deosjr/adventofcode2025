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

func (c connection) connect(circuits map[jbox]int) {
	idFrom, okFrom := circuits[c.from]
	idTo, okTo := circuits[c.to]
	if okFrom && okTo {
		// connect the two circuits
		for k, v := range circuits {
			if v == idTo {
				circuits[k] = idFrom
			}
		}
		return
	}
	if okFrom {
		circuits[c.to] = idFrom
		return
	}
	if okTo {
		circuits[c.from] = idTo
		return
	}
	newid := len(circuits)
	circuits[c.from] = newid
	circuits[c.to] = newid
}

func sizes(circuits map[jbox]int) []int {
	sizes := map[int]int{}
	for _, v := range circuits {
		sizes[v] += 1
	}
	sortFunc := func(a, b int) int {
		return cmp.Compare(b, a)
	}
	return slices.SortedFunc(maps.Values(sizes), sortFunc)
}

func main() {
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

	p1split := 1000
	circuits := map[jbox]int{}
	for _, c := range connections[:p1split] {
		c.connect(circuits)
	}
	values := sizes(circuits)
	lib.WritePart1("%v", values[0] * values[1] * values[2])

	for _, box := range boxes {
		if _, ok := circuits[box]; !ok {
			circuits[box] = len(circuits)
		}
	}
	for _, c := range connections[p1split:] {
		c.connect(circuits)
		if len(sizes(circuits)) == 1 {
			lib.WritePart2("%d", c.from.x * c.to.x)
			break
		}
	}
}
