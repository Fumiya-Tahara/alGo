package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	// buf := make([]byte, 0, 64*1024)
	// sc.Buffer(buf, 1024*1024)
	sc.Split(bufio.ScanWords)

	n := scanInt()

	sList := make([]string, n)
	sLenList := make([]int, n)
	for i := 0; i < n; i++ {
		sList[i] = scanStr()
		sLenList[i] = len(sList[i])
	}

	sort.Ints(sLenList)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if len(sList[j]) == sLenList[i] {
				fmt.Printf("%s", sList[j])
			}
		}
	}
}
