package main

func Solve23() {
	var field [][]byte
	var s string
	for readLine(&s) {
		b := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			b[i] = s[i]
		}
		field = append(field, b)
	}

	nextSteps := func(p [2]int, symbol byte) [][2]int {
		var dirs [][2]int
		switch symbol {
		case '>':
			dirs = [][2]int{{0, 1}}
		case '<':
			dirs = [][2]int{{0, -1}}
		case 'v':
			dirs = [][2]int{{1, 0}}
		case '^':
			dirs = [][2]int{{-1, 0}}
		case '.':
			dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		}

		var nextSteps [][2]int
		for _, dir := range dirs {
			next := [2]int{p[0] + dir[0], p[1] + dir[1]}
			if next[0] < 0 || next[0] >= len(field) {
				continue
			}
			if next[1] < 0 || next[1] >= len(field[0]) {
				continue
			}
			if field[next[0]][next[1]] == '#' {
				continue
			}
			if field[next[0]][next[1]] == 'O' {
				continue
			}

			nextSteps = append(nextSteps, next)
		}

		return nextSteps
	}

	var dfs func(start [2]int, length int) int
	dfs = func(start [2]int, length int) int {
		prev := field[start[0]][start[1]]
		field[start[0]][start[1]] = 'O'
		defer func() { field[start[0]][start[1]] = prev }()

		if start[0] == len(field)-1 {
			return length
		}

		var longest int
		for _, next := range nextSteps(start, prev) {
			longest = max(longest, dfs(next, length+1))
		}
		return longest
	}

	for i := 0; i < len(field[0]); i++ {
		if field[0][i] == '#' {
			continue
		}
		print(dfs([2]int{0, i}, 0))
	}
}

func Solve23_2() {
	var field [][]byte
	var s string
	for readLine(&s) {
		b := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			if s[i] != '#' {
				b[i] = '.'
				continue
			}
			b[i] = s[i]
		}
		field = append(field, b)
	}

	nextSteps := func(p [2]int, visited map[[2]int]struct{}) [][2]int {
		var nextSteps [][2]int
		for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := [2]int{p[0] + dir[0], p[1] + dir[1]}
			if next[0] < 0 || next[0] >= len(field) {
				continue
			}
			if next[1] < 0 || next[1] >= len(field[0]) {
				continue
			}
			if field[next[0]][next[1]] == '#' {
				continue
			}
			if _, ok := visited[next]; ok {
				continue
			}

			nextSteps = append(nextSteps, next)
		}

		return nextSteps
	}

	var start, end [2]int
	for i := 0; i < len(field[0]); i++ {
		if field[0][i] != '#' {
			start = [2]int{0, i}
		}
		if field[len(field)-1][i] != '#' {
			end = [2]int{len(field) - 1, i}
		}
	}

	distances := make(map[[2]int]map[[2]int]int)
	queue := [][2]int{start}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if len(distances[top]) > 0 {
			continue
		}
		distances[top] = make(map[[2]int]int)

		v := make(map[[2]int]struct{})
		for _, next := range nextSteps(top, v) {
			v = map[[2]int]struct{}{
				top: {},
			}
			length := 0
			cur := next

			for {
				v[cur] = struct{}{}
				length++

				nexts := nextSteps(cur, v)
				if len(nexts) == 1 {
					cur = nexts[0]
					continue
				}

				distances[top][cur] = length
				queue = append(queue, cur)
				break
			}
		}
	}

	var dfs func(start [2]int, length int) int
	dfs = func(start [2]int, length int) int {
		if start == end {
			return length
		}

		if field[start[0]][start[1]] == 'O' {
			return 0
		}

		field[start[0]][start[1]] = 'O'
		var res int
		for next, dist := range distances[start] {
			res = max(res, dfs(next, length+dist))
		}
		field[start[0]][start[1]] = '.'
		return res
	}

	print(dfs(start, 0))
}
