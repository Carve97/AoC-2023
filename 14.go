package main

func Solve14() {
	var ss []string
	var s string
	for read(&s) {
		ss = append(ss, s)
	}

	var total int64 = 0
	for j := 0; j < len(ss[0]); j++ {
		var shiftTo = len(ss)
		for i := 0; i < len(ss); i++ {
			if ss[i][j] == '#' {
				shiftTo = len(ss) - i - 1
			} else if ss[i][j] == 'O' {
				total += int64(shiftTo)
				shiftTo--
			}
		}
	}
	print(total)
}

func Solve14_2() {
	var ss [][]byte
	var s string
	for read(&s) {
		ss = append(ss, []byte(s))
	}

	shiftUD := func(startI int, endI int, shiftI int) {
		for j := 0; j < len(ss[0]); j++ {
			moveTo := startI
			for i := startI; i != endI; i += shiftI {
				if ss[i][j] == '#' {
					moveTo = i + shiftI
				} else if ss[i][j] == 'O' {
					if moveTo != i {
						ss[moveTo][j] = 'O'
						ss[i][j] = '.'
					}
					moveTo += shiftI
				}
			}
		}
	}
	shiftLR := func(startJ int, endJ, shiftJ int) {
		for i := 0; i < len(ss); i++ {
			moveTo := startJ
			for j := startJ; j != endJ; j += shiftJ {
				if ss[i][j] == '#' {
					moveTo = j + shiftJ
				} else if ss[i][j] == 'O' {
					if moveTo != j {
						ss[i][moveTo] = 'O'
						ss[i][j] = '.'
					}
					moveTo += shiftJ
				}
			}
		}
	}

	calcLoad := func() int64 {
		var total int64 = 0
		for j := 0; j < len(ss[0]); j++ {
			for i := 0; i < len(ss); i++ {
				if ss[i][j] == 'O' {
					total += int64(len(ss) - i)
				}
			}
		}
		return total
	}

	var res []int64
	checkForPeriod := func() ([]int64, int) {
		if len(res)%2 == 0 {
			return nil, 0
		}

		// period counter
		for k := 2; k <= len(res)/3; k++ {
			periodic := true
			for i := 0; i < k; i++ {
				// check for 3 consecutive string
				if res[len(res)-1-i] != res[len(res)-1-k-i] {
					periodic = false
					break
				}
				if res[len(res)-1-i] != res[len(res)-1-(2*k)-i] {
					periodic = false
					break
				}
			}
			if periodic {
				return res[len(res)-k:], len(res) - 3*k
			}
		}
		return nil, 0
	}

	// run until periodic state is reached
	for {
		shiftUD(0, len(ss), 1)
		shiftLR(0, len(ss[0]), 1)
		shiftUD(len(ss)-1, -1, -1)
		shiftLR(len(ss[0])-1, -1, -1)
		res = append(res, calcLoad())

		period, start := checkForPeriod()
		if period == nil {
			continue
		}

		requiredIteration := 1000000000
		print(period[(requiredIteration-start-1)%len(period)])
		return
	}
}
