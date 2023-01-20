package uiprogress

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/rockwell-uk/csync/waitgroup"
)

func TestStoppingPrintout(t *testing.T) {
	p := New()
	p.SetRefreshInterval(time.Millisecond * 10)

	var buffer = &bytes.Buffer{}
	p.SetOut(buffer)

	bar := p.AddBar(100)
	p.Start()

	var wg *waitgroup.WaitGroup = waitgroup.New()

	wg.Add(1)

	go func() {
		for i := 0; i <= 80; i = i + 10 {
			bar.Set(i)
			time.Sleep(time.Millisecond * 5)
		}

		wg.Done()
	}()

	wg.Wait()

	p.Stop()
	fmt.Fprintf(buffer, "foo")

	var wantSuffix = "[==============================>-------]\nfoo"

	if !strings.HasSuffix(buffer.String(), wantSuffix) {
		t.Errorf("Content that should be printed after stop not appearing on buffer. [%v]", buffer.String())
	}
}
