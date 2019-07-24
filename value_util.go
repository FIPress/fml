package fml

import (
	"reflect"
	"strings"
	"time"
)

func getString(rawVal interface{}) (val string, err error) {
	switch v := rawVal.(type) {
	case string:
		val = v
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getBool(rawVal interface{}) (val bool, err error) {
	switch v := rawVal.(type) {
	case bool:
		val = v
	case string:
		var ok bool
		val, ok = evalBool(v)
		if !ok {
			err = errTypeMismatch
		}
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getInt(rawVal interface{}) (val int, err error) {
	switch v := rawVal.(type) {
	case int:
		val = v
	case string:
		var ok bool
		val, ok = evalInt(v)
		if !ok {
			err = errTypeMismatch
		}
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getFloat(rawVal interface{}) (val float64, err error) {
	switch v := rawVal.(type) {
	case float64:
		val = v
	case string:
		var ok bool
		val, ok = evalFloat(v)
		if !ok {
			err = errTypeMismatch
		}
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getTime(rawVal interface{}) (val time.Time, err error) {
	switch v := rawVal.(type) {
	case time.Time:
		val = v
	case string:
		var ok bool
		val, ok = evalDatetime(v)
		if !ok {
			err = errTypeMismatch
		}

	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getStruct(node *FML, val interface{}) (err error) {
	rv := reflect.ValueOf(val)
	//rv.Kind() == Struct?
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errTypeMismatch
	}

	el := rv.Elem()
	if el.Kind() != reflect.Struct {
		return errTypeMismatch
	}

	for k, v := range node.dict {
		field := el.FieldByName(k)
		if !field.IsValid() {
			field = el.FieldByName(strings.Title(k))
		}
		if !field.IsValid() {
			continue
		}
		switch field.Kind() {
		case reflect.String:
			s, err := getString(v)
			if err == nil {
				field.SetString(s)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := getInt(v)
			if err == nil {
				field.SetInt(int64(i))
			}
		case reflect.Float32, reflect.Float64:
			fl, err := getFloat(v)
			if err == nil {
				field.SetFloat(fl)
			}
		case reflect.Bool:
			b, err := getBool(v)
			if err == nil {
				field.SetBool(b)
			}
		case reflect.Struct:
			if field.Type().Name() == "time.Time" {
				dt, err := getTime(v)
				if err == nil {
					field.Set(reflect.ValueOf(dt))
				}
			} else {
				switch fml := v.(type) {
				case *FML:
					sub := reflect.New(field.Type())
					fml.GetStruct(k, sub.Interface())
				}

			}

		}
	}
	return
}

func getStringArray(rawVal interface{}) (arr []string, err error) {
	switch a := rawVal.(type) {
	case []string:
		arr = a
		return
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getBoolArray(rawVal interface{}) (arr []bool, err error) {
	switch a := rawVal.(type) {
	case []bool:
		arr = a
		return
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getIntArray(rawVal interface{}) (arr []int, err error) {
	switch a := rawVal.(type) {
	case []int:
		arr = a
		return
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getFloatArray(rawVal interface{}) (arr []float64, err error) {
	switch a := rawVal.(type) {
	case []float64:
		arr = a
		return
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getTimeArray(rawVal interface{}) (arr []time.Time, err error) {
	switch a := rawVal.(type) {
	case []time.Time:
		arr = a
		return
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func getStructArray(rawVal interface{}, out interface{}) (err error) {
	/*match := func() {
		switch o := out.(type) {

		}
	}
	switch a := rawVal.(type) {

	case []time.Time:
		arr = a
		return
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}*/
	return
}
