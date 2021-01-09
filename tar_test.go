package press

import (
	"fmt"
	"testing"
)

func TestTar(t *testing.T) {
	source := []string{"test1", "test2"}
	target := ""
	filename := "test.tar"
	err := Tar(source, target, filename)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUnTar(t *testing.T) {
	tarball := "test.tar"
	err := UnTar(tarball, "test_tar")
	if err != nil {
		fmt.Println(err)
	}
}
