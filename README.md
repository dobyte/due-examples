# due-examples

### 1.介绍

due-examples项目为due框架的示例项目，由于due暂未发布正式的v2版本，因而本项目采用go work的方式引入due框架代码。

### 2. 创建workspace工作目录

```shell
$ cd /home
$ mkdir workspace
```

### 3.克隆due框架项目，并切换到main分支

```shell
$ cd /home/workspace
$ git clone https://github.com/dobyte/due.git
$ cd due/
$ git checkout main
```

### 4.克隆due-examples项目

```shell
$ cd /home/workspace
$ git clone https://github.com/dobyte/due-examples.git
```

### 5.下载依赖包

```shell
$ cd /home/workspace/due-examples
$ go mod tidy
```

### 6.运行Gate网关

```shell
$ cd /home/workspace/due-examples/cluster/gate
$ go run main.go
```

### 7.运行Node节点

```shell
$ cd /home/workspace/due-examples/cluster/node
$ go run main.go
```