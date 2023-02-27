package md5sum

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMd5Sum(t *testing.T) {

}

func MustMd5(s string) string {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func MustMd5File(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return MustMd5(string(data))
}
