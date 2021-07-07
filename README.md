# Goal

Make a really simple bloom filter which only has put() and get()

# Struct Element

- filters int -> filters refers to bit array length (default 1000)
- hashs uint32 -> the number of hash function (default 3)
- bs \*BitSet -> the bit array
- count int -> the number of elements in the bit array

# Methods

1. put(data []byte) : put any data in byte array format.
2. check(data []byte) bool : to check if the data exists in Bloom Filter(false positive happens).

# Resource

I did a BF in Java previously ([here](https://github.com/mchen81/DFS-java-netty/blob/main/src/main/java/edu/usfca/cs/dfs/utils/BloomFilter.java)), I'd like to achieve it again in Golang
