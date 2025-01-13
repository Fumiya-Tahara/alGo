// 二分探索Ver
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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scan()
	}

	var answer int64
	for i := 0; i < n; i++ {
		limit := a[i] / 2
		count := sort.Search(len(a), func(i int) bool {
			return a[i] > limit
		})
		answer += int64(count)
	}

	fmt.Println(answer)
}
