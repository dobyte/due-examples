package zookeeper

import (
	"context"
	"github.com/dobyte/due/v2/config"
)

const Name = "zookeeper"

type Source struct {
}

func NewSource() *Source {
	return &Source{}
}

// Name 配置源名称
func (s *Source) Name() string {
	return Name
}

// Load 加载配置项
func (s *Source) Load(ctx context.Context, file ...string) ([]*config.Configuration, error) {
	return nil, nil
}

// Store 保存配置项
func (s *Source) Store(ctx context.Context, file string, content []byte) error {
	return nil
}

// Watch 监听配置项
func (s *Source) Watch(ctx context.Context) (config.Watcher, error) {
	return &Watcher{}, nil
}

// Close 关闭配置源
func (s *Source) Close() error {
	return nil
}

type Watcher struct {
}

// Next 返回配置列表
func (w *Watcher) Next() ([]*config.Configuration, error) {
	return nil, nil
}

// Stop 停止监听
func (w *Watcher) Stop() error {
	return nil
}
