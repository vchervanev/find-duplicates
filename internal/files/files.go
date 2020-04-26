package files

import (
	"fmt"
)

// Enumerate lists files starting from root
func Enumerate(root string) error {
	files := make(chan record, 10000)
	newScanner(root, files, 100).perform()
	for path := range files {
		fmt.Println(path)
	}
	return nil
}
