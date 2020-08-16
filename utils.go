package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Cast2Float64 cast database driver interface{} to float64
func Cast2Float64(t interface{}) float64 {
	switch v := t.(type) {
	case int64:
		return float64(v)
	case float64:
		return v
	case time.Time:
		return float64(v.Unix())
	case time.Duration:
		return float64(v.Nanoseconds())
	case []byte:
		strV := string(v)
		result, err := strconv.ParseFloat(strV, 64)
		if err != nil {
			return math.NaN()
		}
		return result
	case string:
		result, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return math.NaN()
		}
		return result
	case bool:
		if v {
			return 1.0
		}
		return 0.0
	case nil:
		return math.NaN()
	default:
		return math.NaN()
	}
}

// Cast2Float64 cast database driver interface{} to string
func Cast2string(t interface{}) string {
	switch v := t.(type) {
	case int64:
		return fmt.Sprintf("%v", v)
	case float64:
		return fmt.Sprintf("%v", v)
	case time.Time:
		return fmt.Sprintf("%v", v.Unix())
	case nil:
		return ""
	case []byte:
		return string(v)
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return ""
	}
}

func Contains(arr []string, x string) bool {
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}
