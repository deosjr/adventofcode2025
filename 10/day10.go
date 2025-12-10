package main

import (
	"math"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type machine struct {
	lights  int64
	buttons []int64
	joltage []int64
}

// TODO: no button should be pressed twice in a row
func (m machine) solve() int64 {
	if m.lights == 0 {
		return 0
	}
	set := map[int64]struct{}{0: struct{}{}}
	return m.solveBFS(1, set)
}

func (m machine) solveBFS(depth int64, set map[int64]struct{}) int64 {
	newset := map[int64]struct{}{}
	for k := range set {
		for _, b := range m.buttons {
			// flip bits in k according to mask b
			f := ^k & b
			v := k & ^b
			flipped := f | v
			if flipped == m.lights {
				return depth
			}
			newset[flipped] = struct{}{}
		}
	}
	return m.solveBFS(depth+1, newset)
}

func parseMachine(s string) machine {
	s1 := strings.Split(s, "] ")
	rawLights := s1[0][1:]
	lights := parseLights(s1[0][1:])
	s2 := strings.Split(s1[1], " {")
	buttons := parseButtons(s2[0], len(rawLights))
	joltage := parseJoltage(s2[1])
	return machine{lights, buttons, joltage}
}

// assumes bit is previously zero
func setBit(n int64, b int) int64 {
	return n + int64(math.Pow(2, float64(b)))
}

func parseLights(s string) int64 {
	var lights int64
	for i:=0; i<len(s); i++ {
		if s[i] == '.' {
			continue
		}
		lights = setBit(lights, len(s) - i - 1)
	}
	return lights
}

func parseButtons(s string, length int) []int64 {
	var buttons []int64
	for _, sb := range strings.Split(s, " ") {
		var button int64
		for _, sn := range strings.Split(sb[1:len(sb)-1], ",") {
			n := int(lib.MustParseInt(sn))
			button = setBit(button, length - n - 1)
		}
		buttons = append(buttons, button)
	}
	return buttons
}

func parseJoltage(s string) []int64 {
	return nil
}

func main() {
	//lib.Test()
	var machines []machine
	lib.ReadFileByLine(10, func(line string) {
		machines = append(machines, parseMachine(line))
	})
	var p1 int64
	for _, m := range machines {
		p1 += m.solve()
	}
	lib.WritePart1("%d", p1)
}
