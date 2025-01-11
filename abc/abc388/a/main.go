package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() string {
	sc.Scan()

	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	s := scan()

	answer := string(s[0]) + "UPC"

	fmt.Println(answer)
}
