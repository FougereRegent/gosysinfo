package main

import (
	"testing"
)

func TestOpenFile(test *testing.T) {
	if _, err := getContentFile("/proc/stat"); err != nil {
		test.Fatal()
	}
}

func BenchmarkOpenFile(test *testing.B) {
	for i := 0; i < test.N; i++ {
		if _, err := getContentFile("/proc/stat"); err != nil {
			test.Fatal(err)
		}
	}
}
