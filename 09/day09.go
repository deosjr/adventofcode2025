package main

import (
	"image"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

func main() {
	var tiles []image.Point
	var rects []image.Rectangle
	var p1 int
	lib.ReadFileByLine(9, func(line string) {
		split := strings.Split(line, ",")
		x := lib.MustParseInt(split[0])
		y := lib.MustParseInt(split[1])
		tile := image.Pt(int(x), int(y))
		for _, t := range tiles {
			r := image.Rect(tile.X, tile.Y, t.X+1, t.Y+1)
			area := r.Size().X * r.Size().Y
			if area > p1 {
				p1 = area
			}
			rects = append(rects, r)
		}
		tiles = append(tiles, tile)
	})
	lib.WritePart1("%d", p1)

	var lines []image.Rectangle
	from := tiles[len(tiles)-1]
	for _, to := range tiles {
		r := image.Rect(from.X, from.Y, to.X, to.Y)
		if from.X == to.X {
			r.Max.X += 1
		} else {
			r.Max.Y += 1
		}
		lines = append(lines, r)
		from = to
	}
	
	var p2 int
Loop:
	for _, r := range rects {
		area := r.Size().X * r.Size().Y
		if p2 >= area {
			continue
		}
		inset := r.Inset(1)
		for _, line := range lines {
			if inset.Overlaps(line) {
				continue Loop
			}
		}
		if area > p2 {
			p2 = area
		}
	}
	lib.WritePart2("%d", p2)
}
