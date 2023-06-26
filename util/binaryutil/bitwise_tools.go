package binaryutil

import (
	"fmt"
	"math/bits"
	"strconv"
)

func GetBitStringFromInt(n int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(bits.UintSize)+"b", n)
}

func IntShiftLeft(n int, shift int) int {
	return n << shift
}

func IntSetBit(positions ...int) int {
	result := 0
	markBit := 1
	for _, position := range positions {
		result |= markBit << position
	}
	return result
}

func IntGetBit(data, position int) int {
	masking := 1
	masking = masking << position
	maskResult := data & masking
	return maskResult >> position
}
