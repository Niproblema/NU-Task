package counter

import (
	"bufio"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type searchParams struct {
	Word         string
	RegexChecker *regexp.Regexp
}

type result struct {
	FilePath  string
	WordCount int
}

func CountRepository(root string, word string, caseSensitive bool, wholeOnly bool) int {
	var discoveryWG sync.WaitGroup
	var processWG sync.WaitGroup
	countChan := make(chan result)
	totalCount := 0
	searchParams := getSearchParams(word, caseSensitive, wholeOnly)

	// Start traversing recursively
	discoveryWG.Add(1)
	go discover(root, &searchParams, &discoveryWG, &processWG, countChan)

	// Await completion of discovery and processing before closing counting channel.
	go func() {
		discoveryWG.Wait()
		processWG.Wait()
		close(countChan)
	}()

	// While counting channel is still active, keep logging counted files.
	for result := range countChan {
		log.Printf("The word %q appears %d times in file %v", word, result.WordCount, result.FilePath)
		totalCount += result.WordCount
	}

	log.Printf("The word %q appears %d times in total under the directory %q", word, totalCount, root)

	return totalCount
}

func discover(root string, searchParams *searchParams, discoveryWG *sync.WaitGroup, processWG *sync.WaitGroup, countChan chan result) {
	defer discoveryWG.Done()

	walk := func(path string, d fs.DirEntry, err error) error {
		// Error expected in some scenarios like lack of permissions
		if err != nil {
			log.Printf("Error; error browsing path [%v] - [%v]", path, err)
			return nil
		}

		if d.IsDir() {
			// If directory, launch goroutine for further discovery
			if path != root {
				discoveryWG.Add(1)
				go discover(path, searchParams, discoveryWG, processWG, countChan)
				return filepath.SkipDir
			}
			return nil
		} else {
			// If file, launch goroutine for processing
			processWG.Add(1)
			go processFile(path, searchParams, processWG, countChan)
			return nil
		}
	}

	err := filepath.WalkDir(root, walk)
	if err != nil {
		log.Printf("error discovering the path %v: %v", root, err)
	}
}

func processFile(filePath string, searchParams *searchParams, processWG *sync.WaitGroup, countChan chan result) {
	defer processWG.Done()

	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("error opening file %v: %v", filePath, err)
		return
	}
	defer f.Close()
	fstat, err := f.Stat()
	if err != nil {
		log.Printf("error opening file %v: %v", filePath, err)
		return
	}
	wordLen := len((*searchParams).Word)
	fileSize := fstat.Size()
	count := 0

	// Only if file is larger than the actual
	if fileSize > int64(wordLen) {

		// Prepare a buffer, at least the size of the sought word, at most size of the file.
		buffSize := min(max(65536, int64(wordLen)), fileSize)
		buf := make([]byte, buffSize)
		// Have a support buffer for words that might come in between two reads
		var prefBuf string

		reader := bufio.NewReader(f)
		for {
			n, err := reader.Read(buf)

			if n > 0 {
				var textContent string
				if len(prefBuf) == 0 {
					textContent = string(buf[:n])
				} else {
					textContent = prefBuf + string(buf[:n])
				}

				count += len((*searchParams).RegexChecker.FindAllString(textContent, -1))
				prefBuf = textContent[max(0, len(textContent)-wordLen):]
			}

			if err != nil && err != io.EOF {
				log.Printf("error reading file %v: %v", filePath, err)
				return
			}
			if err == io.EOF {
				break
			}
		}
	}

	countChan <- result{FilePath: filePath, WordCount: count}
}

func getSearchParams(word string, caseSensitive bool, wholeOnly bool) searchParams {
	var regStr string
	if wholeOnly {
		regStr = `\b(` + word + `)\b`
	} else {
		regStr = word
	}

	if !caseSensitive {
		regStr = `(?i)` + regStr
	}
	cReg := regexp.MustCompile(regStr)
	return searchParams{Word: word, RegexChecker: cReg}
}
