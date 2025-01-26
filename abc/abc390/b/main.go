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

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = int64(scan())
	}

	if n == 2 {
		fmt.Println("Yes")
		return
	}

	for i := 1; i < n-1; i++ {
		if a[i]*a[i] != a[i-1]*a[i+1] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
