package hash

import (
	"crypto/sha256"
	"encoding/binary"
)

// SHA256Hash 实现了 Hash 接口
type SHA256Hash struct {
	name string
}

func (h *SHA256Hash) Name() string {
	return h.name
}

// Hash 计算输入数据的 SHA256 哈希值，并返回一个 uint64
func (h *SHA256Hash) Hash(data []byte) uint64 {
	hash := sha256.Sum256(data)
	return binary.BigEndian.Uint64(hash[:8])
}

// NewSHA256Hash 创建并返回一个新的 SHA256Hash 实例
func NewSHA256Hash() Hash {
	return &SHA256Hash{
		name: "SHA256Hash",
	}
}
