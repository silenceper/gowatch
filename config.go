package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

type config struct {
	//执行的app名字，默认当前目录文字
	AppName string `yaml:"appname"`
	//指定ouput执行的程序路径
	Output string `yaml:"output"`
	//需要追加监听的文件后缀名字，默认是'.go'，
	WatchExts []string `yaml:"watch_exts"`
	//执行时的额外参数
	CmdArgs []string `yaml:"cmd_args"`
	//执行时追加的环境变量
	Envs []string `yaml:"envs"`
	//vendor 目录下的文件是否也监听
	VendorWatch bool `yaml:"vendor_watch"`
	//不需要监听的目录
	ExcludedPaths []string `yaml:"excluded_paths"`
	//在go build 时期接收的-tags参数
	BuildTags string `yaml:"build_tags"`
}

func parseConfig() *config {
	c := &config{}
	filename, _ := filepath.Abs(configFile)
	if !fileExist(filename) {
		return c
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}
	return c
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
