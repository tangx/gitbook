package linuxgosnippet

import (
	"fmt"
	"os"
	"testing"
)

func OsWalk(name string) {
	entries, err := os.ReadDir(name)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			fmt.Printf("%s 是文件夹\n", name)
		}

		// 如果是文件
		fi, err2 := os.Stat(name)
		if err2 != nil {
			panic(err)
		}

		fmt.Printf("%s  %d  %s\n", fi.Mode().String(), fi.Size(), name)
	}
}

func Test_ls(t *testing.T) {
	OsWalk(".")
}
