package fml

func Unmarshal(data []byte, v interface{}) (err error) {
	node, err := Parse(data)
	if err != nil {
		return
	}

	return getStruct(node, v)
}

func UnmarshalFile(path string, v interface{}) (err error) {
	node, err := Load(path)
	if err != nil {
		return
	}

	return getStruct(node, v)
}

func UnmarshalString(input string, v interface{}) (err error) {
	return Unmarshal([]byte(input), v)
}
