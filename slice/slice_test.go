package main

import (
	"testing"
)

const (
	size int = 1000
)

func BenchmarkEmptyInit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

func BenchmarkKnownAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0, size)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

func BenchmarkKnownDirectAccess(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, size, size)
		for k := 0; k < size; k++ {
			data[k] = k
		}
	}
}
