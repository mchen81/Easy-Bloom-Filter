package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filters int
	hashs   uint32
	bs      *bitset.BitSet
	count   int
}

func NewBloomFilter() *BloomFilter {
	var b bitset.BitSet
	return &BloomFilter{10000, 3, &b, 0}
}

func (bf *BloomFilter) put(data []byte) {
	bf.count++
	var seed uint32
	for seed = 0; seed < bf.hashs; seed++ {
		p := murmur3.Sum32WithSeed(data, seed) % uint32(bf.filters)
		bf.bs.Set(uint(p))
	}
}

func (bf *BloomFilter) check(data []byte) bool {
	contain := true
	var seed uint32
	for seed = 0; seed < bf.hashs; seed++ {
		p := murmur3.Sum32WithSeed(data, seed) % uint32(bf.filters)
		contain = contain && bf.bs.Test(uint(p))
	}
	return contain
}

func main() {
	bf := NewBloomFilter()

	s1 := "test1"
	s2 := "test2"
	s3 := "test3"
	s4 := "test4"

	bf.put([]byte(s1))
	bf.put([]byte(s2))
	bf.put([]byte(s3))
	bf.put([]byte(s4))

	fmt.Println(bf.check([]byte(s1)))
	fmt.Println(bf.check([]byte(s2)))
	fmt.Println(bf.check([]byte(s3)))
	fmt.Println(bf.check([]byte(s4)))

}
