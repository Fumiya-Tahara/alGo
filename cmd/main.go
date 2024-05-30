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

func main() {
	sc.Buffer(buffer, maxN*12)

	nm := scanLineInt()

	a := scanStr()
	strListA := strings.Split(a, " ")
	sliceA := make([][]int, nm[0])
	for i, v := range strListA {
		sliceA[i] = make([]int, 2)
		num, _ := strconv.Atoi(string(v))
		sliceA[i][0] = num
		sliceA[i][1] = 0
	}

	b := scanStr()
	strListB := strings.Split(b, " ")
	sliceB := make([][]int, nm[1])
	for i, v := range strListB {
		sliceB[i] = make([]int, 2)
		num, _ := strconv.Atoi(string(v))
		sliceB[i][0] = num
		sliceB[i][1] = 1
	}

	sliceC := append(sliceA, sliceB...)
	sort.Slice(sliceC, func(i, j int) bool {
		return sliceC[i][0] < sliceC[j][0]
	})

	flag := sliceC[0][1]
	for i := 0; i < len(sliceC); i++ {
		if i == 0 {
			continue
		}

		if flag == sliceC[i][1] && flag == 0 {
			fmt.Println("Yes")
			return
		}

		flag = sliceC[i][1]
	}

	fmt.Println("No")
}
