package rotation

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func fetchFileContents(oriFileName string) []string {
	readFile, err := os.Open(oriFileName)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var filelines []string

	for fileScanner.Scan() {
		line := strings.Trim(fileScanner.Text(), " ")
		filelines = append(filelines, line)
	}
	readFile.Close()

	return filelines
}
