package convutil

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

func Uintptr2Int(uintptrValue uintptr) (int, error) {
	uintStr := fmt.Sprintf("%d", uintptrValue)
	uintValue, err := strconv.ParseUint(uintStr, 10, bits.UintSize)
	if err != nil {
		return -1 * math.MaxInt, err
	}
	return int(uintValue), nil
}
