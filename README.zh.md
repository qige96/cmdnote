# CmdNote

### 介绍
一个简单的命令行笔记工具。

[English Docs](./README.md)

### 安装

1. 安装Go语言：https://golang.google.cn/

2. 下载本项目源代码

```shell
$ git clone https://gitee.com/qige96/cmdnote.git
```

3. 编译（并安装）

```shell
$ cd cmdnote
$ go build   # 此过程回安装第三方依赖bleve，需要联网或本地GOPATH已有此依赖包
$ go install # 安装到GOPATH
```

### 使用

你可以对笔记进行常规的增删改查操作，所有笔记都保存在本地的笔记仓库中。

```shell
$ cmdnote -w hello.txt # 调用你的编辑器去编辑笔记
$ cmdnote -r hello.txt # 调用你的阅读器去查阅笔记
$ cmdnote -l # 列出所有笔记
$ cmdnote --rename hello.txt:world.txt # 重命名笔记, 用“：”分割原笔记名和新笔记名
$ cmdnote --remove world.txt # 删除笔记
```

**还能通过关键词搜索的方式找到现有的笔记。**

```shell
$ cmdnote -l
go.txt
hello.txt
lang.txt
lesson1.txt

$ cmdnote -s "hello"
go.txt
hello.txt

```

**--list**和**--search** 都纸质交互式查询

```shell
$ cmdnote -l -i 
    0) go.txt
    1) hello.txt
    2) lang.txt
    3) lesson1.txt
> Which note would you like to check? 0
> Which program would you like to use? cat
hello golang!
> Which note would you like to check? 1
> Which program would you like to use? cat
hello world!
> Which note would you like to check? 2
> Which program would you like to use? cat
Other langs: 

- C
- C++
- C#
- Objectve-C
- Lisp
- PHP
- JavaScript
- TypeScript

> Which note would you like to check? 3
> Which program would you like to use? cat
This is lesson 1.

Fundamental syntax.
> Which note would you like to check?

$
```

### 配置

程序的配置文件`conf.json`就在程序可执行文件的目录下，Windows的默认配置如下：

```json
{
        "editor": "notepad",
        "browser": "notepad",
        "localRepoDir": "F:\\cmdnote\\cmd_notes",
        "remoteRepoDir": ""
}
```

Linux和MacOS(Darwin)系统的默认配置如下：

```json
{
        "editor": "nano",
        "browser": "nano",
        "localRepoDir": "/home/ubuntu/cmdnote",
        "remoteRepoDir": ""
}
```



### License

[MIT License](https://mit-license.org/)
