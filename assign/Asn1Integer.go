package assign

import (
	"hasparse/unmarshal"
	"hasparse/utils"
)

type Asn1Integer struct {
	ValueBytes []byte // multi bytes for int32
	fileInfo   []int
	position   []byte
}

func (this *Asn1Integer) Init() {
	this.position = []byte{0, 0, 0}
	this.fileInfo = []int{0}
}

func (this *Asn1Integer) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	remainingBytes := parseResult.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1Integer) Value() int32 {
	return utils.BytesToInt(this.ValueBytes)
}
