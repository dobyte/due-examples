package main

import (
	"fmt"
	"github.com/dobyte/due/v2/etc"
	"github.com/dobyte/due/v2/log"
)

type config struct {
	DSN             string `json:"dsn"`
	LogLevel        string `json:"logLevel"`
	SlowThreshold   int    `json:"slowThreshold"`
	MaxIdleConns    int    `json:"maxIdleConns"`
	MaxOpenConns    int    `json:"maxOpenConns"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
}

func main() {
	// 读取单个配置参数
	logLevel := etc.Get("db.mysql.logLevel").String()

	fmt.Printf("mysql log-level: %s\n", logLevel)

	// 读取多个配置参数
	conf := &config{}

	if err := etc.Get("db.mysql").Scan(conf); err != nil {
		log.Errorf("get mysql error: %s", err)
	}

	fmt.Printf("mysql config: %+v\n", conf)

	// 修改配置参数
	if err := etc.Set("db.mysql.logLevel", "info"); err != nil {
		log.Errorf("set mysql log-level error: %s", err)
	}

	// 读取单个配置参数
	logLevel = etc.Get("db.mysql.logLevel").String()

	fmt.Printf("mysql log-level: %s\n", logLevel)
}
