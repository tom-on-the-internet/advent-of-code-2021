package main

import (
	"fmt"
	"strings"
)

type item struct {
	signals []string
	output  []string
}

func (i item) number() int {
	// I want to output the number
	// Given a string, I need to know what number it is.
	// This has to be done per item because the wiring changes.
	// if I can pass in the correct wiring to a function, it can tell me the number.
	// get the correct wiring.
	m := make(map[string]string)
	m["a"] = i.a()
	m["b"] = i.b()
	m["d"] = i.d()
	m["g"] = i.g()
	m["c"] = i.c()
	m["f"] = i.f()
	m["e"] = i.e()

	fmt.Println(m)

	return 0
}

func (i item) a() string {
	output1 := i.output1()
	output7 := i.output7()
	res := strings.Replace(output7, string(output1[0]), "", 1)
	res = strings.Replace(res, string(output1[1]), "", 1)

	return res
}

func (i item) b() string {
	output1 := i.output1()
	output4 := i.output4()
	output3 := i.output3()

	res := strings.Replace(output4, string(output1[0]), "", 1)
	res = strings.Replace(res, string(output1[1]), "", 1)

	for _, v := range output3 {
		res = strings.Replace(res, string(v), "", 1)
	}

	return res
}

func (i item) c() string {
	output1 := i.output1()

	if strings.Contains(i.output6(), string(output1[0])) {
		return string(output1[1])
	}

	return string(output1[0])
}

func (i item) d() string {
	output1 := i.output1()
	output4 := i.output4()

	res := strings.Replace(output4, string(output1[0]), "", 1)
	res = strings.Replace(res, string(output1[1]), "", 1)

	if string(res[0]) == i.b() {
		return string(res[1])
	}

	return string(res[0])
}

func (i item) e() string {
	output8 := i.output8()
	output9 := i.output9()

	for _, v := range output8 {
		if !strings.Contains(output9, string(v)) {
			return string(v)
		}
	}

	return ""
}

func (i item) f() string {
	output1 := i.output1()

	if strings.Contains(i.output6(), string(output1[0])) {
		return string(output1[0])
	}

	return string(output1[1])
}

func (i item) g() string {
	output3 := i.output3()
	output4 := i.output4()
	res := strings.Replace(output3, string(output4[0]), "", 1)
	res = strings.Replace(res, string(output4[1]), "", 1)
	res = strings.Replace(res, string(output4[2]), "", 1)
	res = strings.Replace(res, string(output4[3]), "", 1)
	res = strings.Replace(res, i.a(), "", 1)

	return res
}

func (i item) output0() string {
	for _, v := range i.signals {
		if len(v) == 6 {
			if !strings.Contains(v, i.d()) {
				return v
			}
		}
	}

	return ""
}

func (i item) output1() string {
	for _, v := range i.signals {
		if len(v) == 2 {
			return v
		}
	}

	return ""
}

func (i item) output3() string {
	for _, v := range i.signals {
		if len(v) == 5 {

			output1 := i.output1()
			res := strings.Replace(v, string(output1[0]), "", 1)
			res = strings.Replace(res, string(output1[1]), "", 1)

			if len(res) == 3 {
				return v
			}
		}
	}

	return ""
}

func (i item) output4() string {
	for _, v := range i.signals {
		if len(v) == 4 {
			return v
		}
	}

	return ""
}

func (i item) output6() string {
	// if has len 6 but not both of 1
	output1 := i.output1()
	for _, v := range i.signals {
		if len(v) == 6 {
			if !strings.Contains(v, string(output1[0])) || !strings.Contains(v, string(output1[1])) {
				return v
			}
		}
	}

	return ""
}

func (i item) output7() string {
	for _, v := range i.signals {
		if len(v) == 3 {
			return v
		}
	}

	return ""
}

func (i item) output8() string {
	for _, v := range i.signals {
		if len(v) == 7 {
			return v
		}
	}

	return ""
}

func (i item) output9() string {
	for _, v := range i.signals {
		if len(v) == 6 && v != i.output0() && v != i.output6() {
			return v
		}
	}

	return ""
}

func Dec8A(input []string) int {
	easyNumberCount := 0

	items := toItems(input)

	for _, v := range items {
		for _, val := range v.output {
			if len(val) == 2 || len(val) == 3 || len(val) == 4 || len(val) == 7 {
				easyNumberCount++
			}
		}
	}

	return easyNumberCount
}

func Dec8B(input []string) int {
	sum := 0

	items := toItems(input)

	for _, v := range items {
		sum += v.number()
	}

	return sum
}

func toItems(input []string) []item {
	items := []item{}

	for _, v := range input {
		x := strings.Split(v, " | ")
		items = append(items, item{signals: strings.Fields(x[0]), output: strings.Fields(x[1])})
	}

	return items
}
