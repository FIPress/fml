package fml

import (
	"testing"
)

type Student struct {
	Name string
	Age  string
}

func TestUnmarshalString(t *testing.T) {
	input := `name=Jimmy
Age=15`
	s := new(Student)
	UnmarshalString(input, s)
	t.Log(s.Name)
	t.Log(s.Age)
}
