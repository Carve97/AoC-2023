package main

func Solve11() {
	var ss []string
	var s string
	for readLine(&s) {
		ss = append(ss, s)
	}

	var galaxies [][2]int64
	h := make([]int64, len(ss))
	v := make([]int64, len(ss[0]))
	for i := 0; i < len(ss); i++ {
		v[i] = 1000000
		for j := 0; j < len(ss[0]); j++ {
			if ss[i][j] == '#' {
				galaxies = append(galaxies, [2]int64{int64(i), int64(j)})
				v[i] = 1
			}
		}
	}
	for j := 0; j < len(ss[0]); j++ {
		h[j] = 1000000
		for i := 0; i < len(ss); i++ {
			if ss[i][j] == '#' {
				h[j] = 1
			}
		}
	}

	diff := func(x int64, y int64, horizontal bool) int64 {
		mx, mi := max(x, y), min(x, y)
		s := v
		if horizontal {
			s = h
		}
		var res int64 = 0
		for i := mx; i > mi; i-- {
			res += s[i]
		}
		return res
	}

	var res int64 = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			res += diff(galaxies[i][0], galaxies[j][0], false)
			res += diff(galaxies[i][1], galaxies[j][1], true)
		}
	}
	print(res)
}
