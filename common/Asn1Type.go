package common

import (
	"bytes"
)

type EncodingType int

const (
	BER EncodingType = 1
)

type Asn1Type interface {
	UsePrimitive(isPrimitive bool)
	IsPrimitive() bool
	UseDefinitiveLength(isDefinitiveLength bool)
	IsDefinitiveLength() bool
	UseImplicit(isImplicit bool)
	IsImplicit() bool
	UseBER()
	IsBER() bool
	DcodeBytes(context []byte)
	DcodeBuffer(buffer bytes.Buffer)
}
