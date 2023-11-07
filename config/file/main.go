package main

import (
	"context"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/config/file"
	"github.com/dobyte/due/v2/log"
	"time"
)

func init() {
	// 设置全局配置器
	config.SetConfigurator(config.NewConfigurator(config.WithSources(file.NewSource())))
}

func main() {
	var (
		ctx      = context.Background()
		name     = file.Name
		filepath = "config.toml"
	)

	// 更新配置
	if err := config.Store(ctx, name, filepath, map[string]interface{}{
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
	if err := config.Store(ctx, name, filepath, map[string]interface{}{
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
