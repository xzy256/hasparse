package main

import (
	"bytes"
	"fmt"
)

func KrbDecodeMessage(buffer bytes.Buffer)  {
	parseResult := Asn1ParserBuffer(buffer)
	tag := parseResult.Tag
	msgType := KrbMessageTypeFromValue(tag.TagNo)
	if msgType  == int(AS_REP){
		fmt.Println(parseResult)
	}
}
