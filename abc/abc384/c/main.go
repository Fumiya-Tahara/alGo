package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	n := 5

	abcde := make([]int, n)
	for i := 0; i < n; i++ {
		abcde[i] = scan()
	}

	strs := []string{"A", "B", "C", "D", "E"}

	scoreMap := make(map[int][]string)
	for bits := 0; bits < (1 << n); bits++ {
		var score int
		var str string
		for i := 0; i < n; i++ {
			if bits&(1<<i) != 0 {
				score += abcde[i]
				str += strs[i]
			}
		}

		scoreMap[score] = append(scoreMap[score], str)
	}

	keys := make([]int, 0, len(scoreMap))
	for k := range scoreMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	for _, v := range keys {
		sort.Strings(scoreMap[v])
		for _, name := range scoreMap[v] {
			fmt.Println(name)
		}
	}
}
