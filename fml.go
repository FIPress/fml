package fml

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	errValueNotFound = errors.New("value not found")
	errTypeMismatch  = errors.New("type mismatch")
	errNoKey         = errors.New("no key name")
)

// A FML represents a fml node
type FML struct {
	dict map[string]interface{}
}

// NewFml creates an empty node
func NewFml() *FML {
	return &FML{make(map[string]interface{})}
}

// GetStringOrError gets a string value from the node.
func (f *FML) GetStringOrError(key string) (val string, err error) {
	raw, err := f.getRawVal(key)

	return getString(raw)
}

// GetString gets a string value from the node,
// or return zero value of string, "", if the key does not exist or other error happens.
func (f *FML) GetString(key string) (val string) {
	return f.GetStringOrDefault(key, "")
}

// GetStringOrDefault gets a string value from the node,
// or return defaultVal if the key does not exist or other error happens.
func (f *FML) GetStringOrDefault(key string, defaultVal string) string {
	raw, err := f.getRawVal(key)
	if err != nil {
		return defaultVal
	}

	val, err := getString(raw)
	if err != nil {
		return defaultVal
	} else {
		return val
	}
}

// GetBoolOrError gets a bool value from the node.
func (f *FML) GetBoolOrError(key string) (val bool, err error) {
	raw, err := f.getRawVal(key)
	if err != nil {
		return
	}

	return getBool(raw)
}

// GetBoolOrZero gets a bool value from the node,
// or return zero value of bool, false, if the key does not exist or other error happens.
func (f *FML) GetBool(key string) (val bool) {
	return f.GetBoolOrDefault(key, false)
}

// GetBoolOrDefault gets a bool value from the node,
// or return defaultVal, if the key does not exist or other error happens.
func (f *FML) GetBoolOrDefault(key string, defaultVal bool) bool {
	raw, err := f.getRawVal(key)
	if err != nil {
		return defaultVal
	}

	val, err := getBool(raw)
	if err != nil {
		return defaultVal
	} else {
		return val
	}
}

// GetIntOrError gets a int value from the node.
func (f *FML) GetIntOrError(key string) (val int, err error) {
	raw, err := f.getRawVal(key)
	if err != nil {
		return
	}

	return getInt(raw)
}

// GetIntOrZero gets a int value from the node,
// or return zero value of int, 0, if the key does not exist or other error happens.
func (f *FML) GetInt(key string) (val int) {
	return f.GetIntOrDefault(key, 0)
}

// GetIntOrDefault gets a int value from the node,
// or return defaultVal, if the key does not exist or other error happens.
func (f *FML) GetIntOrDefault(key string, defaultVal int) int {
	raw, err := f.getRawVal(key)
	if err != nil {
		return defaultVal
	}

	val, err := getInt(raw)
	if err != nil {
		return defaultVal
	} else {
		return val
	}
}

// GetFloatOrError gets a float value from the node.
func (f *FML) GetFloatOrError(key string) (val float64, err error) {
	raw, err := f.getRawVal(key)
	if err != nil {
		return
	}

	return getFloat(raw)
}

// GetFloat gets a float value from the node,
// or return zero value of int, 0, if the key does not exist or other error happens.
func (f *FML) GetFloat(key string) (val float64) {
	return f.GetFloatOrDefault(key, 0)
}

// GetFloatOrDefault gets a float value from the node,
// or return defaultVal, if the key does not exist or other error happens.
func (f *FML) GetFloatOrDefault(key string, defaultVal float64) float64 {
	raw, err := f.getRawVal(key)
	if err != nil {
		return defaultVal
	}

	val, err := getFloat(raw)
	if err != nil {
		return defaultVal
	} else {
		return val
	}
}

// GetTimeOrError gets a time.Time value from the node.
func (f *FML) GetTimeOrError(key string) (val time.Time, err error) {
	raw, err := f.getRawVal(key)
	if err != nil {
		return
	}

	return getTime(raw)
}

// GetDatetime gets a time.Time value from the node,
// or return zero value of Datetime, time.Time{}, if the key does not exist or other error happens.
func (f *FML) GetDatetime(key string) (val time.Time) {
	return f.GetDatetimeOrDefault(key, time.Time{})
}

// GetDatetimeOrDefault gets a time.Time value from the node,
// or return defaultVal, if the key does not exist or other error happens.
func (f *FML) GetDatetimeOrDefault(key string, defaultVal time.Time) time.Time {
	raw, err := f.getRawVal(key)
	if err != nil {
		return defaultVal
	}

	val, err := getTime(raw)
	if err != nil {
		return defaultVal
	} else {
		return val
	}
}

