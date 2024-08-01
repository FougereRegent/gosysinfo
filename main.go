package main

import (
	"bufio"
	"os"
	"strings"

	hardware "github.com/FougereRegent/gosysinfo/Hardware"
	cpu "github.com/FougereRegent/gosysinfo/Hardware/Cpu"
	mem "github.com/FougereRegent/gosysinfo/Hardware/Mem"
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
		result.WriteByte('\n')
	}

	res := result.String()
	res = res[:len(res)-1]
	return &res, nil
}

func GetLoadAvg() (*cpu.AvgStat, error) {
	var hard hardware.Hardware
	hard = &cpu.AvgParse{}

	if content, err := getContentFile(_PATH_AVG_LOAD); err != nil {
		return nil, err
	} else {
		resultInteface := hard.ParseContent(*content)
		result, _ := resultInteface.(cpu.AvgStat)
		return &result, nil
	}
}

func GetCpuInfo() (*cpu.CpuStats, error) {
	var hard hardware.Hardware
	hard = &cpu.CpuParse{}

	if content, err := getContentFile(_PATH_PROC_INFO); err != nil {
		return nil, err
	} else {
		resultInteface := hard.ParseContent(*content)
		result, _ := resultInteface.(cpu.CpuStats)
		return &result, nil
	}
}

func GetMemInfo() (*mem.MemStats, error) {
	var hard hardware.Hardware
	hard = &mem.MemParse{}

	if content, err := getContentFile(_PATH_MEM_INFO); err != nil {
		return nil, err
	} else {
		resultInteface := hard.ParseContent(*content)
		result, _ := resultInteface.(mem.MemStats)
		return &result, nil
	}
}
