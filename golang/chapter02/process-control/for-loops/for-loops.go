package forloops

import (
	"fmt"
	"time"
)

func counter() {

	for i := 0; i <= 1; i++ {
		fmt.Println("i=", i)
	}

}

func rangeSlice() {
	names := []string{"zhangsan", "lisi", "wangwu", "zhaoliu"}

	for idx, name := range names {
		fmt.Println(names[idx], name)
	}
}

func rangeMap() {
	users := map[string]int{
		"zhangsan": 20,
		"lisi":     30,
		"wangwu":   40,
	}

	for key, value := range users {
		fmt.Println(key, value)
	}
}

func deadLoop() {
	for {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
