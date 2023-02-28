package chapter02

import "fmt"

const (
	Monday   = 0
	Tuesday  = iota // 2
	_               // 3。 周 3 跳过, 但使用站位符是 iota 递增
	Thursday = iota // 4
	// Friday
	// Saturday
	Sunday = iota + 2 // 7。 5和6不存在， 因此 Sunday 是 iota+2
)

var abc int8 = 100 // 指定 abc 的类型是 int8
var def = 100      // 推定类型为 int

func demo() {
	a1 := 100
	var a2 = 100
	fmt.Println(a1, a2)
}

func demo2(gg int) {
	fmt.Println(gg)
}

func assignment() {
	abc := "word"
	// abc = 123 // cannot use 123 (untyped int constant) as string value

	fmt.Println(abc)
}

func assignmentNumber() {
	// var n int = 100
	var n8 int8
	// n8 = n // cannot use n (variable of type int) as int8 value
	fmt.Println(n8)
}
