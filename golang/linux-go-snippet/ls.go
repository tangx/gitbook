package linuxgosnippet

import (
	"fmt"
	"os"
)

func OsWalk(name string) {
	entries, err := os.ReadDir(name)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("%s 是文件夹\n", entry.Name())
		}
		fmt.Printf("%s 是文件\n", entry.Name())
	}
}
