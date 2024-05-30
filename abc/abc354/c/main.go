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

// ポイント：
// ・問題文をビジュアル化する
// ・順番をスライスの最後に追加する
func main() {
	sc.Buffer(buffer, maxN*12)

	n := scanInt()
	cardsList := make([][]int, n)

	for i := 0; i < n; i++ {
		cardsList[i] = scanLineInt()
		cardsList[i] = append(cardsList[i], i+1)
	}

	sort.Slice(cardsList, func(i, j int) bool {
		return cardsList[i][1] < cardsList[j][1]
	})

	var v int
	S := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if cardsList[i][0] >= v {
			v = cardsList[i][0]
			S = append(S, cardsList[i][2])
		}
	}

	sort.Ints(S)

	fmt.Println(len(S))
	for _, v := range S {
		fmt.Printf("%d ", v)
	}
}
