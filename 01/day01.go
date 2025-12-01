package main

import "github.com/deosjr/adventofcode2025/lib"

// negative modulo always trips me up in Go
func posmod(n, m int64) int64 {
	return (n + m) % m
}

func main() {
	var dial int64 = 50
	var p1, p2 int64 = 0, 0
	lib.ReadFileByLine(1, func(line string) {
		n := lib.MustParseInt(line[1:]) // note utf8!
		p2 += n / 100
		n = posmod(n, 100)
		prev := dial
		switch line[0] {
		case 'L':
			dial = posmod(dial - n, 100)
			if prev != 0 && n > prev && dial != 0 {
				p2++
			}
		case 'R':
			dial = posmod(dial + n, 100)
			if prev != 0 && dial < prev && dial != 0 {
				p2++
			}
		}
		if dial == 0 {
			p1++
			p2++
		}
	})
	lib.WritePart1("%d", p1)
	lib.WritePart2("%d", p2)
}
