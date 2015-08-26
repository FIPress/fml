package fipml

import (
	"testing"
	"fiputil"
)

func TestParse(t *testing.T) {
/*

*/

	input := `title: Scala Basic
author: Chunni Deng
toclevel: 2
defaultLang: en
desc: This is a book about basic knowledge of Scala.
	#cover = "static/img/cover.jpg"
sequenceDigitsToRemove: 2
forceProcess: true
#lang: [zh,en]
	#[lang]
	#zh: "简体中文"
	#en: "English"

content: [
$cover,
content/$lang/preface,
$toc,
content/$lang/*
]

[item]
- name:abc
  age: 123
- name: abcd
  age: 45

  port: 7001

[fiplog]
level: Debug
file: rotor.log
pattern: %date [%level] <%file> %msg

[fiplog.item]
- name: ab
 path: cd
- name: ef
 path: gh

 [lang]
- code: zh
  name: Scala基础教程
- code: en
  name: Scala Basic
`
	doc,err := ParseString(input)
	if err != nil {
		t.Error("Should be ok, err",err)
	}
	IterateFimlDoc(doc)
}

/*func TestA(t *testing.T) {
	fml := NewFml()
	fml.SetValue("img/class-hierarchy.png",fiputil.MinTime)
	tm, err:= fml.GetDatetimeEx("img/class-hierarchy.png")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(tm)
	}
}*/
