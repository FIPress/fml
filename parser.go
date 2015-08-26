package fipml

import "io/ioutil"

import (
	"log"
	"errors"
)

func Load(path string) (doc *FML, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	doc, err = Parse(bytes)

	return
}

func ParseString(input string) (doc *FML, err error) {
	return Parse([]byte(input))
}

func Parse(input []byte) (doc *FML, err error) {
	defer func() {
		if r:=recover();r!=nil {
			log.Println(r)
			err = errors.New("parse fipml failed")
		}
	}()

	doc = NewFml()
	idx,delta := 0,0

	for idx < len(input) {
		idx += skipLeft(input[idx:])
		if idx >= len(input) {
			return
		}

		if input[idx] == '[' {
			delta = extractNode(input[idx:],doc)
		} else {
			//delta = extractBlock(input[idx:],doc)
			delta = extractKeyValue(input[idx:],doc)
		}
		idx += delta
		//idx += extractBlock(input[idx:],doc)
	}

	return
}
