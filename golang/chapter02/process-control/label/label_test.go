package labeldemo

import (
	"fmt"
	"testing"
)

func continueBreakDemo() {

Outer:
	for m := 1; m < 10; m++ {

	Inner:
		for n := 1; n < 10; n++ {
			fmt.Printf("%d*%d=%d ", m, n, m*n)

			if n == 3 {
				break Outer
			}

			if n%4 == 0 {
				break Inner
			}

		}
		fmt.Println("")
	}
}

func TestBreakContinue(t *testing.T) {
	continueBreakDemo()
}

func gotoLabel(n int) {

	fmt.Println("刚开始")

	goto End
	fmt.Println("是奇数") // 永远不会执行， IDE 也提示对应提示

End:
	fmt.Println("结束啦")
}

func TestGotoLabel(t *testing.T) {
	gotoLabel(1)
	gotoLabel(2)
}
