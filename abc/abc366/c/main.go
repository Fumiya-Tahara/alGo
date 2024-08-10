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

	q := scanInt()

	balls := map[int]int{}
	output := make([]int, 0, q)
	for i := 0; i < q; i++ {
		query := scanLineInt()

		queryNum := query[0]

		if queryNum == 3 {
			output = append(output, len(balls))
			continue
		}

		num := query[1]

		if queryNum == 1 {
			balls[num]++
		} else if queryNum == 2 {
			balls[num]--
			if balls[num] == 0 {
				delete(balls, num)
			}
		}
	}

	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}

}
