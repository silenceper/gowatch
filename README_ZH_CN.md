# gowatch
![Go](https://github.com/silenceper/gowatch/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/silenceper/gowatch)](https://goreportcard.com/report/github.com/silenceper/gowatch)


Go 程序热编译工具，提升开发效率

通过监听当前目录下的相关文件变动，进行实时编译


### 安装使用
使用 `go insall` 命令安装 `gowatch`

```go
go install github.com/flyhope/gowatch@latest
```

安装完成之后，即可使用`gowatch`命令，在当前文件执行:

![gowatch](./screenshot/gowatch.png)


### 命令行参数

- -o : 非必须，指定build的目标文件路径
- -p : 非必须，指定需要build的package（也可以是单个文件）
- -args: 非必须，指定程序运行时参数，例如：-args='-host=:8080,-name=demo'
- -v: 非必须，显示gowatch版本信息
- -h: 非必须，帮助信息

例子:

`gowatch -o ./bin/demo -p ./cmd/demo`

### 配置文件
`gowatch.yml`

大部分情况下，不需要更改配置，直接执行`gowatch`命令就能满足的大部分的需要，但是也提供了一些配置用于自定义

可以通过以下命令生成默认配置
```
gowatch init
```
会在执行目录下创建`gowatch.yml`文件
```
# gowatch.yml 配置示例

# 当前目录执行下生成的可执行文件的名字，默认是当前目录名
appname: "test"

# 指定编译完成后执行的命令，可指定自定义脚本
run_cmd: "./run.sh"

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

# 在构建命令时，需要增加的其他参数
build_args:
    - -race

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

#在build app执行的命令 ，例如 swag init	
#prev_build_cmds:	
#  - swag init

# 是否禁止自动运行
disable_run: false

# 编译时使用GOGC，false强制不使用，true根据环境变量配置来，默认使用
build_go_gc: false

# 日志级别, 支持 debug, info, warn, error, fatal
log_level: "debug"

```

