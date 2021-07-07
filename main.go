package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filters int
	hashs   int
	bs      *bitset.BitSet
	seed    uint32
	count   int
}

func NewBloomFilter() *BloomFilter {
	var b bitset.BitSet
	return &BloomFilter{10000, 3, &b, 3, 0}
}

func (bf *BloomFilter) put(data []byte) {
	bf.count++
	var i uint32
	for i = 0; i < bf.seed; i++ {
		p := murmur3.Sum32WithSeed(data, i) % uint32(bf.filters)
		bf.bs.Set(uint(p))
	}
}

func (bf *BloomFilter) check(data []byte) bool {
	contain := true
	var i uint32
	for i = 0; i < bf.seed; i++ {
		p := murmur3.Sum32WithSeed(data, i) % uint32(bf.filters)
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
