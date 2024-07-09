package main

/**
Cho một mảng các số nguyên, sắp xếp các phần tử trong mảng sao cho, các phần tử lẻ đứng trước và giảm dần, các phần tử chẵn đứng sau và tăng dần. Xem thêm ví dụ để hiểu rõ hơn yêu cầu.

Đầu vào
Dòng đầu tiên là số lượng phần tử trong mảng.

Dòng thứ 2 là N phần tử trong mảng.

Giới hạn
1≤n≤10^6

1≤ai≤10^9

Đầu ra
In ra dãy đã được sắp xếp

Ví dụ :
Input 01
10
1 2 3 9 7 4 8 6 10 5
Output 01
9 7 5 3 1 2 4 6 8 10
*/
import (
	"fmt"
	"sort"
)

type Cmp []int

func (a Cmp) Len() int      { return len(a) }
func (a Cmp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Cmp) Less(i, j int) bool {
	if a[i]%2 == 0 && a[j]%2 == 0 {
		return a[i] < a[j]
	}
	if a[i]%2 == 1 && a[j]%2 == 1 {
		return a[i] > a[j]
	}
	if a[i]%2 == 0 && a[j]%2 == 1 {
		return false
	}
	return true
}
func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Sort(Cmp(a))
	fmt.Println("Sắp xếp:", a)
}
