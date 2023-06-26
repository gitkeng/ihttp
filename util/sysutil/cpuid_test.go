package sysutil_test

import (
	. "github.com/klauspost/cpuid/v2"
	"strings"
	"testing"
)

func TestCPUDescrption(t *testing.T) {
	// Print basic CPU information:
	t.Log("Name:", CPU.BrandName)
	t.Log("VendorString:", CPU.VendorString)
	t.Log("PhysicalCores:", CPU.PhysicalCores)
	t.Log("ThreadsPerCore:", CPU.ThreadsPerCore)
	t.Log("LogicalCores:", CPU.LogicalCores)
	t.Log("Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	t.Log("Features:", strings.Join(CPU.FeatureSet(), ","))
	t.Log("Cacheline bytes:", CPU.CacheLine)
	t.Log("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	t.Log("L1 Instruction Cache:", CPU.Cache.L1I, "bytes")
	t.Log("L2 Cache:", CPU.Cache.L2, "bytes")
	t.Log("L3 Cache:", CPU.Cache.L3, "bytes")
	t.Log("Frequency", CPU.Hz, "hz")

	// Test if we have these specific features:
	if CPU.Supports(SSE, SSE2) {
		t.Log("We have Streaming SIMD 2 Extensions")
	}
}
