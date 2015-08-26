package fipml

import (
	. "fiputil"
	"strings"
	"time"
//	"log"
	"log"
)

const (
	invalidNodeName = "invalid node name"
	invalidKeyName = "invalid key name"
	duplicateNode = "duplicate node: "
	duplicateKey = "duplicate key: "
	unsupportedValue = "unsurported value for key: "
	multipleLineClosingErr = "multiple line literal not close properly"
	literalClosingErr = "literal not close properly"
	invalidArray = "invalid aray"
)

var (
	multipleLiteral = []byte {'`','`','`'}
)

func extractNode(input []byte, doc *FML) (idx int) {
	prefixes,name, idx := extractNodeName(input)
	pDoc := getPNodeByName(prefixes,doc)
	if pDoc.dict[name] != nil {
		panic(duplicateNode+name)
	}

	//array, err :=doc.GetTableArray(name)
	isList,delta := isListPrefix(input[idx:])
	if isList {
		list := make([]*FML,0)
		idx += delta
		for idx < len(input) {
			subDoc := NewFml()
			//log.Println("extract node list:",string(input[idx:]))
			idx += extractKeyValueBlock(input[idx:],subDoc)
			//log.Println("subdoc,name:",subDoc.GetString("name",""))
			list = append(list,subDoc)
			goon,delta := isListPrefix(input[idx:])
			idx += delta
			//log.Println("extract node list,is end:",string(input[idx:]))
			//log.Printf("extract node list,idx=%d,c=%c,goon=%v,delta=%d\n",idx,input[idx],goon,delta)
			if !goon {
				break
			}
		}
		//log.Print("extractNode, name:",name,",len:",len(list))
		if len(list) != 0 {
			pDoc.dict[name] = list
		}
	} else {
		subDoc := NewFml()
		idx += extractKeyValueBlock(input[idx:],subDoc)
		pDoc.dict[name] = subDoc
	}
	return
}

func getPNodeByName(names []string, doc *FML) *FML {
	/*l := len(names)
	if l == 0 {
		return nil
	} else if l == 1 {
		return doc
	}*/
	pNode := doc
	for i:=0;i<len(names);i++ {
		if pNode == nil {
			return nil
		}
		switch p := pNode.dict[names[i]].(type){
		case *FML:
			pNode = p
		case nil:
			pNode = NewFml()
		default:
			panic(duplicateKey+names[i])
		}
	}
	return pNode
}

func extractKeyValueBlock(input []byte,doc *FML) (idx int) {
	for idx < len(input) {
		idx += extractKeyValue(input[idx:], doc)
		end,delta := isKeyValueBlockEnd(input[idx:])
		if end {
			idx += delta
			return
		}
	}
	return
}

/*func extractBlock(input []byte, doc *FML) (idx int) {
	for idx < len(input) {
		idx += skipLeft(input[idx:])
		if idx >= len(input) {
			return
		}

		if input[idx] == '[' {
			idx += extractNode(input[idx:],doc)
		} else {
			//delta = extractBlock(input[idx:],doc)
			idx += extractKeyValue(input[idx:],doc)
		}

		//log.Printf("extractBlock,idx=%d,c=%c\n",idx,input[idx])
		//idx += extractKeyValue(input[idx:], doc)
		end,delta := isBlockEnd(input[idx:])
		//log.Printf("extractBlock,idx=%d,c=%c,end=%v,delta=%d\n",idx,input[idx],end,delta)
		if end {
			idx += delta
			return
		}
	}
	return
}*/

func extractKeyValue(input []byte, doc *FML) (idx int) {
	idx = SkipSpace(input)

	key, delta := extractKey(input[idx:])
	//log.Println("extractKeyValue,key:",key,"delta:",delta)
	idx += delta

	if doc.dict[key] != nil {
		panic(duplicateKey+key)
	}

	idx += skipRest(input[idx:])
	val, delta := extractValue(input[idx:])
	//log.Println("extractKeyValue,value:",val,",idx:",idx,",delta 1:",delta)
	//log.Printf("c1=%c\n",input[idx])
	idx += delta
	//log.Printf("c2=%c,c2+1=%c\n",input[idx],input[idx+1])
	if val == nil {
		panic(unsupportedValue+key)
	}

	doc.dict[key] = val

	//idx += skipRest(input[idx:])
	//log.Println("extractKeyValue,value:",val,",idx:",idx,",delta 2:",delta)
	//log.Printf("c2=%c,c2+1=%c\n",input[idx],input[idx+1])
	return
}

