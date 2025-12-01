package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dialAt := 50
	zeros := 0

	path := "./i"
	in, err := os.ReadFile(path)
	check(err)
	dials := strings.Fields(string(in))

	for _, dial := range dials {
		mul := 1
		if dial[0:1] == "L" {
			mul = -1
		}
		num, _ := strconv.ParseInt(dial[1:], 10, 64)
		dialAt = (dialAt + int(num)*mul) % 100
		if dialAt == 0 {
			zeros++
		}
	}

	fmt.Println(zeros)
}
