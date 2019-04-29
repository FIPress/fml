package fml

import (
	"bufio"
	"bytes"
	"testing"
)

func TestGet(t *testing.T) {
	fml := NewFml()
	fml.SetValue("name", "Tony")
	fml.SetValue("age", 13)
	fml.SetValue("score", 89.5)
	fml.SetValue("passed", true)

	t.Log(fml.GetString("name"))
	t.Log(fml.GetString("age"))
	t.Log(fml.GetString("score"))
	t.Log(fml.GetString("passed"))
}

func TestFML_GetStruct(t *testing.T) {
	type Stu struct {
		Name string
		Age  int
	}

	fml := NewFml()
	sub := NewFml()
	sub.SetValue("Name", "Jimmy")
	sub.SetValue("Age", 12)
	fml.SetValue("Stu", sub)

	s := new(Stu)
	fml.GetStruct("Stu", s)
	t.Log(s.Name)
}

func TestWrapVal(t *testing.T) {
	fml := NewFml()
	fml.SetValue("name", "Tony")
	fml.SetValue("age", 13)
	fml.SetValue("score", 89.5)
	fml.SetValue("passed", true)

	//sw := &fiputil.StringWriter{}
	buf := &bytes.Buffer{}
	bw := bufio.NewWriter(buf)

	fml.WriteTo(bw)
	bw.Flush()
	output := buf.String()

	/*shouldBe := `name : Tony
	age : 13
	score : 89.5
	passed : true
	`
		if output != shouldBe {
			t.Error("Write fml error, output:",output)
		}*/
	t.Log(output)
}

func TestWriterTo(t *testing.T) {
	fml := NewFml()
	fml.SetValue("title", "Test")

	sub1 := NewFml()
	sub1.SetValue("host", "local")
	sub1.SetValue("port", 1433)
	sub11 := NewFml()
	sub11.SetValue("loc", "a")
	sub11.SetValue("time", 1)
	sub12 := NewFml()
	sub12.SetValue("loc", "b")
	sub12.SetValue("time", 2)
	sub1.SetValue("bk", []*FML{sub11, sub12})

	fml.SetValue("database", sub1)

	sub2 := NewFml()
	sub2.SetValue("id", 1)
	sub2.SetValue("name", "phone")

	sub3 := NewFml()
	sub3.SetValue("id", 2)
	sub3.SetValue("name", "tv")

	fml.SetValue("items", []*FML{sub2, sub3})

	//sw := &fiputil.StringWriter{}
	buf := &bytes.Buffer{}
	bw := bufio.NewWriter(buf)

	fml.WriteTo(bw)
	bw.Flush()
	output := buf.String()

	/*shouldBe := `title:Test

			[database]
			host:local
			port:1433

			[database.bk]
			- loc:a
			time:1
			- loc:b
			time:2



			[items]
			- name:phone
			id:1
			- id:2
			name:tv
	`
		if output != shouldBe {
			t.Error("Write fml error, output:",output)
		}*/
	t.Log(output)
}
