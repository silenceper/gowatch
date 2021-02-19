package main

import (
	"flag"
	"os"
	path "path/filepath"
	"runtime"
	"strings"
)

var (
	cfg         *config
	currpath    string
	exit        chan bool
	output      string
	buildPkg    string
	cmdArgs     string
	showVersion bool

	started chan bool
)

func init() {
	flag.StringVar(&output, "o", "", "go build output")
	flag.StringVar(&buildPkg, "p", "", "go build packages")
	flag.StringVar(&cmdArgs, "args", "", "app run args,separated by commas. like: -args='-host=:8080,-name=demo'")
	flag.BoolVar(&showVersion, "v", false, "show version")
}

var ignoredFilesRegExps = []string{
	`.#(\w+).go$`,
	`.(\w+).go.swp$`,
	`(\w+).go~$`,
	`(\w+).tmp$`,
}

func main() {
	flag.Parse()

	if showVersion {
		printVersion()
		os.Exit(0)
	}

	cfg = parseConfig()
	currpath, _ = os.Getwd()
	if cfg.AppName == "" {
		// The app name defaults to the directory name
		if output == "" {
			cfg.AppName = path.Base(currpath)
		} else {
			cfg.AppName = path.Base(output)
		}
	}

	if output != "" {
		cfg.Output = output
	}

	// If output is not specified, it is "./appname"
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

	// File suffix to be watched
	cfg.WatchExts = append(cfg.WatchExts, ".go")

	runApp()
}

func runApp() {
	var paths []string
	readAppDirectories(currpath, &paths)
	// In addition to the current directory, add additional watch directories
	for _, path := range cfg.WatchPaths {
		readAppDirectories(path, &paths)
	}

	files := []string{}
	if buildPkg == "" {
		buildPkg = cfg.BuildPkg
	}
	if buildPkg != "" {
		files = strings.Split(buildPkg, ",")
	}
	NewWatcher(paths, files)
	Autobuild(files)
	<-exit
	runtime.Goexit()
}
