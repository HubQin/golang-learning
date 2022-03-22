package main

import (
	"testing"
)

func BenchmarkDownFileWithReadAll(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		DownFileWithReadAll()
	}
}

func BenchmarkDownFileWithCopy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		DownFileWithCopy()
	}
}
