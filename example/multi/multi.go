package main

import (
	"time"

	"github.com/rockwell-uk/csync/waitgroup"
	"github.com/rockwell-uk/uiprogress"
)

func main() {
	waitTime := time.Millisecond * 100
	uiprogress.Start()

	var wg *waitgroup.WaitGroup = waitgroup.New()

	bar1 := uiprogress.AddBar(20).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func(w *waitgroup.WaitGroup) {
		defer wg.Done()
		for bar1.Incr() {
			time.Sleep(waitTime)
		}
	}(wg)

	bar2 := uiprogress.AddBar(40).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func(w *waitgroup.WaitGroup) {
		defer wg.Done()
		for bar2.Incr() {
			time.Sleep(waitTime)
		}
	}(wg)

	time.Sleep(time.Second)
	bar3 := uiprogress.AddBar(20).PrependElapsed().AppendCompleted()
	wg.Add(1)
	go func(w *waitgroup.WaitGroup) {
		defer wg.Done()
		for bar3.Incr() {
			time.Sleep(waitTime)
		}
	}(wg)

	wg.Wait()
}
