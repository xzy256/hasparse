package main

import (
	"bytes"
	"container/list"
)

type Asn1Container struct {
	Asn1PR   *Asn1ParseResult
	Children *list.List
}

func NewAsn1Container(header Asn1Header, bodyStart int, buffer bytes.Buffer) *Asn1Container {
	return &Asn1Container{
		Asn1PR:   NewAsn1ParseResult(header, bodyStart, buffer),
		Children: list.New(),
	}
}



