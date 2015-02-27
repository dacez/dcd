package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
)

type Mode int

const (
	HisMode    Mode = 0
	GlobalMode Mode = 1
)

type Config struct {
	ContainDirs []string
	IgnoreDirs  string
	HisCount    int
	Home        string
	M           Mode
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
