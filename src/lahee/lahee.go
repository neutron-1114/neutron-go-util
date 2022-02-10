package lahee

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type JsonNode struct {
	isNull   bool
	isArr    bool
	objValue map[string]interface{}
	arrValue []interface{}
	eleValue interface{}
}

type NotJsonError struct {
}

func (*NotJsonError) Error() string {
	return "It's not a json!!!"
}

func New(jsonStr string) (*JsonNode, error) {
	if jsonStr == "" {
		return nil, &NotJsonError{}
	}
	firstChar := jsonStr[0:1]
	if firstChar == "[" {
		var a []interface{}
		err := json.Unmarshal([]byte(jsonStr), &a)
		if err != nil {
			return nil, err
		}
		return &JsonNode{false, true, nil, a, a}, nil
	} else if firstChar == "{" {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(jsonStr), &m)
		if err != nil {
			return nil, err
		}
		return &JsonNode{false, false, m, nil, m}, nil
	}
	return nil, &NotJsonError{}
}

func (j *JsonNode) GetKey(key string) *JsonNode {
	if j.isNull || j.isArr || j.objValue == nil || len(j.objValue) == 0 {
		return nil
	}
	v, e := j.objValue[key]
	if !e {
		return nil
	}
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return &JsonNode{false, true, nil, v.([]interface{}), v}
	} else if reflect.TypeOf(v).Kind() == reflect.Map {
		return &JsonNode{false, false, v.(map[string]interface{}), nil, v}
	} else {
		return &JsonNode{false, false, nil, nil, v}
	}
}

func (j *JsonNode) GetIndex(index int) *JsonNode {
	if j.isNull || !j.isArr || j.arrValue == nil || index >= len(j.arrValue) {
		return nil
	}
	v := j.arrValue[index]
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return &JsonNode{false, true, nil, v.([]interface{}), v}
	} else if reflect.TypeOf(v).Kind() == reflect.Map {
		return &JsonNode{false, false, v.(map[string]interface{}), nil, v}
	} else {
		return &JsonNode{false, false, nil, nil, v}
	}
}

func (j *JsonNode) PathKey(key string) *JsonNode {
	if j.isNull || j.isArr || j.objValue == nil || len(j.objValue) == 0 {
		return &JsonNode{true, false, nil, nil, nil}
	}
	v, e := j.objValue[key]
	if !e {
		return &JsonNode{true, false, nil, nil, nil}
	}
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return &JsonNode{false, true, nil, v.([]interface{}), v}
	} else if reflect.TypeOf(v).Kind() == reflect.Map {
		return &JsonNode{false, false, v.(map[string]interface{}), nil, v}
	} else {
		return &JsonNode{false, false, nil, nil, v}
	}
}

func (j *JsonNode) PathIndex(index int) *JsonNode {
	if j.isNull || !j.isArr || j.arrValue == nil || index >= len(j.arrValue) {
		return &JsonNode{true, true, nil, nil, nil}
	}
	v := j.arrValue[index]
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return &JsonNode{false, true, nil, v.([]interface{}), v}
	} else if reflect.TypeOf(v).Kind() == reflect.Map {
		return &JsonNode{false, false, v.(map[string]interface{}), nil, v}
	} else {
		return &JsonNode{false, false, nil, nil, v}
	}
}

func (j *JsonNode) BoolOrDefault(defaultValue bool) bool {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Bool:
		return j.eleValue.(bool)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Bool.String()))
}

func (j *JsonNode) StringOrDefault(defaultValue string) string {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.String:
		return j.eleValue.(string)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.String.String()))
}

func (j *JsonNode) Int64OrDefault(defaultValue int64) int64 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Int64:
		return j.eleValue.(int64)
	case reflect.Int32:
		return int64(j.eleValue.(int32))
	case reflect.Int16:
		return int64(j.eleValue.(int16))
	case reflect.Int8:
		return int64(j.eleValue.(int8))
	case reflect.Uint32:
		return int64(j.eleValue.(uint32))
	case reflect.Uint16:
		return int64(j.eleValue.(uint16))
	case reflect.Uint8:
		return int64(j.eleValue.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int64.String()))
}

func (j *JsonNode) Int32OrDefault(defaultValue int32) int32 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Int32:
		return j.eleValue.(int32)
	case reflect.Int16:
		return int32(j.eleValue.(int16))
	case reflect.Int8:
		return int32(j.eleValue.(int8))
	case reflect.Uint16:
		return int32(j.eleValue.(uint16))
	case reflect.Uint8:
		return int32(j.eleValue.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int32.String()))
}

func (j *JsonNode) Int16OrDefault(defaultValue int16) int16 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Int16:
		return j.eleValue.(int16)
	case reflect.Int8:
		return int16(j.eleValue.(int8))
	case reflect.Uint8:
		return int16(j.eleValue.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int16.String()))
}

func (j *JsonNode) Int8OrDefault(defaultValue int8) int8 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Int8:
		return j.eleValue.(int8)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int8.String()))
}

func (j *JsonNode) UInt64OrDefault(defaultValue uint64) uint64 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Uint64:
		return j.eleValue.(uint64)
	case reflect.Uint32:
		return uint64(j.eleValue.(uint32))
	case reflect.Uint16:
		return uint64(j.eleValue.(uint16))
	case reflect.Uint8:
		return uint64(j.eleValue.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint64.String()))
}

func (j *JsonNode) UInt32OrDefault(defaultValue uint32) uint32 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Uint32:
		return j.eleValue.(uint32)
	case reflect.Uint16:
		return uint32(j.eleValue.(uint16))
	case reflect.Uint8:
		return uint32(j.eleValue.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint32.String()))
}

func (j *JsonNode) UInt16OrDefault(defaultValue uint16) uint16 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Uint16:
		return j.eleValue.(uint16)
	case reflect.Uint8:
		return uint16(j.eleValue.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint16.String()))
}

func (j *JsonNode) UInt8OrDefault(defaultValue uint8) uint8 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Uint8:
		return j.eleValue.(uint8)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint8.String()))
}

func (j *JsonNode) Float64OrDefault(defaultValue float64) float64 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Float64:
		return j.eleValue.(float64)
	case reflect.Float32:
		return float64(j.eleValue.(float32))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Float64.String()))
}

func (j *JsonNode) Float32OrDefault(defaultValue float32) float32 {
	if j.isNull {
		return defaultValue
	}
	kind := reflect.TypeOf(j.eleValue).Kind()
	switch kind {
	case reflect.Float32:
		return j.eleValue.(float32)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Float32.String()))
}

func (j *JsonNode) Size() int {
	if j.isNull {
		return 0
	}
	if j.arrValue != nil {
		return len(j.arrValue)
	}
	if j.objValue != nil {
		return len(j.objValue)
	}
	return 0
}

func (j *JsonNode) IsNullNode() bool {
	return j.isNull
}

func (j *JsonNode) IsArrNode() bool {
	return j.isArr && j.arrValue != nil
}

func (j *JsonNode) IsObjNode() bool {
	return !j.isArr && j.objValue != nil
}
