# gowatch

Golang program real-time compilation tools to enhance development efficiency

Real-time compilation by monitoring file changes in a specified directory

### install

```go
go get github.com/silenceper/gowatch
```

After the installation is complete, you can use the `gowatch` command to execute in the main package:

**screenshot：**
![gowatch](./screenshot/gowatch.png)


### Command line parameters

- -o : Optional，specify the target build path
- -p : Optional，Specified the need to build the package, which can also be a single file

**example:**

`gowatch -o ./bin/demo -p ./cmd/demo`

### config file

`gowatch.yml`

Most of the time, you do not need to change the configuration to do most of what you need to do with the `gowatch` command, but it also provides some configuration for customizing. Create the` gowatch.yml` file in the executable:

```
# gowatch.yml Configuration example

# The name of the generated executable file, the default is the current directory name
appname: "test"
# specify the compiled target file directory
output: /bin/demo
# Need to monitor the additional file name suffix, the default only '.go' file
watch_exts:
    - .yml
# need to monitor the directory, the default only the current directory
watch_paths:
    - ../pk
# when the program is executed, additional parameters need to be added
cmd_args:
    - arg1=val1
# need to increase the additional environment variables, the default has been loaded by the current environment variables
envs:
    - a=b
# whether to monitor the 'vendor' folder under the file changes
vendor_watch: false
# do not need to listen to the directory
excluded_paths:
    - path
# main package path, can also be a single file, multiple files separated by commas
build_pkg: ""
# build tags
build_tags: ""

```




>This project refers to bee run in the beego / bee project
