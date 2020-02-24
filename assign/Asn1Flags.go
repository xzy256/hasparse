package assign

import (
	"github.com/xzy256/hasparse/unmarshal"
	"github.com/xzy256/hasparse/utils"
	"log"
)

/**
KrbFlags   ::= BIT STRING (SIZE (32..MAX))
-- minimum number of bits shall be sent,
-- but no fewer than 32
*/
type Asn1Flags struct {
	TagNo      int // 0x03
	Padding    int
	ValueBytes []byte
	Flags      int

	Name       string
}

func (this *Asn1Flags) SetName(name string) {
	this.Name = name
}

func (this *Asn1Flags) Init() {
}

func (this *Asn1Flags) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	body := parseResult
	for body.GetIndex() != 0 {
		body = body.Children[0]
	}
	remainingBytes := body.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ToValue(remainingBytes)
	}
}

func (this *Asn1Flags) ToValue(val []byte) {
	if len(val) < 1 {
		log.Fatal("Bad stream, zero bytes found for bitstring")
	}
	paddingBits := val[0]
	validatePaddingBits(int(paddingBits))
	this.Padding = int(paddingBits)

	newBytes := make([]byte, len(val)-1)
	if len(val) > 1 {
		utils.ArrayCopy(val, 1, newBytes, 0, len(val)-1)
	}
	this.ValueBytes = newBytes

	if this.Padding != 0 || len(this.ValueBytes) != 4 {
		log.Fatal("Bad bitstring decoded as invalid krb flags")
	}

	this.Value2Flags()
}

func (this *Asn1Flags) Value2Flags() {
	valueBytes := this.ValueBytes
	this.Flags = (int(valueBytes[0]&0xFF) << 24) | (int(valueBytes[1]&0xFF) << 16) | (int(valueBytes[2]&0xFF) << 8) | int(0xFF&valueBytes[3])
}

func validatePaddingBits(paddingBits int) {
	if paddingBits < 0 || paddingBits > 7 {
		log.Fatal("Bad padding number: ", paddingBits, ", should be in [0, 7]")
	}
}
