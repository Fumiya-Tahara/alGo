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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isOneEditDistance(s, t string) bool {
	lenS, lenT := len(s), len(t)

	if abs(lenS-lenT) > 1 {
		return false
	}

	if lenS > lenT {
		s, t = t, s
		lenS, lenT = lenT, lenS
	}

	for i := 0; i < lenS; i++ {
		if s[i] != t[i] {
			if lenS == lenT {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}

	return lenS+1 == lenT
}

func main() {
	buf := make([]byte, 0, 64*1024)
	sc.Buffer(buf, 1024*1024)
	sc.Split(bufio.ScanWords)

	_ = scan()
	s := scan()
	t := scan()

	if s == t {
		fmt.Println("Yes")
		return
	}

	if isOneEditDistance(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
