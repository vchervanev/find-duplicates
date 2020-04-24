package main

import (
	"os"

	files "github.com/vchervanev/find-dup-files/internal"
)

func main() {
	files.Files{Root: "/Users/vladimir.chervanev/go/src", Destination: os.Stdout}.Process()
}
