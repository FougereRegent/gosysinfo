package main

import (
	"fmt"
	"testing"
)

func TestOpenFile(test *testing.T) {
	if _, err := getContentFile("/proc/stat"); err != nil {
		test.Fatal()
	}
}

func TestCpuParse(test *testing.T) {
	if result, err := GetCpuInfo(); err != nil {
		test.Fatal(err.Error())
	} else {
		fmt.Println(result)
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
