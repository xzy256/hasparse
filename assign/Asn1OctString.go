package assign

import "hasparse/unmarshal"

type Asn1OctString struct {
	TagNo int

	ValueBytes []byte
	Name       string
}

func (this *Asn1OctString) Init() {
}

func (this *Asn1OctString) SetName(name string) {
	this.Name = name
}

func (this *Asn1OctString) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	body := parseResult
	for body.GetIndex() == 1 {
		body = body.Children[0]
	}
	if body.GetIndex() == 0 {
		remainingBytes := body.GetBodyBuffer().Bytes()
		if len(remainingBytes) > 0 {
			this.ValueBytes = remainingBytes
		}
	} else if body.GetIndex() > 1 {
		for i := 0; i < body.GetIndex(); i++ {
			remainingBytes := body.Children[i].GetBodyBuffer().Bytes()
			if len(remainingBytes) > 0 {
				if this.ValueBytes == nil {
					this.ValueBytes = remainingBytes
				} else {
					tmpStr := []byte(string(this.ValueBytes) + "/" + string(remainingBytes))
					this.ValueBytes = tmpStr
				}
			}
		}
	}
}
