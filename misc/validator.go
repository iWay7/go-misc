package misc

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidateError struct {
	Name  string
	Value string
	Desc  string
}

type ValidateBuilder struct {
	Errors []ValidateError
}

func (validateBuilder *ValidateBuilder) AddError(name string, value interface{}, desc string) {
	validateBuilder.Errors = append(validateBuilder.Errors, ValidateError{
		Name:  name,
		Value: fmt.Sprint(value),
		Desc:  desc,
	})
}

func (validateBuilder *ValidateBuilder) Count() int {
	return len(validateBuilder.Errors)
}

func (validateBuilder *ValidateBuilder) Desc() string {
	desc := ""
	for index, validateError := range validateBuilder.Errors {
		desc += strconv.Itoa(index) + ". " + validateError.Name + " " + validateError.Desc + "\n"
	}
	return strings.TrimSuffix(desc, "\n")
}

type ValidateParam struct {
	Name   string
	Params []string
}

func (validateParam *ValidateParam) GetIntParam(index int) (int64, bool) {
	if len(validateParam.Params) > index {
		if i, err := strconv.ParseInt(validateParam.Params[index], 10, 64); err == nil {
			return i, true
		}
	}
	return 0, false
}

func (validateParam *ValidateParam) GetFloatParam(index int) (float64, bool) {
	if len(validateParam.Params) > index {
		if i, err := strconv.ParseFloat(validateParam.Params[index], 64); err == nil {
			return i, true
		}
	}
	return 0, false
}

func (validateParam *ValidateParam) GetUintParam(index int) (uint64, bool) {
	if len(validateParam.Params) > index {
		if i, err := strconv.ParseUint(validateParam.Params[index], 10, 64); err == nil {
			return i, true
		}
	}
	return 0, false
}

func parseValidateParams(valid string) []ValidateParam {
	validateParams := []ValidateParam{}
	validItems := strings.Split(valid, " ")
	for _, item := range validItems {
		if item == "" {
			continue
		}
		item = strings.Replace(item, "(", " ", -1)
		item = strings.Replace(item, ")", " ", -1)
		item = strings.Replace(item, ",", " ", -1)
		nameAndParams := strings.Split(item, " ")
		if len(nameAndParams) > 0 {
			validateParams = append(validateParams, ValidateParam{
				Name:   nameAndParams[0],
				Params: nameAndParams[1:],
			})
		}
	}
	return validateParams
}

