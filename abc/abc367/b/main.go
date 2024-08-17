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
	maxVal = 1000000
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

	x := scanStr()

	split := strings.Split(x, ".")

	small := split[1]

	if small == "000" {
		fmt.Println(split[0])

		return
	}

	var taleIsZero bool
	var count int
	for i := 0; i < 3; i++ {
		if i == 0 && (string(small[2-i]) == "0") {
			taleIsZero = true
			count++

			continue
		}

		if (string(small[2-i]) == "0") && taleIsZero {
			count++
		}
	}

	small = small[0 : 3-count]

	fmt.Println(split[0] + "." + small)

}
