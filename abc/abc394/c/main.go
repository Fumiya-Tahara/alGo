package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
	sc.Split(bufio.ScanWords)

	s := scanStr()

	var builder strings.Builder
	var wCount int
	for _, v := range s {
		if v == 'W' {
			wCount++
		} else if v == 'A' && wCount != 0 {
			builder.WriteString("A" + strings.Repeat("C", wCount))
			wCount = 0
		} else {
			builder.WriteString(strings.Repeat("W", wCount) + string(v))
			wCount = 0
		}
	}

	if wCount != 0 {
		builder.WriteString(strings.Repeat("W", wCount))
		wCount = 0
	}

	answer := builder.String()
	fmt.Println(answer)

}
