package main

import (
	"golang.org/x/exp/maps"
	"slices"
	"strings"
)

func Solve20() {
	type module struct {
		name string
		typ  byte
		out  []string
		in   map[string]bool
	}

	modules := make(map[string]*module)
	var s string
	for readLine(&s) {
		s = strings.ReplaceAll(s, " ", "")
		splits := strings.Split(s, "->")
		var typ byte
		var inMap map[string]bool
		if splits[0][0] == '%' || splits[0][0] == '&' {
			typ = splits[0][0]
			splits[0] = splits[0][1:]
			inMap = make(map[string]bool)
		}
		if typ == '%' {
			inMap[""] = false
		}

		m := &module{
			name: splits[0],
			typ:  typ,
			in:   inMap,
		}
		for _, out := range strings.Split(splits[1], ",") {
			m.out = append(m.out, out)
		}
		modules[m.name] = m
	}

	for _, m := range modules {
		for _, out := range m.out {
			if modules[out] != nil && modules[out].typ == '&' {
				modules[out].in[m.name] = false
			}
		}
	}

	type pulse struct {
		from string
		to   string
		hilo bool
	}

	var hi, lo int64 = 0, 0
	for i := 0; i < 1000; i++ {
		var queue = []pulse{{from: "button", to: "broadcaster", hilo: false}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p.hilo {
				hi++
			} else {
				lo++
			}

			m, ok := modules[p.to]
			if !ok {
				continue
			}

			if m.typ == '%' {
				if p.hilo == true {
					continue
				}
				m.in[""] = !m.in[""]
				p.hilo = m.in[""]
			}
			if m.typ == '&' {
				m.in[p.from] = p.hilo
				p.hilo = slices.Contains(maps.Values(m.in), false)
			}

			for _, o := range m.out {
				queue = append(queue, pulse{from: m.name, to: o, hilo: p.hilo})
			}
		}
	}

	print(hi * lo)
}

func Solve20_2() {
	type module struct {
		name string
		typ  byte
		out  []string
		in   map[string]bool
	}

	modules := make(map[string]*module)
	var s string
	for readLine(&s) {
		s = strings.ReplaceAll(s, " ", "")
		splits := strings.Split(s, "->")
		var typ byte
		var inMap map[string]bool
		if splits[0][0] == '%' || splits[0][0] == '&' {
			typ = splits[0][0]
			splits[0] = splits[0][1:]
			inMap = make(map[string]bool)
		}
		if typ == '%' {
			inMap[""] = false
		}

		m := &module{
			name: splits[0],
			typ:  typ,
			in:   inMap,
		}
		for _, out := range strings.Split(splits[1], ",") {
			m.out = append(m.out, out)
		}
		modules[m.name] = m
	}

	for _, m := range modules {
		for _, out := range m.out {
			if modules[out] != nil && modules[out].typ == '&' {
				modules[out].in[m.name] = false
			}
		}
	}

	type pulse struct {
		from string
		to   string
		hilo bool
	}

	expectedLen := len(modules["xm"].in)
	periods := make(map[string]int64, expectedLen)

	for i := 0; len(periods) < expectedLen; i++ {
		var queue = []pulse{{from: "button", to: "broadcaster", hilo: false}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			m, ok := modules[p.to]
			if !ok {
				continue
			}

			if m.typ == '%' {
				if p.hilo == true {
					continue
				}
				m.in[""] = !m.in[""]
				p.hilo = m.in[""]
			}
			if m.typ == '&' {
				m.in[p.from] = p.hilo
				p.hilo = slices.Contains(maps.Values(m.in), false)

				if m.name == "xm" {
					if slices.Contains(maps.Values(m.in), true) {
						print(i+1, m.in)
					}
				}
			}

			// to reach rx, we need to all hi's on xm
			// therefore find the periods where the inputs for
			// xm are true
			if m.name == "xm" && m.in[p.from] {
				if _, in := periods[p.from]; !in {
					periods[p.from] = int64(i + 1)
				}
			}

			for _, o := range m.out {
				queue = append(queue, pulse{from: m.name, to: o, hilo: p.hilo})
			}
		}
	}

	var gcd func(int64, int64) int64
	gcd = func(a, b int64) int64 {
		if a%b == 0 {
			return b
		}
		return gcd(b, a%b)
	}

	lcm := func(a, b int64) int64 {
		g := gcd(max(a, b), min(a, b))
		return (a / g) * (b / g)
	}

	var res int64 = 1
	for _, val := range periods {
		res = lcm(res, val)
	}
	print(res)
}
