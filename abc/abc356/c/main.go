package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	maxN   = 2 * 100000
	maxVal = 1000000000
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	buffer = make([]byte, maxN*12)
)

func scanStr() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func scanLineStr() []string {
	s := scanStr()

	return strings.Split(s, " ")
}

func scanLineInt() []int {
	s := scanStr()
	strList := strings.Split(s, " ")

	slice := make([]int, len(strList))
	for i, v := range strList {
		num, _ := strconv.Atoi(string(v))
		slice[i] = num
	}

	return slice
}

func main() {
	sc.Buffer(buffer, maxN*12)
	sc.Split(bufio.ScanWords)

	n := scanInt()
	m := scanInt()
	k := scanInt()

	as := make([]int, m)
	r := make([]string, m)
	var c int
	for i := 0; i < m; i++ {
		c = scanInt()
		for j := 0; j < c; j++ {
			a := scanInt()
			a--
			as[i] |= 1 << a
		}
		r[i] = scanStr()
	}

	var ans int
	for i := 0; i < 1<<n; i++ {
		ok := true
		for j := 0; j < m; j++ {
			var cnt int
			check := as[j] & i
			for check != 0 {
				cnt += check & 1
				check >>= 1
			}
			if (cnt >= k) != (r[j] == "o") {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}

	fmt.Println(ans)
}
