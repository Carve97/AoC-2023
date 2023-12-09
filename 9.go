package main

import (
	"slices"
	"strings"
)

func Solve9() {
	var a [][]int64
	var s string
	for readLine(&s) {
		var x int64
		var b []int64
		r := strings.NewReader(s)
		for readFromReader(r, &x) {
			b = append(b, x)
		}
		a = append(a, b)
	}

	var r int64 = 0
	for _, v := range a {
		b := append([][]int64{}, v)
		for i := 0; i < len(b); i++ {
			var bb []int64
			f := false
			for j := 1; j < len(b[i]); j++ {
				val := b[i][j] - b[i][j-1]
				bb = append(bb, val)
				f = f || val != 0
			}
			if !f {
				b[i] = append(b[i], lastElem(b[i]))
				break
			}

			b = append(b, bb)
		}

		for i := len(b) - 1; i >= 1; i-- {
			val := lastElem(b[i]) + lastElem(b[i-1])
			if i == 1 {
				r += val
				break
			}
			b[i-1] = append(b[i-1], val)
		}
	}
	print(r)
}

func Solve9_2() {
	var a [][]int64
	var s string
	for readLine(&s) {
		var x int64
		var b []int64
		r := strings.NewReader(s)
		for readFromReader(r, &x) {
			b = append(b, x)
		}
		slices.Reverse(b)
		a = append(a, b)
	}

	var r int64 = 0
	for _, v := range a {
		b := append([][]int64{}, v)
		for i := 0; i < len(b); i++ {
			var bb []int64
			f := false
			for j := 1; j < len(b[i]); j++ {
				val := b[i][j] - b[i][j-1]
				bb = append(bb, val)
				f = f || val != 0
			}
			if !f {
				b[i] = append(b[i], lastElem(b[i]))
				break
			}

			b = append(b, bb)
		}

		for i := len(b) - 1; i >= 1; i-- {
			val := lastElem(b[i]) + lastElem(b[i-1])
			if i == 1 {
				r += val
				break
			}
			b[i-1] = append(b[i-1], val)
		}
	}
	print(r)
}
