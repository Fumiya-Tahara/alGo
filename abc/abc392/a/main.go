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

	a1 := scanInt()
	a2 := scanInt()
	a3 := scanInt()

	if a1*a2 == a3 {
		fmt.Println("Yes")
	} else if a1*a3 == a2 {
		fmt.Println("Yes")
	} else if a2*a3 == a1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
