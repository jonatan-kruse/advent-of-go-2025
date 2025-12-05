package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	from int64
	to   int64
}

func main() {
	ans := 0
	data, _ := io.ReadAll(os.Stdin)
	D := string(data)
	DL := strings.Split(D, "\n\n")
	lines1 := strings.FieldsSeq(DL[0])
	ranges := []Range{}

	for line := range lines1 {
		split := strings.Split(line, "-")
		f, _ := strconv.ParseInt(split[0], 10, 64)
		t, _ := strconv.ParseInt(split[1], 10, 64)
		newr := Range{from: f, to: t}
		toDelete := []int{}
		for i, r := range ranges {
			delete := false
			if r.from <= newr.from && newr.from <= r.to {
				newr.from = r.from
				delete = true
			}
			if r.from <= newr.to && newr.to <= r.to {
				delete = true
				newr.to = r.to
			}
			if newr.from < r.from && newr.to > r.to {
				delete = true
			}
			if delete {
				toDelete = append(toDelete, i)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(toDelete)))
		for _, i := range toDelete {
			ranges = delete_at_index(ranges, i)
		}
		ranges = append(ranges, newr)
	}
	for _, r := range ranges {
		ans += int(r.to - r.from) + 1
	}
	fmt.Println(ans)
}

func delete_at_index(slice []Range, index int) []Range {
	return append(slice[:index], slice[index+1:]...)
}
