package main

import (
	"flag"
	"os"
	path "path/filepath"
	"runtime"
	"strings"
)

var (
	cfg      *config
	currpath string
	exit     chan bool
	output   string
	buildPkg string
	cmdArgs  string

	started chan bool
)

func init() {
	flag.StringVar(&output, "o", "", "go build output")
	flag.StringVar(&buildPkg, "p", "", "go build packages")
	flag.StringVar(&cmdArgs, "args", "", "app run args,separated by commas. like: -args='-host=:8080,-name=demo'")
}

var ignoredFilesRegExps = []string{
	`.#(\w+).go`,
	`.(\w+).go.swp`,
	`(\w+).go~`,
	`(\w+).tmp`,
}

func main() {
	flag.Parse()
	cfg = parseConfig()

	currpath, _ = os.Getwd()
	if cfg.AppName == "" {
		//app名默认取目录名
		if output == "" {
			cfg.AppName = path.Base(currpath)
		} else {
			cfg.AppName = path.Base(output)
		}
	}

	if output != "" {
		cfg.Output = output
	}

	//如果未指定output则为"./appname"
	if cfg.Output == "" {
		outputExt := ""
		if runtime.GOOS == "windows" {
			outputExt = ".exe"
		}
		cfg.Output = "./" + cfg.AppName + outputExt
	}

	if cmdArgs != "" {
		cfg.CmdArgs = strings.Split(cmdArgs, ",")
	}

	//监听的文件后缀
	cfg.WatchExts = append(cfg.WatchExts, ".go")

	runApp()
}

func runApp() {
	var paths []string
	readAppDirectories(currpath, &paths)
	//除了当前目录，增加额外监听的目录
	for _, path := range cfg.WatchPaths {
		readAppDirectories(path, &paths)
	}

	files := []string{}
	if buildPkg != "" {
		files = strings.Split(buildPkg, ",")
	}
	NewWatcher(paths, files)
	Autobuild(files)
	for {
		select {
		case <-exit:
			runtime.Goexit()
		}
	}
}
