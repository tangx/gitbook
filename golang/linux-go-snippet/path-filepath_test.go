package linuxgosnippet

import (
	"fmt"
	"net/url"
	"path"
	"testing"
)

func Test_PathJoinURL(t *testing.T) {
	domain := "http://www.baidu.com/chatgpt"
	res := "/api/v1/demo"

	s := path.Join(domain, res)
	fmt.Println(s)

	// Output:
	// http:/www.baidu.com/chatgpt/api/v1/demo
}

func Test_URLJoinURL(t *testing.T) {
	domain := "http://www.baidu.com/chatgpt"
	res := "/api/v1/demo"

	s, err := url.JoinPath(domain, res)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)

	// Output:
	// http://www.baidu.com/chatgpt/api/v1/demo
}
