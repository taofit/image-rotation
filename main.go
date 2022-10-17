package main

import (
	"flag"
	"image-rotate/rotation"
)

var (
	oriFileName string
	angle       float64
)

func main() {
	flag.StringVar(&oriFileName, "fileName", "bitmap.pbm", "generated pbm image file name")
	flag.Float64Var(&angle, "degree", 90, "the number of degrees to rotate")
	flag.Parse()
	rotation.Process(oriFileName, angle)
}
