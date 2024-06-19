package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxN   = 2 * 100000
	maxVal = 1000000000
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	buffer = make([]byte, maxN*12)
)

func scanStr() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func scanLineStr() []string {
	s := scanStr()

	return strings.Split(s, " ")
}

func scanLineInt() []int {
	s := scanStr()
	strList := strings.Split(s, " ")

	slice := make([]int, len(strList))
	for i, v := range strList {
		num, _ := strconv.Atoi(string(v))
		slice[i] = num
	}

	return slice
}

func main() {
	sc.Buffer(buffer, maxN*12)

	n := scanInt()
	grid := make([][]string, 729)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, 729)
	}

	l := 1
	grid[0][0] = "#"
	for k := 0; k < n; k++ {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x == 0 && y == 0 {
					continue
				}
				if x == 1 && y == 1 {
					for i := 0; i < l; i++ {
						for j := 0; j < l; j++ {
							grid[x*l+i][y*l+j] = "."
						}
					}
					continue
				}
				for i := 0; i < l; i++ {
					for j := 0; j < l; j++ {
						grid[x*l+i][y*l+j] = grid[i][j]
					}
				}
			}
		}
		l = 3 * l
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			fmt.Printf(grid[i][j])
		}
		fmt.Println()
	}
}
