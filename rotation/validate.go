package rotation

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func fetchPBM(fileContents []string) (PBM, error) {
	var isValid bool
	isValid, err := validateHeader(fileContents[0])
	if !isValid {
		return PBM{}, err
	}
	lineIndex, err := validateComment(fileContents)
	if err != nil {
		return PBM{}, err
	}
	isValid = validateSize(fileContents[lineIndex])
	if !isValid {
		return PBM{}, errors.New("file size is not correct")
	}
	wAndHArr := getWAndH(fileContents[lineIndex])
	lineIndex++
	pixelsInfo, succeed := compressPixels(fileContents, lineIndex, wAndHArr)
	if !succeed {
		return PBM{}, errors.New(pixelsInfo)
	}
	pbm := PBM{wAndHArr[0], wAndHArr[1], pixelsInfo}

	return pbm, nil
}

func validateHeader(fileHeader string) (bool, error) {
	if fileHeader != header {
		return false, errors.New("file header is not correct")
	}
	return true, nil
}

func validateComment(fileContents []string) (int, error) {
	lineIndex := 1
	for len(fileContents[lineIndex]) > 0 && fileContents[lineIndex][0] == '#' {
		lineIndex++
	}
	if lineIndex == 1 {
		return lineIndex, errors.New("not comment in the file")
	}

	return lineIndex, nil
}

func validateSize(size string) bool {
	match, _ := regexp.MatchString(`^[1-9][0-9]*\s{1,}[1-9][0-9]*$`, size)

	return match
}

func getWAndH(lineSize string) []int {
	var wAndHArr []int
	sizeArr := strings.Fields(lineSize)

	for _, number := range sizeArr {
		aSNumber, _ := strconv.Atoi(number)
		wAndHArr = append(wAndHArr, aSNumber)
	}

	return wAndHArr
}

func compressPixels(fileContents []string, lineIndex int, wAndHArr []int) (string, bool) {
	pixels := strings.Join(fileContents[lineIndex:], "")
	pixels = strings.ReplaceAll(pixels, " ", "")
	requiredPixelsLen := wAndHArr[0] * wAndHArr[1]

	if len(pixels) < requiredPixelsLen {
		return "pixel length is smaller than the size defined", false
	}
	for _, c := range pixels {
		if c != '0' && c != '1' {
			return "pixels contain illegal characters", false
		}
	}

	return pixels, true
}