// GetArrayOrError gets an array from the node.
func (f *FML) GetArrayOrError(key string) (array interface{}, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	switch arr := doc.dict[fKey].(type) {
	case []string:
		array = arr
	case []bool:
		array = arr
	case []int:
		array = arr
	case []float64:
		array = arr
	case []time.Time:
		array = arr
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

// GetStringArrayOrError gets an array of string from the node.
func (f *FML) GetStringArrayOrError(key string) (arr []string, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return
	}

	return getStringArray(doc.dict[fKey])
}

// GetStringArray gets an array of string from the node.
// It returns nil if the key does not exist or other error happens.
func (f *FML) GetStringArray(key string) []string {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return nil
	}

	arr, _ := getStringArray(doc.dict[fKey])
	return arr
}

// GetBoolArrayOrError gets an array of bool from the node.
func (f *FML) GetBoolArrayOrError(key string) (arr []bool, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return
	}

	return getBoolArray(doc.dict[fKey])
}

// GetBoolArray gets an array of bool from the node.
// It returns nil if the key does not exist or other error happens.
func (f *FML) GetBoolArray(key string) []bool {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return nil
	}

	arr, _ := getBoolArray(doc.dict[fKey])
	return arr
}

// GetIntArrayOrError gets an array of int from the node.
func (f *FML) GetIntArrayOrError(key string) (arr []int, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return
	}

	return getIntArray(doc.dict[fKey])
}

// GetIntArray gets an array of int from the node.
// It returns nil if the key does not exist or other error happens.
func (f *FML) GetIntArray(key string) []int {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return nil
	}

	arr, _ := getIntArray(doc.dict[fKey])
	return arr
}

// GetFloatArrayOrError gets an array of float from the node.
func (f *FML) GetFloatArrayOrError(key string) (arr []float64, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return
	}

	return getFloatArray(doc.dict[fKey])
}

// GetFloatArray gets an array of float from the node.
// It returns nil if the key does not exist or other error happens.
func (f *FML) GetFloatArray(key string) []float64 {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		return nil
	}

	arr, _ := getFloatArray(doc.dict[fKey])
	return arr
}

// GetDatetimeArray gets an array of time.Time from the node.
func (f *FML) GetTimeArrayOrError(key string) (arr []time.Time, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		//fmt.Println("key,err:",key,err)
		return
	}

	return getTimeArray(doc.dict[fKey])
}

// GetDatetimeArray gets an array of time.Time from the node.
// It returns nil if the key does not exist or other error happens.
func (f *FML) GetTimeArray(key string) []time.Time {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	if err != nil {
		//fmt.Println("key,err:",key,err)
		return nil
	}

	arr, _ := getTimeArray(doc.dict[fKey])
	return arr
}

// GetStruct get a struct from the node
func (f *FML) GetStruct(key string, v interface{}) (err error) {
	//fKey, node, err := getFinalKeyAndNode(key, f)
	node, err := f.GetNode(key)
	if err != nil {
		return
	}

	return getStruct(node, v)
}

// GetNode gets a sub-node from the node
func (f *FML) GetNode(key string) (node *FML, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	switch v := doc.dict[fKey].(type) {
	case *FML:
		node = v
	case nil:
		//err = errValueNotFound
		node = nil
	default:
		err = errTypeMismatch
	}
	return
}

// GetNodeList gets a node list from the node
func (f *FML) GetNodeList(key string) (list []*FML, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	switch v := doc.dict[fKey].(type) {
	case []*FML:
		list = v
	case nil:
		err = errValueNotFound
		//table = nil
	default:
		err = errTypeMismatch
	}
	return
}

func (f *FML) getRawVal(key string) (val interface{}, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, f)
	//fmt.Println("getRawVal:fkey:",fKey,"doc:",doc,",err:",err)
	if err != nil {
		return
	}

	if doc == nil || doc.dict == nil {
		err = errValueNotFound
		return
	}

	val, ok := doc.dict[fKey]
	if !ok {
		err = errValueNotFound
	}
	return
}

