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

	n := scanInt()
	a := make([]string, n)
	b := make([]string, n)

	for i := 0; i < n; i++ {
		a[i] = scanStr()
	}
	for i := 0; i < n; i++ {
		b[i] = scanStr()
	}

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			for j := 0; j < n; j++ {
				if a[i][j] != b[i][j] {
					fmt.Println(i+1, j+1)
					return
				}
			}
		}
	}
}
