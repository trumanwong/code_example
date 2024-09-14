package main

import "testing"

const size = 10000

func BenchmarkMakeSliceWithoutCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0)
		for j := 0; j < size; j++ {
			arr = append(arr, j)
		}
	}
}

func BenchmarkMakeSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, size)
		for j := 0; j < size; j++ {
			arr = append(arr, j)
		}
	}
}
