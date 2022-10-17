package rotation

import (
	"fmt"
	"log"
	"os"
)

func writeToFile(rstPixelsArr [][]byte) {
	f, err := os.Create("rst-bitmap.pbm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	width := len(rstPixelsArr[0])
	height := len(rstPixelsArr)

	f.WriteString(header + "\n")

	comment := fmt.Sprintln(`# This is a rotated bitmap`)
	f.WriteString(comment)

	// write size
	_, err = f.WriteString(fmt.Sprintf("%d %d\n", width, height))
	if err != nil {
		log.Fatal(err)
	}

	writePixels(f, rstPixelsArr)
	fmt.Println("done ....")
}

func writePixels(f *os.File, rstPixelsArr [][]byte) {
	var newline byte = 10
	var emptySpace byte = 32

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
