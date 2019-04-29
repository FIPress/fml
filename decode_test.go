package fml

import (
	"testing"
)

type Student struct {
	Name string
	Age  int
}

func TestUnmarshalString(t *testing.T) {
	input := `name="Jimmy"
Age=15`
	s := new(Student)
	UnmarshalString(input, s)
	t.Log(s.Name)
}
