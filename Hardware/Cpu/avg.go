package cpu

import (
	"strconv"
	"strings"
)

type AvgParse struct {
}

type AvgStat struct {
	InstantAvg, OneAvg, FiveAvg float32
}

func (o *AvgParse) ParseContent(content string) interface{} {
	var result AvgStat
	values := strings.Split(content, " ")

	InstantAvg, _ := strconv.ParseFloat(values[0], 64)
	OneAvg, _ := strconv.ParseFloat(values[1], 64)
	FiveAvg, _ := strconv.ParseFloat(values[2], 64)

	result.InstantAvg = float32(InstantAvg)
	result.OneAvg = float32(OneAvg)
	result.FiveAvg = float32(FiveAvg)

	return result
}