func extractKey(input []byte) (key string, idx int) {
	delta,found := SkipUntilOrStopAtLineEnd(input,':')
	if !found || delta == 0 {
		log.Println("extract key:",string(input))
		log.Println("found:",found,"delta:",delta)
		panic(invalidKeyName)
	}
	idx = delta
	key = string(input[:idx])
	key = strings.TrimSpace(key)
	idx++ //:
	idx += skipRest(input[idx:])
	return
}

func extractValue(input []byte) (val interface{}, idx int) {
	switch input[0] {
	case '`':
		val,idx = extractLiteral(input)
	case '[':
		val, idx = extractArray(input)
	default:
		var raw string
		raw,idx = getRawValue(input)
		val = eval(raw)
	}
	return
}

func extractNodeName(input []byte) (prefixes []string, name string,idx int) {
	delta,found := SkipUntilOrStopAtLineEnd(input[1:],']')
	if !found || delta == 0 {
		panic(invalidNodeName)
	}
	idx = delta + 1 //plus '['
	str := string(input[1:idx])
	idx++ //plus ']'
	str = strings.TrimSpace(str)
	names := strings.Split(str,".")
	l := len(names)
	if l == 0 {
		return
	} else {
		name = names[l-1]
		if l > 1 {
			prefixes = names[:l-1]
		}
	}
	idx += skipRest(input[idx:])
	return
}

func extractLiteral(input []byte) (val string,idx int) {
	if len(input)>6 && SliceEquals(input[:3],multipleLiteral) {
		delta,found := SkipUntilArray(input[3:],multipleLiteral)
		if !found {
			panic(multipleLineClosingErr)
		}
		end := 3 + delta
		val = string(input[3:end])
		idx = end + 3
	} else {
		delta,found := SkipUntilOrStopAtLineEnd(input[1:],'`')
		if !found {
			panic(literalClosingErr)
		}
		idx = delta + 1
		val = string(input[1:idx])
		idx++
	}
	return
}

func extractArray(input []byte) (val interface{}, idx int) {
	//i := 1 + skipLeft(input[1:])
	items,idx := getArrayItems(input)
	length := len(items)
	if length == 0 {
		return
	}

	val0 := eval(items[0])

	switch v0 := val0.(type) {
	case string:
		arr := make([]string, length, length)
		arr[0] = v0
		for i := 1; i < length; i++ {
			arr[i] = evalString(items[i])
		}
		val = arr
	case int:
		arr := make([]int, length, length)
		arr[0] = v0
		for i := 1; i < length; i++ {
			vi,ok := evalInt(items[i])
			if !ok {
				panic(invalidArray)
			}
			arr[i] = vi
		}
		val = arr
	case float64:
		arr := make([]float64, length, length)
		for i := 1; i < length; i++ {
			vi,ok := evalFloat(items[i])
			if !ok {
				panic(invalidArray)
			}
			arr[i] = vi
		}
		val = arr
	case bool:
		arr := make([]bool, length, length)
		arr[0] = v0
		for i := 1; i < length; i++ {
			vi,ok := evalBool(items[i])
			if !ok {
				panic(invalidArray)
			}
			arr[i] = vi
		}
		val = arr
	case time.Time:
		arr := make([]time.Time, length, length)
		arr[0] = v0
		for i := 1; i < length; i++ {
			vi,ok := evalDatetime(items[i])
			if !ok {
				panic(invalidArray)
			}
			arr[i] = vi
		}
		val = arr
	}
	return
}

func getArrayItems(input []byte) (items []string,idx int) {
	from := -1
	add := func(to int) {
		val := string(input[from:to])
		val = strings.TrimSpace(val)
		items = append(items,val)
	}
	//log.Printf("getArrayItems,c0=%c\n",input[0])
	for i:=1; i<len(input);i++ {
		switch input[i] {
		case ' ','\t','\n','\r','\f':
			//do nothing
		case ',':
			if from == -1 {
				panic(invalidArray)
			}
			add(i)
			from = -1
		case ']':
			if from != -1 {
				add(i)
			}
			return items,i+1
		default:
			if from == -1 {
				from = i
			}
		}
	}
	panic(invalidArray)
}

