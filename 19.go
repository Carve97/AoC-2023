package main

import (
	"maps"
	"strconv"
	"strings"
)

func Solve19() {
	type part struct {
		values map[byte]int64
	}

	type checkFunc func(part) string

	workflows := make(map[string][]checkFunc)
	var s string
	for readLine(&s) && len(s) > 0 {
		wfName := s[:strings.Index(s, "{")]
		rules := strings.Split(s[strings.Index(s, "{")+1:len(s)-1], ",")

		for _, rule := range rules {
			ruleParts := strings.Split(rule, ":")
			if len(ruleParts) == 1 {
				workflows[wfName] = append(workflows[wfName], func(part) string { return rule })
				continue
			}

			var num int64
			num, _ = strconv.ParseInt(ruleParts[0][2:], 10, 64)

			f := func(part part) string {
				if ruleParts[0][1] == '<' && part.values[ruleParts[0][0]] < num {
					return ruleParts[1]
				}
				if ruleParts[0][1] == '>' && part.values[ruleParts[0][0]] > num {
					return ruleParts[1]
				}
				return ""
			}
			workflows[wfName] = append(workflows[wfName], f)
		}
	}

	var res int64 = 0
	for readLine(&s) {
		values := strings.Split(s[1:len(s)-1], ",")
		var num int64
		p := part{values: make(map[byte]int64)}
		for _, value := range values {
			num, _ = strconv.ParseInt(value[2:], 10, 64)
			p.values[value[0]] = num
		}

		curWf := "in"
		for curWf != "R" && curWf != "A" {
			for _, rule := range workflows[curWf] {
				if nextWf := rule(p); nextWf != "" {
					curWf = nextWf
					break
				}
			}
		}
		if curWf == "A" {
			for _, v := range p.values {
				res += v
			}
		}
	}
	print(res)
}

func Solve19_2() {
	type interval struct {
		from int64
		to   int64
	}
	intervalValid := func(i interval) bool {
		return i.from < i.to
	}

	type configurations map[byte]interval
	configurationsValid := func(c configurations) bool {
		for _, i := range c {
			if !intervalValid(i) {
				return false
			}
		}
		return true
	}
	countValidConfigurations := func(c configurations) int64 {
		var res int64 = 1
		for _, i := range c {
			res *= 1 + i.to - i.from
		}
		return res
	}

	type splitFunc func(configurations) (pass configurations, dontPass configurations, passDest string)
	workflows := make(map[string][]splitFunc)
	var s string

	for readLine(&s) && len(s) > 0 {
		wfName := s[:strings.Index(s, "{")]
		rules := strings.Split(s[strings.Index(s, "{")+1:len(s)-1], ",")

		for _, rule := range rules {
			ruleParts := strings.Split(rule, ":")
			if len(ruleParts) == 1 {
				workflows[wfName] = append(workflows[wfName], func(configurations configurations) (
					pass configurations, dontPass configurations, passDest string) {
					return configurations, nil, ruleParts[0]
				})
				continue
			}

			var num int64
			num, _ = strconv.ParseInt(ruleParts[0][2:], 10, 64)
			var splitF splitFunc = func(configurations configurations) (pass configurations, dontPass configurations, passDest string) {
				pass = maps.Clone(configurations)
				dontPass = maps.Clone(configurations)

				value := ruleParts[0][0]
				if ruleParts[0][1] == '<' {
					pass[value] = interval{to: num - 1, from: pass[value].from}
					dontPass[value] = interval{to: dontPass[value].to, from: num}
				}
				if ruleParts[0][1] == '>' {
					pass[value] = interval{to: pass[value].to, from: num + 1}
					dontPass[value] = interval{to: num, from: dontPass[value].from}
				}
				return pass, dontPass, ruleParts[1]
			}
			workflows[wfName] = append(workflows[wfName], splitF)
		}
	}

	var dfs func(configuration configurations, workflow string) int64
	dfs = func(configuration configurations, workflow string) int64 {
		if workflow == "A" {
			return countValidConfigurations(configuration)
		}

		if workflow == "R" {
			return 0
		}

		var res int64 = 0
		curConfiguration := configuration
		for _, rule := range workflows[workflow] {
			canPass, dontPass, nextPassWF := rule(curConfiguration)
			if configurationsValid(canPass) {
				res += dfs(canPass, nextPassWF)
			}
			curConfiguration = dontPass
		}
		return res
	}

	fullConfiguration := map[byte]interval{
		'x': {from: 1, to: 4000},
		'm': {from: 1, to: 4000},
		'a': {from: 1, to: 4000},
		's': {from: 1, to: 4000},
	}

	print(dfs(fullConfiguration, "in"))
}
