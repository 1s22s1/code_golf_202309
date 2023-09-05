package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var strs []string

	for scanner.Scan() {
		strs = append(strs, strings.ReplaceAll(scanner.Text(), " ", ""))
	}

	for _, str := range strs {
		fmt.Println(str)
	}
}
