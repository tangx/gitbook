package linuxgosnippet

import (
	"os"
	"testing"
)

func Test_Mkdir(t *testing.T) {
	Mkdir()
}

func Mkdir() {
	os.Mkdir("path", os.ModePerm)
	os.MkdirAll("path/to/abc", os.ModePerm)
}
