package main

func Solve16() {
	type move struct {
		from [2]int
		dir  [2]int
	}

	var ss []string
	var s string
	for readLine(&s) {
		ss = append(ss, s)
	}

	addXY := func(a, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}

	nextMoves := func(m move) []move {
		b := ss[m.from[0]][m.from[1]]

		var newDirs [][2]int
		switch {
		case b == '|' && m.dir[0] == 0:
			newDirs = [][2]int{{-1, 0}, {1, 0}}
		case b == '-' && m.dir[1] == 0:
			newDirs = [][2]int{{0, -1}, {0, 1}}
		case b == '/':
			newDirs = [][2]int{{m.dir[1] * -1, m.dir[0] * -1}}
		case b == '\\':
			newDirs = [][2]int{{m.dir[1], m.dir[0]}}
		default:
			newDirs = [][2]int{m.dir}
		}

		var nextMoves []move
		for _, newDir := range newDirs {
			nextMoves = append(nextMoves, move{from: addXY(m.from, newDir), dir: newDir})
		}
		return nextMoves
	}

	seen := make(map[move]struct{})
	var dfs func(m move)
	dfs = func(m move) {
		if m.from[0] < 0 || m.from[0] >= len(ss) {
			return
		}
		if m.from[1] < 0 || m.from[1] >= len(ss[0]) {
			return
		}
		if _, ok := seen[m]; ok {
			return
		}
		seen[m] = struct{}{}

		for _, nextMove := range nextMoves(m) {
			dfs(nextMove)
		}
	}

	dfs(move{from: [2]int{0, 0}, dir: [2]int{0, 1}})

	energized := make(map[[2]int]struct{})
	for k := range seen {
		energized[k.from] = struct{}{}
	}
	print(len(energized))
}

func Solve16_2() {
	type move struct {
		from [2]int
		dir  [2]int
	}

	var ss []string
	var s string
	for readLine(&s) {
		ss = append(ss, s)
	}

	addXY := func(a, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}

	nextMoves := func(m move) []move {
		b := ss[m.from[0]][m.from[1]]

		var newDirs [][2]int
		switch {
		case b == '|' && m.dir[0] == 0:
			newDirs = [][2]int{{-1, 0}, {1, 0}}
		case b == '-' && m.dir[1] == 0:
			newDirs = [][2]int{{0, -1}, {0, 1}}
		case b == '/':
			newDirs = [][2]int{{m.dir[1] * -1, m.dir[0] * -1}}
		case b == '\\':
			newDirs = [][2]int{{m.dir[1], m.dir[0]}}
		default:
			newDirs = [][2]int{m.dir}
		}

		var nextMoves []move
		for _, newDir := range newDirs {
			nextMoves = append(nextMoves, move{from: addXY(m.from, newDir), dir: newDir})
		}
		return nextMoves
	}

	seen := make(map[move]struct{})
	var dfs func(m move)
	dfs = func(m move) {
		if m.from[0] < 0 || m.from[0] >= len(ss) {
			return
		}
		if m.from[1] < 0 || m.from[1] >= len(ss[0]) {
			return
		}
		if _, ok := seen[m]; ok {
			return
		}
		seen[m] = struct{}{}

		for _, nextMove := range nextMoves(m) {
			dfs(nextMove)
		}
	}

	var best int
	for i := 0; i < len(ss); i++ {
		for j := 0; j < len(ss[0]); j++ {
			var nextDirs [][2]int

			if i == 0 {
				nextDirs = append(nextDirs, [2]int{1, 0})
			}
			if i == len(ss)-1 {
				nextDirs = append(nextDirs, [2]int{-1, 0})
			}
			if j == 0 {
				nextDirs = append(nextDirs, [2]int{0, 1})
			}
			if j == len(ss[0])-1 {
				nextDirs = append(nextDirs, [2]int{0, -1})
			}

			for _, nextDir := range nextDirs {
				seen = make(map[move]struct{})
				dfs(move{from: [2]int{i, j}, dir: nextDir})
				energized := make(map[[2]int]struct{})
				for k := range seen {
					energized[k.from] = struct{}{}
				}
				best = max(best, len(energized))
			}
		}
	}
	print(best)
}
