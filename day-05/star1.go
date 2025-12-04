package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ans := 0

	data, _ := io.ReadAll(os.Stdin)
	D := string(data)
	lines := strings.Fields(D)
	G := [][]rune{}
	for _, line := range strings.Fields(D) {
		G = append(G, []rune(line))
	}
	R, C := len(G), len(G[0])
	_, _, _ = lines, R, C

	fmt.Println(ans)
}
