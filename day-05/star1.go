package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from int64
	to int64
}

func main() {
	ans := 0
	data, _ := io.ReadAll(os.Stdin)
	D := string(data)
	DL := strings.Split(D, "\n\n")
	lines1 := strings.FieldsSeq(DL[0])
	lines2 := strings.FieldsSeq(DL[1])
	ranges := []Range{}

	for line := range lines1 {
		split := strings.Split(line, "-")
		f, _ := strconv.ParseInt(split[0], 10, 64)
		t, _ := strconv.ParseInt(split[1], 10, 64)
		ranges = append(ranges, Range{from:f, to: t})
	}

	for line := range lines2 {
		v, _ := strconv.ParseInt(line, 10, 64)
		for _, r := range ranges {
			if r.from < v && v <= r.to {
				ans++
				break
			}
		}
	}
	fmt.Println(ans)
}
