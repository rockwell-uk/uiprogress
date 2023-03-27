package main

import (
	"time"

	"github.com/rockwell-uk/csync/waitgroup"
	"github.com/rockwell-uk/uiprogress"
)

func main() {
	uiprogress.Start()

	var wg *waitgroup.WaitGroup = waitgroup.New()

	bar := uiprogress.AddBar(100)

	bar.AppendCompleted()
	bar.PrependElapsed()

	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 50)
		wg.Add(1)
		go func(j int) {
			bar.Set(j)
			wg.Done()
		}(i)
	}
	bar.Set(100)
	time.Sleep(time.Millisecond * 50)

	wg.Wait()

	uiprogress.Stop()

	uiprogress.Start()

	bar1 := uiprogress.AddBar(100)

	bar1.AppendCompleted()
	bar1.PrependElapsed()

	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 50)
		wg.Add(1)
		go func(j int) {
			bar1.Set(j)
			wg.Done()
		}(i)
	}
	bar1.Set(100)
	time.Sleep(time.Millisecond * 50)

	wg.Wait()

	uiprogress.Stop()
}
