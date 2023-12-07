package main

import "sort"

func Solve7() {
	m := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	getType := func(s string) int {
		m := make(map[rune]int)
		for _, v := range s {
			m[v]++
		}
		r := 0
		for _, v := range m {
			r += v * v
		}
		return r
	}

	type hand struct {
		cards string
		bet   int64
		typ   int
	}
	var a []hand
	var s string
	var x int64
	for read(&s, &x) {
		a = append(a, hand{s, x, getType(s)})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].typ != a[j].typ {
			return a[i].typ < a[j].typ
		}
		for k := 0; k < len(a[k].cards); k++ {
			if m[a[i].cards[k]] == m[a[j].cards[k]] {
				continue
			}
			return m[a[i].cards[k]] < m[a[j].cards[k]]
		}
		return false
	})

	var r int64
	for i := 0; i < len(a); i++ {
		r += a[i].bet * int64(i+1)
	}
	print(r)
}

func Solve7_2() {
	m := map[byte]int{
		'J': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	getType := func(s string) int {
		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			m[s[i]]++
		}
		mv := 0
		mk := byte('A')
		for k, v := range m {
			if v > mv && k != 'J' {
				mv = v
				mk = k
			}
		}
		m[mk] += m['J']
		m['J'] = 0
		r := 0
		for _, v := range m {
			r += v * v
		}
		return r
	}

	type hand struct {
		cards string
		bet   int64
		typ   int
	}
	var a []hand
	var s string
	var x int64
	for read(&s, &x) {
		a = append(a, hand{s, x, getType(s)})
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].typ != a[j].typ {
			return a[i].typ < a[j].typ
		}
		for k := 0; k < len(a[k].cards); k++ {
			if m[a[i].cards[k]] == m[a[j].cards[k]] {
				continue
			}
			return m[a[i].cards[k]] < m[a[j].cards[k]]
		}
		return false
	})

	var r int64
	for i := 0; i < len(a); i++ {
		r += a[i].bet * int64(i+1)
	}
	print(r)
}
