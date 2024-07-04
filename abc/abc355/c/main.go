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

func main() {
	sc.Buffer(buffer, maxN*12)

	nt := scanLineInt()
	n := nt[0]
	t := nt[1]

	a := scanLineInt()

	row := make([]int, n)
	col := make([]int, n)

	diagonal := make([]int, 2)

	for i := 0; i < t; i++ {
		x := (a[i] - 1) / n
		y := (a[i] - 1) % n

		row[x]++
		if row[x] == n {
			fmt.Println(i + 1)
			return
		}

		col[y]++
		if col[y] == n {
			fmt.Println(i + 1)
			return
		}

		if x == y {
			diagonal[0]++
			if diagonal[0] == n {
				fmt.Println(i + 1)
				return
			}
		}

		if x+y == n-1 {
			diagonal[1]++
			if diagonal[1] == n {
				fmt.Println(i + 1)
				return
			}
		}
	}

	fmt.Println(-1)
}
