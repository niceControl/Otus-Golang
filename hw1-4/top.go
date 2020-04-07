package main

import (
	"sort"
	"strings"
)

func uniqueMap(pre []string) map[string]int {
	res := map[string]int{}
	for i := range pre {
		res[pre[i]] = 0
	}
	return res
}

func ShowTop(input string) (result []string) {
	sorted := strings.Split(input, " ")
	type Pair struct {
		Word   string
		Amount int
	}
	sort.Strings(sorted)
	preresult := uniqueMap(sorted)
	for i := 0; i < len(sorted); i++ {
		if _, ok := preresult[sorted[i]]; ok {
			preresult[sorted[i]] += 1
		}
	}
	var tops []Pair
	for w, a := range preresult {
		tops = append(tops, Pair{w, a})
	}
	sort.Slice(tops, func(i, j int) bool {
		return tops[i].Amount > tops[j].Amount
	})
	for i := 0; i < 10; i++ {
		result = append(result,tops[i].Word )
	}

	return
}

func main() {
}
