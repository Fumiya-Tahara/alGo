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

	nm := scanLineInt()
	n := nm[0]
	m := nm[1]
	a := scanLineInt()
	sum := make([]int, m)

	for i := 0; i < n; i++ {
		num := scanLineInt()
		for j := 0; j < m; j++ {
			sum[j] += num[j]
		}
	}

	for i := 0; i < nm[1]; i++ {
		if sum[i] < a[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}
