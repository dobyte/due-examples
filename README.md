# due-examples

### 一、项目介绍

due-examples项目为due框架的完整示例项目。

### 二、项目说明

```shell
cluster/            # 集群示例
    actor/          # actor使用示例
    channel/        # 频道功能示例
        tcp/        # tcp示例
        ws/         # websocket示例
    service/        # 微服务功能示例
        grpc/       # grpc示例
        rpcx/       # rpcx示例
    simple/         # 简单示例
        tcp/        # tcp示例
        web/        # web示例
        ws/         # websocket示例
config/             # 配置中心示例
    consul/         # consul示例
    etcd/           # etcd示例
    file/           # 文件示例
    nacos/          # nacos示例
errors/             # 错误示例
etc/                # etc启动配置示例
registry/           # 注册中心示例
    consul/         # consul示例
    etcd/           # etcd示例
    nacos/          # nacos示例
lock/               # 分布式锁示例
    memcache/       # memcache示例
    redis/          # redis示例
```

### 三、快速开始

由于篇幅关系，此处以一个简单的分布式示例项目作为介绍，具体代码详见 [simple/ws](./cluster/simple/ws/)。

#### 1. 创建workspace工作目录

```shell
$ cd /home
$ mkdir workspace
```

#### 2.克隆due-examples项目

```shell
$ cd /home/workspace
$ git clone https://github.com/dobyte/due-examples.git
```

#### 3.启动docker
```shell
$ cd /home/workspace/due-examples/docker
$ docker-compose up -d
```

#### 4.下载依赖包

```shell
$ cd /home/workspace/due-examples
$ go mod tidy
```

#### 5.运行WS网关

```shell
$ cd /home/workspace/due-examples/cluster/simple/ws/gate
$ go run main.go
```

#### 6.运行节点

```shell
$ cd /home/workspace/due-examples/cluster/simple/ws/node
$ go run main.go
```

#### 7.运行WS客户端

```shell
$ cd /home/workspace/due-examples/cluster/simple/ws/client
$ go run main.go
```

#### 示例输出结果：

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