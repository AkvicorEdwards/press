package press

import (
	"fmt"
	"testing"
)

func TestTgz(t *testing.T) {
	source := []string{"test1", "test2"}
	target := ""
	filename := "test.tgz"
	err := Tgz(source, target, filename)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUnTgz(t *testing.T) {
	tarball := "test.tgz"
	err := UnTgz(tarball, "test_tgz")
	if err != nil {
		fmt.Println(err)
	}
}