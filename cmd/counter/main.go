package main

import (
	"flag"
	"fmt"

	"github.com/Niproblema/NU-Task/internal/counter"
)

func main() {
	// Define and parse command line flags
	path := flag.String("dir", ".", "Path to scan for word occurrences. Current folder as default value.")
	word := flag.String("word", "Corpus", "The word to count in the specified directory. Default value \"Corpus\"")
	caseSensitive := flag.Bool("case", false, "Whether or not word comparison should be case sensitive. Default false")
	whole := flag.Bool("whole", false, "Whether only standalone, whole words should be counted. Default false")
	flag.Parse()

	fmt.Println(counter.CountRepository(*path, *word, *caseSensitive, *whole))
}
