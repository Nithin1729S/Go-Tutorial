package main

import (
	"fmt"
	"hash"
	"math/rand"

	//"strings"
	"github.com/google/uuid"
	"github.com/spaolacci/murmur3"
)	

var hashfns [] hash.Hash32

func init(){
	hashfns = make([]hash.Hash32, 0)
	for i:=0;i<100;i++{
		hashfns=append(hashfns, murmur3.New32WithSeed(uint32(rand.Uint32())))
	}
}

func murmurHash(key string,size int,hashFnIdx int) int {
	hashfns[hashFnIdx].Write([]byte(key))
	result := hashfns[hashFnIdx].Sum32() % uint32(size)
	hashfns[hashFnIdx].Reset()
	return int(result)
}

type BloomFilter struct{
	filter []uint8
	size int
}

func NewBloomFilter(size int) *BloomFilter{
	return &BloomFilter{filter:make([]uint8,size),size:size}
}

func (b *BloomFilter) Add(key string,numHashFns int){
	for i:=0;i<numHashFns;i++{
		idx :=murmurHash(key,b.size,i) % b.size
		arrIdx:=idx/8
		bitIdx:=idx%8
		b.filter[arrIdx] = b.filter[arrIdx] | (1<<bitIdx)
	}
}

func (b *BloomFilter) Print(){
	fmt.Println(b.filter)
}

func (b *BloomFilter) Exists(key string,numHashFns int) bool {
	for i:=0;i<numHashFns;i++{
		idx :=murmurHash(key,b.size,i) % b.size
		arrIdx:=idx/8
		bitIdx:=idx%8
		exist:=b.filter[arrIdx] & (1<<bitIdx) != 0
		if !exist{
			return false
		}
	}
	return true
}


func main(){
	dataset:= make([]string, 0, 1000)
	dataset_exists:=make(map[string]bool)
	dataset_notexists:=make(map[string]bool)

	for i:=0;i<500;i++{
		u:=uuid.New()
		dataset=append(dataset,u.String() )
		dataset_exists[u.String()]=true
	}

	for i:=0;i<500;i++{
		u:=uuid.New()
		dataset=append(dataset,u.String() )
		dataset_notexists[u.String()]=false
	}

	for j:=1;j<len(hashfns);j++ {
		bloom:=NewBloomFilter(10000)

		for key := range dataset_exists{
			bloom.Add(key,j)
		}

		falsePositive:=0
		for _,key :=range dataset{
			exits:=bloom.Exists(key,j)
			if exits{
				
				if _,ok:=dataset_notexists[key];ok {
					falsePositive++;  //which is ok in a bloom filter
				}
			}
		}
		fmt.Println(float64(falsePositive)/float64(len(dataset)))
	}
}