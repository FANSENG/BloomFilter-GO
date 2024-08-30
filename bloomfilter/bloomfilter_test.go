package bloomfilter

import (
	"context"
	"fmt"
	"fs1n/bloomfilter/dao"
	"fs1n/bloomfilter/hash"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	ctx := context.Background()

	// 64 * 2^7 * 2^10 * 100 bits = 100 MB
	size := uint64(2 ^ 17*100)
	hashFuncs := []hash.Hash{
		hash.NewSHA256Hash(),
		hash.NewFNV1aHash(),
		hash.NewMurmur3Hash(),
	}

	memoryDAO := dao.NewMemoryBloomFilterDAO(size, hashFuncs)
	bf := NewBloomFilter(memoryDAO, hashFuncs, dao.Config{})

	t.Run("Add and Contains", func(t *testing.T) {
		err := bf.Add(ctx, "test1")
		if err != nil {
			t.Fatalf("Failed to add item: %v", err)
		}

		exists, err := bf.Contains(ctx, "test1")
		if err != nil {
			t.Fatalf("Failed to check item: %v", err)
		}
		if !exists {
			t.Errorf("Item 'test1' should exist in the filter")
		}

		exists, err = bf.Contains(ctx, "test2")
		if err != nil {
			t.Fatalf("Failed to check item: %v", err)
		}
		if exists {
			t.Errorf("Item 'test2' should not exist in the filter")
		}
	})

	t.Run("Multiple Adds", func(t *testing.T) {
		items := []string{"item1", "item2", "item3"}
		for _, item := range items {
			err := bf.Add(ctx, item)
			if err != nil {
				t.Fatalf("Failed to add item %s: %v", item, err)
			}
		}

		for _, item := range items {
			exists, err := bf.Contains(ctx, item)
			if err != nil {
				t.Fatalf("Failed to check item %s: %v", item, err)
			}
			if !exists {
				t.Errorf("Item '%s' should exist in the filter", item)
			}
		}
	})

	t.Run("False Positive Test", func(t *testing.T) {
		// 添加大量元素以增加假阳性的可能性
		for i := 0; i < 1000; i++ {
			err := bf.Add(ctx, fmt.Sprintf("item%d", i))
			if err != nil {
				t.Fatalf("Failed to add item: %v", err)
			}
		}

		falsePositives := 0
		tests := 1000
		for i := 0; i < tests; i++ {
			exists, err := bf.Contains(ctx, fmt.Sprintf("nonexistent%d", i))
			if err != nil {
				t.Fatalf("Failed to check item: %v", err)
			}
			if exists {
				falsePositives++
			}
		}

		falsePositiveRate := float64(falsePositives) / float64(tests)
		t.Logf("False positive rate: %f", falsePositiveRate)
		bf.PrintStorageSpaceUsage()
		// 这里可以根据您的具体实现和需求设置一个合理的阈值
		if falsePositiveRate > 0.1 {
			t.Errorf("False positive rate too high: %f", falsePositiveRate)
		}
	})
}
