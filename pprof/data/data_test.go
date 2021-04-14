package data

import (
	"testing"
)

const str = "Hello, World"

//go test -bench=. -cpuprofile=cpu.prof
//go tool pprof -http=:8080 cpu.prof
//pprof -http=:8080 cpu.prof

func TestAdd(t *testing.T) {
	s := Add(str)
	if s == "" {
		t.Errorf("Test.Add error")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(str)
	}
}
