package main

import (
	"fmt"
	"strings"
)

func Solve6() {
	var a, b []int64
	var x, y int64
	var s1, s2 string
	readLine(&s1)
	readLine(&s2)
	s1, s2 = strings.TrimPrefix(s1, "Time:"), strings.TrimPrefix(s2, "Distance:")
	r1, r2 := strings.NewReader(s1), strings.NewReader(s2)
	for _, err := fmt.Fscan(r1, &x); err == nil; _, err = fmt.Fscan(r1, &x) {
		fmt.Fscan(r2, &y)
		a = append(a, x)
		b = append(b, y)
	}

	var res int64 = 1
	for i := 0; i < len(a); i++ {
		var j, cnt int64 = 0, 0
		for ; j <= a[i]; j++ {
			if j*(a[i]-j) > b[i] {
				cnt++
			}
		}
		res *= cnt
	}
	print(res)
}

func Solve6_2() {
	var x, y int64
	var s1, s2 string
	readLine(&s1)
	readLine(&s2)
	s1, s2 = strings.TrimPrefix(s1, "Time:"), strings.TrimPrefix(s2, "Distance:")
	s1, s2 = strings.ReplaceAll(s1, " ", ""), strings.ReplaceAll(s2, " ", "")
	r1, r2 := strings.NewReader(s1), strings.NewReader(s2)
	fmt.Fscan(r1, &x)
	fmt.Fscan(r2, &y)
	print(x, y)

	var i int64
	for ; i <= x; i++ {
		if i*(x-i) > y {
			print(((x - i) - i) + 1)
			return
		}
	}

}
