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
	// buf := make([]byte, 0, 64*1024)
	// sc.Buffer(buf, 1024*1024)
	sc.Split(bufio.ScanWords)

	s := scanStr()
	var answer string
	for _, v := range s {
		if string(v) == "2" {
			answer += string(v)
		}
	}
	fmt.Println(answer)
}
