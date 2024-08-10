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
	maxN   = 100000
	maxVal = 100000
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
	sc.Buffer(buffer, maxVal)

	n := scanInt()

	var maxLen int
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = scanStr()

		if len(s[i]) > maxLen {
			maxLen = len(s[i])
		}
	}

	for i := 0; i < maxLen; i++ {
		var newStr string
		var emptyCount int
		for j := 0; j < n; j++ {
			got := s[n-1-j]

			if len(got) >= i+1 {
				newStr += strings.Repeat("*", emptyCount)
				newStr += string(got[i])
				emptyCount = 0

				continue
			}

			emptyCount++
		}

		fmt.Println(newStr)
	}

}
