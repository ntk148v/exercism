package flatten

import "reflect"

// Flatten takes a nested list and return a single flattened list with all values except nil/null
func Flatten(input interface{}) []interface{} {
	result := []interface{}{}
	nestedResult, ok := input.([]interface{})
	if !ok {
		return result
	}
	for _, v := range nestedResult {
		rt := reflect.TypeOf(v)
		if rt == nil {
			continue
		}
		switch rt.Kind() {
		case reflect.Slice:
			result = append(result, Flatten(v)...)
		case reflect.Int:
			result = append(result, v)
		}
	}
	return result
}
