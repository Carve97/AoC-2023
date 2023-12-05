package main

import (
	"fmt"
	"strings"
)

func Solve4() {
	var s string
	var i int
	var res int
	for readLine(&s) {
		idx := strings.Index(s, ":")
		s = s[idx+1:]
		splits := strings.Split(s, "|")

		m := make(map[int]struct{})
		r := strings.NewReader(splits[0])
		for _, err := fmt.Fscan(r, &i); err == nil; _, err = fmt.Fscan(r, &i) {
			m[i] = struct{}{}
		}
		r = strings.NewReader(splits[1])
		cnt := -1
		for _, err := fmt.Fscan(r, &i); err == nil; _, err = fmt.Fscan(r, &i) {
			if _, ok := m[i]; ok {
				cnt++
			}
		}
		res += 1 << max(cnt, 0)
	}
	print(res)
}

func Solve4_2() {
	var s string
	var i int
	var res int

	mr := make(map[int]int)

	for gn := 1; readLine(&s); gn++ {
		idx := strings.Index(s, ":")
		s = s[idx+1:]
		splits := strings.Split(s, "|")

		m := make(map[int]struct{})
		r := strings.NewReader(splits[0])
		for _, err := fmt.Fscan(r, &i); err == nil; _, err = fmt.Fscan(r, &i) {
			m[i] = struct{}{}
		}
		r = strings.NewReader(splits[1])
		cnt := 0
		for _, err := fmt.Fscan(r, &i); err == nil; _, err = fmt.Fscan(r, &i) {
			if _, ok := m[i]; ok {
				cnt++
			}
		}
		mr[gn]++
		for ; cnt > 0; cnt-- {
			mr[gn+cnt] += mr[gn]
		}
	}
	for _, v := range mr {
		res += v
	}
	print(res)
}
