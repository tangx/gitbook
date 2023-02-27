package switchdemo

import (
	"fmt"
	"testing"
)

func demo(n int) {
	switch n {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2, 3:
		fmt.Println("2 或 3")
	default:
		fmt.Println("其他")
	}
}

func Test_Demo(t *testing.T) {
	demo(1)
}
