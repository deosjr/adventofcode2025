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
}

// observation: ranges length diff between from/to is max 1
func parseIDrange(s string) IDrange {
	split := strings.Split(strings.TrimSpace(s), "-")
	from := lib.MustParseInt(split[0])
	to := lib.MustParseInt(split[1])
	r := IDrange{from:from, to:to}
	fromLen, toLen := len(split[0]), len(split[1])
	fromEven, toEven := fromLen % 2 == 0, toLen % 2 == 0
	if !fromEven && !toEven {
		return r
	}
	if !fromEven {
		from = int64(math.Pow10(fromLen))
		fromLen++
	}
	if !toEven {
		to = int64(math.Pow10(fromLen)) - 1
	}
	cutOff := int64(math.Pow10(fromLen/2))
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

func (r IDrange) findInvalidIDs() []int64 {
	ids := []int64{}
	if r.fromHigh == 0 {
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
	lib.WritePart2("%v", r)
	lib.WritePart2("%d", ids)
	return ids
}

func main() {
	//lib.Test()
	in := lib.ReadFile(2)
	ranges := []IDrange{}
	for _, s := range strings.Split(in, ",") {
		ranges = append(ranges, parseIDrange(s))
	}
	var p1 int64
	for _, r := range ranges {
		invalids := r.findInvalidIDs()
		for _, n := range invalids {
			p1 += n
		}
	}
	lib.WritePart1("%d", p1)
}
