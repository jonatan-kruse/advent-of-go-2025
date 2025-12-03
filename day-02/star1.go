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
			if isDouble(strconv.Itoa(int(i))) {
				ans += int(i)
			}
		}
	}

	fmt.Println(ans)
}

func isDouble(test string) bool {
	return len(test)%2 == 0 && test[0:len(test)/2] == test[len(test)/2:]
}
