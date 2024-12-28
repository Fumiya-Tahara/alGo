package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() string {
	sc.Scan()
	s := sc.Text()

	return s
}

func main() {
	sc.Split(bufio.ScanWords)

	s := scan()

	sList := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		intChar, _ := strconv.Atoi(string(s[i]))
		sList[i] = intChar
	}

	var count int
	for i := 0; i < len(s); i++ {
		if sList[i] == 0 && i < len(s)-1 {
			if sList[i+1] == 0 {
				count++
				i++

				continue
			}
		}

		count++
	}

	fmt.Println(count)
}
