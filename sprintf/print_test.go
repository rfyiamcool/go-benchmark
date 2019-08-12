package main

import (
	"fmt"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := fmt.Sprintf("xiaorui.cc-%s", "a")
		_ = a
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := "xiaorui.cc-" + "a"
		_ = a
	}
}
