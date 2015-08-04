package cd

import (
	"encoding/json"
	"github.com/dacez/dcd/config"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func GetAllDir(root string, dirs *[]string) {
	if root[0:1] == "~" {
		root = strings.Replace(root, "~", config.GetConfig().Home, 1)
	}
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

func GetHis(dirs *[]string) {
	tmpDirs := make([]string, 0)
	content, _ := ioutil.ReadFile(config.GetConfig().Home + "/.dacecd/.dacecdhis")
	json.Unmarshal(content, &tmpDirs)
	curDir, _ := os.Getwd()
	for _, v := range tmpDirs {
		if v != curDir {
			*dirs = append(*dirs, v)
		}
	}
}

func RmHis(dir string) {
	if dir == "" {
		return
	}
	var dirs []string
	content, _ := ioutil.ReadFile(config.GetConfig().Home + "/.dacecd/.dacecdhis")
	json.Unmarshal(content, &dirs)
	var tmpDirs []string
	for _, v := range dirs {
		if v != dir && v != "" {
			tmpDirs = append(tmpDirs, v)
		}
	}
	if len(tmpDirs) > config.GetConfig().HisCount {
		tmpDirs = tmpDirs[0:config.GetConfig().HisCount]
	}
	w, _ := json.Marshal(tmpDirs)
	ioutil.WriteFile(config.GetConfig().Home+"/.dacecd/.dacecdhis", w, os.ModePerm)
}

func PushHis(dir string) {
	if dir == "" {
		return
	}
	var dirs []string
	content, _ := ioutil.ReadFile(config.GetConfig().Home + "/.dacecd/.dacecdhis")
	json.Unmarshal(content, &dirs)
	var tmpDirs []string
	tmpDirs = append(tmpDirs, dir)
	for _, v := range dirs {
		if v != dir && v != "" {
			tmpDirs = append(tmpDirs, v)
		}
	}
	if len(tmpDirs) > config.GetConfig().HisCount {
		tmpDirs = tmpDirs[0:config.GetConfig().HisCount]
	}
	w, _ := json.Marshal(tmpDirs)
	ioutil.WriteFile(config.GetConfig().Home+"/.dacecd/.dacecdhis", w, os.ModePerm)
}

func GetDirs() []string {
	var retDirs []string
	var dirs []string
	if config.GetConfig().DirectoryMode == config.HisMode {
		GetHis(&dirs)
		retDirs = append(retDirs, dirs...)
	} else if config.GetConfig().DirectoryMode == config.AllMode {
		for _, v := range config.GetConfig().ContainDirs {
			dirs = make([]string, 0)
			GetAllDir(v, &dirs)
			retDirs = append(retDirs, dirs...)
		}
	}
	return retDirs
}
