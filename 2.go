package main

import (
	"strconv"
	"strings"
)

func Solve2() {
	var s string
	var cnt int64 = 0
	for i := 1; readLine(&s); i++ {
		idx := strings.Index(s, ":")
		s = s[idx+1:]
		cnt += gamePossible(s)
	}
	print(cnt)
}

func gamePossible(gameString string) int64 {
	minCnt := map[string]int{
		"green": 0,
		"blue":  0,
		"red":   0,
	}
	subGames := strings.Split(gameString, ";")
	for _, subGame := range subGames {
		entries := strings.Split(subGame, ",")
		for _, entry := range entries {
			cVal := strings.Split(entry[1:], " ")
			val, _ := strconv.Atoi(cVal[0])
			minCnt[cVal[1]] = max(minCnt[cVal[1]], val)
		}
	}

	var cnt int64 = 1
	for _, val := range minCnt {
		cnt *= int64(val)
	}
	return cnt
}
