package main

import (
	"bufio"
	"os"
	"strings"

	cpu "github.com/FougereRegent/gosysinfo/Cpu"
)

const (
	_PATH_PROC_INFO string = "/proc/stat"
	_PATH_AVG_LOAD         = "/proc/loadavg"
	_PATH_MEM_INFO         = "/proc/meminfo"
)

func getContentFile(path string) (*string, error) {
	result := strings.Builder{}
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result.WriteString(scanner.Text())
	}

	res := result.String()
	return &res, nil
}

func GetLoadAvg() error {
	return nil
}

func GetProcStat() (*cpu.CpuStats, error) {
	return nil, nil
}

func GetMemStat() error {
	return nil
}
