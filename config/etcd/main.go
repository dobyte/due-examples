package main

import (
	"context"
	"github.com/dobyte/due/config/etcd/v2"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/log"
	"time"
)

func init() {
	// 设置全局配置器
	config.SetConfigurator(config.NewConfigurator(config.WithSources(etcd.NewSource())))
}

func main() {
	var (
		ctx  = context.Background()
		name = etcd.Name
		file = "config.toml"
	)

	// 更新配置
	if err := config.Store(ctx, name, file, map[string]interface{}{
		"timezone": "Local",
	}); err != nil {
		log.Errorf("store config failed: %v", err)
		return
	}

	time.Sleep(5 * time.Millisecond)

	// 读取配置
	timezone := config.Get("config.timezone", "UTC").String()
	log.Infof("timezone: %s", timezone)

	// 更新配置
	if err := config.Store(ctx, name, file, map[string]interface{}{
		"timezone": "UTC",
	}); err != nil {
		log.Errorf("store config failed: %v", err)
		return
	}

	time.Sleep(5 * time.Millisecond)

	// 读取配置
	timezone = config.Get("config.timezone", "UTC").String()
	log.Infof("timezone: %s", timezone)
}
