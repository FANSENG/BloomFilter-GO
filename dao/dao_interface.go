package dao

import (
	"context"
)

type Config map[string]string

// BloomFilterDAO 定义了布隆过滤器存储层的接口
type BloomFilterDAO interface {
	Init(ctx context.Context, config Config)
	// Set 设置指定位置的位为 1
	Set(ctx context.Context, key string, position uint64) error

	// Get 获取指定位置的位的值
	Get(ctx context.Context, key string, position uint64) (bool, error)

	// Clear 清空布隆过滤器
	Clear(ctx context.Context, key string) error

	// Capacity 获取布隆过滤器的容量(bits)
	Capacity() uint64

	// Occupy 获取某个 hash 对应的存储占用了多少 bits
	Occupy(key string) uint64
}
