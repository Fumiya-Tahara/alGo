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
	a := scanLineInt()

	aMap := make(map[int]int)
	for i, v := range a {
		aMap[v] = i + 1
	}

	count := 0
	result := make([]string, 0, n-1)
	for i := 0; i < n; i++ {
		if aMap[i+1] == i+1 {
			continue
		}

		ansI := i + 1
		ansJ := aMap[i+1]
		ans := strconv.Itoa(ansI) + " " + strconv.Itoa(ansJ)
		result = append(result, ans)

		// マップ、配列は互いに依存しているため入れ替える前の値を使って入れ替える
		aMap[i+1], aMap[a[i]] = aMap[a[i]], aMap[i+1]
		a[i], a[ansJ-1] = a[ansJ-1], a[i]

		count++
	}

	fmt.Println(count)
	for _, v := range result {
		fmt.Println(v)
	}
}
