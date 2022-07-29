package encryption

import (
	"fmt"
	"math"
	"strings"
)

func Encryption(input string) string {
	var output string

	inputWithoutSpaces := getStringWithoutSpaces(input)
	sizeOfInputWithoutSpaces := len(inputWithoutSpaces)

	lenSqrt := math.Sqrt(float64(sizeOfInputWithoutSpaces))
	axis_x := math.Floor(lenSqrt)
	axis_y := math.Ceil(lenSqrt)

	L := int(axis_x) * int(axis_y)
	for L < sizeOfInputWithoutSpaces {
		fmt.Println("etrou")
		if axis_x < axis_y {
			axis_x = axis_x + 1
		} else {
			axis_y = axis_y + 1
		}
		L = int(axis_x) * int(axis_y)
	}

	fmt.Println("axis_x:", axis_x)
	fmt.Println("axis_y:", axis_y)

	inputGrid := make([]string, int(axis_x))

	for i := 1; i <= int(axis_x); i++ {
		var text string
		if i != int(axis_x) {
			text = inputWithoutSpaces[(i-1)*int(axis_y) : i*int(axis_y)]
		} else {
			text = inputWithoutSpaces[(i-1)*int(axis_y):]
		}
		inputGrid[i-1] = text
	}
	fmt.Println(inputGrid)

	for i := 0; i < int(axis_y); i++ {
		for j := 0; j < int(axis_x); j++ {
			if j < int(axis_x)-1 {
				output += string(inputGrid[j][i])
			} else {
				if len(inputGrid[j]) > i {
					output += string(inputGrid[j][i])
				}
			}
		}
		output += " "
	}
	fmt.Println(output)

	return strings.TrimSpace(output)
}

func getStringWithoutSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}
