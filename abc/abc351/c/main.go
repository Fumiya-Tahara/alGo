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

func checkBalls(balls []int) []int {
	if len(balls) <= 1 {
		return balls
	}
	if balls[len(balls)-1] == balls[len(balls)-2] {
		newBall := balls[len(balls)-1] + 1
		balls = balls[:len(balls)-2]
		balls = append(balls, newBall)
		return checkBalls(balls)
	}
	return balls
}

func main() {
	sc.Buffer(buffer, maxN*12)

	n := scanInt()
	a := scanLineInt()

	balls := make([]int, 0, n)
	for _, ball := range a {
		balls = append(balls, ball)
		balls = checkBalls(balls)
	}

	fmt.Println(len(balls))
}
