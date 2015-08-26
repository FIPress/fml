package fipml

import (
	. "fiputil"
	"strings"
	"unicode"
)

//Skip spaces and comments
func skipLeft(input []byte) (skip int) {
	i := 0
	for i < len(input) {
		if input[i] == '#' {
			i += skipComments(input[i:])
		} else if IsSpaceOrLineEnd(input[i]) {
			i++
		} else {
			return i
		}
	}
	return i
}

//Skip comments or spaces until line end
func skipRest(input []byte) (skip int) {
	i := 0
	for i < len(input) {
		if input[i] == '#' {
			i += skipComments(input[i:])
		} else if IsSpace(input[i]) {
			i++
		} else if IsLineEnd(input[i]) {
			return i+1
		} else {
			return i
		}
	}
	return i
}

func skipComments(input []byte) int {
	return SkipUntilFunc(input, IsLineEnd, true)
}
//A node end by blank lines.
//A blank line means it contains only spaces or comments
func isBlankLine(input []byte) (bool, int) {
	i:=0
	for ;i < len(input); i++ {
		switch input[i] {
		case '\n', '\r', '\f':
			return true, i + 1
		case ' ', '\t':
		case '#':
			return true, i+skipComments(input[i:])
		default:
			return false, 0
		}
	}
	return true, i
}

//A block end by blank lines, or next line starts a list block
func isKeyValueBlockEnd(input []byte) (bool, int) {
	isList,_ := isListPrefix(input)
	if isList {
		return true, 0
	}

	return isBlockEnd(input)
}

func isBlockEnd(input []byte) (bool, int) {
	i:=0
	for ;i < len(input); i++ {
		switch input[i] {
		case '\n', '\r', '\f':
			return true, i + 1
		case ' ', '\t':
		case '#':
			return true, i+skipComments(input[i:])
		default:
			return false, 0
		}
	}
	return true, i
}

func isListPrefix(input []byte) (bool,int) {
	idx := SkipSpace(input)
	if idx + 2 < len(input) &&
		input[idx] == '-' &&
			IsSpace(input[idx+1]) {
		return true,idx+2
	}
	return false,0
}

//value end at line end or comment
func skipUntilValueEnd(input []byte) int {
	i:=0
	for ;i<len(input);i++ {
		if input[i] == '#' && IsSpace(input[i-1]) {
			return i-1
		}
		if IsLineEnd(input[i]) {
			return i
		}
	}
	return i
}

//get value in string
func getRawValue(input []byte) (string,int) {
	start := SkipSpace(input)
	end := start + skipUntilValueEnd(input[start:])
	val := string(input[start:end])
	val = strings.TrimRightFunc(val,unicode.IsSpace)

	idx := end + skipRest(input[end:])
	return val, idx
}

