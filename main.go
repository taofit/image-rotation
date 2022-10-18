package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PBM struct {
	width  int
	height int
	Pixels string
}

var header = "P1"

func process(fileContents []string) {
	pbm, err := fetchPBM(fileContents)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
	// fmt.Println(pbm)
	rstPixelsArr := rotate(pbm)
	writeToFile(rstPixelsArr)
}

func rotate(pbm PBM) [][]byte {
	oriPixelsArr := fetchPixelsArr(pbm)
	width := pbm.width
	height := pbm.height
	rstPixelsArr := initRstPixelsArr(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rstXInx := height - y - 1
			rstYInx := x
			rstPixelsArr[rstYInx][rstXInx] = oriPixelsArr[y][x]
		}
	}

	for _, line := range rstPixelsArr {
		fmt.Println(string(line))
	}

	return rstPixelsArr
}

func writeToFile(rstPixelsArr [][]byte) {
	f, err := os.Create("rst-bitmap.pbm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	width := len(rstPixelsArr[0])
	height := len(rstPixelsArr)

	f.WriteString(header + "\n")

	comment := fmt.Sprintln(`# This is the bitmap of a letter rotated 90 degrees clockwise`)
	f.WriteString(comment)

	// write size
	_, err = f.WriteString(fmt.Sprintf("%d %d\n", width, height))
	if err != nil {
		log.Fatal(err)
	}

	writePixels(f, rstPixelsArr)
	fmt.Println("done")
}

func writePixels(f *os.File, rstPixelsArr [][]byte) {
	newline := byte(10)
	emptySpace := byte(32)

	for _, pixelsLine := range rstPixelsArr {
		rstPixelsLine := []byte{}
		for _, pixel := range pixelsLine {
			rstPixelsLine = append(rstPixelsLine, pixel, emptySpace)
		}
		rstPixelsLine = rstPixelsLine[:len(rstPixelsLine)-1]
		rstPixelsLine = append(rstPixelsLine, newline)
		f.Write(rstPixelsLine)
	}
}

func initRstPixelsArr(height int, width int) [][]byte {
	rstPixelsArr := make([][]byte, height)
	for i := 0; i < height; i++ {
		rstPixelsArr[i] = make([]byte, width)
	}
	return rstPixelsArr
}

func fetchPixelsArr(pbm PBM) (pixelsArr [][]byte) {
	width := pbm.width
	height := pbm.height
	pixels := []byte(pbm.Pixels)

	for y := 0; y < height; y++ {
		pixelLineArr := []byte{}
		for x := 0; x < width; x++ {
			tempX := y*width + x
			pixelLineArr = append(pixelLineArr, pixels[tempX:tempX+1]...)
		}
		pixelsArr = append(pixelsArr, pixelLineArr)
	}
	return
}

func fetchPBM(fileContents []string) (PBM, error) {
	var isValid bool
	isValid, err := validatHeader(fileContents[0])
	if !isValid {
		return PBM{}, err
	}
	lineIndex, err := validatComment(fileContents)
	if err != nil {
		return PBM{}, err
	}
	isValid = validatSize(fileContents[lineIndex])
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

func validatHeader(fileHeader string) (bool, error) {
	if fileHeader != header {
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
	match, _ := regexp.MatchString(`^[1-9][0-9]*\s{1,}[1-9][0-9]*$`, size)

	return match
}

func getWAndH(lineSize string) []int {
	var wAndH []int
	sizeArr := strings.Fields(lineSize)

	for _, number := range sizeArr {
		aSNumber, _ := strconv.Atoi(number)
		wAndH = append(wAndH, aSNumber)
	}

	return wAndH
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

func main() {
	var oriFileName string
	flag.StringVar(&oriFileName, "oriFileName", "bitmap.pbm", "generated pbm image file name")
	flag.Parse()
	var fileContents = fetchFileContents(oriFileName)
	process(fileContents)
}
