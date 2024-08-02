package gosysinfo_test

import (
	"testing"

	"github.com/FougereRegent/gosysinfo"
)

func TestCpuParse(test *testing.T) {
	if _, err := gosysinfo.GetCpuInfo(); err != nil {
		test.Fatal(err.Error())
	}
}

func TestAvgParse(test *testing.T) {
	if _, err := gosysinfo.GetLoadAvg(); err != nil {
		test.Fatal(err.Error())
	}
}

func TestMemInfo(test *testing.T) {
	if _, err := gosysinfo.GetMemInfo(); err != nil {
		test.Fatal(err.Error())
	}
}

func BenchmarkGetCpuInfo(test *testing.B) {
	for i := 0; i < test.N; i++ {
		if _, err := gosysinfo.GetCpuInfo(); err != nil {
			test.Fatal(err.Error())
		}
	}
}

func BenchmarkGetAvgLoadInfo(test *testing.B) {
	for i := 0; i < test.N; i++ {
		if _, err := gosysinfo.GetLoadAvg(); err != nil {
			test.Fatal(err.Error())
		}
	}
}
