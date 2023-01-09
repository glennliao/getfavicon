package getfavicon

import (
	"fmt"
	"testing"
)

func TestFromUrl(t *testing.T) {
	favicon, err := FromUrl("https://gitee.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println(favicon)
}

func TestGet(t *testing.T) {
	favicon, err := Get("https://www.github.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println(favicon)
}
