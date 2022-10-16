package rotation

import (
	"fmt"
	"log"
	"math"
	"os"
)

type PBM struct {
	width  int
	height int
	Pixels string
}

var header = "P1"

func Process(oriFileName string) {
	var fileContents = fetchFileContents(oriFileName)
	pbm, err := fetchPBM(fileContents)
	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}
	// fmt.Println(pbm)
	rstPixelsArr := rotate(pbm, 60)
	writeToFile(rstPixelsArr)
}

func rotate(pbm PBM, angle float32) [][]byte {
	oriPixelsArr := fetchOriPixelsArr(pbm)
	oriWidth := pbm.width
	oriHeight := pbm.height
	rstWidth, rstHeight, minX, minY, maxX, maxY := getRstPixelsSize(angle, oriWidth, oriHeight)

	fmt.Println(rstWidth, rstHeight, minX, minY, maxX, maxY)
	rstPixelsArr := initRstPixelsArr(rstHeight, rstWidth)

	for y := 0; y < int(rstHeight); y++ {
		for x := 0; x < int(rstWidth); x++ {

		}
	}

	for y := 0; y < oriHeight; y++ {
		for x := 0; x < oriWidth; x++ {
			rstXInx := oriHeight - y - 1
			rstYInx := x
			rstPixelsArr[rstYInx][rstXInx] = oriPixelsArr[y][x]
		}
	}

	for _, line := range rstPixelsArr {
		fmt.Println(string(line))
	}

	return rstPixelsArr
}

func getRstPixelsSize(angle float32, oriWidth int, oriHeight int) (int32, int32, float32, float32, float32, float32) {
	sine := getSine(angle)
	cosine := getCosine(angle)

	p1x := float32(oriWidth) * cosine
	p1y := -float32(oriWidth) * sine
	p2x := (float32(oriHeight) * sine) + (float32(oriWidth) * cosine)
	p2y := (float32(oriHeight) * cosine) - (float32(oriWidth) * sine)
	p3x := float32(oriHeight) * sine
	p3y := float32(oriHeight) * cosine
	// fmt.Println(oriWidth, oriHeight, p1x, p1y, p2x, p2y, p3x, p3y)
	minX := min(0, min(p3x, min(p1x, p2x)))
	minY := min(0, min(p3y, min(p1y, p2y)))
	maxX := max(p3x, max(p1x, p2x))
	maxY := max(p3y, max(p1y, p2y))

	rstWidth := int32(math.Ceil(float64(maxX - minX)))
	rstHeight := int32(math.Ceil(float64(maxY - minY)))

	return rstWidth, rstHeight, minX, minY, maxX, maxY
}

func getSine(angle float32) float32 {
	// TODO
	radian := degreeToRadian((angle))
	return float32(math.Sin(float64(radian)))
}

func getCosine(angle float32) float32 {
	radian := degreeToRadian((angle))
	return float32(math.Cos(float64(radian)))
}

func degreeToRadian(angle float32) float32 {
	return angle * (math.Pi / 180)
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

func initRstPixelsArr(height int32, width int32) [][]byte {
	rstPixelsArr := make([][]byte, height)
	for i := 0; i < int(height); i++ {
		rstPixelsArr[i] = make([]byte, width)
	}
	return rstPixelsArr
}

func fetchOriPixelsArr(pbm PBM) (pixelsArr [][]byte) {
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

func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}
