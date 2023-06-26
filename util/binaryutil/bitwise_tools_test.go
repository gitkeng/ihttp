package binaryutil_test

import (
	"fmt"
	"github.com/gitkeng/ihttp/util/binaryutil"
	"testing"
)

func TestGetBitStringFromInt(t *testing.T) {
	fmt.Println(binaryutil.GetBitStringFromInt(10))
}

func TestIntShiftLeft(t *testing.T) {
	data := 1
	shiftResult := binaryutil.IntShiftLeft(data, 7)
	fmt.Println(binaryutil.GetBitStringFromInt(shiftResult))
}

func TestIntSetBit(t *testing.T) {
	setBitResult := binaryutil.IntSetBit(0, 1, 5, 30)
	fmt.Println(binaryutil.GetBitStringFromInt(setBitResult))
}

func TestIntGetBit(t *testing.T) {
	data := 0xFFFF
	bitString := binaryutil.GetBitStringFromInt(data)
	fmt.Printf("Data: %s\n", binaryutil.GetBitStringFromInt(data))
	for i, _ := range bitString {
		fmt.Printf("Bit %d: %d\n", i, binaryutil.IntGetBit(data, i))
	}
}
