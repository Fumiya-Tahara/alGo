// 尺取り法Var
package main

import (
	"bufio"
	"fmt"
	"os"
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
	var j int
	for i := 0; i < n; i++ {
		limit := a[i] / 2

		for j < n && a[j] <= limit {
			j++
		}

		answer += int64(j)
	}

	fmt.Println(answer)
}
