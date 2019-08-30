package main

import "testing"

type InterfaceA interface {
	AA()
}

type A struct {
	v int
}

func (a *A) AA() {
	a.v += 1
}

func TypeSwitch(v interface{}) {
	switch v.(type) {
	case InterfaceA:
		v.(InterfaceA).AA()
	}
}

func NormalSwitch(a *A) {
	a.AA()
}

func InterfaceSwitch(v interface{}) {
	v.(InterfaceA).AA()
}

func Benchmark_NormalSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		NormalSwitch(a)
	}
}

func Benchmark_InterfaceSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		InterfaceSwitch(a)
	}
}

func Benchmark_TypeSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		TypeSwitch(a)
	}
}
