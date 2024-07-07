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

func scanInt64() int64 {
	sc.Scan()
	i, err := strconv.ParseInt(sc.Text(), 10, 64)
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

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	sc.Buffer(buffer, maxVal)
	sc.Split(bufio.ScanWords)

	sx := scanInt64()
	sy := scanInt64()
	tx := scanInt64()
	ty := scanInt64()

	if (sx+sy)%2 == 1 {
		sx -= 1
	}

	if (tx+ty)%2 == 1 {
		tx -= 1
	}

	x := abs(sx - tx)
	y := abs(sy - ty)

	var answer int64
	if x < y {
		answer = y
	} else {
		answer = (x + y) / 2
	}

	fmt.Println(answer)
}
