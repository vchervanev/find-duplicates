package main

import (
	"github.com/vchervanev/find-dup-files/internal/files"
)

func main() {
	files.New("/Users/vladimir.chervanev/go/src").Enumerate()
}
