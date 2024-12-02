package main

import (
	"fmt"
	"runtime"
	"testing"
)

// BenchmarkRun is a benchmark test function that measures the performance.
//
// Parameters:
// t - A pointer to a testing.B object, which provides benchmarking functionality.
//
// Return value:
// None. The benchmark results are printed to the standard output.
func BenchmarkRun_1k(b *testing.B) {

	mg := uint64(1)

	var memStats runtime.MemStats
	fmt.Println("Memory stats (byte):")
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc / mg

	b.ReportMetric(float64(memStats.Alloc/mg), "Alloc_before")

	var n uint64 = 1_000

	for i := uint64(0); i < uint64(b.N); i++ {
		Run(n)
	}

	fmt.Println("Memory stats after goroutines:")
	runtime.ReadMemStats(&memStats)
	b.ReportMetric(float64(memStats.Alloc/mg), "Alloc_after")

	b.ReportMetric(float64(memStats.Alloc/mg-before), "Alloc_delta")
}

func BenchmarkRun_1m(b *testing.B) {

	mg := uint64(1024)

	var memStats runtime.MemStats
	fmt.Println("Memory stats (Kb):")
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc / mg

	b.ReportMetric(float64(memStats.Alloc/mg), "Alloc_before")

	var n uint64 = 1_000_000

	for i := uint64(0); i < uint64(b.N); i++ {
		Run(n)
	}

	fmt.Println("Memory stats after goroutines:")
	runtime.ReadMemStats(&memStats)
	b.ReportMetric(float64(memStats.Alloc/mg), "Alloc_after")

	b.ReportMetric(float64(memStats.Alloc/mg-before), "Alloc_delta")
}

func BenchmarkRun_1b(b *testing.B) {

	mg := uint64(1024 * 1024)

	var memStats runtime.MemStats
	fmt.Println("Memory stats (Mb):")
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc / mg

	b.ReportMetric(float64(memStats.Alloc/mg), "Alloc_before")

	var n uint64 = 1_000_000_000

	for i := uint64(0); i < uint64(b.N); i++ {
		Run(n)
	}

	fmt.Println("Memory stats after goroutines:")
	runtime.ReadMemStats(&memStats)
	b.ReportMetric(float64(memStats.Alloc/mg), "Alloc_after")

	b.ReportMetric(float64(memStats.Alloc/mg-before), "Alloc_delta")
}
