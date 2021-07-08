package bloomfilter

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filters uint32
	hashs   uint32
	bs      *bitset.BitSet
	count   uint
}

func NewBloomFilter() *BloomFilter {
	var b bitset.BitSet
	return &BloomFilter{10000, 3, &b, 0}
}

func (bf *BloomFilter) put(data []byte) {
	bf.count++
	var seed uint32
	for seed = 0; seed < bf.hashs; seed++ {
		p := murmur3.Sum32WithSeed(data, seed) % bf.filters
		bf.bs.Set(uint(p))
	}
}

func (bf *BloomFilter) check(data []byte) bool {
	isContain := true
	var seed uint32
	for seed = 0; seed < bf.hashs; seed++ {
		p := murmur3.Sum32WithSeed(data, seed) % bf.filters
		isContain = isContain && bf.bs.Test(uint(p))
	}
	return isContain
}
