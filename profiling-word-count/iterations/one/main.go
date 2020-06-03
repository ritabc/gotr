// Iterative approach, (with no concurrency) takes about 190ms to process 1 file
// Takes about 2.4s to process all files in testdata
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/SpirentOrion/logrus"
)

func main() {
	if len(os.Args) == 1 {
		log.Error("No file given")
		return
	}
	// Otherwise, we'll assume files are given

	result := make(map[string]int)

	// Start a timer to time the work of this program
	start := time.Now()

	// Iterate over files given, process the files
	for _, fn := range os.Args[1:] {
		processFile(result, fn)
	}

	// get time the program took till now, but print the processing time last
	defer fmt.Printf("Processing took: %v\n", time.Since(start))

	printResult(result)

}

func processFile(result map[string]int, fn string) {
	var w string
	r, err := os.Open(fn)
	if err != nil {
		log.Warn(err)
		return
	}

	defer r.Close()

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		w = strings.ToLower(sc.Text())
		result[w] = result[w] + 1
	}
}

func printResult(result map[string]int) {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")
	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
