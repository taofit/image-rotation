package main

import (
	"flag"
	"image-rotate/rotation"
)

func main() {
	var oriFileName string
	flag.StringVar(&oriFileName, "oriFileName", "bitmap.pbm", "generated pbm image file name")
	flag.Parse()
	rotation.Process(oriFileName)
}
