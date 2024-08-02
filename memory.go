package main

import (
	"fmt"
	"reflect"
	"strings"
)

type memParse struct{}

type MemStats struct {
	Total     uint64 `fieldname:"MemTotal"`
	Free      uint64 `fieldname:"MemFree"`
	Available uint64 `fieldname:"MemAvailable"`
	SwapFree  uint64 `fieldname:"SwapFree"`
	SwapTotal uint64 `fieldname:"SwapTotal"`
}

func (o *memParse) ParseContent(content string) interface{} {
	result := MemStats{}
	dictValue := make(map[string]uint64)
	content = strings.ReplaceAll(content, ":", "")
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		var name string
		var value uint64
		if n, _ := fmt.Sscanf(line, "%s %d kB", &name, &value); n == 2 {
			dictValue[name] = value
		}
	}
	mapValue(dictValue, &result)
	return result
}

func mapValue(dict map[string]uint64, stats *MemStats) {
	t := reflect.TypeOf(*stats)
	v := reflect.ValueOf(stats)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv := v.Elem().Field(i)
		value := dict[field.Tag.Get("fieldname")]

		switch fv.Kind() {
		case reflect.Uint64:
			fv.SetUint(value)
			break
		default:
		}
	}
}
