package main

import (
	"math"
	"strings"
	"github.com/deosjr/adventofcode2025/lib"
)

type IDrange struct {
	from, to int64
	fromHigh, toHigh int64
	fromLow, toLow int64
	pow int64
	evenLen int
	unevenLen int
	unevenFrom, unevenTo int64
}

// observation: ranges length diff between from/to is max 1
// observation: 10 digits is the max for any number
// p1 checks all even repetitions, p2 adds uneven for length 3, 5 and 7
// 3 has to be checked in ranges of len 3, 6 and 9, 5 in 5 and 10, and 7 in 7 only
func parseIDrange(s string) IDrange {
	split := strings.Split(strings.TrimSpace(s), "-")
	from := lib.MustParseInt(split[0])
	to := lib.MustParseInt(split[1])
	r := IDrange{from:from, to:to}
	fromLen, toLen := len(split[0]), len(split[1])
	fromEven, toEven := fromLen % 2 == 0, toLen % 2 == 0
	if !fromEven && !toEven {
		r.unevenFrom = from
		r.unevenTo = to
		r.unevenLen = fromLen
		return r
	}
	if !fromEven {
		r.unevenFrom = from
		r.unevenTo = int64(math.Pow10(fromLen)) - 1
		r.unevenLen = fromLen
		from = int64(math.Pow10(fromLen))
		r.evenLen = toLen
	}
	if !toEven {
		r.unevenFrom = int64(math.Pow10(fromLen))
		r.unevenTo = to
		r.unevenLen = toLen
		to = int64(math.Pow10(fromLen)) - 1
		r.evenLen = fromLen
	}
	if fromEven && toEven {
		r.evenLen = fromLen
	}
	cutOff := int64(math.Pow10(r.evenLen/2))
	r.fromHigh = from / cutOff
	r.toHigh = to / cutOff
	r.fromLow = from % cutOff
	r.toLow = to % cutOff
	r.pow = cutOff
	return r
}

func pasteNum(high, low, pow int64) int64 {
	return high * pow + low
}

func (r IDrange) findEvenInvalidIDs() []int64 {
	ids := []int64{}
	if r.evenLen == 0 {
		return ids
	}
	if r.fromHigh == r.toHigh {
		if r.fromLow <= r.fromHigh && r.toLow >= r.fromHigh {
			ids = append(ids, pasteNum(r.fromHigh, r.fromHigh, r.pow))
		} 
		return ids
	}
	if r.fromLow <= r.fromHigh {
		ids = append(ids, pasteNum(r.fromHigh, r.fromHigh, r.pow))
	}
	for n:=r.fromHigh+1; n<r.toHigh; n++ {
		ids = append(ids, pasteNum(n, n, r.pow))
	}
	if r.toLow >= r.toHigh {
		ids = append(ids, pasteNum(r.toHigh, r.toHigh, r.pow))
	}
	return ids
}

func (r IDrange) find3In3() []int64 {
	ids := []int64{}
	for _, n := range []int64{ 111, 222, 333, 444, 555, 666, 777, 888, 999 } {
		if n < r.unevenFrom || n > r.unevenTo {
			continue
		}
		ids = append(ids, n)
	}
	return ids
}

func (r IDrange) find3In6() []int64 {
	ids := []int64{}
	fromHigh := r.fromHigh / 10
	toHigh := r.toHigh / 10
	targetFunc := func(n int64) int64 {
		return n * 10000 + n * 100 + n
	}
	from := pasteNum(r.fromHigh, r.fromLow, r.pow)
	to := pasteNum(r.toHigh, r.toLow, r.pow)

	if fromHigh == toHigh {
		target := targetFunc(fromHigh)
		if target >= from && target <= to {
			ids = append(ids, target)
		}
		return ids
	}
	if target := targetFunc(fromHigh); target >= from {
		ids = append(ids, target)
	}
	for n := fromHigh+1; n<toHigh; n++ {
		ids = append(ids, targetFunc(n))
	}
	if target := targetFunc(toHigh); target <= to {
		ids = append(ids, target)
	}
	return ids
}

func (r IDrange) find3In9() []int64 {
	ids := []int64{}
	fromHigh := r.unevenFrom / 1000000
	toHigh := r.unevenTo / 1000000
	targetFunc := func(n int64) int64 {
		return n * 1000000 + n * 1000 + n
	}
	from := r.unevenFrom
	to := r.unevenTo

	if fromHigh == toHigh {
		target := targetFunc(fromHigh)
		if target >= from && target <= to {
			ids = append(ids, target)
		}
		return ids
	}
	if target := targetFunc(fromHigh); target >= from {
		ids = append(ids, target)
	}
	for n := fromHigh+1; n<toHigh; n++ {
		ids = append(ids, targetFunc(n))
	}
	if target := targetFunc(toHigh); target <= to {
		ids = append(ids, target)
	}
	return ids
}

func (r IDrange) find5In5() []int64 {
	ids := []int64{}
	for _, n := range []int64{ 11111, 22222, 33333, 44444, 55555, 66666, 77777, 88888, 99999 } {
		if n < r.unevenFrom || n > r.unevenTo {
			continue
		}
		ids = append(ids, n)
	}
	return ids
}

func (r IDrange) find5In10() []int64 {
	ids := []int64{}
	subr := IDrange{
		evenLen: 4,
		fromHigh: r.fromHigh / 1000,
		fromLow: (r.fromHigh / 10) % 100,
		toHigh: r.toHigh / 1000,
		toLow: (r.toHigh / 10) % 100,
		pow: 100,
	}
	for _, dupe := range subr.findEvenInvalidIDs() {
		n := dupe % 100
		target := dupe * 1000000 + dupe * 100 + n
		if target < r.from || target > r.to {
			continue
		}
		ids = append(ids, target)
	}
	return ids
}

func (r IDrange) find7In7() []int64 {
	ids := []int64{}
	for _, n := range []int64{ 1111111, 2222222, 3333333, 4444444, 5555555, 6666666, 7777777, 8888888, 9999999 } {
		if n < r.unevenFrom || n > r.unevenTo {
			continue
		}
		ids = append(ids, n)
	}
	return ids
}

func main() {
	in := lib.ReadFile(2)
	ranges := []IDrange{}
	for _, s := range strings.Split(in, ",") {
		ranges = append(ranges, parseIDrange(s))
	}
	var p1, p2 int64
	for _, r := range ranges {
		p2set := map[int64]struct{}{}
		invalids := r.findEvenInvalidIDs()
		for _, id := range invalids {
			p1 += id
			p2set[id] = struct{}{}
		}
		if r.unevenLen == 3 {
			for _, id := range r.find3In3() {
				p2set[id] = struct{}{}
			}
		}
		if r.evenLen == 6 {
			for _, id := range r.find3In6() {
				p2set[id] = struct{}{}
			}
		}
		if r.unevenLen == 9 {
			for _, id := range r.find3In9() {
				p2set[id] = struct{}{}
			}
		}
		if r.unevenLen == 5 {
			for _, id := range r.find5In5() {
				p2set[id] = struct{}{}
			}
		}
		if r.evenLen == 10 {
			for _, id := range r.find5In10() {
				p2set[id] = struct{}{}
			}
		}
		if r.unevenLen == 7 {
			for _, id := range r.find7In7() {
				p2set[id] = struct{}{}
			}
		}
		for id, _ := range p2set {
			p2 += id
		}
	}
	lib.WritePart1("%d", p1)
	lib.WritePart2("%d", p2)
}
