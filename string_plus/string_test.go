package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var m = [...]string{
	"AAAAAAAAA",
	"BBBBBBBBB",
	"CCCCCCCCC",
	"DDDDDDDDD",
	"EEEEEEEEE",
	"FFFFFFFFF",
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(m[:], "")
	}
}

func BenchmarkAppendOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 string
		for _, v := range m {
			m2 += m2 + v
		}
	}
}

func BenchmarkHardCoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m[0] + m[1] + m[2] + m[3] + m[4] + m[5]
	}
}

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s%s%s", m[0], m[1], m[2], m[3], m[4], m[5])
	}
}

func BenchmarkByteArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b []byte
		for _, v := range m {
			b = append(b, v...)
		}
		_ = string(b)
	}
}

func BenchmarkCapByteArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b = make([]byte, 0, 100)
		for _, v := range m {
			b = append(b, v...)
		}
		_ = string(b)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for _, v := range m {
			b.Write([]byte(v))
		}
		_ = b.String()
	}
}

func BenchmarkCapBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b = bytes.NewBuffer(make([]byte, 0, 100))
		for _, v := range m {
			b.WriteString(v)
		}
		_ = b.String()
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for _, v := range m {
			b.WriteString(v)
		}
		_ = b.String()
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		for _, v := range m {
			b.WriteString(v)
		}
		_ = b.String()
	}
}
