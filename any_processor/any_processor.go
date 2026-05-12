package any_processor

import (
	"fmt"
	"strconv"
)

func ProcessAny(val any) any {
	if val == nil {
		return "NULL"
	}

	switch v := val.(type) {

	case int, int8, int16, int32, int64:
		// All integers are in int64 and multiplied by 10
		return toInt64(v) * 10

	case uint, uint8, uint16, uint32, uint64:
		return toUint64(v) * 10

	case string:
		// trying to convert a string to a number
		if num, ok := tryParseInt(v); ok {
			return num
		}
		return "String - " + v

	case float32, float64:
		return toFloat64(v) * 2.5

	case bool:
		return !v // inverting bool

	case []any, []int, []string:
		return fmt.Sprintf("Slice[%d elements]", getLen(v))

	default:
		return v // return it as for unknown types

	}
}
func tryParseInt(s string) (int, bool) {
	num, err := strconv.Atoi(s)
	return num, err == nil
}

func toInt64(v any) int64 {
	switch x := v.(type) {
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	default:
		return 0
	}
}

func toUint64(v any) uint64 {
	switch x := v.(type) {
	case uint:
		return uint64(x)
	case uint8:
		return uint64(x)
	case uint16:
		return uint64(x)
	case uint32:
		return uint64(x)
	case uint64:
		return x
	default:
		return 0
	}
}
func toFloat64(v any) float64 {
	switch x := v.(type) {
	case float32:
		return float64(x)
	case float64:
		return x
	default:
		return 0
	}
}

func getLen(v any) int {
	switch x := v.(type) {
	case []any:
		return len(x)
	case []int:
		return len(x)
	case []string:
		return len(x)
	default:
		return 0
	}
}
