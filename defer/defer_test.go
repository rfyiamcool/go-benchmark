package main

import (
	"testing"
)

func BenchmarkDefer(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(1, &t)
	}
}

func BenchmarkDeferLoop3(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(3, &t)
	}
}

func BenchmarkLastCallLoop3(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doLastCall(3, &t)
	}
}

func BenchmarkDeferLoop5(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(5, &t)
	}
}

func BenchmarkLastCallLoop5(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doLastCall(5, &t)
	}
}

func BenchmarkDeferLoop10(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doDefer(10, &t)
	}
}

func BenchmarkLastCallLoop10(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		doLastCall(10, &t)
	}
}

func BenchmarkDeferLoop20(b *testing.B) {
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

func doLastCall(loop int, t *int) {
	f := func() {
		for index := 0; index < loop; index++ {
			*t++
		}
	}
	f()
}
