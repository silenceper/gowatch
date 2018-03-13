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
# gowatch.yml 配置示例

# 当前目录执行下生成的可执行文件的名字，默认是当前目录名
appname: "test"
# 指定编译后的目标文件目录
output: /bin/demo
# 需要追加监听的文件名后缀，默认只有'.go'文件
watch_exts:
    - .yml
# 需要监听的目录，默认只有当前目录
watch_paths:
    - ../pk
# 在执行命令时，需要增加的其他参数
cmd_args:
    - arg1=val1
# 需要增加环境变量，默认已加载当前环境变量
envs:
    - a=b
# 是否监听 ‘vendor’ 文件夹下的文件改变
vendor_watch: false
# 不需要监听的目录名字
excluded_paths:
    - path
# main 包路径，也可以是单个文件，多个文件使用逗号分隔
build_pkg: ""
# build tags
build_tags: ""

```




>This project refers to bee run in the beego / bee project
