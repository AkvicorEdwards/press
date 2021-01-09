package press

import (
	"fmt"
	"testing"
)

func TestGzip(t *testing.T) {
	source := "testFile1"
	target := ""
	filename := "test.gz"
	err := Gzip(source, target, filename)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUnGzip(t *testing.T) {
	tarball := "test.gz"
	err := UnGzip(tarball, "test.ungz")
	if err != nil {
		fmt.Println(err)
	}
}
