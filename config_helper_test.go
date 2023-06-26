package ihttp_test

import (
	"fmt"
	"github.com/gitkeng/ihttp"
	"os"
	"testing"
)

func TestReadEnv(t *testing.T) {
	// setting up env variables
	os.Setenv("TGC_AUTO_1", "an automatic env 1")
	os.Setenv("TGC_AUTO_2", "another automatic env 2")
	os.Setenv("TGC_AUTO_3", "another automatic env 3")
	fmt.Println("output 1: ", ihttp.Getenv("auto_1", "TGC"))
	fmt.Println("output 2: ", ihttp.Getenv("AUTO_1", "TGC"))
	fmt.Println("output 3: ", ihttp.Getenv("auto_2", "TGC"))
	fmt.Println("output 3: ", ihttp.Getenv("3", "TGC", "AUTO"))

	// when not using prefix, Get by full name
	fmt.Println("output 4: ", ihttp.Getenv("TGC_AUTO_1"))
	fmt.Println("output 5: ", ihttp.Getenv("tgc_auto_1")) // case insensitive
	fmt.Println("output 6: ", ihttp.Getenv("TGC_AUTO_2"))
}
