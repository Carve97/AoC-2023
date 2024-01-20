package main

import (
	"fmt"
	"math"
	"strings"
)

func Solve24() {
	type hailstone struct {
		pos   [3]float64
		speed [3]float64
	}

	var hs []hailstone
	var s string
	for readLine(&s) {
		h := hailstone{}
		fmt.Fscanf(strings.NewReader(s), "%f, %f, %f @ %f, %f, %f",
			&h.pos[0], &h.pos[1], &h.pos[2], &h.speed[0], &h.speed[1], &h.speed[2])
		hs = append(hs, h)
	}

	calcIntersectionTime := func(h1, h2 hailstone) (float64, float64) {
		xdiff := h2.pos[0] - h1.pos[0]
		normXDiff := xdiff / h1.speed[0]
		normJSpeed := h2.speed[0] / h1.speed[0]

		nums := h1.pos[1] + h1.speed[1]*normXDiff - h2.pos[1]
		vs := h2.speed[1] - h1.speed[1]*normJSpeed

		v := nums / vs
		u := normXDiff + normJSpeed*v

		return u, v
	}

	equalFloats := func(a, b float64) bool {
		return math.Abs(a-b) < 0.0000001
	}

	var from, to int64 = 200000000000000, 400000000000000
	var res int64
	for i := 0; i < len(hs); i++ {
		for j := i + 1; j < len(hs); j++ {
			if equalFloats(hs[i].speed[0]/hs[j].speed[0], hs[i].speed[1]/hs[j].speed[1]) {
				continue
			}

			u, v := calcIntersectionTime(hs[i], hs[j])

			if u < 0 || v < 0 {
				continue
			}

			intX := hs[j].pos[0] + v*hs[j].speed[0]
			intY := hs[j].pos[1] + v*hs[j].speed[1]

			if intX < float64(from) || intX > float64(to) {
				continue
			}
			if intY < float64(from) || intY > float64(to) {
				continue
			}

			res++
		}
	}

	print(res)
}
