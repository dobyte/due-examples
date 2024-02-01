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

### 5.启动docker
```shell
$ cd /home/workspace/due-examples/docker
$ docker-compose up -d
```

### 6.下载依赖包

```shell
$ cd /home/workspace/due-examples
$ go mod tidy
```

### 7.运行Gate网关

```shell
$ cd /home/workspace/due-examples/cluster/gate
$ go run main.go
```

### 8.运行Node节点

```shell
$ cd /home/workspace/due-examples/cluster/node
$ go run main.go
```

### 9.运行Client客户端

```shell
$ cd /home/workspace/due-examples/cluster/client
$ go run main.go
```

### 示例输出结果：

Node服务端：

```shell
DEBU[2024/02/01 14:57:35.733166] E:/workspace/due/log/log.go:47 [Welcome to the due framework v2.0.0, Learn more at https://github.com/dobyte/due]
DEBU[2024/02/01 14:57:35.773292] E:/workspace/due/log/log.go:47 [redis channel subscribe succeeded, due:locate:cluster:gate:event]
DEBU[2024/02/01 14:57:35.773830] E:/workspace/due/log/log.go:47 [redis channel subscribe succeeded, due:locate:cluster:node:event]
DEBU[2024/02/01 14:57:35.776918] E:/workspace/due/log/log.go:47 [node server startup successful]
DEBU[2024/02/01 14:57:35.776918] E:/workspace/due/log/log.go:47 [rpcx server listen on 0.0.0.0:53444]
INFO[2024/02/01 14:59:40.133013] E:/workspace/due/log/log.go:54 [I'm client, and the current time is: 2024-02-01 14:59:40]
INFO[2024/02/01 14:59:41.142905] E:/workspace/due/log/log.go:54 [I'm client, and the current time is: 2024-02-01 14:59:41]
INFO[2024/02/01 14:59:42.149013] E:/workspace/due/log/log.go:54 [I'm client, and the current time is: 2024-02-01 14:59:42]
INFO[2024/02/01 14:59:43.159088] E:/workspace/due/log/log.go:54 [I'm client, and the current time is: 2024-02-01 14:59:43]
INFO[2024/02/01 14:59:44.177531] E:/workspace/due/log/log.go:54 [I'm client, and the current time is: 2024-02-01 14:59:44]
```

Client客户端：

```shell
DEBU[2024/02/01 14:59:40.129873] E:/workspace/due/log/log.go:47 [Welcome to the due framework v2.0.0, Learn more at https://github.com/dobyte/due]
INFO[2024/02/01 14:59:40.134595] E:/workspace/due/log/log.go:54 [I'm server, and the current time is: 2024-02-01 14:59:40]
INFO[2024/02/01 14:59:41.145629] E:/workspace/due/log/log.go:54 [I'm server, and the current time is: 2024-02-01 14:59:41]
INFO[2024/02/01 14:59:42.150618] E:/workspace/due/log/log.go:54 [I'm server, and the current time is: 2024-02-01 14:59:42]
INFO[2024/02/01 14:59:43.161436] E:/workspace/due/log/log.go:54 [I'm server, and the current time is: 2024-02-01 14:59:43]
INFO[2024/02/01 14:59:44.179827] E:/workspace/due/log/log.go:54 [I'm server, and the current time is: 2024-02-01 14:59:44]
```