package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	d := scan()

	tl := make([][2]int, n)
	for i := 0; i < n; i++ {
		tl[i][0] = scan()
		tl[i][1] = scan()
	}

	for i := 1; i <= d; i++ {
		weight := make([]int, n)
		for j := 0; j < n; j++ {
			weight[j] = tl[j][0] * (tl[j][1] + i)
		}

		sort.Ints(weight)
		fmt.Println(weight[len(weight)-1])
	}
}
