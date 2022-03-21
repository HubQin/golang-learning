package main

import (
	"strings"
	"testing"
)

func BenchmarkIOReadAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IOReadAll(strings.NewReader(bigStr))
	}
}

func BenchmarkIOCopy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		IOCopy(strings.NewReader(bigStr))
	}
}