package version

import (
	"fmt"
	"os"
)

var (
	Version = ""
)

func PrintVersion() {
	fmt.Printf("version: %s\n", Version)
	os.Exit(0)
}
