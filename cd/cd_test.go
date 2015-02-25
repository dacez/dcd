package cd

import (
	"fmt"
	"testing"
)

func TestGetAllDir(t *testing.T) {
	var dirs []string
	GetAllDir("E:\\GoProject\\src", &dirs)
	for _, v := range dirs {
		fmt.Println(v)
	}
}
