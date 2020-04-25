package files

import (
	"fmt"
	"os"
	"sync"
)

type dirs struct {
	queue chan string
	wg    sync.WaitGroup
}

func (d *dirs) pushDir(path string) {
	d.wg.Add(1)
	if d.queue == nil {
		d.queue = make(chan string, 10000)
	}
	if len(d.queue) == cap(d.queue) {
		fmt.Fprintln(os.Stderr, "Warning: Too many files")
	}
	d.queue <- path
}

func (d *dirs) done() {
	d.wg.Done()
}

func (d *dirs) wait() {
	d.wg.Wait()
}
