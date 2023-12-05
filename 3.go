package main

import "unicode"

func Solve3() {
	var a []string
	var s string
	for read(&s) {
		a = append(a, s)
	}

	check := func(i, j int) bool {
		if i < 0 || i >= len(a) {
			return false
		}
		if j < 0 || j >= len(a[0]) {
			return false
		}
		return !unicode.IsDigit(rune(a[i][j])) && a[i][j] != '.'
	}

	symbolAround := func(i, j int) bool {
		res := false
		for _, ii := range []int{-1, 0, 1} {
			for _, jj := range []int{-1, 0, 1} {
				res = check(i+ii, j+jj) || res
			}
		}
		return res
	}

	var cnt int64
	for i, line := range a {
		var num int64
		var include bool
		for j, b := range line {
			if unicode.IsDigit(b) {
				num *= 10
				num += int64(b - '0')
				include = symbolAround(i, j) || include
			}
			if !unicode.IsDigit(b) || j == len(a[0])-1 {
				if include {
					cnt += num
				}
				num = 0
				include = false
			}
		}
	}
	print(cnt)
}

func Solve3_2() {
	var a []string
	var s string
	for read(&s) {
		a = append(a, s)
	}

	backtraceNum := func(i, j int) (int, int) {
		for ; j >= 0; j-- {
			if !unicode.IsDigit(rune(a[i][j])) {
				return i, j + 1
			}
		}
		return i, j + 1
	}

	followNum := func(i, j int) int64 {
		var num int64
		for ; j < len(a[0]); j++ {
			if !unicode.IsDigit(rune(a[i][j])) {
				break
			}
			num *= 10
			num += int64(a[i][j] - '0')
		}
		return num
	}

	getNum := func(i, j int) (int64, int) {
		if i < 0 || i >= len(a) {
			return 1, 0
		}
		if j < 0 || j >= len(a[0]) {
			return 1, 0
		}
		if !unicode.IsDigit(rune(a[i][j])) {
			return 1, 0
		}
		return followNum(backtraceNum(i, j)), 1
	}

	checkDiag := func(i, j int) (int64, int) {
		var num int64 = 1
		var cnt int
		if i < 0 || i >= len(a) {
			return 1, 0
		}
		if unicode.IsDigit(rune(a[i][j])) {
			return followNum(backtraceNum(i, j)), 1
		}

		if val, nums := getNum(i, j-1); nums != 0 {
			num *= val
			cnt += nums
		}
		if val, nums := getNum(i, j+1); nums != 0 {
			num *= val
			cnt += nums
		}
		return num, cnt
	}

	var cnt int64
	for i, line := range a {
		for j, b := range line {
			if b != '*' {
				continue
			}

			var val int64 = 1
			var nums = 0

			for _, f := range []func() (int64, int){
				func() (int64, int) { return checkDiag(i-1, j) },
				func() (int64, int) { return getNum(i, j-1) },
				func() (int64, int) { return getNum(i, j+1) },
				func() (int64, int) { return checkDiag(i+1, j) },
			} {
				v, c := f()
				val *= v
				nums += c
			}

			if nums == 2 {
				cnt += val
			}
		}
	}
	print(cnt)
}
