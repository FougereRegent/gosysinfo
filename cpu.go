package gosysinfo

import (
	"strconv"
	"strings"
	"time"
)

const (
	_CPU           string = "cpu"
	_INTR                 = "intr"
	_CTXT                 = "ctxt"
	_BTIME                = "btime"
	_PROCESSES            = "processes"
	_PROCS_RUNNING        = "procs_running"
	_PROCS_BLOCKED        = "procs_blocked"
	_SOFTIRQ              = "softirq"
)

const (
	_CPU_SIZE           uint = 3
	_INTR_SIZE               = 4
	_CTXT_SIZE               = 4
	_BTIME_SIZE              = 5
	_PROCESSES_SIZE          = 9
	_PROCS_RUNNING_SIZE      = 12
	_PROCS_BLOCKED_SIZE      = 13
	_SOFTIRQ_SIZE            = 7
)

type cpuParse struct {
}

type CpuStats struct {
	NumberCpu    uint16
	Cpus         []cpuLoad
	Cpu          cpuLoad
	Context      uint64
	BootTime     time.Time
	ProcessSize  uint64
	ProcsRunning uint64
	ProcsBlocked uint64
	Softirq      uint64
}

type cpuLoad struct {
	UserPercent   float64
	NicePercent   float64
	SystemPercent float64
	IdlePercent   float64
	Irq           float64
	Softirq       float64
}

func (o *cpuParse) ParseContent(content string) interface{} {
	var result CpuStats
	stringTab := strings.Split(content, "\n")

	index := &result.NumberCpu
	result.Cpus = make([]cpuLoad, 1)

	for _, line := range stringTab {
		lineSplit := strings.Split(line, " ")
		switch {
		case lineSplit[0] == _CPU:
			cpuMapping(&result.Cpu, lineSplit[1:])
			break
		case line[:_CPU_SIZE] == _CPU:
			cpuMapping(&result.Cpus[*index], lineSplit[1:])
			*index++
			result.Cpus = append(result.Cpus, cpuLoad{})
			break
		case line[:_CTXT_SIZE] == _CTXT:
			result.Context, _ = strconv.ParseUint(lineSplit[1], 10, 64)
			break
		case line[:_BTIME_SIZE] == _BTIME:
			btime, _ := strconv.ParseInt(lineSplit[1], 10, 64)
			result.BootTime = time.Unix(btime, 0)
			break
		case line[:_PROCESSES_SIZE] == _PROCESSES:
			result.ProcessSize, _ = strconv.ParseUint(lineSplit[1], 10, 64)
			break
		case line[:_PROCS_RUNNING_SIZE] == _PROCS_RUNNING:
			result.ProcsRunning, _ = strconv.ParseUint(lineSplit[1], 10, 64)
			break
		case line[:_PROCS_BLOCKED_SIZE] == _PROCS_BLOCKED:
			result.ProcsBlocked, _ = strconv.ParseUint(lineSplit[1], 10, 64)
			break
		case line[:_SOFTIRQ_SIZE] == _SOFTIRQ:
			result.Softirq, _ = strconv.ParseUint(lineSplit[1], 10, 64)
			break
		default:
			break
		}
	}

	return result
}

func cpuMapping(cpu *cpuLoad, value []string) {
	cpu.UserPercent, _ = strconv.ParseFloat(value[0], 64)
	cpu.NicePercent, _ = strconv.ParseFloat(value[1], 64)
	cpu.SystemPercent, _ = strconv.ParseFloat(value[2], 64)
	cpu.IdlePercent, _ = strconv.ParseFloat(value[3], 64)
	cpu.Irq, _ = strconv.ParseFloat(value[4], 64)
	cpu.Softirq, _ = strconv.ParseFloat(value[5], 64)
}
