package main

import (
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func paths(from string, graph map[string][]string, mem map[string]int64) int64 {
	if v, ok := mem[from]; ok {
		return v
	}
	directs := graph[from]
	if len(directs) == 1 && directs[0] == "out" {
		mem[from] = 1
		return 1
	}
	var sum int64
	for _, w := range graph[from] {	
		v, ok := mem[w]
		if !ok {
			v = paths(w, graph, mem)
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
	p1 := paths("you", graph, pathMem)
	lib.WritePart1("%d", p1)
}
