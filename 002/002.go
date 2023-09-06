package main

import (
	"fmt"
	"sort"
)

func main() {
	var str string
	asciiCount := map[string]int{}

	fmt.Scan(&str)

	for _, c := range str {
		asciiCount[string(c)] += 1
	}

	var ks []string
	for k := range asciiCount {
		ks = append(ks, k)
	}

	sort.Strings(ks)

	for _, k := range ks {
		fmt.Println(k, " => ", asciiCount[k])
	}
}
