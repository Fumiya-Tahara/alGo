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
	m := scanInt()

	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
	}

	var count int
	answerList := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		ok := true
		for j := 0; j < m; j++ {
			if i == a[j] {
				ok = false
			}
		}
		if ok {
			answerList = append(answerList, i)
			count++
		}
	}

	fmt.Println(count)
	for _, v := range answerList {
		fmt.Printf("%d ", v)
	}
}
