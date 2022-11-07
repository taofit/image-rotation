# Image rotation

## Project description

The purpose of this project is to rotate a PBM format image in any anger.
There is a link that explains what a PBM format image is.
Basically, a simple example of the PBM format is as follows (there is a newline character at the end of each line): 
```
P1
# This is an example bitmap of the letter "J"
6 10
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0
0 0 0 0 0 0
```

It is not required that pixels are nicely lined up, the format ignores whitespaces and linefeeds in the data section.
The following displays the same image:
```
P1
# This is an example bitmap of the letter "J"
6 10
000010000010000010000010000010000010100010011100000000000000
```
The program is to generate a new sequence of 0 and 1 after it turns certain degree.

## Run the program

In project root folder, run command: `go run main.go -fileName bitmap.pbm -degree 45` bitmap.pbm is the generated file name after rotation, 45 is the degree the original image rotates anti-clockwise. If want the image to rotate clockwise, add minus sign before 45, like `-45`.

`bitmap.pbm` file that is located in the root folder, is the original PBM format image that will be rotated, it of course can be modified by placing 0 or 1 in different location, or change the size of the image.
The generated resultant image is saved in file `rst-bitmap.pbm` in root folder.

## Run the benchmark test

Under rotation folder, run command: `go test -bench=.` When writing pixels to the resultant image, in general running a go routine on each sub array has shown the best performance comparing without using go routine and using go routine on each pixels. The code for without using go routine and go routine on array's each element are in the previous commits.

## Solution

The calculation of the rotated pixels are based on the formala below:

newX = x cos(@) + y sin(@)  
newY = -x sin(@) + y cos(@)  

Some reference: https://datagenetics.com/blog/august32013/index.html  
The code should be self-explanatory

Thanks
