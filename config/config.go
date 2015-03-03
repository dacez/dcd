package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
)

type DirMode int
type FindMode int

const (
	HisMode DirMode = 0
	AllMode DirMode = 1
)

const (
	NameMode FindMode = 0
	PathMode FindMode = 1
)

type Config struct {
	ContainDirs   []string
	IgnoreDirs    string
	HisCount      int
	Home          string
	DirectoryMode DirMode
	FuzzyFindMode FindMode
}

var conf Config

func init() {
	home := ""
	confPath := ""
	if runtime.GOOS == "windows" {
		home = "C:" + os.Getenv("HOMEPATH")
		confPath = home + "\\.dacecd\\.dacecdrc"
	} else {
		home = os.Getenv("HOME")
		confPath = home + "/.dacecd/.dacecdrc"
	}
	content, _ := ioutil.ReadFile(confPath)
	json.Unmarshal(content, &conf)
	conf.Home = home
}

func GetConfig() *Config {
	return &conf
}

func GetStateLine() string {
	retStr := "Directory["
	if GetConfig().DirectoryMode == AllMode {
		retStr += "  All  "
	} else if GetConfig().DirectoryMode == HisMode {
		retStr += "History"
	} else {
		return ""
	}
	retStr += "] (Ctrl-X)        FindMode["

	if GetConfig().FuzzyFindMode == NameMode {
		retStr += "Name"
	} else if GetConfig().FuzzyFindMode == PathMode {
		retStr += "Path"
	} else {
		return ""
	}
	retStr += "] (Ctrl-D)"
	return retStr
}
