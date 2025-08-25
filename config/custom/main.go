package main

import (
	"context"

	"github.com/dobyte/due-examples/config/custom/zookeeper"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/log"
)

const filename = "config.toml"

func main() {
	// 设置zookeeper配置源
	config.SetConfigurator(config.NewConfigurator(config.WithSources(zookeeper.NewSource())))

	// 更新配置
	if err := config.Store(context.Background(), zookeeper.Name, filename, map[string]interface{}{
		"timezone": "Local",
	}); err != nil {
		log.Errorf("store config failed: %v", err)
		return
	}

	// 读取配置
	timezone := config.Get("config.timezone", "UTC").String()
	log.Infof("timezone: %s", timezone)

	// 更新配置
	if err := config.Store(context.Background(), zookeeper.Name, filename, map[string]interface{}{
		"timezone": "UTC",
	}); err != nil {
		log.Errorf("store config failed: %v", err)
		return
	}

	// 读取配置
	timezone = config.Get("config.timezone", "UTC").String()
	log.Infof("timezone: %s", timezone)
}
