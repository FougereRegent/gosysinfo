package cpu

import "time"

type parsingMethodDict map[string]func(string) []string

const (
	_CPU           string = "cpu"
	_INTR                 = "intr"
	_CTXT                 = "ctxt"
	_BTIME                = "btime"
	_PROCESSES            = "processes"
	_PROCS_RUNING         = "procs_running"
	_PROCS_BLOCKED        = "procs_blocked"
	_SOFTIRQ              = "softirq"
)

const (
	_CPU_SIZE           uint = 3
	_INTR_SIZE               = 4
	_CTXT_SIZE               = 4
	_BTIME_SIZE              = 5
	_PROCESSES_SIZE          = 9
	_PROCS_RUNING_SIZE       = 12
	_PROCS_BLOCKED_SIZE      = 13
	_SOFTIRQ_SIZE            = 7
)

var _dict = parsingMethodDict{
	_CPU:           parseMultipleValue(_CPU_SIZE, 10),
	_CTXT:          parseSingleValue(_CTXT_SIZE),
	_BTIME:         parseSingleValue(_BTIME_SIZE),
	_PROCESSES:     parseSingleValue(_PROCESSES_SIZE),
	_PROCS_RUNING:  parseSingleValue(_PROCS_RUNING_SIZE),
	_PROCS_BLOCKED: parseSingleValue(_PROCS_BLOCKED_SIZE),
	_SOFTIRQ:       parseMultipleValue(_SOFTIRQ_SIZE, 11),
}

type CpuParse struct {
}

type CpuStats struct {
	NumberCpu uint16
	Cpus      []cpuLoad
	Cpu       cpuLoad
	Context   uint64
	BootTime  time.Time
}

type cpuLoad struct {
	UserPercent   float64
	NicePercent   float64
	SystemPercent float64
	IdlePercent   float64
	Irq           float64
	Softirq       float64
}

func (o *CpuParse) ParseContent() interface{} {
	result := new(CpuStats)
	return result
}

func parseMultipleValue(paramSize, nbValue uint) func(string) []string {
	return func(content string) []string {
		return nil
	}
}

func parseSingleValue(paramSize uint) func(string) []string {
	return func(content string) []string {
		return nil
	}
}
