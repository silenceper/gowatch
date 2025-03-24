package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

var configFile = "./gowatch.yml"

type config struct {
	// The name of the running app, the default current directory name
	AppName string `yaml:"appname"`
	// Specify the program path for output execution
	Output string `yaml:"output"`
	// Need to add watch file suffix name, the default is'.go',
	WatchExts []string `yaml:"watch_exts"`
	// Need to add watch directory, the default is the current folder
	WatchPaths []string `yaml:"watch_paths"`
	// Additional commands executed before the build command
	PrevBuildCmds []string `yaml:"prev_build_cmds"`
	// Extra parameters when running the application
	CmdArgs []string `yaml:"cmd_args"`
	// Additional parameters during build
	BuildArgs []string `yaml:"build_args"`
	// use GOGC on build
	BuildGOGC bool `yaml:"build_go_gc"`
	// Environment variables added when running the application
	Envs []string `yaml:"envs"`
	// Specify whether the files in the vendor directory are also watched
	VendorWatch bool `yaml:"vendor_watch"`
	// Specify a directory that does not require watch
	ExcludedPaths []string `yaml:"excluded_paths"`
	// For packages or files that need to be compiled, use the -p parameter first
	BuildPkg string `yaml:"build_pkg"`
	// -tags parameter accepted during go build
	BuildTags string `yaml:"build_tags"`
	// Specify whether the program runs automatically
	DisableRun bool `yaml:"disable_run"`
	// commands when build finished to run
	RunCmd string `yaml:"run_cmd"`
	// log level, support: debug, info, warn, error, fatal
	LogLevel string `yaml:"log_level"`
}

func parseConfig(confPath string) *config {
	c := &config{}
	if confPath != "" {
		configFile = confPath
	}
	filename, _ := filepath.Abs(configFile)
	if !fileExist(filename) {
		return c
	}
	yamlFile, err := os.ReadFile(filename)
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
