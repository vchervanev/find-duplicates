package files

import (
	"fmt"
)

// Enumerate lists files starting from root
func Enumerate(root string) error {
	files := make(chan record, 10000)
	newScanner(root, files, 100).perform()
	pathsBySize := newStringsByInt(1000)
	for rec := range files {
		pathsBySize.register(rec)
	}

	var uniq, multi int64
	for size, paths := range pathsBySize {
		switch len(paths) {
		case 1:
			uniq++
		default:
			multi++
			fmt.Println("Size:", size)
			for _, path := range paths {
				fmt.Println(path, md5sum(path))
			}
		}
	}

	fmt.Printf("Uniq: %d, Multi: %d\n", uniq, multi)

	return nil
}
