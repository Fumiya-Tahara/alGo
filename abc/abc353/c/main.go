package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	maxN = 3 * 100000
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

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sc.Buffer(buffer, maxN*8)

	const x = 100000000

	n := scanInt()
	a := scanLineInt()

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	var sum int
	for i := 0; i < n; i++ {
		sum += a[i] * (n - 1)
	}

	var count, result int
	r := n
	for i := 0; i < n; i++ {
		r = bigger(r, i+1)
		for r > i+1 && a[i]+a[r-1] >= x {
			r--
		}
		count += n - r
	}

	result = sum - count*x

	fmt.Println(result)
}
