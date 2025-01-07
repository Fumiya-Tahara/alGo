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

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) push(i int) {

	q.data = append(q.data, i)
}

func (q *Queue) pop() int {
	if len(q.data) == 0 {
		return -1
	}

	data := q.data[0]
	q.data = q.data[1:]

	return data
}

func main() {
	sc.Split(bufio.ScanWords)

	h := scanInt()
	w := scanInt()

	s := make([]int, 2)
	s[0] = scanInt()
	s[1] = scanInt()
	start := s[0]*w + s[1]

	g := make([]int, 2)
	g[0] = scanInt()
	g[1] = scanInt()
	goal := g[0]*w + g[1]

	maze := make([][]string, h+1)
	for i := 1; i <= h; i++ {
		inputLine := scanStr()
		mazeLine := make([]string, w+1)
		for j := 1; j <= w; j++ {
			mazeLine[j] = string(inputLine[j-1])
		}

		maze[i] = mazeLine
	}

	// グラフを隣接リストとして構築
	G := make(map[int][]int)
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if maze[i][j] == "." {
				idx := i*w + j
				// 横方向の辺
				if j < w && maze[i][j+1] == "." {
					rightIdx := i*w + (j + 1)
					G[idx] = append(G[idx], rightIdx)
					G[rightIdx] = append(G[rightIdx], idx)
				}
				// 縦方向の辺
				if i < h && maze[i+1][j] == "." {
					downIdx := (i+1)*w + j
					G[idx] = append(G[idx], downIdx)
					G[downIdx] = append(G[downIdx], idx)
				}
			}
		}
	}

	dist := make(map[int]int)
	for i := 1; i <= h*w; i++ {
		dist[i] = -1
	}
	dist[start] = 0

	queue := NewQueue()
	queue.push(start)

	for !queue.isEmpty() {
		pos := queue.pop()

		for _, nex := range G[pos] {
			if dist[nex] == -1 {
				dist[nex] = dist[pos] + 1
				queue.push(nex)
			}
		}
	}

	fmt.Println(dist[goal])
}
