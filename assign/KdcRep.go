package assign

import (
	"hasparse/unmarshal"
	"log"
)

type KdcRepField int

const (
	PVNO     KdcRepField = 0
	MSG_TYPE KdcRepField = 1
	PADATA   KdcRepField = 2
	CREALM   KdcRepField = 3
	CNAME    KdcRepField = 4
	TICKET   KdcRepField = 5
	ENC_PART KdcRepField = 6
)

type KdcRep struct {
	Pvno       int
	MsgType    int
	Padata     *PAData //optional
	Crealm     string
	Cname      *PrincipalName
	Ticket     *KdcTicket
	EncData    *EncryptedData
	fieldInfos []int
	position   []byte // same as 0x532, 5-Ticket 3-Ticket.enc_part 2-Ticket.enc_part.cipher, init 0x000
}

func (this *KdcRep) Init() {
	this.fieldInfos = []int{0, 1, 2, 3, 4, 5, 6}
	this.position = []byte{255, 255, 255}
}

func (this *KdcRep) Decode(parseResult *unmarshal.Asn1ParseResult) {
	if parseResult.Tag.IsNested() {
		this.TaggingDecode(parseResult)
	} else {
		this.UniversalDecode(parseResult)
	}

}

func (this *KdcRep) UniversalDecode(parseResult *unmarshal.Asn1ParseResult) {
	kdcPos := int(this.position[0])
	kdcDirectChildPos := int(this.position[1])
	kdcGrandsonPos := int(this.position[2])
	switch kdcPos {
	case int(PVNO):
		interger := &Asn1Integer{}
		interger.DecodeBody(parseResult)
		this.Pvno = int(interger.Value())
	case int(MSG_TYPE):
		interger := &Asn1Integer{}
		interger.DecodeBody(parseResult)
		this.MsgType = int(interger.Value())
	case int(PADATA):
		if kdcDirectChildPos == int(PADATA_TYPE) {
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.Padata.PadataType = int(interger.Value())
		} else if kdcDirectChildPos == int(PADATA_VALUE) {

		} else {
			log.Fatal("Can't unmarshal PADATA<type,value>, over 2 items", )
		}
	case int(CREALM):
		octstring := &Asn1OctetString{}
		octstring.DecodeBody(parseResult)
		this.Crealm = octstring.Value()
	case int(CNAME):
		if kdcDirectChildPos == int(NAME_TYPE) {
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.Padata.PadataType = int(interger.Value())
		} else if kdcDirectChildPos == int(NAME_STRING) {

		} else {
			log.Fatal("Can't unmarshal PADATA<type,value>, over 2 items", )
		}
	case int(TICKET):
		switch kdcDirectChildPos {
		case TKT_VNO:
		case REALM:
		case SNAME:
		case TK_ENC_PART:
			if kdcGrandsonPos == int(ETYPE) {
				interger := &Asn1Integer{}
				interger.DecodeBody(parseResult)
				this.Padata.PadataType = int(interger.Value())
			} else if kdcGrandsonPos == int(KVNO) {

			} else if kdcGrandsonPos == int(CIPHER) {

			} else {
				log.Fatal("Can't unmarshal PADATA<type,value>, over 2 items", )
			}
		default:
			log.Fatal("Can't unmarshal TICKET<vno,realm,sname,Encry>, over 4 items", )
		}
	case int(ENC_PART):
		if kdcDirectChildPos == int(ETYPE) {
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.Padata.PadataType = int(interger.Value())
		} else if kdcDirectChildPos == int(KVNO) {

		} else if kdcDirectChildPos == int(CIPHER) {

		} else {
			log.Fatal("Can't unmarshal PADATA<type,value>, over 2 items", )
		}
	default:
		log.Fatal("Can't unmarshal item:", parseResult)
	}
}

func (this *KdcRep) TaggingDecode(parseResult *unmarshal.Asn1ParseResult) {
	body := parseResult.Children[0]
	if len(body.Children) == 1 {
		this.Decode(body)
	}
	this.ActDecodeBody(body, 0)
}

func (this *KdcRep) ActDecodeBody(parseResult *unmarshal.Asn1ParseResult, level int) {
	children := parseResult.Children
	lastPos := -1
	foundPos := -1
	for _, parseItem := range children {
		if parseItem.Tag.IsEOC() {
			continue
		}
		foundPos = this.match(lastPos, parseItem)
		if foundPos == -1 {
			log.Fatal("Unexpected item:", parseItem)
		}
		lastPos = foundPos
		this.position[level] = byte(foundPos)
		this.attemptBinding(parseItem, foundPos, level)
	}
}

func (this *KdcRep) match(lastPos int, parseItem *unmarshal.Asn1ParseResult) int {
	foundPos := -1
	for i := lastPos + 1; i < len(this.fieldInfos); i++ {
		fieldInfo := this.fieldInfos[i]
		if fieldInfo != -1 {
			if !parseItem.Tag.IsContextSpecific() {
				continue
			}
			if fieldInfo == parseItem.Tag.TagNo {
				foundPos = i
				break
			}
		} else {
			log.Fatal("Unsupported item:", parseItem)
		}
	}
	return foundPos
}

func (this *KdcRep) attemptBinding(parseItem *unmarshal.Asn1ParseResult, foundPos int, level int) {
	body := parseItem.Children[0]
	length := body.GetIndex()
	isnest := body.Tag.IsNested()
	if length == 0 && !isnest {
		this.UniversalDecode(body)
	} else {
		this.findElement(level)
	}
}

func (this *KdcRep) findElement(level int)  {

}
