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

	filesByMD5 := newFilesByMD5(1000)
	for _, paths := range pathsBySize {
		if len(paths) > 0 {
			for _, path := range paths {
				filesByMD5.register(path)
			}
		}
	}

	var dups, totalDups int

	for md5value, paths := range(filesByMD5) {
		if len(paths) == 1 {
			continue
		}

		fmt.Println("\nMD5:", md5value)
		dups++
		totalDups += len(paths)
		for _, path := range paths {
			fmt.Println("\t", path)
		}
	}

	fmt.Printf("dups: %d, total dups: %d\n", dups, totalDups)

	return nil
}
