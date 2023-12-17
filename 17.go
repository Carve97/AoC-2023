package main

import "container/heap"

func Solve17() {
	var m [][]int
	var s string
	for i := 0; readLine(&s); i++ {
		vals := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			vals[j] = int(s[j] - '0')
		}
		m = append(m, vals)
	}

	type move struct {
		tile       [2]int
		dir        [2]int
		sameDirCnt int
	}

	addXY := func(a [2]int, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}

	validTile := func(t [2]int) bool {
		if t[0] < 0 || t[0] >= len(m) {
			return false
		}
		if t[1] < 0 || t[1] >= len(m[0]) {
			return false
		}
		return true
	}

	priorityQ := intHeap[move]{
		slice: []keyVal[int, move]{{
			key: 0,
			val: move{
				tile:       [2]int{0, 0},
				dir:        [2]int{0, 0},
				sameDirCnt: 0,
			},
		}},
	}
	heap.Init(&priorityQ)

	seen := make(map[move]struct{})

	for v := heap.Pop(&priorityQ); v != nil; v = heap.Pop(&priorityQ) {
		curMove := v.(keyVal[int, move])

		// bottom-right corner reached
		if curMove.val.tile == [2]int{len(m) - 1, len(m[0]) - 1} {
			print(curMove.key)
			return
		}

		if _, ok := seen[curMove.val]; ok {
			continue
		}
		seen[curMove.val] = struct{}{}

		for _, nextDir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextTile := addXY(curMove.val.tile, nextDir)
			if !validTile(nextTile) {
				continue
			}

			// prevent going back
			if curMove.val.dir == [2]int{nextDir[0] * -1, nextDir[1] * -1} {
				continue
			}

			sameDirCount := 0
			if curMove.val.dir == nextDir {
				sameDirCount = curMove.val.sameDirCnt + 1
			}

			if sameDirCount >= 3 {
				continue
			}

			nextHeat := curMove.key + m[nextTile[0]][nextTile[1]]
			nextMove := move{
				tile:       nextTile,
				dir:        nextDir,
				sameDirCnt: sameDirCount,
			}

			heap.Push(&priorityQ, keyVal[int, move]{key: nextHeat, val: nextMove})
		}
	}
}

func Solve17_2() {
	var m [][]int
	var s string
	for i := 0; readLine(&s); i++ {
		vals := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			vals[j] = int(s[j] - '0')
		}
		m = append(m, vals)
	}

	type move struct {
		tile       [2]int
		dir        [2]int
		sameDirCnt int
	}

	addXY := func(a [2]int, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}

	validTile := func(t [2]int) bool {
		if t[0] < 0 || t[0] >= len(m) {
			return false
		}
		if t[1] < 0 || t[1] >= len(m[0]) {
			return false
		}
		return true
	}

	priorityQ := intHeap[move]{}
	heap.Init(&priorityQ)

	// add the initial moves in both possible directions (down and right)
	heap.Push(&priorityQ, keyVal[int, move]{key: 0, val: move{tile: [2]int{0, 0}, dir: [2]int{1, 0}, sameDirCnt: 0}})
	heap.Push(&priorityQ, keyVal[int, move]{key: 0, val: move{tile: [2]int{0, 0}, dir: [2]int{0, 1}, sameDirCnt: 0}})

	seen := make(map[move]struct{})

	for v := heap.Pop(&priorityQ); v != nil; v = heap.Pop(&priorityQ) {
		curMove := v.(keyVal[int, move])

		// bottom-right corner reached
		if curMove.val.tile == [2]int{len(m) - 1, len(m[0]) - 1} && curMove.val.sameDirCnt >= 4 {
			print(curMove.key)
			return
		}

		if _, ok := seen[curMove.val]; ok {
			continue
		}
		seen[curMove.val] = struct{}{}

		for _, nextDir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextTile := addXY(curMove.val.tile, nextDir)
			if !validTile(nextTile) {
				continue
			}

			// prevent going back
			if curMove.val.dir == [2]int{nextDir[0] * -1, nextDir[1] * -1} {
				continue
			}

			sameDirCnt := 1
			if curMove.val.dir == nextDir {
				sameDirCnt = curMove.val.sameDirCnt + 1
			}

			if sameDirCnt > 10 {
				continue
			}

			if curMove.val.dir != nextDir && curMove.val.sameDirCnt < 4 {
				continue
			}

			nextHeat := curMove.key + m[nextTile[0]][nextTile[1]]
			nextMove := move{
				tile:       nextTile,
				dir:        nextDir,
				sameDirCnt: sameDirCnt,
			}

			heap.Push(&priorityQ, keyVal[int, move]{key: nextHeat, val: nextMove})
		}
	}
}
