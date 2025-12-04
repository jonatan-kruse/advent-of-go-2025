package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ans := 0

	path := "./i"
	in, _ := os.ReadFile(path)
	lines := strings.Fields(string(in))
	m := map[int]map[int]rune{}
	jMax := len(lines[0]) - 1
	iMax := len(lines) - 1

	for i, line := range lines {
		if m[i] == nil {
			m[i] = map[int]rune{}
		}
		for j, c := range line {
			m[i][j] = c
		}
	}
	lastAns := -1
	for lastAns != ans {
		lastAns = ans
		for i := 0; i <= iMax; i++ {
			for j := 0; j <= jMax; j++ {
				if m[i][j] != '@' {
					continue
				}
				rolls := 0
				for ii := -1; ii < 2; ii++ {
					for jj := -1; jj < 2; jj++ {
						i2 := i + ii
						j2 := j + jj
						if inside(i2, j2, iMax, jMax) {
							if m[i2][j2] == '@' || m[i2][j2] == 'x' {
								rolls++
							}
						}
					}
				}
				if rolls < 5 {
					m[i][j] = 'x'
					ans++
				}
			}
		}
		for i := 0; i <= iMax; i++ {
			for j := 0; j <= jMax; j++ {
				if m[i][j] == 'x' {
					m[i][j] = '.'
				}
			}
		}
	}

	fmt.Println(ans)
}

func inside(x int, y int, width int, height int) bool {
	return x >= 0 && x <= width && y >= 0 && y <= width
}
