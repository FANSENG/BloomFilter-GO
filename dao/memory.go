package dao

import (
	"context"

	"fs1n/bloomfilter/hash"
)

type hashData struct {
	data []uint64

	// bits count
	capacity uint64
	occupy   uint64
}

// MemoryBloomFilterDAO 是一个基于内存的 BloomFilterDAO 实现
type MemoryBloomFilterDAO struct {
	data  map[string]hashData
	hashs []hash.Hash
	size  uint64
}

func (m *MemoryBloomFilterDAO) Init(ctx context.Context, config Config) {
	for _, val := range m.hashs {
		m.data[val.Name()] = hashData{
			data:     make([]uint64, (m.size+63)/64),
			capacity: m.size << 3,
			occupy:   0,
		}
	}
}

// NewMemoryBloomFilterDAO 创建一个新的 MemoryBloomFilterDAO 实例
func NewMemoryBloomFilterDAO(size uint64, hashs []hash.Hash) *MemoryBloomFilterDAO {
	return &MemoryBloomFilterDAO{
		data:  make(map[string]hashData),
		hashs: hashs,
		size:  size,
	}
}

func (m *MemoryBloomFilterDAO) Set(ctx context.Context, key string, position uint64) error {
	var val hashData
	var ok bool
	if val, ok = m.data[key]; !ok {
		panic("key not found")
	}
	position %= m.size
	idx, bit := position/64, position%64
	if val.data[idx]>>bit == 0 {
		val.occupy++
		m.data[key] = val
	}
	val.data[idx] |= 1 << bit
	return nil
}

func (m *MemoryBloomFilterDAO) Get(ctx context.Context, key string, position uint64) (bool, error) {
	var val hashData
	var ok bool
	if val, ok = m.data[key]; !ok {
		panic("key not found")
	}
	position %= m.size
	idx, bit := position/64, position%64
	return (val.data[idx] & (1 << bit)) != 0, nil
}

func (m *MemoryBloomFilterDAO) Clear(ctx context.Context, key string) error {
	if _, ok := m.data[key]; !ok {
		return nil // 如果键不存在，无需清除
	}

	// 将所有位重置为 0
	for i := range m.data[key].data {
		m.data[key].data[i] = 0
	}

	return nil
}

func (m *MemoryBloomFilterDAO) Capacity() uint64 {
	return m.size << 3
}

func (m *MemoryBloomFilterDAO) Occupy(key string) uint64 {
	return m.data[key].occupy
}
