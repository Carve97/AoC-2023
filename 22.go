package main

import (
	"fmt"
	"sort"
	"strings"
)

func Solve22() {
	type brick struct {
		from [3]int
		to   [3]int
	}

	var bricks []brick
	var s string
	var mx, my, mz int
	for readLine(&s) {
		var a, b [3]int
		fmt.Fscanf(strings.NewReader(s), "%d,%d,%d~%d,%d,%d",
			&a[0], &a[1], &a[2],
			&b[0], &b[1], &b[2])
		bricks = append(bricks, brick{from: a, to: b})
		mx = max(mx, b[0])
		my = max(my, b[1])
		mz = max(mz, a[2])
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].from[2] < bricks[j].from[2]
	})

	field := make([][][]int16, mz+1)
	for z := 0; z <= mz; z++ {
		field[z] = make([][]int16, mx+1)
		for x := 0; x <= mx; x++ {
			field[z][x] = make([]int16, my+1)
		}
	}

	supportFor := make([]map[int16]struct{}, len(bricks)+1)
	supportedBy := make([]map[int16]struct{}, len(bricks)+1)
	for i := range supportFor {
		supportFor[i] = make(map[int16]struct{})
		supportedBy[i] = make(map[int16]struct{})
	}

	supporting := make(map[int16]struct{})
	isOnBrick := func(b brick, z int, idx int16) bool {
		for x := b.from[0]; x <= b.to[0]; x++ {
			if onBrick := field[z][x][b.from[1]]; onBrick != 0 {
				supportedBy[idx][onBrick] = struct{}{}
				supportFor[onBrick][idx] = struct{}{}
			}
		}
		for y := b.from[1]; y <= b.to[1]; y++ {
			if onBrick := field[z][b.from[0]][y]; onBrick != 0 {
				supportedBy[idx][onBrick] = struct{}{}
				supportFor[onBrick][idx] = struct{}{}
			}
		}

		if len(supportedBy[idx]) == 1 {
			for k := range supportedBy[idx] {
				supporting[k] = struct{}{}
			}
		}

		return len(supportedBy[idx]) > 0
	}

	for i, brick := range bricks {
		zf, zt := brick.from[2], brick.to[2]
		for zf > 1 {
			if isOnBrick(brick, zf-1, int16(i+1)) {
				break
			}

			zf--
			zt--
		}

		for z := zf; z <= zt; z++ {
			for x := brick.from[0]; x <= brick.to[0]; x++ {
				for y := brick.from[1]; y <= brick.to[1]; y++ {
					field[z][x][y] = int16(i + 1)
				}
			}
		}
	}

	willFall := func(brickIdx int16, gone map[int16]struct{}) bool {
		for supportBrick := range supportedBy[brickIdx] {
			if _, isGone := gone[supportBrick]; !isGone {
				return false
			}
		}
		return true
	}

	var p2Res int64
	for k := range supporting {
		queue := []int16{k}
		var cur int64
		gone := make(map[int16]struct{})
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]
			gone[top] = struct{}{}
			for brick := range supportFor[top] {
				if willFall(brick, gone) {
					cur++
					queue = append(queue, brick)
				}
			}
		}
		p2Res += cur
	}

	print("Part 1: ", len(bricks)-len(supporting))
	print("Part 2: ", p2Res)
}
