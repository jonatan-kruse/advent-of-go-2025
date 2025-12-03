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
	banks := strings.FieldsSeq(string(in))

	for bank := range banks {
		joltage := ""
		largestIndex := -1
		for j := 11; j >= 0; j-- {
			largest := 0
			for i := largestIndex + 1; i < len(bank)-j; i++ {
				value := int(bank[i] - '0')
				if value > largest {
					largest = value
					largestIndex = i
				}
			}
			joltage = joltage + strconv.Itoa(largest)
		}
		value, _ := strconv.ParseInt(joltage, 10, 64)
		ans += int(value)
	}

	fmt.Println(ans)
}
