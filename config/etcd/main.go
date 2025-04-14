package main

import (
	"context"
	"github.com/dobyte/due/config/etcd/v2"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/log"
)

const filename = "config.toml"

func main() {
	// 设置etcd配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(etcd.NewSource())))

	// 更新配置
	if err := config.Store(context.Background(), etcd.Name, filename, map[string]interface{}{
		"timezone": "Local",
	}); err != nil {
		log.Errorf("store config failed: %v", err)
		return
	}

	// 读取配置
	timezone := config.Get("config.timezone", "UTC").String()
	log.Infof("timezone: %s", timezone)

	// 更新配置
	if err := config.Store(context.Background(), etcd.Name, filename, map[string]interface{}{
		"timezone": "UTC",
	}); err != nil {
		log.Errorf("store config failed: %v", err)
		return
	}

	// 读取配置
	timezone = config.Get("config.timezone", "UTC").String()
	log.Infof("timezone: %s", timezone)
}
