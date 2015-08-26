package fipml

import "testing"

func TestIsListPrefix(t *testing.T) {
	input := "  - a list"
	isList,idx := isListPrefix([]byte(input))
	if !isList || idx != 4 {
		t.Log("Should be a list, and idx should be 4, idx:",idx)
		t.FailNow()
	}

	input = "-a list"
	isList,idx = isListPrefix([]byte(input))
	if isList || idx != 0 {
		t.Error("Should NOT be a list, and idx should be 0, idx:",idx)
	}
}

func TestSkipComments(t *testing.T) {
	input := "#comments"
	idx := skipComments([]byte(input))
	if idx != len(input) {
		t.Error("should skip whole comments,idx:",idx,"len:",len(input))
	}

	input = "#comments\n"
	idx = skipComments([]byte(input))
	if idx != len(input) {
		t.Error("should skip whole comments,idx:",idx,"len:",len(input))
	}
}

func TestIsBlockEnd(t *testing.T) {
	input := "  # lines contains only space and comments considered a block end"
	isEnd,idx := isBlockEnd([]byte(input))
	if !isEnd || idx != len(input) {
		t.Error("Should be a block end, and idx should be the length, idx:",idx,"len:",len(input))
	}

	input = "  "
	isEnd,idx = isBlockEnd([]byte(input))
	if !isEnd || idx != 2 {
		t.Error("Should be a block end, and idx should be 2, idx:",idx)
	}

	input = "  not block end"
	isEnd,idx = isBlockEnd([]byte(input))
	if isEnd || idx != 0 {
		t.Error("Should NOT be a block end, and idx should be 0, idx:",idx)
	}
}

func TestKeyValueBlockEnd(t *testing.T) {
	input := "  # lines contains only space and comments considered a block end"
	isEnd,idx := isKeyValueBlockEnd([]byte(input))
	if !isEnd || idx != len(input) {
		t.Error("Should be a block end, and idx should be the length, idx:",idx,"len:",len(input))
	}

	input = "  - a line starts a list considered a block end"
	isEnd,idx = isKeyValueBlockEnd([]byte(input))
	if !isEnd || idx != 0 {
		t.Error("Should be a block end, and idx should be 0, idx:",idx)
	}

	input = "  not block end"
	isEnd,idx = isKeyValueBlockEnd([]byte(input))
	if isEnd || idx != 0 {
		t.Error("Should NOT be a block end, and idx should be 0, idx:",idx)
	}
}

func TestIsNodeEnd(t *testing.T) {
	input := "  # lines contains only space and comments considered a node end"
	isEnd,idx := isBlankLine([]byte(input))
	if !isEnd || idx != len(input) {
		t.Error("Should be a node end, and idx should be the length, idx:",idx,"len:",len(input))
	}

	input = "  not node end"
	isEnd,idx = isBlankLine([]byte(input))
	if isEnd || idx != 0 {
		t.Error("Should NOT be a node end, and idx should be 0, idx:",idx)
	}
}

func TestSkipRest(t *testing.T) {
	input := "  # Skip rest will skip spaces and comments, and one line end\n"
	idx := skipRest([]byte(input))
	if idx != len(input) {
		t.Error("Should only skip one line end, idx:",idx,"len:",len(input))
	}

	input = "  \n\n"
	idx = skipRest([]byte(input))
	if idx != len(input) - 1 {
		t.Error("Should only skip one line end, idx:",idx,"len:",len(input))

	}
}

func TestSkipLeft(t *testing.T) {
	input := "  # Skip rest will skip spaces and comments, and one line end\n\n  real"
	idx := skipLeft([]byte(input))
	if idx != len(input) - 4 {
		t.Error("Should skip all spaces and line ends and comments, idx:",idx,"len:",len(input))
	}
}

func TestSkipUntilValueEnd(t *testing.T) {
	input := "abc # comments"
	idx := skipUntilValueEnd([]byte(input))
	if idx != 3 {
		t.Error("Should skip 3,idx:",idx)
	}

	input = "abc \n"
	idx = skipUntilValueEnd([]byte(input))
	if idx != 4 {
		t.Error("Should skip 4,idx:",idx)
	}

	input = "abc"
	idx = skipUntilValueEnd([]byte(input))
	if idx != 3 {
		t.Error("Should skip 3,idx:",idx)
	}
}

func TestGetRawValue(t *testing.T) {
	input := "abc # comment"
	val, idx := getRawValue([]byte(input))
	if val != "abc" || idx != len(input) {
		t.Error("Should get abc,val:",val,"idx:",idx)
	}

	input = "abc  \n"
	val, idx = getRawValue([]byte(input))
	if val != "abc" || idx != len(input) {
		t.Error("Should get abc,val:",val,"idx:",idx)
	}
}


