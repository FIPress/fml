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
	errValueNotFound = errors.New("Value not found")
	errTypeMismatch  = errors.New("Type mismatch")
	errNoKey         = errors.New("No key name")
)

type FML struct {
	dict map[string]interface{}
}

func NewFml() *FML {
	return &FML{make(map[string]interface{})}
}

func (t *FML) GetStringEx(key string) (val string, err error) {
	raw, err := t.getRawVal(key)

	switch v := raw.(type) {
	case string:
		val = v
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func (t *FML) GetString(key string, dflt string) string {
	raw, err := t.getRawVal(key)
	if err != nil {
		return dflt
	}

	switch v := raw.(type) {
	case string:
		return v
	case nil:
		return dflt
	default:
		return wrapVal(v)
	}
}

func (t *FML) GetBoolEx(key string) (val bool, err error) {
	raw, err := t.getRawVal(key)
	if err != nil {
		return
	}

	switch v := raw.(type) {
	case bool:
		val = v
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func (t *FML) GetBool(key string, dflt bool) bool {
	raw, err := t.getRawVal(key)
	if err != nil {
		return dflt
	}

	switch v := raw.(type) {
	case bool:
		return v
	default:
		return dflt
	}
}

func (t *FML) GetIntEx(key string) (val int, err error) {
	raw, err := t.getRawVal(key)
	if err != nil {
		return
	}

	switch v := raw.(type) {
	case int:
		val = v
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func (t *FML) GetInt(key string, dflt int) int {
	raw, err := t.getRawVal(key)
	if err != nil {
		return dflt
	}

	switch v := raw.(type) {
	case int:
		return v
	default:
		return dflt
	}
}

func (t *FML) GetFloatEx(key string) (val float64, err error) {
	raw, err := t.getRawVal(key)
	if err != nil {
		return
	}

	switch v := raw.(type) {
	case float64:
		val = v
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func (t *FML) GetFloat(key string, dflt float64) float64 {
	raw, err := t.getRawVal(key)
	if err != nil {
		return dflt
	}

	switch v := raw.(type) {
	case float64:
		return v
	default:
		return dflt
	}
}

func (t *FML) GetDatetimeEx(key string) (val time.Time, err error) {
	raw, err := t.getRawVal(key)
	if err != nil {
		return
	}

	switch v := raw.(type) {
	case time.Time:
		val = v
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}

func (t *FML) GetDatetime(key string, dflt time.Time) time.Time {
	raw, err := t.getRawVal(key)
	if err != nil {
		return dflt
	}

	switch v := raw.(type) {
	case time.Time:
		return v
	default:
		return dflt
	}
}

func (t *FML) GetArrayEx(key string) (array interface{}, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, t)
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

func (t *FML) GetStringArray(key string) []string {
	fKey, doc, err := getFinalKeyAndNode(key, t)
	if err != nil {
		return nil
	}

	switch arr := doc.dict[fKey].(type) {
	case []string:
		return arr
	default:
		return nil
	}
}

func (t *FML) GetBoolArray(key string) []bool {
	fKey, doc, err := getFinalKeyAndNode(key, t)
	if err != nil {
		return nil
	}

	switch arr := doc.dict[fKey].(type) {
	case []bool:
		return arr
	default:
		return nil
	}
}

func (t *FML) GetIntArray(key string) []int {
	fKey, doc, err := getFinalKeyAndNode(key, t)
	if err != nil {
		return nil
	}

	switch arr := doc.dict[fKey].(type) {
	case []int:
		return arr
	default:
		return nil
	}
}

func (t *FML) GetFloatArray(key string) []float64 {
	fKey, doc, err := getFinalKeyAndNode(key, t)
	if err != nil {
		return nil
	}

	switch arr := doc.dict[fKey].(type) {
	case []float64:
		return arr
	default:
		return nil
	}
}

func (t *FML) GetDatetimeArray(key string) []time.Time {
	fKey, doc, err := getFinalKeyAndNode(key, t)
	if err != nil {
		//fmt.Println("key,err:",key,err)
		return nil
	}

	switch arr := doc.dict[fKey].(type) {
	case []time.Time:
		return arr
	default:
		return nil
	}
}

func (t *FML) GetNode(key string) (node *FML, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, t)
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

func (t *FML) GetNodeList(key string) (list []*FML, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, t)
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

func (t *FML) getRawVal(key string) (val interface{}, err error) {
	fKey, doc, err := getFinalKeyAndNode(key, t)
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

/*
func (t toml) getStruct(key string, st interface {}) ( err error) {
	switch doc := t.dict[key].(type){
	case *toml:
		v := reflect.ValueOf(st).Elem()
		tp := v.Type()
		for i:=0; i< v.NumField(); i++ {
			f := v.Field(i)
			name := tp.Field(i).Name
			switch f.Type() {
			case reflect.String:
				s, e := doc.GetString(name)
				if e != nil {
					f.SetString(s)
				}
				case reflect.Int


			}
		}

		for key := range doc.dict {
			v.FieldByName(key) = doc[key] //by type recursive
		}
	case nil:
		err = errValueNotFound
	default:
		err = errTypeMismatch
	}
	return
}
*/

func (t *FML) SetValue(key string, v interface{}) {
	t.dict[key] = v
}

func (t *FML) RemoveItem(key string) {
	delete(t.dict, key)
}

func (t *FML) ValueSet() (values []interface{}) {
	values = make([]interface{}, len(t.dict))

	idx := 0
	for _, v := range t.dict {
		values[idx] = v
		idx++
	}
	return
}

func (t *FML) KeySet() (keys []string) {
	keys = make([]string, len(t.dict))

	idx := 0
	for k := range t.dict {
		keys[idx] = k
		idx++
	}
	return
}

/*
func (t *Toml) SetTable(key string, v *Toml) {

}*/

func (doc *FML) WriteToFile(path string) (err error) {
	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return
	}

	writer := bufio.NewWriter(file)
	doc.WriteTo(writer)
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
