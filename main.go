package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func process() {
	// var fileContents = fetchFileContents()
	// var bitmap [][]int

}

func validateFile(fileContents []string) (bool, error) {
	var isValid bool
	isValid, err := validatHeader(fileContents[0])
	if !isValid {
		return false, err
	}
	lineIndex, err := validatComment(fileContents)
	if err != nil {
		return false, err
	}
	isValid = validatSize(fileContents[lineIndex])
	if !isValid {
		return false, errors.New("file size is not correct")
	}
	wAndHArr := getWAndH(fileContents[lineIndex])
	lineIndex++
	isValid = validatData(fileContents, lineIndex, wAndHArr)
	if !isValid {
		return false, errors.New("file data is not correct")
	}
	return true, nil
}

func validatHeader(fileHeader string) (bool, error) {
	var header = fileHeader
	if header != "P1" {
		return false, errors.New("file header is not correct")
	}
	return true, nil
}

func validatComment(fileContents []string) (int, error) {
	lineIndex := 1
	for len(fileContents[lineIndex]) > 0 && fileContents[lineIndex][0] == '#' {
		lineIndex++
	}
	if (lineIndex == 1) {
		return lineIndex, errors.New("not comment in the file")
	}
	
	return lineIndex, nil
}

func validatSize(size string) bool {
	match, _ := regexp.MatchString("^[1-9][0-9]* [1-9][0-9]*$", size)

	return match
}

func getWAndH(lineSize string) []int {
	var wAndH []int
	sizeArr := strings.Split(lineSize, " ")

	for _, number := range sizeArr {
		aSNumber, _ := strconv.Atoi(number)
		wAndH = append(wAndH, aSNumber)
	}

	return wAndH
}

func validatData(fileContents []string, lineIndex int, wAndHArr []int) bool {
	width := wAndHArr[0]
	height := wAndHArr[1]
	
	if len(fileContents[lineIndex:]) < height {
		return false
	}

	for i, line := range fileContents[lineIndex:] {
		pattern := fmt.Sprintf("^([0,1]\\s){%d}[0,1]$", width-1)
		match, _ := regexp.MatchString(pattern, line)
		if !match {
			return false
		}
		if i + 1 == height {
			break
		} 
	}

	return true
}

func fetchFileContents() []string {
	readFile, err := os.Open("bitmap")

	if err != nil {
		fmt.Println(err)
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

func main() {
	var fileContents = fetchFileContents()
	isValid, err := validateFile(fileContents)
	fmt.Println(isValid, err)
}
