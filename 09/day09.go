package main

import (
	"math"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type coord struct {
	x, y int64
}

func main() {
	var tiles []coord
	var p1 int64
	lib.ReadFileByLine(9, func(line string) {
		split := strings.Split(line, ",")
		x := lib.MustParseInt(split[0])
		y := lib.MustParseInt(split[1])
		tile := coord{x, y}
		for _, t := range tiles {
			dx := int64(math.Abs(float64(tile.x - t.x))) + 1
			dy := int64(math.Abs(float64(tile.y - t.y))) + 1
			if area := dx * dy; area > p1 {
				p1 = area
			}
		}
		tiles = append(tiles, tile)
	})
	lib.WritePart1("%d", p1)
}
