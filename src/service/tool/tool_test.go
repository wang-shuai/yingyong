package tool

import (
	"testing"
	"fmt"
)

func TestHandIP(t *testing.T) {
	r := HandIP("202.192.13.32")
	fmt.Println(r)

	HandIP("202.-192.13.32") // err
	HandIP("202.1192.13.32") // err
}

func TestHandTimeStr(t *testing.T) {
	r := HandTimeStr("2016-05-13T18:24:28.123Z")
	fmt.Println(r)
}
