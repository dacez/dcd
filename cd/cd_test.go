package cd

import (
	"fmt"
	"testing"
)

func TestGetAllDir(t *testing.T) {
	dirs := GetDirs()
	for _, v := range dirs {
		fmt.Println(v)
	}
}
