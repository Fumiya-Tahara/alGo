package toolset

import (
	"sort"
)

// 整数のスライスをソートする。
func Sort() {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	// 昇順
	sort.Ints(slice)
}
