package main

import (
	"bufio"
	"fmt"
	"os"
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

	n := scan()

	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = scan()
	}

	var answer int = 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if h[i] != h[j] {
				continue
			}

			count := 1
			skip := j - i
			for k := j; k < n; k = k + skip {
				if h[i] != h[k] {
					break
				}
				count++
			}
			if count > answer {
				answer = count
			}
		}
	}

	fmt.Println(answer)
}
