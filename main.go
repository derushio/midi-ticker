package main

import (
	"fmt"
	"time"
)

func main() {
	ticker()
}

func average(xs []int64) int64 {
	total := int64(0)
	for _, v := range xs {
		total += v
	}
	return total / int64(len(xs))
}

func ticker() {
	ticker := time.NewTicker(100 * time.Microsecond)
	done := make(chan bool)

	var laptimes []int64
	go func() {
		lap := time.Now().UnixNano()

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				laptime := (t.UnixNano() - lap) / 1000
				laptimes = append(laptimes, laptime)
				lap = t.UnixNano()
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)
	ticker.Stop()
	fmt.Println(average(laptimes))
	done <- true
	fmt.Println("Ticker stopped")
}
