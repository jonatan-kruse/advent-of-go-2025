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
		dialAtPre := (dialAt + int(num)*mul)

		if dialAtPre > 99 {
			zeros += dialAtPre / 100
		}
		if dialAtPre < 1 {
			if dialAt != 0 {
				zeros++
			}
			zeros -= dialAtPre / 100
		}
		if dialAtPre < 0 {
			dialAt = (100 + (dialAtPre % 100)) % 100
		} else {
			dialAt = dialAtPre % 100
		}
	}

	fmt.Println(zeros)
}
