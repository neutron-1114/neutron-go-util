package myjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Json struct {
	root interface{}
}

type JsonNode struct {
	Null  bool
	Value interface{}
}

type NotJsonError struct {
}

func (*NotJsonError) Error() string {
	return "It's not a json!!!"
}

func NotJson() error {
	return &NotJsonError{}
}

func New(jsonStr string) (*Json, error) {
	if jsonStr == "" {
		return nil, NotJson()
	}
	firstChar := jsonStr[0:1]
	if firstChar == "[" {
		var a []interface{}
		err := json.Unmarshal([]byte(jsonStr), &a)
		if err != nil {
			return nil, err
		}
		return &Json{a}, nil
	} else if firstChar == "{" {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(jsonStr), &m)
		if err != nil {
			return nil, err
		}
		return &Json{m}, nil
	}
	return nil, NotJson()
}

func (j *Json) Get(key string) *JsonNode {
	if key == "" {
		return &JsonNode{false, j.root}
	}
	fields := strings.Split(key, ".")
	var node = j.root
	for _, v := range fields {
		if node == nil {
			return &JsonNode{true, nil}
		}
		kind := reflect.TypeOf(node).Kind()
		if kind == reflect.Invalid {
			return &JsonNode{true, nil}
		} else if kind == reflect.Map {
			value, e := node.(map[string]interface{})[v]
			if !e {
				return &JsonNode{true, nil}
			}
			node = value
		} else if kind == reflect.Slice {
			vv, e := strconv.ParseInt(v, 10, 64)
			if e != nil {
				return &JsonNode{true, node}
			}
			if int(vv) >= len(node.([]interface{})) {
				return &JsonNode{true, node}
			}
			value := node.([]interface{})[int(vv)]
			node = value
		} else {
			return &JsonNode{true, node}
		}
	}
	return &JsonNode{false, node}
}

func (j *JsonNode) BoolOrDefault(defaultValue bool) bool {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Bool:
		return j.Value.(bool)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Bool.String()))
}

func (j *JsonNode) StringOrDefault(defaultValue string) string {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.String:
		return j.Value.(string)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.String.String()))
}

func (j *JsonNode) Int64OrDefault(defaultValue int64) int64 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Int64:
		return j.Value.(int64)
	case reflect.Int32:
		return int64(j.Value.(int32))
	case reflect.Int16:
		return int64(j.Value.(int16))
	case reflect.Int8:
		return int64(j.Value.(int8))
	case reflect.Uint32:
		return int64(j.Value.(uint32))
	case reflect.Uint16:
		return int64(j.Value.(uint16))
	case reflect.Uint8:
		return int64(j.Value.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int64.String()))
}

func (j *JsonNode) Int32OrDefault(defaultValue int32) int32 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Int32:
		return j.Value.(int32)
	case reflect.Int16:
		return int32(j.Value.(int16))
	case reflect.Int8:
		return int32(j.Value.(int8))
	case reflect.Uint16:
		return int32(j.Value.(uint16))
	case reflect.Uint8:
		return int32(j.Value.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int32.String()))
}

func (j *JsonNode) Int16OrDefault(defaultValue int16) int16 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Int16:
		return j.Value.(int16)
	case reflect.Int8:
		return int16(j.Value.(int8))
	case reflect.Uint8:
		return int16(j.Value.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int16.String()))
}

func (j *JsonNode) Int8OrDefault(defaultValue int8) int8 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Int8:
		return j.Value.(int8)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Int8.String()))
}

func (j *JsonNode) UInt64OrDefault(defaultValue uint64) uint64 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Uint64:
		return j.Value.(uint64)
	case reflect.Uint32:
		return uint64(j.Value.(uint32))
	case reflect.Uint16:
		return uint64(j.Value.(uint16))
	case reflect.Uint8:
		return uint64(j.Value.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint64.String()))
}

func (j *JsonNode) UInt32OrDefault(defaultValue uint32) uint32 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Uint32:
		return j.Value.(uint32)
	case reflect.Uint16:
		return uint32(j.Value.(uint16))
	case reflect.Uint8:
		return uint32(j.Value.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint32.String()))
}

func (j *JsonNode) UInt16OrDefault(defaultValue uint16) uint16 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Uint16:
		return j.Value.(uint16)
	case reflect.Uint8:
		return uint16(j.Value.(uint8))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint16.String()))
}

func (j *JsonNode) UInt8OrDefault(defaultValue uint8) uint8 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Uint8:
		return j.Value.(uint8)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Uint8.String()))
}

func (j *JsonNode) Float64OrDefault(defaultValue float64) float64 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Float64:
		return j.Value.(float64)
	case reflect.Float32:
		return float64(j.Value.(float32))
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Float64.String()))
}

func (j *JsonNode) Float32OrDefault(defaultValue float32) float32 {
	if j.Null {
		return defaultValue
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Float32:
		return j.Value.(float32)
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Float32.String()))
}

func (j *JsonNode) Map() map[string]interface{} {
	if j.Null {
		return nil
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Map:
		return j.Value.(map[string]interface{})
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Map.String()))
}

func (j *JsonNode) List() []interface{} {
	if j.Null {
		return nil
	}
	kind := reflect.TypeOf(j.Value).Kind()
	switch kind {
	case reflect.Slice:
		return j.Value.([]interface{})
	}
	panic(fmt.Sprintf("value existed... But can not transform [%v] to [%v]", kind.String(), reflect.Slice.String()))
}
