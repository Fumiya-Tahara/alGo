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

type Node struct {
	i, j int
}
type Queue struct {
	data []Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) push(i, j int) {
	node := Node{
		i: i,
		j: j,
	}

	q.data = append(q.data, node)
}

func (q *Queue) pop() Node {
	if len(q.data) == 0 {
		return Node{
			i: -1,
			j: -1,
		}
	}

	node := q.data[0]
	q.data = q.data[1:]

	return node
}

func main() {
	sc.Split(bufio.ScanWords)

	h := scanInt()
	w := scanInt()
	d := scanInt()

	queue := NewQueue()

	s := make([][]string, h)
	distance := make([][]int, h)
	var count int
	for i := 0; i < h; i++ {
		inputStr := scanStr()

		sLine := make([]string, w)
		distanceLine := make([]int, w)
		for j := 0; j < w; j++ {
			sLine[j] = string(inputStr[j])

			if sLine[j] == "H" {
				distanceLine[j] = 0
				queue.push(i, j)
				count++
			} else {
				distanceLine[j] = -1
			}
		}
		s[i] = sLine
		distance[i] = distanceLine
	}

	for !queue.isEmpty() {
		node := queue.pop()
		if node.i < 0 || node.j < 0 {
			break
		}
		i := node.i
		j := node.j

		if (i-1 >= 0) && (s[i-1][j] == ".") && (distance[i-1][j] < 0) {
			distance[i-1][j] = distance[i][j] + 1
			if distance[i-1][j] <= d {
				count++
			}
			queue.push(i-1, j)
		}
		if (i+1 < h) && (s[i+1][j] == ".") && (distance[i+1][j] < 0) {
			distance[i+1][j] = distance[i][j] + 1
			if distance[i+1][j] <= d {
				count++
			}
			queue.push(i+1, j)
		}
		if (j-1 >= 0) && (s[i][j-1] == ".") && (distance[i][j-1] < 0) {
			distance[i][j-1] = distance[i][j] + 1
			if distance[i][j-1] <= d {
				count++
			}
			queue.push(i, j-1)
		}
		if (j+1 < w) && (s[i][j+1] == ".") && (distance[i][j+1] < 0) {
			distance[i][j+1] = distance[i][j] + 1
			if distance[i][j+1] <= d {
				count++
			}
			queue.push(i, j+1)
		}
	}

	fmt.Println(count)
}
