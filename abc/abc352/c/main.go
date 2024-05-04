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

	l := make([][]int, n)

	topFaceSize := 0
	topIndex := 0
	for i := 0; i < n; i++ {
		cl := scanLineInt()
		faceSize := cl[1] - cl[0]
		if faceSize > topFaceSize {
			topFaceSize = faceSize
			topIndex = i
		}
		l[i] = cl
	}

	height := 0
	for i := 0; i < n; i++ {
		if i == topIndex {
			continue
		}
		height += l[i][0]
	}

	height += l[topIndex][1]
	fmt.Println(height)
}
