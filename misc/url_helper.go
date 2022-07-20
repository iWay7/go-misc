package misc

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func MakeUrl(proto string, ip string, port int, parts ...string) string {
	result := proto + "://" + ip + ":" + strconv.Itoa(port)
	for _, item := range parts {
		result += "/" + strings.Trim(item, "/")
	}
	return result
}

func urlQueryUnescape(fieldType reflect.Type, fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.Ptr:
		urlQueryUnescape(fieldType.Elem(), fieldValue.Elem())
	case reflect.Struct:
		numField := fieldType.NumField()
		for i := 0; i < numField; i++ {
			urlQueryUnescape(fieldType.Field(i).Type, fieldValue.Field(i))
		}
	case reflect.String:
		value := fieldValue.String()
		if fieldValue.CanSet() {
			unescapedValue, err := url.QueryUnescape(value)
			if err == nil {
				fieldValue.SetString(unescapedValue)
			}
		}
	case reflect.Array, reflect.Slice:
		arrLen := fieldValue.Len()
		for i := 0; i < arrLen; i++ {
			urlQueryUnescape(fieldType.Elem(), fieldValue.Index(i))
		}
	case reflect.Map:
		for _, mapKey := range fieldValue.MapKeys() {
			urlQueryUnescape(fieldType.Elem(), fieldValue.MapIndex(mapKey))
		}
	}
}

func UrlQueryUnescape(v interface{}) {
	dataType := reflect.TypeOf(v)
	dataValue := reflect.ValueOf(v)
	urlQueryUnescape(dataType, dataValue)
}

func urlQueryEscape(fieldType reflect.Type, fieldValue reflect.Value) {
	switch fieldValue.Kind() {
	case reflect.Ptr:
		urlQueryEscape(fieldType.Elem(), fieldValue.Elem())
	case reflect.Struct:
		numField := fieldType.NumField()
		for i := 0; i < numField; i++ {
			urlQueryEscape(fieldType.Field(i).Type, fieldValue.Field(i))
		}
	case reflect.String:
		value := fieldValue.String()
		if fieldValue.CanSet() {
			fieldValue.SetString(url.QueryEscape(value))
		}
	case reflect.Array, reflect.Slice:
		arrLen := fieldValue.Len()
		for i := 0; i < arrLen; i++ {
			urlQueryEscape(fieldType.Elem(), fieldValue.Index(i))
		}
	case reflect.Map:
		for _, mapKey := range fieldValue.MapKeys() {
			urlQueryEscape(fieldType.Elem(), fieldValue.MapIndex(mapKey))
		}
	}
}

func UrlQueryEscape(v interface{}) {
	dataType := reflect.TypeOf(v)
	dataValue := reflect.ValueOf(v)
	urlQueryEscape(dataType, dataValue)
}
