package main

import (
	"fmt"
	"strings"

	"github.com/deosjr/adventofcode2025/lib"
)

func main() {
	input := lib.ReadFile(12)
	segments := strings.Split(input, "\n\n")
	last := segments[len(segments)-1]
	regions := strings.Split(strings.TrimSpace(last), "\n")
	
	// incredibly lame but apparently good enough...
	sum := 0
	for _, s := range regions {
		var x, y, p1, p2, p3, p4, p5, p6 int
		fmt.Sscanf(s, "%dx%d: %d %d %d %d %d %d", &x, &y, &p1, &p2, &p3, &p4, &p5, &p6)
		pieces := p1 + p2 + p3 + p4 + p5 + p6
		fit := (x/3) * (y/3)
		if pieces <= fit {
			sum++
		}
	}

	lib.WritePart1("%d", sum)
}
