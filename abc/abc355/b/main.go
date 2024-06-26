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

	a := scanLineInt()
	b := scanLineInt()

	aMap := make(map[int]bool, nm[0])
	for i := 0; i < nm[0]; i++ {
		aMap[a[i]] = true
	}

	c := append(a, b...)

	sort.Ints(c)

	var conA int
	for i := 0; i < nm[0]+nm[1]; i++ {
		if aMap[c[i]] {
			conA++
		} else {
			conA = 0
		}
		if conA >= 2 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
