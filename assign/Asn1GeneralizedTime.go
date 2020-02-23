package assign

import (
	"fmt"
	"hasparse/unmarshal"
	"log"
	"strings"
	"time"
)

type Asn1GeneralizedTime struct {
	TagNo      int // 24
	ValueBytes []byte

	Name string
}

func (this *Asn1GeneralizedTime) SetName(name string) {
	this.Name = name
}

func (this *Asn1GeneralizedTime) Init() {

}

func (this *Asn1GeneralizedTime) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	body := parseResult
	for body.GetIndex() != 0 {
		body = body.Children[0]
	}
	remainingBytes := body.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1GeneralizedTime) ParseGeneralizedTime(bytes []byte) (ret time.Time, err error) {
	const formatStr = "20060102150405Z0700"
	s := string(bytes)

	if ret, err = time.Parse(formatStr, s); err != nil {
		return
	}

	if serialized := ret.Format(formatStr); serialized != s {
		err = fmt.Errorf("asn1: time did not serialize back to the original value and may be invalid: given %q, but serialized as %q", s, serialized)
	}

	return
}

func getMillSeconds(dateStr string) string {
	millDigits := []byte{48, 48, 48} // 000
	iPos := strings.Index(dateStr, ".")
	if iPos > 0 { // uncheck
		if iPos != 14 {
			log.Fatal("Bad generalized time string, with improper milli seconds.", dateStr)
		}

		j := 0
		for i := 15; i < len(dateStr) && j < len(millDigits); i++ {
			chr := dateStr[i]
			if 0 <= chr && chr <= 9 {
				millDigits[j] = chr
				j++
			} else {
				break
			}
		}
	}

	return string(millDigits)
}

func getTimeZonePart(dateStr string) string {
	iPos := strings.Index(dateStr, "+")
	if iPos == -1 {
		iPos = strings.Index(dateStr, "-")
	}
	if iPos > 0 && iPos != len(dateStr)-5 {
		log.Fatal("Bad generalized time string, with improper timezone part " + dateStr)
	}

	if iPos > 0 {
		return dateStr[iPos:]
	}
	return ""
}
