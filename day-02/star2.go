package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := 0

	path := "./i"
	in, _ := os.ReadFile(path)
	ranges := strings.SplitSeq(string(in), ",")

	for r := range ranges {
		r := strings.SplitN(r, "-", 2)
		left := r[0]
		right := r[1]
		lv, _ := strconv.ParseInt(left, 10, 64)
		rv, _ := strconv.ParseInt(right, 10, 64)
		for i := lv; i <= rv; i++ {
			test := false
			for _, v := range getDivisors(len(strconv.Itoa(int(i)))) {
				if repeats(strconv.Itoa(int(i)), v) {
					test = true
				}
			}
			if test {
				fmt.Println(i)
				ans += int(i)
			}
		}
	}

	fmt.Println(ans)
}

func repeats(test string, l int) bool {
	list := []string{}
	for i := 0; i < len(test); i += l {
		list = append(list, test[i:i+l])
	}
	first := list[0]
	for _, v := range list {
		if v != first {
			return false
		}
	}
	return true
}

func getDivisors(test int) []int {
	list := []int{}
	for i := 1; i < test; i++ {
		if test%i == 0 {
			list = append(list, i)
		}
	}
	return list
}
