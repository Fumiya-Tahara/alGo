package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() string {
	sc.Scan()

	return sc.Text()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	h, _ := strconv.Atoi(scan())
	w, _ := strconv.Atoi(scan())

	minX := 1000
	minY := 1000
	maxX := -1
	maxY := -1

	grid := make([][]string, 0, 1000)
	for i := 0; i < h; i++ {
		input := scan()
		line := make([]string, 0, 1000)
		for j := 0; j < w; j++ {
			str := string(input[j])
			if str == "#" {
				minX = min(minX, j)
				minY = min(minY, i)
				maxX = max(maxX, j)
				maxY = max(maxY, i)
			}
			line = append(line, str)
		}
		grid = append(grid, line)
	}

	flag := true

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if grid[i][j] == "." {
				flag = false
				break
			}
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
