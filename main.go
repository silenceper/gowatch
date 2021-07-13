package main

import (
	"flag"
	"fmt"
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
	showHelp    bool

	started chan bool
)

func init() {
	flag.StringVar(&output, "o", "", "go build output")
	flag.StringVar(&buildPkg, "p", "", "go build packages")
	flag.StringVar(&cmdArgs, "args", "", "app run args,separated by commas. like: -args='-host=:8080,-name=demo'")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "help")
}

var ignoredFilesRegExps = []string{
	`.#(\w+).go$`,
	`.(\w+).go.swp$`,
	`(\w+).go~$`,
	`(\w+).tmp$`,
}

var defaultYml = `
# gowatch.yml configuration example

# The name of the executable file generated under the current directory execution. The default is the current directory name.
appname: "test"

# Specify the command to run after builds done
run_cmd: "./run.sh"

# Specify the directory where the compiled object files are stored
output: /bin/demo
# The file name suffix that needs to be monitored. By default, there is only a '.go' file.
watch_exts:
    - .yml

# The directory that needs to listen for file changes. By default, only the current directory.
watch_paths:
    - ../pk

# Additional parameters that need to be added when running the application
cmd_args:
    - arg1=val1

# Additional parameters that need to be added when building the application
build_args:
    - -race

# Need to increase environment variables, the current environment variables are loaded by default
envs:
    - a=b

# Whether to listen to file changes in the 'vendor' folder
vendor_watch: false

# Directory that do not need to listen for file changes
excluded_paths:
    - path

# main package path, can also be a single file, multiple files separated by commas
build_pkg: ""

# build tags
build_tags: ""

# Commands that can be executed before build the app
#prev_build_cmds:
#  - swag init

# Whether to prohibit automatic operation
disable_run: false
`

func main() {
	flag.Parse()

	// init gowatch.yml
	if len(os.Args) > 1 && os.Args[1] == "init" {
		if _, err := os.Stat("gowatch.yml"); os.IsNotExist(err) {
			_ = ioutil.WriteFile("gowatch.yml", []byte(defaultYml), 0755)
			fmt.Println("gowatch.yml file created to the current directory with the default settings")
		} else {
			fmt.Println("gowatch.yml has been exists")
		}
		os.Exit(0)
	}

	if showHelp {
		fmt.Println("Usage of gowatch:\n\nIf no command is provided gowatch will start the runner with the provided flags\n\nCommands:\n  init  creates a gowatch.yml file with default settings to the current directory\n\nFlags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

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
