package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func paths(from, to, avoid string, graph map[string][]string, mem map[string]int64) int64 {
	if v, ok := mem[from]; ok {
		return v
	}
	directs := graph[from]
	for _, w := range directs {
		if w == to {
			mem[from] = 1
			return 1
		}
	}
	var sum int64
	for _, w := range directs {
		if w == avoid {
			continue
		}
		v, ok := mem[w]
		if !ok {
			v = paths(w, to, avoid, graph, mem)
			mem[w] = v
		}
		sum += v
	}
	mem[from] = sum
	return sum
}

func main() {
	graph := map[string][]string{}
	lib.ReadFileByLine(11, func(line string) {
		s := strings.Split(line, ": ")
		for _, w := range strings.Fields(s[1]) {
			graph[s[0]] = append(graph[s[0]], w)
		}
	})
	pathMem := map[string]int64{}
	p1 := paths("you", "out", "", graph, pathMem)
	lib.WritePart1("%d", p1)
	
	// p2 amounts to svr->fft->dac->out paths
	// + svr->dac->fft->out paths
	svr2fft := paths("svr", "fft", "dac", graph, map[string]int64{})
	fft2dac := paths("fft", "dac", "out", graph, map[string]int64{})
	dac2out := paths("dac", "out", "fft", graph, map[string]int64{})

	svr2dac := paths("svr", "dac", "fft", graph, map[string]int64{})
	dac2fft := paths("dac", "fft", "out", graph, map[string]int64{})
	fft2out := paths("fft", "out", "dac", graph, map[string]int64{})

	p2 := svr2fft * fft2dac * dac2out + svr2dac * dac2fft * fft2out
	lib.WritePart2("%d", p2)
}
