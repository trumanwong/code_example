package main

import "testing"

func BenchmarkSingleSumString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleSumString(base)
	}
	b.StopTimer()
}

func BenchmarkSingleSprintfString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleSprintfString(base)
	}
	b.StopTimer()
}

func BenchmarkSingleBuilderString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleBuilderString(base)
	}
	b.StopTimer()
}

func BenchmarkSingleBytesBuffString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleBytesString(base)
	}
	b.StopTimer()
}

func BenchmarkSingleJoinstring(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleJoinstring()
	}
	b.StopTimer()
}

func BenchmarkSingleByteSliceString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SingleByteSliceString(base)
	}
	b.StopTimer()
}

func BenchmarkSumString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumString()
	}
	b.StopTimer()
}

func BenchmarkSprintfString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SprintfString()
	}
	b.StopTimer()
}

func BenchmarkBuilderString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuilderString()
	}
	b.StopTimer()
}

func BenchmarkBytesBufferString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytesString()
	}
	b.StopTimer()
}

func BenchmarkJoinstring(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Joinstring()
	}
	b.StopTimer()
}

func BenchmarkByteSliceString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byteSliceString()
	}
	b.StopTimer()
}
