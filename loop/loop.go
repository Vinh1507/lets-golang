package loop

import "fmt"

func LoopHandler() {
	//
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{2, 4, 6, 8, 10}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
