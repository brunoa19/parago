package utils

import (
	"encoding/json"
	"strings"
)

func ParseValues(values string) (out []string) {
	tokens := strings.Split(values, ",")
	for _, t := range tokens {
		val := strings.TrimSpace(t)
		if val != "" {
			out = append(out, val)
		}
	}
	return
}

func CopyJsonData(src, dst interface{}) {
	data, _ := json.Marshal(src)
	json.Unmarshal(data, dst)
}
