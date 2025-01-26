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

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	correctSlice := []int{1, 2, 3, 4, 5}
	wrong := make([]int, 0, 2)
	var wrongCount int
	for i := 0; i < 5; i++ {
		input := scan()
		if input != correctSlice[i] {
			wrong = append(wrong, i)
			wrongCount++

			if wrongCount > 2 {
				fmt.Println("No")
				return
			}
		}
	}

	if wrongCount == 0 || abs(wrong[0]-wrong[1]) > 1 {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
