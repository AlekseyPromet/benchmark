package main

import (
	"testing"
	"unique"
)

// go test -bench=. main_test.go

func BenchmarkCompareString(b *testing.B) {

	var b1 = make([]byte, 1<<24)

	s1 := string(b1)
	s2 := string(b1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1 == s2
	}
}

func BenchmarkCompareString2Point(b *testing.B) {
	var b1 = make([]byte, 1<<24)

	s1 := string(b1)
	s2 := &s1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1 == *s2
	}
}

func BenchmarkCompareStruct(b *testing.B) {
	type S struct {
		s string
	}

	b1 := make([]byte, 10<<20)

	s1 := S{s: string(b1)}
	s2 := S{s: string(b1)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1.s == s2.s
	}
}

func BenchmarkCompareStruct2Point(b *testing.B) {
	type S struct {
		s string
	}

	b1 := make([]byte, 10<<20)

	s1 := &S{s: string(b1)}
	s2 := &S{s: string(b1)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1.s == s2.s
	}
}

func BenchmarkCompareInterface(b *testing.B) {
	type S struct {
		s string
	}

	b1 := make([]byte, 10<<20)

	s1 := S{s: string(b1)}
	s2 := S{s: string(b1)}

	var i1 interface{} = s1
	var i2 interface{} = s2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = i1.(S).s == i2.(S).s
	}
}

func BenchmarkCompareInterface2Point(b *testing.B) {
	type S struct {
		s string
	}

	b1 := make([]byte, 10<<20)

	s1 := &S{s: string(b1)}
	s2 := &S{s: string(b1)}

	var i1 interface{} = s1
	var i2 interface{} = s2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = i1.(*S).s == i2.(*S).s
	}
}

func BenchmarkCompareMap(b *testing.B) {
	b1 := make([]byte, 10<<20)

	s1 := map[string]string{"key": string(b1)}
	s2 := map[string]string{"key": string(b1)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1["key"] == s2["key"]
	}
}

func BenchmarkCompareMap2Point(b *testing.B) {
	b1 := make([]byte, 10<<20)

	s1 := &map[string]string{"key": string(b1)}
	s2 := &map[string]string{"key": string(b1)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = (*s1)["key"] == (*s2)["key"]
	}
}

func BenchmarkCompareUnique(b *testing.B) {
	b1 := make([]byte, 10<<20)

	s1 := string(b1)
	s2 := string(b1)
	u1 := unique.Make(s1)
	u2 := unique.Make(s2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = u1 == u2
	}
}
