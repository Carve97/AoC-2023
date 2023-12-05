package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

func Solve5() {
	var s string
	var i int64
	var seeds []int64
	readLine(&s)
	idx := strings.Index(s, ":")
	s = s[idx+1:]
	r := strings.NewReader(s)
	for _, err := fmt.Fscan(r, &i); err == nil; _, err = fmt.Fscan(r, &i) {
		seeds = append(seeds, i)
	}
	for readLine(&s) {
		if idx := strings.Index(s, ":"); len(s) == 0 || idx != -1 {
			continue
		}

		var vals [][3]int64
		var dst, src, ln int64
		for len(s) > 0 {
			r = strings.NewReader(s)
			fmt.Fscan(r, &dst, &src, &ln)
			vals = append(vals, [3]int64{src, src + ln - 1, dst})
			readLine(&s)
		}
		sort.Slice(vals, func(i, j int) bool {
			return vals[i][0] < vals[j][0]
		})
		for i := 0; i < len(seeds); i++ {
			idx := sort.Search(len(vals), func(j int) bool {
				return vals[j][0] > seeds[i]
			})
			if idx == 0 || vals[idx-1][1] < seeds[i] {
				continue
			}
			idx--
			diff := seeds[i] - vals[idx][0]
			seeds[i] = vals[idx][2] + diff
		}
	}
	print(slices.Min(seeds))
}

func Solve5_2() {
	var s string
	var i, j int64
	var seeds [][2]int64
	readLine(&s)
	s = strings.TrimPrefix(s, "seeds:")
	r := strings.NewReader(s)
	for _, err := fmt.Fscan(r, &i, &j); err == nil; _, err = fmt.Fscan(r, &i, &j) {
		seeds = append(seeds, [2]int64{i, i + j - 1})
	}
	for readLine(&s) {
		if idx := strings.Index(s, ":"); len(s) == 0 || idx != -1 {
			continue
		}

		var vals [][3]int64
		var dst, src, ln int64
		for len(s) > 0 {
			r = strings.NewReader(s)
			fmt.Fscan(r, &dst, &src, &ln)
			vals = append(vals, [3]int64{src, src + ln - 1, dst})
			readLine(&s)
		}
		sort.Slice(vals, func(i, j int) bool {
			return vals[i][0] < vals[j][0]
		})
		var newSeeds [][2]int64
		for i := 0; i < len(seeds); i++ {
			from := seeds[i][0]
			initialTo := seeds[i][1]

			for from < initialTo {

				idx := sort.Search(len(vals), func(j int) bool {
					return vals[j][0] > from
				})

				var dest, lowerBound, upperBound int64 = 0, 0, math.MaxInt64
				if idx < len(vals) {
					upperBound = vals[idx][0]
				}
				if idx > 0 && from <= vals[idx-1][1] {
					lowerBound = vals[idx-1][0]
					upperBound = min(initialTo, vals[idx-1][1])
					dest = vals[idx-1][2]
				}

				to := min(initialTo, upperBound)

				diffFrom := from - lowerBound
				diffTo := to - lowerBound
				newSeeds = append(newSeeds, [2]int64{dest + diffFrom, dest + diffTo})

				from = to + 1
			}
		}
		seeds = newSeeds
	}

	var res int64 = math.MaxInt64
	for _, v := range seeds {
		res = min(res, v[0])
	}
	print(res)
}
