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

type Queue struct {
	snakes []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) push(i int) {
	q.snakes = append(q.snakes, i)
}

func (q *Queue) pop() int {
	if len(q.snakes) == 0 {
		return -1
	}

	snake := q.snakes[0]
	q.snakes = q.snakes[1:]

	return snake
}

func main() {
	sc.Split(bufio.ScanWords)

	q := scan()

	snakeQueue := NewQueue()

	queueList := make([]int, 0, q)
	queueList = append(queueList, 0)
	var minusCount int
	var totalLenMinus int
	for i := 0; i < q; i++ {
		queryNum := scan()
		if queryNum == 1 {
			l := scan()
			snakeQueue.push(l)
			queueList = append(queueList, queueList[len(queueList)-1]+l)
		} else if queryNum == 2 {
			minusCount++
			totalLenMinus += snakeQueue.pop()
		} else if queryNum == 3 {
			k := scan()
			fmt.Println(queueList[k+minusCount-1] - totalLenMinus)
		}
	}
}
