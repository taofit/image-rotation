package rotation

import (
	"fmt"
	"log"
	"os"
)

func writeToFile(rstPixelsArr [][]byte) {
	resultantFile := "rst-bitmap.pbm"
	f, err := os.Create(resultantFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	width := len(rstPixelsArr[0])
	height := len(rstPixelsArr)

	comment := fmt.Sprintln(`# This is a rotated bitmap`)
	size := fmt.Sprintf("%d %d\n", width, height)
	_, err = f.WriteString(header + "\n" + comment + size)

	if err != nil {
		log.Fatal(err)
	}

	writePixels(f, rstPixelsArr)
	fmt.Printf("Done! rotated image is saved in %s", resultantFile)
}

func writePixels(f *os.File, rstPixelsArr [][]byte) {
	var newline byte = 10
	var emptySpace byte = 32
	var rstPixels []byte

	for _, pixelsLine := range rstPixelsArr {
		rstPixelsLine := []byte{}
		for _, pixel := range pixelsLine {
			rstPixelsLine = append(rstPixelsLine, pixel, emptySpace)
		}
		rstPixelsLine = rstPixelsLine[:len(rstPixelsLine)-1]
		rstPixelsLine = append(rstPixelsLine, newline)
		rstPixels = append(rstPixels, rstPixelsLine...)
	}
	f.Write(rstPixels)
}
