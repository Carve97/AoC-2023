package main

import (
	"strings"
)

func Solve12() {
	var s string
	var x int
	var res int64 = 0
	for readLine(&s) {
		s = strings.ReplaceAll(s, ",", " ")
		r := strings.NewReader(s)

		var lens []int
		var remVal int
		readFromReader(r, &s)
		for readFromReader(r, &x) {
			lens = append(lens, x)
			remVal += x
		}

		mem := make([][]*int64, len(s)+2)
		for i := range mem {
			mem[i] = make([]*int64, remVal+1)
		}

		var rec func(idx int, rem []int, remVal int) int64
		rec = func(idx int, rem []int, remVal int) int64 {
			if mem[idx][remVal] != nil {
				return *mem[idx][remVal]
			}

			if len(rem) == 0 {
				if idx >= len(s) || strings.Count(s[idx:], "#") == 0 {
					return 1
				}
				return 0
			}
			if len(s) < idx+rem[0] || len(s)-idx < remVal {
				return 0
			}

			placeAt := s[idx : idx+rem[0]]
			freeSpace := strings.Count(placeAt, ".") == 0
			freeNeighborL := idx == 0 || s[idx-1] != '#'
			freeNeighborR := idx+rem[0] == len(s) || s[idx+rem[0]] != '#'

			var subRes int64 = 0
			if freeSpace && freeNeighborL && freeNeighborR {
				subRes = rec(idx+rem[0]+1, rem[1:], remVal-rem[0])
			}
			if s[idx] != '#' {
				subRes += rec(idx+1, rem, remVal)
			}

			mem[idx][remVal] = &subRes
			return subRes
		}
		res += rec(0, lens, remVal)
	}
	print(res)
}

func Solve12_2() {
	var s string
	var x int
	var res int64 = 0
	for readLine(&s) {
		s = strings.ReplaceAll(s, ",", " ")
		r := strings.NewReader(s)

		var lens []int
		var remVal int
		readFromReader(r, &s)
		for readFromReader(r, &x) {
			lens = append(lens, x)
			remVal += x
		}

		var ss []string
		var ll []int
		for i := 0; i < 5; i++ {
			ss = append(ss, s)
			ll = append(ll, lens...)
		}

		s = strings.Join(ss, "?")
		lens = ll
		remVal *= 5

		mem := make([][]*int64, len(s)+2)
		for i := range mem {
			mem[i] = make([]*int64, remVal+1)
		}

		var rec func(idx int, rem []int, remVal int) int64
		rec = func(idx int, rem []int, remVal int) int64 {
			if mem[idx][remVal] != nil {
				return *mem[idx][remVal]
			}

			if len(rem) == 0 {
				if idx >= len(s) || strings.Count(s[idx:], "#") == 0 {
					return 1
				}
				return 0
			}
			if len(s) < idx+rem[0] || len(s)-idx < remVal {
				return 0
			}

			placeAt := s[idx : idx+rem[0]]
			freeSpace := strings.Count(placeAt, ".") == 0
			freeNeighborL := idx == 0 || s[idx-1] != '#'
			freeNeighborR := idx+rem[0] == len(s) || s[idx+rem[0]] != '#'

			var subRes int64 = 0
			if freeSpace && freeNeighborL && freeNeighborR {
				subRes = rec(idx+rem[0]+1, rem[1:], remVal-rem[0])
			}
			if s[idx] != '#' {
				subRes += rec(idx+1, rem, remVal)
			}

			mem[idx][remVal] = &subRes
			return subRes
		}
		res += rec(0, lens, remVal)
	}
	print(res)
}
