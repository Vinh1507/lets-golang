package main

/**
Tìm kiếm nhị phân

8
1 4 8 12 19 20 21 25
10
4
YES
5
NO
6
NO
7
NO
8
YES
9
NO
12
YES
13
NO
25
YES
26
NO
*/
import (
	"fmt"
)

func binarySearch(a []int, target int) int {
	l, r := 0, len(a)-1
	for l <= r {
		var mid int = (l + r) / 2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			l++
		} else {
			r--
		}
	}
	return -1
}
func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var query int
	fmt.Scan(&query)

	for i := 0; i < query; i++ {
		var target int
		fmt.Scan(&target)

		if binarySearch(a, target) > 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