func validatePtr(v interface{}, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "not_nil":
			if v == nil {
				result = false
				desc += "，" + "必须提供"
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validateString(value string, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "not_empty":
			if len(value) == 0 {
				result = false
				desc += "，" + "不能为空"
			}
		case "len_min":
			if l, ok := validParam.GetIntParam(0); ok {
				if len(value) < int(l) {
					result = false
					desc += "，" + "长度不足"
				}
			}
		case "len_max":
			if l, ok := validParam.GetIntParam(0); ok {
				if len(value) > int(l) {
					result = false
					desc += "，" + "长度过长"
				}
			}
		case "ip_v4":
			pattern := `^(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}$`
			match, err := regexp.MatchString(pattern, value)
			if err != nil {
				result = false
				desc += "，" + "验证出错：" + err.Error()
			} else if !match {
				result = false
				desc += "，" + "需要符合正则表达式：" + pattern
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validateInt(value int64, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "positive":
			if value <= 0 {
				result = false
				desc += "，" + "必须为正整数"
			}
		case "negative":
			if value >= 0 {
				result = false
				desc += "，" + "必须为负整数"
			}
		case "min":
			if i, ok := validParam.GetIntParam(0); ok {
				if value < i {
					result = false
					desc += "，" + "必须大于" + strconv.FormatInt(i, 10)
				}
			}
		case "max":
			if i, ok := validParam.GetIntParam(0); ok {
				if value > i {
					result = false
					desc += "，" + "必须小于" + strconv.FormatInt(i, 10)
				}
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validateUint(value uint64, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "min":
			if i, ok := validParam.GetUintParam(0); ok {
				if value < i {
					result = false
					desc += "，" + "必须大于" + strconv.FormatUint(i, 10)
				}
			}
		case "max":
			if i, ok := validParam.GetUintParam(0); ok {
				if value > i {
					result = false
					desc += "，" + "必须小于" + strconv.FormatUint(i, 10)
				}
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validateFloat(value float64, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "positive":
			if value <= 0 {
				result = false
				desc += "，" + "必须为正数"
			}
		case "negative":
			if value >= 0 {
				result = false
				desc += "，" + "必须为负数"
			}
		case "min":
			if f, ok := validParam.GetFloatParam(0); ok {
				if value < f {
					result = false
					desc += "，" + "必须大于" + strconv.FormatFloat(f, 'f', -1, 64)
				}
			}
		case "max":
			if f, ok := validParam.GetFloatParam(0); ok {
				if value > f {
					result = false
					desc += "，" + "必须大于" + strconv.FormatFloat(f, 'f', -1, 64)
				}
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validateNotNil(isNil bool, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "not_nil":
			if isNil {
				result = false
				desc += "，" + "必须提供"
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validateNotEmpty(isEmpty bool, valid string) (bool, string) {
	validParams := parseValidateParams(valid)
	result := true
	desc := ""
	for _, validParam := range validParams {
		switch validParam.Name {
		case "not_empty":
			if isEmpty {
				result = false
				desc += "，" + "不能为空"
			}
		}
	}
	return result, strings.TrimPrefix(desc, "，")
}

func validate(validateBuilder *ValidateBuilder, fieldName string, fieldType reflect.Type, fieldValue reflect.Value, valid string) {
	switch fieldValue.Kind() {
	case reflect.Ptr:
		if fieldValue.IsNil() {
			pass, desc := validatePtr(nil, valid)
			if !pass {
				validateBuilder.AddError(fieldName, "", desc)
			}
		} else {
			validate(validateBuilder, fieldName+".Elem", fieldType.Elem(), fieldValue.Elem(), valid)
		}
	case reflect.Struct:
		numField := fieldValue.NumField()
		for i := 0; i < numField; i++ {
			validate(validateBuilder, fieldName+"."+fieldType.Field(i).Name, fieldType.Field(i).Type, fieldValue.Field(i), fieldType.Field(i).Tag.Get("valid"))
		}
	case reflect.String:
		value := fieldValue.String()
		pass, desc := validateString(value, valid)
		if !pass {
			validateBuilder.AddError(fieldName, value, desc)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value := fieldValue.Int()
		pass, desc := validateInt(value, valid)
		if !pass {
			validateBuilder.AddError(fieldName, value, desc)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value := fieldValue.Uint()
		pass, desc := validateUint(value, valid)
		if !pass {
			validateBuilder.AddError(fieldName, value, desc)
		}
	case reflect.Float32, reflect.Float64:
		value := fieldValue.Float()
		pass, desc := validateFloat(value, valid)
		if !pass {
			validateBuilder.AddError(fieldName, value, desc)
		}
	case reflect.Complex64, reflect.Complex128:
	case reflect.Bool:
	case reflect.Array, reflect.Slice:
		if fieldValue.IsNil() {
			pass, desc := validateNotNil(true, valid)
			if !pass {
				validateBuilder.AddError(fieldName, "", desc)
			} else {
				pass, desc = validateNotEmpty(true, valid)
				if !pass {
					validateBuilder.AddError(fieldName, "", desc)
				}
			}
		} else {
			arrayOrSliceLen := fieldValue.Len()
			if arrayOrSliceLen == 0 {
				pass, desc := validateNotEmpty(true, valid)
				if !pass {
					validateBuilder.AddError(fieldName, "", desc)
				}
			} else {
				for i := 0; i < arrayOrSliceLen; i++ {
					validate(validateBuilder, fieldName+"["+strconv.Itoa(i)+"]", fieldType.Elem(), fieldValue.Index(i), "")
				}
			}
		}
	case reflect.Map:
		if fieldValue.IsNil() {
			pass, desc := validateNotNil(true, valid)
			if !pass {
				validateBuilder.AddError(fieldName, "", desc)
			} else {
				pass, desc = validateNotEmpty(true, valid)
				if !pass {
					validateBuilder.AddError(fieldName, "", desc)
				}
			}
		} else {
			mapLen := fieldValue.Len()
			if mapLen == 0 {
				pass, desc := validateNotEmpty(true, valid)
				if !pass {
					validateBuilder.AddError(fieldName, "", desc)
				}
			} else {
				for _, mapKey := range fieldValue.MapKeys() {
					if mapKey.CanInterface() {
						key := fmt.Sprint(mapKey.Interface())
						validate(validateBuilder, fieldName+"["+key+"]", fieldType.Elem(), fieldValue.MapIndex(mapKey), "")
					} else {
						validate(validateBuilder, fieldName+"[UnknownKey]", fieldType.Elem(), fieldValue.MapIndex(mapKey), "")
					}
				}
			}
		}
	}
}

func Validate(v interface{}) (int, string) {
	validateBuilder := &ValidateBuilder{}
	dataType := reflect.TypeOf(v)
	dataValue := reflect.ValueOf(v)
	validate(validateBuilder, "Req", dataType, dataValue, "")
	return validateBuilder.Count(), validateBuilder.Desc()
}
