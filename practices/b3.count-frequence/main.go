package main

/**
Cho mảng các số nguyên không âm gồm N phần tử, thực hiện đếm tần suất xuất hiện của các phần tử và in theo mẫu.

Đầu vào
Dòng đầu tiên là số lượng phần tử trong mảng.

Dòng thứ 2 là N phần tử trong mảng.

Giới hạn
2≤n≤10^6

0≤ai≤10^7

Đầu ra
In ra tần suất xuất hiện của các phần tử theo thứ tự từ nhỏ tới lớn sau đó bỏ trống 1 dòng và in ra tần suất xuất hiện của các phần tử theo thứ tự xuất hiện trong mảng(mỗi giá trị chỉ liệt kê 1 lần).

Ví dụ :
Input 01

8
2 1 2 3 4 8 2 3
Output 01

1 1
2 3
3 2
4 1
8 1

2 3
1 1
3 2
4 1
8 1
*/
import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	var myMap = make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		myMap[a[i]]++
	}

	var keys []int
	for key := range myMap {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Println(key, myMap[key])
	}

	fmt.Println()

	for _, item := range a {
		value, exists := myMap[item]
		if exists {
			fmt.Println(item, value)
			delete(myMap, item)
		}
	}
}
