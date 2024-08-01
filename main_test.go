package main

import (
	"testing"
)

func TestOpenFile(test *testing.T) {
	if _, err := getContentFile("/proc/stat"); err != nil {
		test.Fatal()
	}
}

func TestCpuParse(test *testing.T) {
	if _, err := GetCpuInfo(); err != nil {
		test.Fatal(err.Error())
	}
}

func TestAvgParse(test *testing.T) {
	if _, err := GetLoadAvg(); err != nil {
		test.Fatal(err.Error())
	}
}

func TestMemInfo(test *testing.T) {
	if _, err := GetMemInfo(); err != nil {
		test.Fatal(err.Error())
	}
}

func BenchmarkOpenFile(test *testing.B) {
	for i := 0; i < test.N; i++ {
		if _, err := getContentFile("/proc/stat"); err != nil {
			test.Fatal(err)
		}
	}
}

func BenchmarkGetCpuInfo(test *testing.B) {
	for i := 0; i < test.N; i++ {
		if _, err := GetCpuInfo(); err != nil {
			test.Fatal(err.Error())
		}
	}
}

func BenchmarkGetAvgLoadInfo(test *testing.B) {
	for i := 0; i < test.N; i++ {
		if _, err := GetLoadAvg(); err != nil {
			test.Fatal(err.Error())
		}
	}
}
