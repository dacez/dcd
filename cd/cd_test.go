package cd

import (
	"fmt"
	"testing"
)

func TestGetDirs(t *testing.T) {
	dirs := GetDirs()
	for _, v := range dirs {
		fmt.Println(v)
	}
}

func TestGetAllDir(t *testing.T) {
	dirs := make([]string, 0)
	GetAllDir("~/", &dirs)
	for _, v := range dirs {
		fmt.Println(v)
	}
}
