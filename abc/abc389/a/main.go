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

	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	s := scan()

	a, _ := strconv.Atoi(string(s[0]))
	b, _ := strconv.Atoi(string(s[2]))

	fmt.Println(a * b)
}
