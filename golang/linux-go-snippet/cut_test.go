package linuxgosnippet

import (
	"fmt"
	"strings"
	"testing"
)

func CutB() {
	s := "abcdefg"

	b1 := s[0]              // (1) var b1 byte
	fmt.Println(string(b1)) // cut -b 1

	fmt.Println(s[0:5]) // cut -b 1-5
}

func TestCutB(t *testing.T) {
	CutB()
}

func CutSplit() {
	s := `a-bc-de-123-a`

	list := strings.Split(s, "-") // (1) var list []string

	fmt.Println(list[1]) // 取第 2 个位置
	// Output: bc

	s2 := strings.Join(list, "/")
	fmt.Println(s2)
	// Output: a/bc/de/123/a
}

func TestCutSplit(t *testing.T) {
	CutSplit()
}
