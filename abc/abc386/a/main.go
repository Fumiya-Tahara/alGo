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

	m := make(map[int]int)
	for i := 0; i < 4; i++ {
		num := scan()
		m[num]++
	}

	if len(m) == 2 {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
