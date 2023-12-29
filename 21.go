package main

import "golang.org/x/exp/maps"

func Solve21() {
	var s string
	var field [][]byte
	var start [2]int
	for readLine(&s) {
		b := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			b[i] = s[i]
			if s[i] == 'S' {
				start = [2]int{len(field), i}
			}
		}
		field = append(field, b)
	}

	valid := func(pos [2]int) bool {
		if pos[0] < 0 || pos[0] >= len(field) {
			return false
		}
		if pos[1] < 0 || pos[1] >= len(field[0]) {
			return false
		}
		return field[pos[0]][pos[1]] != '#'
	}

	queue := [][2]int{start}
	for i := 0; i < 64; i++ {
		newQueue := make(map[[2]int]struct{})
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]

			for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				newPos := [2]int{top[0] + dir[0], top[1] + dir[1]}
				if !valid(newPos) {
					continue
				}

				newQueue[newPos] = struct{}{}
			}
		}

		queue = append(queue, maps.Keys(newQueue)...)
	}

	print(len(queue))
}

func Solve21_2() {
	var s string
	var field [][]byte
	var start [2]int
	for readLine(&s) {
		b := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			b[i] = s[i]
			if s[i] == 'S' {
				start = [2]int{len(field), i}
			}
		}
		field = append(field, b)
	}

	normalize := func(a [2]int) [2]int {
		mod := [2]int{a[0] % len(field), a[1] % len(field[0])}
		if mod[0] < 0 {
			mod[0] += len(field)
		}
		if mod[1] < 0 {
			mod[1] += len(field[0])
		}
		return mod
	}

	getPeriod := func(res []int64) (int, int) {
		for i := 3; i <= len(res)/3; i++ {
			s1, s2, s3 := len(res)-3*i, len(res)-2*i, len(res)-i
			for j := 0; j <= i; j++ {
				if j == i {
					return s1, i
				}
				if res[s2+j]-res[s1+j] != res[s3+j]-res[s2+j] {
					break
				}
			}
		}
		return 0, 0
	}

	var res, diffs []int64
	res = append(res, 0)
	var periodStart, period int

	queue := [][2]int{start}
	for i := 0; periodStart == 0; i++ {
		newQueue := make(map[[2]int]struct{})
		for len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]

			for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				newPos := [2]int{top[0] + dir[0], top[1] + dir[1]}
				normalized := normalize(newPos)
				if field[normalized[0]][normalized[1]] == '#' {
					continue
				}

				newQueue[newPos] = struct{}{}
			}
		}

		queue = append(queue, maps.Keys(newQueue)...)
		diffs = append(diffs, int64(len(queue))-res[len(res)-1])
		res = append(res, int64(len(queue)))
		periodStart, period = getPeriod(diffs)
	}

	var num int64 = 26501365
	shifted := num - int64(periodStart)
	mod := shifted % int64(period)
	times := shifted / int64(period)

	var increases []int64
	var increasesSum int64
	for i := 0; i < period; i++ {
		increase := diffs[periodStart+i+period] - diffs[periodStart+i]
		increases = append(increases, increase)
		increasesSum += increase
	}

	var result = res[periodStart]
	var shiftAdd = res[periodStart+period] - res[periodStart]
	for i := 0; i < int(times); i++ {
		result += shiftAdd
		shiftAdd += increasesSum
	}

	for i := 0; int64(i) < mod; i++ {
		result += diffs[periodStart+i] + increases[i]*times
	}

	print(result)
}
