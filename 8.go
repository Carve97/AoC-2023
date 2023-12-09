package main

import (
	"fmt"
	"strings"
)

func Solve8() {
	var s1, s2 string
	m := make(map[string][2]string)

	readLine(&s1)
	readLine(&s2)
	for readLine(&s2) {
		s2 = strings.Map(func(r rune) rune {
			switch r {
			case '=', ',', '(', ')':
				return ' '
			default:
				return r
			}
		}, s2)
		var x, l, r string
		fmt.Fscan(strings.NewReader(s2), &x, &l, &r)
		m[x] = [2]string{l, r}
	}

	var r int = 0
	var s string = "AAA"
	for s != "ZZZ" {
		idx := 0
		if s1[r%len(s1)] == 'R' {
			idx++
		}
		s = m[s][idx]
		r++
	}
	print(r)
}

func Solve8_2() {
	var s1, s2 string
	var as []string
	m := make(map[string][2]string)

	readLine(&s1)
	readLine(&s2)
	for readLine(&s2) {
		s2 = strings.Map(func(r rune) rune {
			switch r {
			case '=', ',', '(', ')':
				return ' '
			default:
				return r
			}
		}, s2)
		var x, l, r string
		fmt.Fscan(strings.NewReader(s2), &x, &l, &r)
		m[x] = [2]string{l, r}
		if x[len(x)-1] == 'A' {
			as = append(as, x)
		}
	}

	var gcd func(int64, int64) int64
	gcd = func(i, j int64) int64 {
		if i%j == 0 {
			return j
		}
		return gcd(j, i%j)
	}
	lcm := func(i, j int64) int64 {
		return (i * j) / gcd(max(i, j), min(i, j))
	}

	var res int64 = 1
	for _, a := range as {
		var r int64 = 0
		for a[len(a)-1] != 'Z' {
			idx := 0
			if s1[int(r)%len(s1)] == 'R' {
				idx++
			}
			a = m[a][idx]
			r++
		}
		res = lcm(res, r)
	}
	print(res)
}
