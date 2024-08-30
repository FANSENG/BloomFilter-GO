package hash

import (
	"github.com/spaolacci/murmur3"
)

// Murmur3Hash 实现了 Hash 接口
type Murmur3Hash struct {
	name string
}

func (h *Murmur3Hash) Name() string {
	return h.name
}

// Hash 计算输入数据的 Murmur3 哈希值，并返回一个 uint64
func (h *Murmur3Hash) Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}

// NewMurmur3Hash 创建并返回一个新的 Murmur3Hash 实例
func NewMurmur3Hash() Hash {
	return &Murmur3Hash{
		name: "Murmur3Hash",
	}
}
