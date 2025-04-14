package main

import (
	"context"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/registry"
	"github.com/dobyte/due/v2/utils/xuuid"
	"time"
)

func main() {
	var (
		reg   = etcd.NewRegistry()
		id    = xuuid.UUID()
		name  = "game-server"
		alias = "mahjong"
		ins   = &registry.ServiceInstance{
			ID:       id,
			Name:     name,
			Kind:     cluster.Node.String(),
			Alias:    alias,
			State:    cluster.Work.String(),
			Endpoint: "grpc://127.0.0.1:6339",
		}
	)

	// 监听
	watch(reg, name, 1)
	watch(reg, name, 2)

	// 注册服务
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err := reg.Register(ctx, ins)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	// 更新服务
	ins.State = cluster.Busy.String()
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	err = reg.Register(ctx, ins)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	// 解注册服务
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	err = reg.Deregister(ctx, ins)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
}

func watch(reg *etcd.Registry, serviceName string, goroutineID int) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	watcher, err := reg.Watch(ctx, serviceName)
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			services, err := watcher.Next()
			if err != nil {
				log.Fatalf("goroutine %d: %v", goroutineID, err)
				return
			}

			log.Infof("goroutine %d: new event entity", goroutineID)

			for _, service := range services {
				log.Infof("goroutine %d: %+v", goroutineID, service)
			}
		}
	}()
}
