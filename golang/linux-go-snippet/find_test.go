package linuxgosnippet

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func walk(name string, handler func(name string)) {

	entries, err := os.ReadDir(name)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		name := filepath.Join(name, entry.Name())
		if entry.IsDir() {
			// (1) 递归查找
			walk(name, handler)
		}

		// (2) 过滤参数: -type f
		if strings.HasSuffix(name, ".md") {
			// _ = os.Remove(name)

			// (3) 管道应用 | xargs rm -f
			handler(name) //
		}
	}
}

func remove(name string) {
	err := os.Remove(name)
	if err != nil {
		panic(err)
	}
}

func OsWalkDelete() {
	walk("path/2/file", remove)
}

func TestOsWalkDelete(t *testing.T) {
	OsWalkDelete()
}
