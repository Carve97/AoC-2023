package AoC_2023

import "unicode"

func Solve1() {
	defer writer.Flush()
	var cnt int64
	var s string

	digitStrings := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	isValid := func(i int) (int64, bool) {
		if unicode.IsDigit(rune(s[i])) {
			return int64(s[i] - '0'), true
		}
		for idx, ds := range digitStrings {
			if i+len(ds) > len(s) {
				continue
			}
			if ds == s[i:i+len(ds)] {
				return int64(idx + 1), true
			}
		}
		return 0, false
	}

	for {
		ok := read(&s)
		if !ok {
			break
		}

		for i := 0; i < len(s); i++ {
			if num, ok := isValid(i); ok {
				cnt += num * 10
				break
			}
		}
		for i := len(s) - 1; i >= 0; i-- {
			if num, ok := isValid(i); ok {
				cnt += num
				break
			}
		}
	}
	print(cnt)
}
