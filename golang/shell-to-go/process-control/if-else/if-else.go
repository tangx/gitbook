package ifelse

import "fmt"

func if_else(n int) {
	if n >= 100 || n <= -1 {
		fmt.Println("无效")
	} else if n > 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("一般")
	}
}
