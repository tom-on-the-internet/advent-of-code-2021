package main

import (
	"strings"
)

type item struct {
	signals []string
	output  []string
}

func (i item) number() int {
	m := make(map[string]string)
	m["a"] = i.a()

	return 0
}

func (i item) a() string {
	output1 := i.output1()
	output7 := i.output7()
	res := strings.Replace(output7, string(output1[0]), "", 1)
	res = strings.Replace(res, string(output1[1]), "", 1)

	return res
}

func (i item) output1() string {
	for _, v := range i.signals {
		if len(v) == 2 {
			return v
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
