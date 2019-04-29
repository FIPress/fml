package fml

import "testing"

func TestExtractNodeName(t *testing.T) {
	input := "[node]"
	prefixes, name, idx := extractNodeName([]byte(input))
	if prefixes != nil || name != "node" || idx != len(input) {
		t.Error("Should get node name. name:", name, "idx:", idx, "len:", len(input))
	}

	input = "[a.node]  #comment"
	prefixes, name, idx = extractNodeName([]byte(input))
	if len(prefixes) != 1 || prefixes[0] != "a" || name != "node" || idx != len(input) {
		t.Error("Should get node name. prefixes:", prefixes, "name:", name, "idx:", idx, "len:", len(input))
	}
}

func TestExtractKey(t *testing.T) {
	input := "key:"
	name, idx := extractKey([]byte(input))
	if name != "key" || idx != len(input) {
		t.Error("Should get key name. name:", name, "idx:", idx, "len:", len(input))
	}

	input = "key:  #comment"
	name, idx = extractKey([]byte(input))
	if name != "key" || idx != len(input) {
		t.Error("Should get key name. name:", name, "idx:", idx, "len:", len(input))
	}
}

func TestExtractLiteral(t *testing.T) {
	input := "`" + `literal \n` + "`"
	val, idx := extractLiteral([]byte(input))
	if val != "literal \\n" || idx != len(input) {
		t.Error("Should get literal. val:", val, "idx:", idx, "len:", len(input))
	}

	input = "`literal  # not comment`"
	val, idx = extractLiteral([]byte(input))
	if val != "literal  # not comment" || idx != len(input) {
		t.Error("Should get literal. val:", val, "idx:", idx, "len:", len(input))
	}

	input = "```" + `literal line 1
line2 \n # not comment` + "```"
	val, idx = extractLiteral([]byte(input))
	if val != "literal line 1\nline2 \\n # not comment" || idx != len(input) {
		t.Error("Should get literal. val:", val, "idx:", idx, "len:", len(input))
	}

}

func TestGetArrayItems(t *testing.T) {
	input := "[1,2,3]"
	items, idx := getArrayItems([]byte(input))
	if items[0] != "1" || items[2] != "3" || idx != len(input) {
		t.Error("Should get array itmes. items:", items, "idx:", idx, "len:", len(input))
	}

}

func TestExtractArray(t *testing.T) {
	input := "[t,f,t]"
	array, idx := extractArray([]byte(input))
	if idx != len(input) {
		t.Error("should skip all")
	}
	switch a := array.(type) {
	case []bool:
		if len(a) != 3 || a[0] != true || a[1] != false {
			t.Error("Should get array. array:", array)
		}
	default:
		t.Error("Should get bool array")
	}

	input = `[zh,
	en
	]`
	array, idx = extractArray([]byte(input))
	if idx != len(input) {
		t.Error("should skip all")
	}
	switch s := array.(type) {
	case []string:
		if len(s) != 2 || s[0] != "zh" || s[1] != "en" {
			t.Error("Should get array. array:", array)
		}
	default:
		t.Error("should get string array")
	}

}

func TestExtractValue(t *testing.T) {
	input := "abc # comment\n"
	val, idx := extractValue([]byte(input))
	if val != "abc" || idx != len(input) {
		t.Error("Should get value. val:", val, "idx:", idx, "len:", len(input))
	}

	input = "abc\n"
	val, idx = extractValue([]byte(input))
	if val != "abc" || idx != len(input) {
		t.Error("Should get value. val:", val, "idx:", idx, "len:", len(input))
	}

	input = "123  #comment"
	val, idx = extractValue([]byte(input))
	if val != 123 || idx != len(input) {
		t.Error("Should get value. val:", val, "idx:", idx, "len:", len(input))
	}
}

func TestExtractKeyValue(t *testing.T) {
	input := "name:Tony"
	doc := NewFml()
	idx := extractKeyValue([]byte(input), doc)
	if doc.GetString("name") != "Tony" || idx != len(input) {
		t.Error("Should get key value,idx:", idx, "len:", len(input))
	}

}

func TestExtractNode(t *testing.T) {
	input := `[database]
	host:local
	port:1433`
	doc := NewFml()
	idx := extractNode([]byte(input), doc)
	db, err := doc.GetNode("database")
	if err != nil {
		t.Error("Extract node error:", err)
	}
	host := db.GetString("host")
	port := db.GetInt("port")
	if host != "local" || port != 1433 {
		t.Error("Should get host and port, host:", host, "port:", port)
	}

	//IterateTomlDoc(doc)
	if idx != len(input) {
		t.Error("Extract node, idx error,idx:", idx)
	}
}

func TestExtractNodeOfList(t *testing.T) {
	input := `[staff]
	- name: Abby
	age: 12
	- name: Tony
	age: 13`
	doc := NewFml()
	idx := extractNode([]byte(input), doc)
	//IterateFimlDoc(doc)
	if idx != len(input) {
		t.Error("Extract list of node idx error,idx,len:", idx, len(input))
	}

	staff, err := doc.GetNodeList("staff")
	if err != nil {
		t.Error("Should get node list, err:", err)
	}
	if len(staff) != 2 || staff[0].GetString("name") != "Abby" ||
		staff[1].GetInt("age") != 13 {
		t.Error("Get node list error,")
		IterateFimlDoc(doc)
	}
}
