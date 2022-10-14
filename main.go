package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func process(fileContents []string) {
	fileDataStr, err := validateFile(fileContents)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
	fmt.Println(fileDataStr)
}

func validateFile(fileContents []string) (string, error) {
	var isValid bool
	isValid, err := validatHeader(fileContents[0])
	if !isValid {
		return "", err
	}
	lineIndex, err := validatComment(fileContents)
	if err != nil {
		return "", err
	}
	isValid = validatSize(fileContents[lineIndex])
	if !isValid {
		return "", errors.New("file size is not correct")
	}
	wAndHArr := getWAndH(fileContents[lineIndex])
	lineIndex++
	fileDataStr, isValid := validatData(fileContents, lineIndex, wAndHArr)
	if !isValid {
		return "", errors.New("file data is not correct")
	}
	return fileDataStr, nil
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
	if lineIndex == 1 {
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

func validatData(fileContents []string, lineIndex int, wAndHArr []int) (string, bool) {
	fileDataStr := strings.Join(fileContents[lineIndex:], "")
	fileDataStr = strings.ReplaceAll(fileDataStr, " ", "")
	requiredDatalen := wAndHArr[0] * wAndHArr[1]
	
	pattern := fmt.Sprintf("^[0,1]{%d,}$", requiredDatalen)
	match, _ := regexp.MatchString(pattern, fileDataStr)
	if match {
		fileDataStr = fileDataStr[0:requiredDatalen]
	}

	return fileDataStr, match
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
	process(fileContents)

}
