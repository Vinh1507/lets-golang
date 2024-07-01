package main

import (
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)       // có 1 thread cần chờ
	go sendTicket() // tác vụ tốn thời gian
	go sendTicket() // tác vụ tốn thời gian
	wg.Wait()
}
func sendTicket() {
	time.Sleep(2 * time.Second)
	print("Sent!")
	wg.Done()
}
