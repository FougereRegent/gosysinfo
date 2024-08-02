package gosysinfo

import (
	"bufio"
	"os"
	"strings"
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

func GetLoadAvg() (*AvgStat, error) {
	var hard hardware
	hard = &avgParse{}

	if content, err := getContentFile(_PATH_AVG_LOAD); err != nil {
		return nil, err
	} else {
		resultInteface := hard.ParseContent(*content)
		result, _ := resultInteface.(AvgStat)
		return &result, nil
	}
}

func GetCpuInfo() (*CpuStats, error) {
	var hard hardware
	hard = &cpuParse{}

	if content, err := getContentFile(_PATH_PROC_INFO); err != nil {
		return nil, err
	} else {
		resultInteface := hard.ParseContent(*content)
		result, _ := resultInteface.(CpuStats)
		return &result, nil
	}
}

func GetMemInfo() (*MemStats, error) {
	var hard hardware
	hard = &memParse{}

	if content, err := getContentFile(_PATH_MEM_INFO); err != nil {
		return nil, err
	} else {
		resultInteface := hard.ParseContent(*content)
		result, _ := resultInteface.(MemStats)
		return &result, nil
	}
}
