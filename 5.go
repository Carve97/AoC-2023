package AoC_2023

import (
	"fmt"
	"strings"
)

func Solve5() {
	var s string
	var i int
	var a []int
	readLine(&s)
	idx := strings.Index(s, ":")
	s = s[idx+1:]
	r := strings.NewReader(s)
	for _, err := fmt.Fscan(r, &i); err != nil; _, err = fmt.Fscan(r, &i) {

	}
}
