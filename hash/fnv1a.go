package hash

import (
	"hash/fnv"
)

type FNV1aHash struct {
	name string
}

func (h *FNV1aHash) Name() string {
	return h.name
}

// Hash 计算输入数据的 FNV-1a 哈希值，并返回一个 uint64
func (h *FNV1aHash) Hash(data []byte) uint64 {
	val := fnv.New64a()
	// fnv.Write never returns an error
	_, _ = val.Write(data)
	return val.Sum64()
}

// NewFNV1aHash 创建并返回一个新的 FNV1aHash 实例
func NewFNV1aHash() Hash {
	return &FNV1aHash{
		name: "FNV1aHash",
	}
}
