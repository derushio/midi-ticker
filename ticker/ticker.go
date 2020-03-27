package ticker

import (
	"fmt"
	"time"
)

type Note struct {
	id        int32
	deltaTime int64
}

func average(xs []int64) int64 {
	total := int64(0)
	for _, v := range xs {
		total += v
	}
	return total / int64(len(xs))
}

func Ticker() {
	var notes []Note
	for i := 0; i < 100; i++ {
		notes = append(notes, Note{int32(i), int64(100)})
	}

	ticker := time.NewTicker(time.Millisecond)
	done := make(chan bool)

	var laptimes []int64
	go func() {
		lap := time.Now().UnixNano()
		delta := int64(0)
		index := 0

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				laptime := (t.UnixNano() - lap) / 1000
				laptimes = append(laptimes, laptime)
				lap = t.UnixNano()

				delta += laptime
				if notes[index].deltaTime < delta {
					fmt.Println(notes[index])
					index++
					delta = delta - notes[index].deltaTime

					if len(notes)-1 <= index {
						ticker.Stop()
						fmt.Println(string(average(laptimes)) + "microsec")
						fmt.Println("Ticker stopped")
						break
					}
				}
			}
		}
	}()

	time.Sleep(100000000000000000)
}
