package files

import (
	"fmt"
)

// ListDuplicates does everything
func ListDuplicates(root string) {
	files := enumerate(root)
	index := indexFiles(files)
	detectDuplicates(files, index)
	calculateStatuses(files)
}

func enumerate(root string) []record {
	filesChan := make(chan record, 10000)
	newScanner(root, filesChan, 100).perform()
	files := make([]record, 40000)
	for rec := range filesChan {
		files = append(files, rec)
	}
	return files
}

func calculateStatuses(files []record) {
}

func detectDuplicates(files []record, index map[string]record) {
	pathsBySize := newStringsByInt(1000)
	for _, rec := range files {
		pathsBySize.register(rec)
	}
	filesByMD5 := newFilesByMD5(len(pathsBySize) / 2)
	for _, paths := range pathsBySize {
		if len(paths) > 0 {
			for _, path := range paths {
				filesByMD5.register(path)
			}
		}
	}

	// var dups, totalDups int

	for md5value, paths := range filesByMD5 {
		if len(paths) == 1 {
			continue
		}

		fmt.Println("\nMD5:", md5value)
		// 	dups++
		// 	totalDups += len(paths)
		for _, path := range paths {
			// &(index[path]).markAsDuplicate()
			fmt.Println("\t", path)
		}
	}

	// fmt.Printf("dups: %d, total dups: %d\n", dups, totalDups)
}

func indexFiles(files []record) map[string]record {
	index := make(map[string]record, len(files))
	for _, file := range files {
		index[file.path] = file
	}
	return index
}
