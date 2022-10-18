# Image rotation

## Run the program

In project root folder, run command: `go run main.go -fileName bitmap.pbm -degree 45` bitmap.pbm is the generated file name after rotation, 45 is the degree the original image rotates anti-clockwise. If want the image to rotate clockwise, add minus sign before 45, like `-45`.

The generated resultant image is saved as rst-bitmap.pbm in root folder

## Run the benchmark test

Under rotation folder, run command: `go test -bench=.` When writing pixels to the resultant image, in general running a go routine on each sub array has shown the best performance comparing without using go routine and using go routine on each pixels. The code for without using go routine and go routine on array each element are in the previous commits.

## Solution

The calculation of the rotated pixels are based on the formala below:

newX = x cos(@) + y sin(@)
newY = -x sin(@) + y cos(@)

Some reference: https://datagenetics.com/blog/august32013/index.html
The code should be self-explanatory

Thanks
