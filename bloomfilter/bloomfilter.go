package bloomfilter

import (
	"context"
	"fmt"

	"fs1n/bloomfilter/dao"
	"fs1n/bloomfilter/hash"
)

type BloomFilter struct {

	// 用于计算
	hashFuncList []hash.Hash

	// 提供需要的函数
	daoClient dao.BloomFilterDAO
}

func NewBloomFilter(daoClient dao.BloomFilterDAO, hashFuncList []hash.Hash, conf dao.Config) *BloomFilter {
	val := &BloomFilter{
		hashFuncList: hashFuncList,
		daoClient:    daoClient,
	}
	val.daoClient.Init(context.Background(), conf)
	return val
}

func (bf *BloomFilter) Add(ctx context.Context, item string) error {
	for _, hashFunc := range bf.hashFuncList {
		index := hashFunc.Hash([]byte(item))
		if err := bf.daoClient.Set(ctx, hashFunc.Name(), index); err != nil {
			return err
		}
	}
	return nil
}

func (bf *BloomFilter) Contains(ctx context.Context, item string) (bool, error) {
	for _, hashFunc := range bf.hashFuncList {
		index := hashFunc.Hash([]byte(item))
		exists, err := bf.daoClient.Get(ctx, hashFunc.Name(), index)
		if err != nil {
			return false, err
		}
		if !exists {
			return false, nil
		}
	}
	return true, nil
}

func (bf *BloomFilter) PrintStorageSpaceUsage() {
	capacity := bf.daoClient.Capacity()
	for _, hashFunc := range bf.hashFuncList {
		fmt.Printf("%v: %v/%v\n", hashFunc.Name(), bf.daoClient.Occupy(hashFunc.Name()), capacity)
	}
}
