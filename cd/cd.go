package cd

import (
	"io/ioutil"
	"runtime"
)

func GetAllDir(root string, dirs *[]string) {
	ds, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}
	s := "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	for _, v := range ds {
		if v.Name()[0] == '.' || v.IsDir() == false {
			continue
		}
		*dirs = append(*dirs, root+s+v.Name())
		GetAllDir(root+s+v.Name(), dirs)
	}
}
