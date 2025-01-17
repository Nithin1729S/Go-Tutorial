package main

import (
	"fmt"
	"hash"
	"time"
	"github.com/spaolacci/murmur3"
)	

var mHasher hash.Hash32

func init(){
	mHasher = murmur3.New32WithSeed(uint32(time.Now().UnixNano()))
}

func murmurHash(key string,size int) int {
	mHasher.Write([]byte(key))
	result := mHasher.Sum32() % uint32(size)
	mHasher.Reset()
	return int(result)
}

type BloomFilter struct{
	filter []bool
	size int
}

func NewBloomFilter(size int) *BloomFilter{
	return &BloomFilter{filter:make([]bool,size),size:size}
}

func (b *BloomFilter) Add(key string){
	idx :=murmurHash(key,b.size) % b.size
	b.filter[idx] = true
}

func (b *BloomFilter) Exists(key string) bool {
	idx :=murmurHash(key,b.size) % b.size
	return b.filter[idx]
}


func main(){
	bloom:=NewBloomFilter(16)
	keys:=[]string{"a","b","c"}
	for _,key := range keys{
		bloom.Add(key)
	}
	keys=append(keys,"d")
	for _,key :=range keys{
		fmt.Println(key,bloom.Exists(key))
	}
}