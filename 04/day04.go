package main

import "github.com/deosjr/adventofcode2025/lib"

type coord struct {
	x, y int
}

func main() {
	//lib.Test()
	y := 0
	grid := map[coord]bool{}
	neighbours := map[coord]int{}
	lib.ReadFileByLine(4, func(line string) {
		for x, c := range line {
			if c == '.' {
				continue
			}
			grid[coord{x, y}] = true
			neighbours[coord{x-1, y}] += 1
			neighbours[coord{x+1, y}] += 1
			neighbours[coord{x, y-1}] += 1
			neighbours[coord{x, y+1}] += 1
			neighbours[coord{x-1, y-1}] += 1
			neighbours[coord{x-1, y+1}] += 1
			neighbours[coord{x+1, y-1}] += 1
			neighbours[coord{x+1, y+1}] += 1
		}
		y++
	})
	p1 := 0
	for c := range grid {
		if neighbours[c] < 4 {
			p1++
		}
	}
	lib.WritePart1("%d", p1)
}
