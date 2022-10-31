package rotation

import (
	"log"
	"math"
	"sync"
)

type PBM struct {
	width  int
	height int
	Pixels string
}

var header = "P1"

func Process(oriFileName string, angle float64) {
	var fileContents = fetchFileContents(oriFileName)
	pbm, err := fetchPBM(fileContents)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	rstPixelsArr := rotate(pbm, angle)
	writeToFile(rstPixelsArr)
}

func rotate(pbm PBM, angle float64) [][]byte {
	oriPixelsArr := fetchOriPixelsArr(pbm)
	oriWidth := pbm.width
	oriHeight := pbm.height
	sine := getSine(angle)
	cosine := getCosine(angle)

	rstWidth, rstHeight, minX, minY := getRstPixelsInfo(sine, cosine, oriWidth, oriHeight)
	rstPixelsArr := initRstPixelsArr(rstHeight, rstWidth)

	var wg sync.WaitGroup
	for x := 0; x < rstHeight; x++ {
		wg.Add(1)
		go func(x int) {
			for y := 0; y < rstWidth; y++ {
				oriX := int(((float64(x) + minX) * cosine) + ((float64(y) + minY) * sine))
				oriY := int(((float64(y) + minY) * cosine) - ((float64(x) + minX) * sine))
				if (oriX >= 0 && oriX < oriHeight) && (oriY >= 0 && oriY < oriWidth) {
					rstPixelsArr[x][y] = oriPixelsArr[oriX][oriY]
				}
			}
			wg.Done()
		}(x)
	}
	wg.Wait()

	return rstPixelsArr
}

func getRstPixelsInfo(sine, cosine float64, oriWidth, oriHeight int) (int, int, float64, float64) {
	p1x := -float64(oriWidth) * sine
	p1y := float64(oriWidth) * cosine
	p2x := (float64(oriHeight) * cosine) - (float64(oriWidth) * sine)
	p2y := (float64(oriHeight) * sine) + (float64(oriWidth) * cosine)
	p3x := float64(oriHeight) * cosine
	p3y := float64(oriHeight) * sine

	minX := math.Min(0, math.Min(p3x, math.Min(p1x, p2x)))
	minY := math.Min(0, math.Min(p3y, math.Min(p1y, p2y)))
	maxX := math.Max(0, math.Max(p3x, math.Max(p1x, p2x)))
	maxY := math.Max(0, math.Max(p3y, math.Max(p1y, p2y)))

	rstHeight := int(math.Ceil(maxX - minX))
	rstWidth := int(math.Ceil(maxY - minY))

	return rstWidth, rstHeight, minX, minY
}

func getSine(angle float64) float64 {
	radian := degreeToRadian(angle)
	return math.Sin(radian)
}

func getCosine(angle float64) float64 {
	radian := degreeToRadian(angle)
	return math.Cos(radian)
}

func degreeToRadian(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func initRstPixelsArr(height int, width int) [][]byte {
	rstPixelsArr := make([][]byte, height)
	for i := 0; i < int(height); i++ {
		rstPixelsArr[i] = make([]byte, width)
		for j := 0; j < int(width); j++ {
			rstPixelsArr[i][j] = '0'
		}
	}
	return rstPixelsArr
}

func fetchOriPixelsArr(pbm PBM) (pixelsArr [][]byte) {
	width := pbm.width
	height := pbm.height
	pixels := []byte(pbm.Pixels)

	for x := 0; x < height; x++ {
		pixelLineArr := []byte{}
		for y := 0; y < width; y++ {
			tempX := x*width + y
			pixelLineArr = append(pixelLineArr, pixels[tempX:tempX+1]...)
		}
		pixelsArr = append(pixelsArr, pixelLineArr)
	}

	return pixelsArr
}
