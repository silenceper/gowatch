package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

var configFile = "./gowatch.yml"

type config struct {
	//执行的app名字，默认当前目录文字
	AppName string `yaml:"appname"`
	//指定ouput执行的程序路径
	Output string `yaml:"output"`
	//需要追加监听的文件后缀名字，默认是'.go'，
	WatchExts []string `yaml:"watch_exts"`
	//需要追加监听的目录，默认是当前文件夹，
	WatchPaths []string `yaml:"watch_paths"`
	//运行前额外执行的命令
	ExternalCmds []string `yaml:"external_cmds"`
	//执行时的额外参数
	CmdArgs []string `yaml:"cmd_args"`
	//构建时的额外参数
	BuildArgs []string `yaml:"build_args"`
	//执行时追加的环境变量
	Envs []string `yaml:"envs"`
	//vendor 目录下的文件是否也监听
	VendorWatch bool `yaml:"vendor_watch"`
	//不需要监听的目录
	ExcludedPaths []string `yaml:"excluded_paths"`
	//需要编译的包或文件,优先使用-p参数
	BuildPkg string `yaml:"build_pkg"`
	//在go build 时期接收的-tags参数
	BuildTags string `yaml:"build_tags"`
	//程序是否自动运行
	DisableRun bool `yaml:"disable_run"`
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
