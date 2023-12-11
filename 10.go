package main

func Solve10() {
	var a [][]byte
	var s [2]int

	var x string
	for readLine(&x) {
		b := []byte(x)
		for i := 0; i < len(x); i++ {
			if b[i] == 'S' {
				s = [2]int{len(a), i}
			}
		}
		a = append(a, b)
	}

	m := map[byte][2][2]int{
		'|': {{1, 0}, {-1, 0}},
		'-': {{0, 1}, {0, -1}},
		'L': {{-1, 0}, {0, 1}},
		'J': {{-1, 0}, {0, -1}},
		'F': {{1, 0}, {0, 1}},
		'7': {{1, 0}, {0, -1}},
	}

	addXY := func(a [2]int, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}

	valid := func(idx [2]int) bool {
		if idx[0] < 0 || idx[0] >= len(a) {
			return false
		}
		if idx[1] < 0 || idx[1] >= len(a[0]) {
			return false
		}
		return true
	}

	move := func(from [2]int) [][2]int {
		var res [][2]int
		dirs := m[a[from[0]][from[1]]]
		for _, dir := range dirs {
			if !valid(addXY(from, dir)) {
				continue
			}
			res = append(res, addXY(from, dir))
		}
		return res
	}

	validMove := func(from [2]int, to [2]int) bool {
		moves := move(to)
		for _, move := range moves {
			if move == from {
				return true
			}
		}
		return false
	}

	for k, v := range m {
		next1 := addXY(s, v[0])
		next2 := addXY(s, v[1])
		if !valid(next1) || !validMove(s, next1) {
			continue
		}
		if !valid(next2) || !validMove(s, next2) {
			continue
		}
		a[s[0]][s[1]] = k
	}

	v := make(map[[2]int]struct{})
	var res int64 = 0
	var dfs func([2]int)
	dfs = func(start [2]int) {
		if _, visited := v[start]; visited {
			return
		}
		res++
		v[start] = struct{}{}
		moves := move(start)
		for _, mov := range moves {
			dfs(mov)
		}
	}
	dfs(s)
	print(res / 2)
}

func Solve10_2() {
	var a [][]byte
	var s [2]int

	var x string
	for readLine(&x) {
		b := []byte(x)
		for i := 0; i < len(x); i++ {
			if b[i] == 'S' {
				s = [2]int{len(a), i}
			}
		}
		a = append(a, b)
	}

	m := map[byte][2][2]int{
		'|': {{1, 0}, {-1, 0}},
		'-': {{0, 1}, {0, -1}},
		'L': {{-1, 0}, {0, 1}},
		'J': {{-1, 0}, {0, -1}},
		'F': {{1, 0}, {0, 1}},
		'7': {{1, 0}, {0, -1}},
	}

	addXY := func(a [2]int, b [2]int) [2]int {
		return [2]int{a[0] + b[0], a[1] + b[1]}
	}

	valid := func(idx [2]int) bool {
		if idx[0] < 0 || idx[0] >= len(a) {
			return false
		}
		if idx[1] < 0 || idx[1] >= len(a[0]) {
			return false
		}
		return true
	}

	move := func(from [2]int) [][2]int {
		var res [][2]int
		dirs := m[a[from[0]][from[1]]]
		for _, dir := range dirs {
			if !valid(addXY(from, dir)) {
				continue
			}
			res = append(res, addXY(from, dir))
		}
		return res
	}

	validMove := func(from [2]int, to [2]int) bool {
		moves := move(to)
		for _, move := range moves {
			if move == from {
				return true
			}
		}
		return false
	}

	for k, v := range m {
		next1 := addXY(s, v[0])
		next2 := addXY(s, v[1])
		if !valid(next1) || !validMove(s, next1) {
			continue
		}
		if !valid(next2) || !validMove(s, next2) {
			continue
		}
		a[s[0]][s[1]] = k
	}

	b := make([][]byte, 0, len(a))
	for range a {
		bb := make([]byte, len(a[0]))
		for i := 0; i < len(bb); i++ {
			bb[i] = '.'
		}
		b = append(b, bb)
	}

	var dfs func([2]int)
	dfs = func(start [2]int) {
		if b[start[0]][start[1]] == '*' {
			return
		}
		b[start[0]][start[1]] = '*'
		moves := move(start)
		for _, mov := range moves {
			dfs(mov)
		}
	}
	dfs(s)

	for _, bb := range b {
		print(string(bb))
	}

	c := make([][]byte, len(a)+1)
	for i := 0; i <= len(a); i++ {
		c[i] = make([]byte, len(a[0])+1)
		for j := 0; j <= len(a[0]); j++ {
			c[i][j] = '.'
		}
	}

	canMoveLR := func(top byte) bool {
		switch top {
		case '7', 'F', '|':
			return false
		default:
			return true
		}
	}
	canMoveUD := func(left byte) bool {
		switch left {
		case 'L', 'F', '-':
			return false
		default:
			return true
		}
	}

	valid2 := func(idx [2]int) bool {
		if idx[0] < 0 || idx[0] >= len(a)+1 {
			return false
		}
		if idx[1] < 0 || idx[1] >= len(a[0])+1 {
			return false
		}
		return true
	}

	type dir struct {
		move    [2]int
		shift   [2]int
		canMove func(byte) bool
	}

	dirs := []dir{
		{[2]int{-1, 0}, [2]int{-1, -1}, canMoveUD},
		{[2]int{0, -1}, [2]int{-1, -1}, canMoveLR},
		{[2]int{1, 0}, [2]int{0, -1}, canMoveUD},
		{[2]int{0, 1}, [2]int{-1, 0}, canMoveLR},
	}

	var dfs2 func(start [2]int, label byte) bool
	dfs2 = func(start [2]int, label byte) bool {
		if !valid2(start) {
			return true
		}

		if c[start[0]][start[1]] == label {
			return false
		}
		c[start[0]][start[1]] = label

		canReachOut := false
		for _, dir := range dirs {
			s := addXY(start, dir.shift)
			if !valid(s) {
				canReachOut = true
				continue
			}

			if b[s[0]][s[1]] == '*' && !dir.canMove(a[s[0]][s[1]]) {
				continue
			}

			canReachOut = canReachOut || dfs2(addXY(start, dir.move), label)
		}
		return canReachOut
	}

	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(a[0]); j++ {
			if c[i][j] != '.' {
				continue
			}

			canReachOut := dfs2([2]int{i, j}, 'X')
			var label byte = 'I'
			if canReachOut {
				label = 'O'
			}
			dfs2([2]int{i, j}, label)
		}
	}

	for _, cc := range c {
		print(string(cc))
	}

	var res int64 = 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if b[i][j] == '*' {
				continue
			}
			inner := true
			for _, diag := range [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}} {
				shift := addXY([2]int{i, j}, diag)
				inner = inner && c[shift[0]][shift[1]] == 'I'
			}
			if inner {
				res++
			}
		}
	}
	print(res)
}
