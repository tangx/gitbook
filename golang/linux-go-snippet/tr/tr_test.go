package trdemo

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrDemo(t *testing.T) {
	s1 := strings.ToLower("abcDEFg")
	s2 := strings.ToUpper("abcDEFg")
	fmt.Println(s1, s2)
}

func TestTrDeleteDemo(t *testing.T) {

	rep := strings.NewReplacer("a", "", "c", "")
	abc := rep.Replace("abcabcabc")
	fmt.Println(abc)

}
