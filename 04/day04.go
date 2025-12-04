package main

import "github.com/deosjr/adventofcode2025/lib"

type coord struct {
	x, y int
}

func getNeighbours(grid map[coord]bool) map[coord]int {
	neighbours := map[coord]int{}
	for c := range grid {
		x, y := c.x, c.y
		neighbours[coord{x-1, y}] += 1
		neighbours[coord{x+1, y}] += 1
		neighbours[coord{x, y-1}] += 1
		neighbours[coord{x, y+1}] += 1
		neighbours[coord{x-1, y-1}] += 1
		neighbours[coord{x-1, y+1}] += 1
		neighbours[coord{x+1, y-1}] += 1
		neighbours[coord{x+1, y+1}] += 1
	}
	return neighbours
}

func reduceGrid(grid map[coord]bool) map[coord]bool {
	newgrid := map[coord]bool{}
	neighbours := getNeighbours(grid)
	for c := range grid {
		if neighbours[c] < 4 {
			continue
		}
		newgrid[c] = true
	}
	return newgrid
}

func main() {
	y := 0
	grid := map[coord]bool{}
	lib.ReadFileByLine(4, func(line string) {
		for x, c := range line {
			if c == '.' {
				continue
			}
			grid[coord{x, y}] = true
		}
		y++
	})
	p1 := 0
	neighbours := getNeighbours(grid)
	for c := range grid {
		if neighbours[c] < 4 {
			p1++
		}
	}
	lib.WritePart1("%d", p1)
	p2 := 0
	for {
		g := reduceGrid(grid)
		diff := len(grid) - len(g)
		if diff == 0 {
			break
		}
		p2 += diff
		grid = g
	}
	lib.WritePart2("%d", p2)
}
