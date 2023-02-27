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

func breakLoops() {
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			// 当 i 是双数的时候， 退出当前循环
			continue
		}

		if i > 6 {
			// 当 i 大于 6 的时候， 结束所有循环
			break
		}

		fmt.Println(i)
	}
}
