package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanStr() string {
	sc.Scan()

	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scanInt()
	q := scanInt()

	var answer int
	pigeonMap := make(map[int]int)
	sMap := make(map[int]int)
	for i := 1; i < n+1; i++ {
		pigeonMap[i] = i
		sMap[i] = 1
	}

	for i := 0; i < q; i++ {
		num := scanInt()
		switch num {
		case 1:
			p := scanInt()
			h := scanInt()

			curHouse := pigeonMap[p]
			sMap[curHouse]--
			if sMap[curHouse] == 1 {
				answer--
			}

			pigeonMap[p] = h
			sMap[h]++
			if sMap[h] == 2 {
				answer++
			}
		case 2:
			fmt.Println(answer)
		}
	}
}
