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
func New(root string, destination io.StringWriter) files {
	return files{Root: root, Destination: destination}
}

// 	filesQueue  chan os.FileInfo
// 	dirsQueue   chan string
// 	count       int
// 	totalSize   int
// }

// Enumerate lists files starting from Root and writes the results into Destination
func (f files) Enumerate() error {
	fmt.Println("enumerating")
	// f.worker()
	// _, err := f.Destination.WriteString(f.Root + "-> result\n")
	// if err != nil {
	// 	log.Println("Fatal: " + err.Error())
	// 	return err
	// }
	return nil
}

// func (f Files) worker(id int) {
// 	for path := range f.dirsQueue {
// 		fmt.Println("before", len(f.dirsQueue))
// 		f.visit(path)
// 		fmt.Println("after", len(f.dirsQueue))
// 	}
// }

// func formatFileInfo(path string, file os.FileInfo) string {
// 	return fmt.Sprintf("%s\t%s\t%d", path, file.Name(), file.Size())
// }
