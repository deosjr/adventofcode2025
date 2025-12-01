package main

import (
	"github.com/deosjr/adventofcode2024/lib" // todo: 2025
)

func posmod(n, m int64) int64 {
	return (n + m) % m
}

func main() {
	var dial int64 = 50
	p1 := 0
	lib.ReadFileByLine(1, func(line string) {
		n := lib.MustParseInt(line[1:]) // note utf8!
		n = posmod(n, 100)
		switch line[0] {
		case 'L':
			dial = posmod(dial - n, 100)
		case 'R':
			dial = posmod(dial + n, 100)
		}
		if dial == 0 {
			p1++
		}
	})
	lib.WritePart1("%d", p1)
}
