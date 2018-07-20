package main

import (
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
}
