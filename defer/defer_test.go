package main

import (
	"testing"
)

func BenchmarkDeferYes1(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(1, &t)
	}
}

func BenchmarkDeferYes3(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(3, &t)
	}
}

func BenchmarkDeferYes5(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(5, &t)
	}
}

func BenchmarkDeferYes10(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(10, &t)
	}
}

func BenchmarkDeferYes20(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(20, &t)
	}
}

func BenchmarkDeferNo(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doNoDefer(&t)
	}
}

func doNoDefer(t *int) {
	func() {
		*t++
	}()
}

func doDefer(loop int, t *int) {
	for index := 0; index < loop; index++ {
		defer func() {
			*t++
		}()
	}
}
