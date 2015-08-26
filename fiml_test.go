package fipml

/*import (
	"testing"
)


func TestFloat(t *testing.T) {
	f := []float64{123.34,22.2}
	i := []int{2,3}
	s := []string{"a","b"}
	SwitchType(f,t)
	SwitchType(i,t)
	SwitchType(s,t)

}

func SwitchType(in []interface {},t *testing.T) {
	switch in.(type){
	case []interface {}:
		t.Log("array")
	default:
		t.Log("other")
	}
}*/

/*
const (
	times = 500000

	//str2 = "abc"
)
var str = []string{"1dadf","falhuydd","true","123","12.3","abc"}

func TestBenchAtoi1(t *testing.T) {
	var val interface {}
	var ok bool
	for i:=0;i<times;i++ {
		for j:=0;j<len(str);j++ {
			val,ok = evalBool(str[j])
			if ok {
				continue
			}
			val, ok = evalInt(str[j])
			if ok {
				continue
			}
			val, ok = evalFloat(str[j])
			if ok {
				continue
			}
			val, ok = evalDatetime(str[j])
			if ok {
				continue
			}
			//else string
			val = str
			switch val.(type){

			}
		}
	}
}

func TestBenchAtoi2(t *testing.T) {
	var val interface {}
	var ok bool
	for i:=0;i<times;i++ {
		for j:=0;j<len(str);j++ {
			switch str[j][0] {
			case 't','T','f','F':
				val,ok = evalBool(str[j])
				if ok {
					continue
				} else {
					val = str
				}
			case '+', '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				val, ok = evalInt(str[j])
				if ok {
					continue
				}
				val, ok = evalFloat(str[j])
				if ok {
					continue
				}
				val, ok = evalDatetime(str[j])
				if ok {
					continue
				}
				val = str
			default:
				val = str
			}
			switch val.(type){

			}

		}
	}
}*/

