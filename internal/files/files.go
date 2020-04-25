package files

import (
	"fmt"
	"io"
)

type files struct {
	Root        string
	Destination io.StringWriter
}

// New returns instance of file loader
func New(root string) files {
	return files{Root: root}
}

// Enumerate lists files starting from Root and writes the results into Destination
func (f files) Enumerate() error {
	files := make(chan string, 10000)
	newScanner(f.Root, files, 100).perform()
	for path := range files {
		fmt.Println(path)
	}
	return nil
}
