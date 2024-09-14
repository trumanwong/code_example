package main

import "testing"

const mapSize = 10000

func BenchmarkMakMapWithoutCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for j := 0; j < mapSize; j++ {
			m[i] = i
		}
	}
}

func BenchmarkMakMapWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, mapSize)
		for j := 0; j < mapSize; j++ {
			m[i] = i
		}
	}
}
