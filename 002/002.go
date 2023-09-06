package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)

	for _, c := range str {
		fmt.Println(c)
	}
}
