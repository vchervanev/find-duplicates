package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Files load configuration
type Files struct {
	Root        string
	Destination io.StringWriter
	filesQueue  chan os.FileInfo
	dirsQueue   chan string
	count       int
	totalSize   int
}

// Process lists files starting from Root and writes the results into Destination
func (f Files) Process() error {
	f.visit(f.Root)
	// _, err := f.Destination.WriteString(f.Root + "-> result\n")
	// if err != nil {
	// 	log.Println("Fatal: " + err.Error())
	// 	return err
	// }
	return nil
}

func (f Files) visit(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			f.visit(path + "/" + file.Name())
		} else {
			fmt.Println(formatFileInfo(path, file))
		}
	}
}

func formatFileInfo(path string, file os.FileInfo) string {
	return fmt.Sprintf("%s\t%s\t%d", path, file.Name(), file.Size())
}