func getFinalKeyAndNode(key string, doc *FML) (finalKey string, finalToml *FML, err error) {
	if len(key) == 0 {
		err = errNoKey
		return
	}

	keys := strings.SplitN(key, ".", 2)
	//keys[0] is always there
	if len(keys[0]) == 0 {
		err = errNoKey
		return
	}
	if len(keys) == 1 {
		finalKey, finalToml = key, doc
	} else {
		if len(keys[1]) == 0 {
			err = errNoKey
			return
		}
		switch v := doc.dict[keys[0]].(type) {
		case *FML:
			finalKey, finalToml, err = getFinalKeyAndNode(keys[1], v)
		default:
			err = errValueNotFound
		}
	}

	return
}

func (f *FML) SetValue(key string, v interface{}) {
	f.dict[key] = v
}

func (f *FML) RemoveItem(key string) {
	delete(f.dict, key)
}

func (f *FML) ValueSet() (values []interface{}) {
	values = make([]interface{}, len(f.dict))

	idx := 0
	for _, v := range f.dict {
		values[idx] = v
		idx++
	}
	return
}

func (f *FML) KeySet() (keys []string) {
	keys = make([]string, len(f.dict))

	idx := 0
	for k := range f.dict {
		keys[idx] = k
		idx++
	}
	return
}

/*
func (t *Toml) SetTable(key string, v *Toml) {

}*/

func (f *FML) WriteToFile(path string) (err error) {
	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return
	}

	writer := bufio.NewWriter(file)
	f.WriteTo(writer)
	err = writer.Flush()
	return
}

func (f *FML) WriteTo(writer *bufio.Writer) {
	f.writeTo("", writer)
}

func (f *FML) writeTo(parentKey string, writer *bufio.Writer) {
	for key, val := range f.dict {
		fullKey := key
		if len(parentKey) != 0 {
			fullKey = parentKey + "." + key
		}
		switch v := val.(type) {
		case []*FML:
			defer func() {
				writeNodeName(fullKey, writer)
				for i := 0; i < len(v); i++ {
					writer.Write([]byte{'-', ' '})
					v[i].writeTo(fullKey, writer)
				}
			}()
		case *FML:
			defer func() {
				writeNodeName(fullKey, writer)
				v.writeTo(fullKey, writer)
			}()
		default:
			//log.Println("write",key)
			writer.WriteString(key)
			writer.WriteByte(':')
			writer.WriteString(wrapVal(v))
			//writer.WriteString("after key value")
			writer.WriteByte('\n')
		}
	}
}

func writeNodeName(key string, writer *bufio.Writer) {
	if writer.Buffered() > 0 {
		writer.WriteByte('\n')
	}
	writer.WriteByte('[')
	writer.WriteString(key)
	writer.WriteByte(']')
	writer.WriteByte('\n')
}

//func wrapSimpleVal(val interface {}) string {}

func wrapVal(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case time.Time:
		return v.Format(time.RFC3339)
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []time.Time:
		s := "["
		l := len(v)
		for i := 0; i < l; i++ {
			if i > 0 {
				s += ","
			}
			s += v[i].Format(time.RFC3339)
		}
		s += "]"
		return s
	case []int:
		s := "["
		l := len(v)
		for i := 0; i < l; i++ {
			if i > 0 {
				s += ","
			}
			s += strconv.Itoa(v[i])
		}
		s += "]"
		return s
	case []float64:
		s := "["
		l := len(v)
		for i := 0; i < l; i++ {
			if i > 0 {
				s += ","
			}
			s += strconv.FormatFloat(v[i], 'f', -1, 64)
		}
		s += "]"
		return s
	case []bool:
		s := "["
		l := len(v)
		for i := 0; i < l; i++ {
			if i > 0 {
				s += ","
			}
			s += strconv.FormatBool(v[i])
		}
		s += "]"
		return s
	case []string:
		s := "["
		l := len(v)
		for i := 0; i < l; i++ {
			if i > 0 {
				s += ","
			}
			s += v[i]
		}
		s += "]"
		return s
	default:
		return fmt.Sprint(v)
	}
}

//for test
func IterateFimlDoc(doc *FML) {
	dict := doc.dict
	count := len(dict)
	fmt.Println("length of fml:", count)
	for key := range dict {
		switch val := dict[key].(type) {
		case nil:
			fmt.Println("empty value of: ", key)
		case string:
			fmt.Println(key, "=", val, "(string)")
		case *FML:
			fmt.Println("Node:", key)
			IterateFimlDoc(val)
		case []*FML:
			fmt.Println("List of nodes:", key, ",len:", len(val))

			for i, v := range val {
				fmt.Println("List ", i)
				IterateFimlDoc(v)
			}
		default:
			fmt.Println(key, "=", val, "(", reflect.TypeOf(val), ")")
		}
	}
}
