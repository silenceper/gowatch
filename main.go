package main

import (
	"os"
	path "path/filepath"
	"runtime"
)

var (
	cfg      *config
	currpath string
	exit     chan bool

	started chan bool
)

var ignoredFilesRegExps = []string{
	`.#(\w+).go`,
	`.(\w+).go.swp`,
	`(\w+).go~`,
	`(\w+).tmp`,
}

func main() {
	cfg = parseConfig()

	currpath, _ = os.Getwd()
	//app名默认去当前目录名
	if cfg.AppName == "" {
		cfg.AppName = path.Base(currpath)
	}

	//监听的文件后缀
	cfg.WatchExts = append(cfg.WatchExts, ".go")

	runApp()
}

func runApp() {
	var paths []string
	readAppDirectories(currpath, &paths)

	files := []string{}
	NewWatcher(paths, files)
	Autobuild(files)
	for {
		select {
		case <-exit:
			runtime.Goexit()
		}
	}
}
