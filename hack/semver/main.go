package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("semver to parse is required")
	}
	toParse := os.Args[1]
	sv := semver.New(toParse)
	fmt.Printf("%d %d %d", sv.Major, sv.Minor, sv.Patch)
}
