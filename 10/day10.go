package main

import (
	"math"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type machine struct {
	lights  int64
	buttons []button
	joltage []int64
}

type button struct {
	bits int64
	nums []int
}

// TODO: no button should be pressed twice in a row
// TODO: same config at a later depth can be ignored
func (m machine) solveP1() int64 {
	if m.lights == 0 {
		return 0
	}
	set := map[int64]struct{}{0: struct{}{}}
	return m.solveP1BFS(1, set)
}

func (m machine) solveP1BFS(depth int64, set map[int64]struct{}) int64 {
	newset := map[int64]struct{}{}
	for k := range set {
		for _, button := range m.buttons {
			// flip bits in k according to mask b
			b := button.bits
			f := ^k & b
			v := k & ^b
			flipped := f | v
			if flipped == m.lights {
				return depth
			}
			newset[flipped] = struct{}{}
		}
	}
	return m.solveP1BFS(depth+1, newset)
}

func (m machine) solveP2() int64 {
	init := make([]int64, len(m.joltage))
	if m.checkJoltage(init) {
		return 0
	}
	return m.solveP2BFS(1, [][]int64{init})
}

func (m machine) solveP2BFS(depth int64, fringe [][]int64) int64 {
	var newfringe [][]int64
	for _, list := range fringe {
	Loop:
		for _, button := range m.buttons {
			l := make([]int64, len(list))
			copy(l, list)
			for _, i := range button.nums {
				l[i] += 1
				if l[i] > m.joltage[i] {
					continue Loop
				}
			}
			if m.checkJoltage(l) {
				return depth
			}
			// TODO: check duplicates?
			newfringe = append(newfringe, l)
		}
	}
	return m.solveP2BFS(depth+1, newfringe)
}

func (m machine) checkJoltage(list []int64) bool {
	if len(m.joltage) != len(list) {
		return false
	}
	for i, j := range m.joltage {
		if list[i] != j {
			return false
		}
	}
	return true
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

func parseButtons(s string, length int) []button {
	var buttons []button
	for _, sb := range strings.Split(s, " ") {
		var bits int64
		var nums []int
		for _, sn := range strings.Split(sb[1:len(sb)-1], ",") {
			n := int(lib.MustParseInt(sn))
			nums = append(nums, n)
			bits = setBit(bits, length - n - 1)
		}
		buttons = append(buttons, button{bits, nums})
	}
	return buttons
}

func parseJoltage(s string) []int64 {
	var joltage []int64
	for _, sn := range strings.Split(strings.Trim(s, "}"), ",") {
		n := lib.MustParseInt(sn)
		joltage = append(joltage, n)
	}
	return joltage
}

func main() {
	//lib.Test()
	var machines []machine
	lib.ReadFileByLine(10, func(line string) {
		machines = append(machines, parseMachine(line))
	})
	var p1, p2 int64
	for _, m := range machines {
		p1 += m.solveP1()
		p2 += m.solveP2()
	}
	lib.WritePart1("%d", p1)
	lib.WritePart2("%d", p2)
}
