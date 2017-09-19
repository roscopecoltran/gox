package core

import (
	"fmt"
)

// GoxVersion is the version of the build
var GoxVersion = "undefined"

func PrintVersion() int {
	fmt.Printf("Version: %s \n", GoxVersion)
	return 0
}
