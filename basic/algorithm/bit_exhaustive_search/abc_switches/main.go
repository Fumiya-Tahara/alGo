// https://atcoder.jp/contests/abc128/tasks/abc128_c

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
	maxVal = 2 * 10000000000000000
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
	nm := scanLineInt()
	n := nm[0]
	m := nm[1]

	k := make([]int, m)
	s := make([][]int, m)
	for i := 0; i < m; i++ {
		ks := scanLineInt()
		k[i] = ks[0]
		s[i] = ks[1:]
	}

	p := scanLineInt()

	var answer int
	for i := 0; i < (1 << n); i++ {
		var lightOnCount int
		for j := 0; j < m; j++ {
			var onCount int
			for l := 0; l < k[j]; l++ {
				if i&(1<<(s[j][l]-1)) != 0 {
					onCount++
				}
			}

			if onCount%2 != p[j] {
				break
			}

			lightOnCount++
		}
		if lightOnCount == m {
			answer++
		}
	}

	fmt.Println(answer)
}
