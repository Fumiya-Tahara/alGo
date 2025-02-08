package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanStr() string {
	sc.Scan()

	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()

	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = scanInt()
	}

	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = scanInt()
	}

	clothNumKeyMap := make(map[int]int)
	personalNumKeyMap := make(map[int]int)

	for i := 0; i < n; i++ {
		clothNumKeyMap[q[i]] = p[i]
		personalNumKeyMap[i+1] = q[i]
	}

	for i := 1; i <= n; i++ {
		watchedPersonNum := clothNumKeyMap[i]
		fmt.Printf("%d ", personalNumKeyMap[watchedPersonNum])
	}
}
