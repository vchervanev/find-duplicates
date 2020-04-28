package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

type record struct {
	path        string
	size        int64
	isDuplicate bool
}

func (r *record) markAsDuplicate() {
	r.isDuplicate = true
}

type scanner struct {
	filesQueue chan record
	poolSize   int
	dirs       dirs
}

func newScanner(path string, filesQueue chan record, poolSize int) *scanner {
	scn := scanner{filesQueue: filesQueue, poolSize: poolSize}
	scn.dirs.pushDir(path)
	return &scn
}

func (s *scanner) perform() {
	for i := 0; i < s.poolSize; i++ {
		go s.worker()
	}
	go s.waitForDirs()
}

func (s *scanner) waitForDirs() {
	s.dirs.wait()
	close(s.filesQueue)
}

func (s *scanner) worker() {
	for dir := range s.dirs.queue {
		s.visit(dir)
		s.dirs.done()
	}
	fmt.Println("end worker")
}

func (s *scanner) visit(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		path := path.Join(dir, file.Name())
		if file.IsDir() {
			s.dirs.pushDir(path)
		} else {
			s.filesQueue <- record{path: path, size: file.Size()}
		}
	}
}
