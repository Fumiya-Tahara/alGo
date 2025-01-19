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

	x := scan()

	count := 1
	for i := 1; ; i++ {
		count *= i
		if count == x {
			fmt.Println(i)

			return
		}
	}
}
