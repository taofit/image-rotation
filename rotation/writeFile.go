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

	comment := fmt.Sprintln(`# This is the bitmap of a letter rotated 90 degrees clockwise`)
	f.WriteString(comment)

	// write size
	_, err = f.WriteString(fmt.Sprintf("%d %d\n", width, height))
	if err != nil {
		log.Fatal(err)
	}

	writePixels(f, rstPixelsArr)
	fmt.Println("done ....")
}