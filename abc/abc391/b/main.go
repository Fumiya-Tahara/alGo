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
	m := scanInt()

	s := make([]string, 0, 50)
	for i := 0; i < n; i++ {
		input := scanStr()
		s = append(s, input)
	}

	t := make([]string, 0, 50)
	for i := 0; i < m; i++ {
		input := scanStr()
		t = append(t, input)
	}

	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			ok := true
			for k := 0; k < m; k++ {
				for l := 0; l < m; l++ {
					if string(s[i+k][j+l]) != string(t[k][l]) {
						ok = false
					}
				}
			}
			if ok {
				fmt.Println(i+1, j+1)
				return
			}
		}
	}
}
