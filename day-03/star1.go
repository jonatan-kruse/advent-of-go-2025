package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	ans := 0

	path := "./t"
	in, _ := os.ReadFile(path)
	banks := strings.FieldsSeq(string(in))

	for bank := range banks {
		largest := 0
		largestIndex := 0
		for i, r := range bank {
			value := int(r - '0')
			if i == len(bank)-1 {
				continue
			}
			if value > largest {
				largest = value
				largestIndex = i
			}
		}
		largest2 := 0
		for i := largestIndex + 1; i < len(bank); i++ {
			value := int(bank[i] - '0')
			largest2 = int(math.Max(float64(largest2), float64(value)))
		}
		value, _ := strconv.ParseInt(strconv.Itoa(largest)+strconv.Itoa(largest2), 10, 64)
		ans += int(value)
	}

	fmt.Println(ans)
}
