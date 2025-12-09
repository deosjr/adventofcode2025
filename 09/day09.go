package main

import (
	"image"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func main() {
	var tiles []image.Point
	var p1 int
	lib.ReadFileByLine(9, func(line string) {
		split := strings.Split(line, ",")
		x := lib.MustParseInt(split[0])
		y := lib.MustParseInt(split[1])
		tile := image.Pt(int(x), int(y))
		for _, t := range tiles {
			r := image.Rect(tile.X, tile.Y, t.X, t.Y)
			area := (r.Size().X + 1) * (r.Size().Y + 1)
			if area > p1 {
				p1 = area
			}
		}
		tiles = append(tiles, tile)
	})
	lib.WritePart1("%d", p1)
}
