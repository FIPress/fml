package fipml

import (
	"testing"
	"time"
)

func TestEvalBool(t *testing.T) {
	input := "true"
	val,ok := evalBool(input)
	if !val || !ok {
		t.Error("Should ok, the val should be true,val:",val,",ok:",ok)
	}

	input = "False"
	val,ok = evalBool(input)
	if val || !ok {
		t.Error("Should ok, the val should be false,val:",val,",ok:",ok)
	}

	input = "T"
	val,ok = evalBool(input)
	if !val || !ok {
		t.Error("Should ok, the val should be true,val:",val,",ok:",ok)
	}

	input = "Fals"
	val,ok = evalBool(input)
	if ok {
		t.Error("Should NOT be ok, ok:",ok)
	}
}

func TestEvalInt(t *testing.T) {
	input := "123"
	val,ok := evalInt(input)
	if val != 123 || !ok {
		t.Error("Should be ok, the val should be 123,val:",val,",ok:",ok)
	}

	input = "-321"
	val,ok = evalInt(input)
	if val != -321 || !ok {
		t.Error("Should be ok, the val should be -321,val:",val,",ok:",ok)
	}

	input = "123.32"
	val,ok = evalInt(input)
	if ok {
		t.Error("Should NOT be ok,val:",val,",ok:",ok)
	}

	input = "Fals"
	val,ok = evalInt(input)
	if ok {
		t.Error("Should NOT be ok,val:",val,",ok:",ok)
	}
}

func TestEvalFloat(t *testing.T) {
	input := "123"
	val,ok := evalFloat(input)
	if val != 123.0 || !ok {
		t.Error("Should be ok, the val should be 123,val:",val,",ok:",ok)
	}

	input = "-123.32"
	val,ok = evalFloat(input)
	if val != -123.32 || !ok {
		t.Error("Should be ok, the val should be -123.32,val:",val,",ok:",ok)
	}

	input = "12a"
	val,ok = evalFloat(input)
	if ok {
		t.Error("Should NOT be ok,val:",val,",ok:",ok)
	}
}

func TestEvalDatetime(t *testing.T) {
	input := "2014-07-06"
	val,ok := evalDatetime(input)
	tm,_ := time.Parse("2006-01-02","2014-07-06")
	if val != tm || !ok {
		t.Error("should parse short date")
	}

	input = "2014-07-06 12:03:31"
	val,ok = evalDatetime(input)
	tm,_ = time.Parse("2006-01-02 15:04:05","2014-07-06 12:03:31")
	if val != tm || !ok {
		t.Error("should parse long datetime,",val,tm)
	}

	input = "2014-07-06T12:03:31Z"
	val,ok = evalDatetime(input)
	tm,_ = time.Parse(time.RFC3339,"2014-07-06T12:03:31Z")
	if val != tm || !ok {
		t.Error("should parse RFC3399 datetime")
	}

	input = "2014-07-06 2:03:31"
	val,ok = evalDatetime(input)
	if ok {
		t.Error("should NOT parse datetime")
	}
}

func TestEval(t *testing.T) {
	input := "abc"
	val := eval(input)
	if val != "abc" {
		t.Error("Should get string,val")
	}

	input = "123"
	val = eval(input)
	if val != 123 {
		t.Error("Should get int")
	}

	input = "true"
	val = eval(input)
	if val != true {
		t.Error("Should get bool")
	}

	input = "F"
	val = eval(input)
	if val != false {
		t.Error("Should get bool")
	}

	input = "12.345"
	val = eval(input)
	if val != 12.345 {
		t.Error("Should get float")
	}

	input = "2013-12-03 11:08:54"
	val = eval(input)
	tm,_ := time.Parse("2006-01-02 15:04:05",input)
	if val != tm {
		t.Error("Should get datetime,",val)
	}

	input = `123a\n`
	val = eval(input)
	if val != "123a\n" {
		t.Error("Should get string,",val)
	}
}
