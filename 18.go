package main

import (
	"math"
	"sort"
)

func Solve18() {
	type move struct {
		dir      string
		distance int64
		color    string
	}

	var moves []move
	var m move
	for read(&m.dir, &m.distance, &m.color) {
		moves = append(moves, m)
	}

	dirs := map[string][2]int64{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	addXY := func(a, b [2]int64) [2]int64 {
		return [2]int64{a[0] + b[0], a[1] + b[1]}
	}

	type part struct {
		start [2]int64
		len   int64
		done  bool
	}

	var topParts []part
	var cur [2]int64
	for _, m := range moves {
		dir := dirs[m.dir]
		dir = [2]int64{dir[0] * m.distance, dir[1] * m.distance}

		// horizontal part made
		if dir[0] == 0 {
			start := cur
			if dir[1] < 0 {
				start = addXY(cur, dir)
			}

			topParts = append(topParts, part{start: start, len: m.distance, done: false})
		}
		cur = addXY(cur, dir)
	}

	sort.Slice(topParts, func(i, j int) bool {
		if topParts[i].start[1] != topParts[j].start[1] {
			return topParts[i].start[1] < topParts[j].start[1]
		}
		return topParts[i].start[0] < topParts[j].start[0]
	})

	var res int64 = 0
	var curY int64 = 0
	var activeTopParts []part
	for {

		var nextY int64 = math.MaxInt64
		for i := 0; i < len(activeTopParts); i++ {
			nextY = min(nextY, activeTopParts[i].start[1]+activeTopParts[i].len+1)
		}

		start := sort.Search(len(topParts), func(i int) bool {
			return topParts[i].start[1] >= curY
		})

		for i := start; i < len(topParts); i++ {
			if topParts[i].start[1] > curY {
				nextY = min(nextY, topParts[i].start[1])
				break
			}

			nextY = min(nextY, topParts[i].start[1]+topParts[i].len)
			activeTopParts = append(activeTopParts, topParts[i])
		}

		if len(activeTopParts) == 0 {
			break
		}

		sort.Slice(activeTopParts, func(i, j int) bool {
			return activeTopParts[i].start[0] < activeTopParts[j].start[0]
		})

		for i := 1; i < len(activeTopParts); i = i + 2 {
			res += (1 + activeTopParts[i].start[0] - activeTopParts[i-1].start[0]) * (nextY - curY)
		}

		var newActiveTopParts []part
		for i := 0; i < len(activeTopParts); i++ {
			if activeTopParts[i].start[1]+activeTopParts[i].len > nextY {
				newActiveTopParts = append(newActiveTopParts, activeTopParts[i])
			}
		}
		activeTopParts = newActiveTopParts
		curY = nextY
	}
	print(res)
}
