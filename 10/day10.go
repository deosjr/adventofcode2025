package main

import (
	"math"
	"sort"
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

// NOTE: order does _not_ matter for part 2
func (m machine) solveP2() int64 {
	init := make([]int64, len(m.joltage))
	occ := occurances(m.buttons)
	// try buttons in order, pressing them 0-300 times each
	res := m.solveP2rec(init, 0, 0, math.MaxInt64, occ)
	if res == math.MaxInt64 {
		panic("failed to find")
	}
	return res
}

func (m machine) solveP2rec(list []int64, buttonIdx int, total, best int64, occ map[int]int) int64 {
	if total > best {
		return best
	}
	if m.checkJoltage(list) {
		return total
	}
	if buttonIdx == len(m.buttons) {
		return best
	}

	button := m.buttons[buttonIdx]

	// if a button is the only one left for a joltage index, set, dont loop!
	// if multiple nums are the only one left, we only have to try one of them
	for _, n := range button.nums {
		if occ[n] > 1 {
			continue
		}
		if occ[n] < 1 {
			panic("smth has gone wrong")
		}
		// this button is the last one to toggle joltage number n
		presses := m.joltage[n] - list[n]
		if ok := m.updateList(list, list, button, int64(presses)); !ok {
			return best
		}
		useButton(occ, button)
		res := m.solveP2rec(list, buttonIdx+1, total+int64(presses), best, occ)
		undoButton(occ, button)
		if res < best {
			return res
		}
		return best
	}
	

	// otherwise, try in a loop up until max joltage (~300)
	limit := 300
	l := make([]int64, len(list))
	useButton(occ, button)
	for i:=0; i<limit; i++ {
		if ok := m.updateList(l, list, button, int64(i)); !ok {
			undoButton(occ, button)
			return best
		}
		res := m.solveP2rec(l, buttonIdx+1, total+int64(i), best, occ)
		if res < best {
			best = res
		}
	}
	undoButton(occ, button)
	return best
}

func (m machine) updateList(dst, src []int64, button button, numPresses int64) bool {
	copy(dst, src)
	for _, n := range button.nums {
		v := src[n] + numPresses
		if v > m.joltage[n] {
			return false
		}
		dst[n] = v
	}
	return true
}

func useButton(occ map[int]int, button button) {
	for _, n := range button.nums {
		occ[n] -= 1
	}
}

func undoButton(occ map[int]int, button button) {
	for _, n := range button.nums {
		occ[n] += 1
	}
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

func occurances(buttons []button) map[int]int {
	occ := map[int]int{}
	for _, b := range buttons {
		for _, n := range b.nums {
			occ[n] += 1
		}
	}
	return occ
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
	occ := occurances(buttons)
	sort.Slice(buttons, func(i, j int) bool {
		bi := buttons[i].nums
		mini := len(buttons)
		for _, n := range bi {
			if occ[n] < mini {
				mini = occ[n]
			}
		}
		bj := buttons[j].nums
		minj := len(buttons)
		for _, n := range bj {
			if occ[n] < minj {
				minj = occ[n]
			}
		}
		return mini < minj
	})
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
		lib.WritePart1("%v", m)
		p1 += m.solveP1()
		p2 += m.solveP2()
	}
	lib.WritePart1("%d", p1)
	lib.WritePart2("%d", p2)
}
