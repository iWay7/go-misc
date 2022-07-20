package misc

import "reflect"

func StringArrayRemoveDuplicated(array []string) []string {
	result := []string{}
	m := map[string]bool{}
	for _, v := range array {
		m[v] = true
	}
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func IntArrayContains(array []int, value int) bool {
	for _, i := range array {
		if i == value {
			return true
		}
	}
	return false
}

func Int64ArrayContains(array []int64, value int64) bool {
	for _, i := range array {
		if i == value {
			return true
		}
	}
	return false
}

func StringArrayContains(array []string, value string) bool {
	for _, i := range array {
		if i == value {
			return true
		}
	}
	return false
}

func StringArrayContainsAny(array []string, values []string) bool {
	for _, i := range array {
		if StringArrayContains(values, i) {
			return true
		}
	}
	return false
}

func SubList(array interface{}, pageSize, pageIndex int) interface{} {
	arrayValue := reflect.ValueOf(array)
	count := arrayValue.Len()
	arrayType := reflect.TypeOf(array)
	arrayInstance := reflect.New(arrayType).Elem()
	start := CalcStart(pageSize, pageIndex)
	limit := CalcLimit(pageSize, pageIndex)
	end := start + limit
	for i := start; i < end; i++ {
		if i >= 0 && i < count {
			arrayItem := arrayValue.Index(i)
			arrayInstance = reflect.Append(arrayInstance, arrayItem)
		}
	}
	return arrayInstance.Interface()
}

func FillStringSlice(src []string, val string) {
	for i := 0; i < len(src); i++ {
		src[i] = val
	}
}

func FillBoolSlice(src []bool, val bool) {
	for i := 0; i < len(src); i++ {
		src[i] = val
	}
}

func NewStringSlice(len int, def string) []string {
	result := make([]string, len)
	FillStringSlice(result, def)
	return result
}

func NewBoolSlice(len int, def bool) []bool {
	result := make([]bool, len)
	FillBoolSlice(result, def)
	return result
}

func CopyStringSlice(src []string) []string {
	return append([]string{}, src...)
}

func RemoveAtStringSlice(src []string, at int) []string {
	result := []string{}
	result = append(result, src[0:at]...)
	result = append(result, src[at+1:]...)
	return result
}
