package convutil

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

func I2String(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}

func I2Int(src interface{}) (int, error) {
	switch src.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		i32, err := strconv.Atoi(fmt.Sprintf("%d", src))
		if err != nil {
			return 0, err
		}
		return i32, nil
	case float32, float64:
		i32, err := strconv.Atoi(fmt.Sprintf("%0.0f", src))
		if err != nil {
			return 0, err
		}
		return i32, nil
	default:
		return 0, errors.New("I2Int is receive not support interface type")
	}
}

func I2Int64(src interface{}) (int64, error) {
	switch src.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		i64, err := strconv.ParseInt(fmt.Sprintf("%d", src), 10, 64)
		if err != nil {
			return 0, err
		}
		return i64, nil
	case float32, float64:
		i64, err := strconv.ParseInt(fmt.Sprintf("%0.0f", src), 10, 64)
		if err != nil {
			return 0, err
		}
		return i64, nil
	default:
		return 0, errors.New("I2Int64 is receive not support interface type")
	}
}

func I2Float64(src interface{}) (float64, error) {
	switch src.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		f64, err := strconv.ParseFloat(I2String(src), 64)
		if err != nil {
			return 0.0, err
		}
		return f64, nil
	default:
		return 0.0, errors.New("I2Float64 is receive not support interface type")
	}
}

func I2Bool(src interface{}) bool {
	if src == nil {
		return false
	}
	switch i2 := src.(type) {
	default:
		return false
	case bool:
		return i2
	case *bool:
		if i2 == nil {
			return false
		}
		return *i2
	}
}

func Obj2Map[T any](obj T) map[string]any {
	mapResult := map[string]any{}
	jsonStr, _ := json.Marshal(&obj)
	err := json.Unmarshal(jsonStr, &mapResult)
	if err != nil {
		return nil
	}
	return mapResult
}

func Str2Map(data string) map[string]any {
	mapResult := map[string]any{}
	if err := json.Unmarshal([]byte(data), &mapResult); err == nil {
		return mapResult
	}
	return nil
}
