package main

import (
	"github.com/spaolacci/murmur3"
	"hash/crc32"
	"hash/fnv"
	"testing"
)

var (
	data       = []byte("xiaorui.cc")
	dataString = "xiaorui.cc"
)

func fnvHash(b []byte) uint32 {
	fnvHasher := fnv.New32a()
	fnvHasher.Write(b)
	return fnvHasher.Sum32()
}

func crc32Hash(b []byte) uint32 {
	return crc32.ChecksumIEEE(b)
}

func murmurHash(b []byte) uint32 {
	murmurHasher := murmur3.New32()
	murmurHasher.Write(b)
	return murmurHasher.Sum32()
}

func fnvHashString(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}

	return hash
}

func TestFnv(t *testing.T) {
	n1 := fnvHash([]byte("xiaorui.cc"))
	n2 := fnvHash([]byte("xiaorui.cc"))
	if n1 != n2 {
		t.Fatal("not eq")
	}
}

func BenchmarkFnv(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fnvHash(data)
	}
}

func BenchmarkCrc32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		crc32Hash(data)
	}
}

func BenchmarkMurmur32(b *testing.B) {
	for n := 0; n < b.N; n++ {
		murmurHash(data)
	}
}

func BenchmarkFnvString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fnvHashString(dataString)
	}
}
