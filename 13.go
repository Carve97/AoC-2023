package main

import "slices"

func Solve13() {
	var s string
	var ss [][]byte
	var res int64 = 0

	extractColumn := func(slice [][]byte, idx int) []byte {
		b := make([]byte, 0, len(slice[0]))
		for i := 0; i < len(slice); i++ {
			b = append(b, slice[i][idx])
		}
		return b
	}

	solve := func() int64 {
		for i := 1; i < len(ss); i++ {
			mirror := true
			for j := 0; j < min(i, len(ss)-i); j++ {
				if !slices.Equal(ss[i-j-1], ss[i+j]) {
					mirror = false
					break
				}
			}
			if mirror {
				return int64(i) * 100
			}
		}
		for j := 1; j < len(ss[0]); j++ {
			mirror := true
			for i := 0; i < min(j, len(ss[0])-j); i++ {
				if !slices.Equal(extractColumn(ss, j-i-1), extractColumn(ss, j+i)) {
					mirror = false
					break
				}
			}
			if mirror {
				return int64(j)
			}
		}
		return 0
	}

	for readLine(&s) {
		if len(s) == 0 {
			res += solve()
			ss = [][]byte{}
			continue
		}

		ss = append(ss, []byte(s))
	}
	res += solve()
	print(res)
}

func Solve13_2() {
	var s string
	var ss [][]byte
	var res int64 = 0

	extractColumn := func(slice [][]byte, idx int) []byte {
		b := make([]byte, 0, len(slice[0]))
		for i := 0; i < len(slice); i++ {
			b = append(b, slice[i][idx])
		}
		return b
	}

	numberOfDiffs := func(a, b []byte) int {
		r := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				r++
			}
		}
		return r
	}

	solve := func() int64 {
		for i := 1; i < len(ss); i++ {
			r := 0
			for j := 0; j < min(i, len(ss)-i); j++ {
				r += numberOfDiffs(ss[i-j-1], ss[i+j])
			}
			if r == 1 {
				return int64(i) * 100
			}
		}
		for j := 1; j < len(ss[0]); j++ {
			r := 0
			for i := 0; i < min(j, len(ss[0])-j); i++ {
				r += numberOfDiffs(extractColumn(ss, j-i-1), extractColumn(ss, j+i))
			}
			if r == 1 {
				return int64(j)
			}
		}
		return 0
	}

	for readLine(&s) {
		if len(s) == 0 {
			res += solve()
			ss = [][]byte{}
			continue
		}

		ss = append(ss, []byte(s))
	}
	res += solve()
	print(res)
}
