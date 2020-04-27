package files

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func md5sum(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
