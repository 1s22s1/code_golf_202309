package main

import "fmt"

func main() {
	var str string
	asciiCount := map[string]int{}

	fmt.Scan(&str)

	for _, c := range str {
		asciiCount[string(c)] += 1
	}

	// TODO: 並び順を固定にしたい
	for key, value := range asciiCount {
		fmt.Println(key, value)
	}
}
