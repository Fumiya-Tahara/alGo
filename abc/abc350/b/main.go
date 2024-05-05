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

	nq := scanLineInt()
	n := nq[0]
	q := nq[1]
	t := scanLineInt()

	teeth := make([]int, n)

	for i := 0; i < q; i++ {
		target := teeth[t[i]-1]
		if target == 0 {
			teeth[t[i]-1] = 1
		} else {
			teeth[t[i]-1] = 0
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if teeth[i] == 1 {
			continue
		}
		count++
	}

	fmt.Println(count)
}
