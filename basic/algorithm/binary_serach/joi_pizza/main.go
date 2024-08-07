// https://atcoder.jp/contests/joi2009ho/tasks/joi2009ho_b

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc.Buffer(buffer, maxVal)

	d := scanInt()
	n := scanInt()
	m := scanInt()

	dBranch := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		dBranch[i+1] = scanInt()
	}
	dBranch[n] = d

	sort.Ints(dBranch)

	deli := make([]int, m)
	for i := 0; i < m; i++ {
		deli[i] = scanInt()
	}

	var answer int
	for i := 0; i < m; i++ {
		left, right := 0, n

		for right-left > 1 {
			mid := (left + right) / 2

			if dBranch[mid] == deli[i] {
				break
			}

			if deli[i] < dBranch[mid] {
				right = mid
			} else {
				left = mid
			}

			if right-left == 1 {
				answer += min(deli[i]-dBranch[left], dBranch[right]-deli[i])
			}
		}
	}

	fmt.Println(answer)
}
