package assign

import (
	"container/list"
	"fmt"
	"hasparse/unmarshal"
	"hasparse/utils"
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

/**
KdcRep         ::= SEQUENCE {
	pvno            [0] INTEGER (5),
	msg-type        [1] INTEGER (11 -- AS -- | 13 -- TGS --),
	padata          [2] SEQUENCE OF PA-DATA OPTIONAL
	-- NOTE: not empty --,
	crealm          [3] Realm,
	cname           [4] PrincipalName,
	ticket          [5] Ticket,
	enc-part        [6] EncryptedData
	-- EncASRepPart or EncTGSRepPart,
	-- as appropriate
}
*/

type KdcRep struct {
	Pvno        int
	MsgType     int
	Padata      *PAData //optional
	Crealm      string
	Cname       *PrincipalName
	Ticket      *KdcTicket
	EncData     *EncryptedData
	fieldInfos  []int
	position    []byte // same as 0x532, 5-Ticket 3-Ticket.enc_part 2-Ticket.enc_part.cipher, default [255,255,255]
	taggingList *list.List
}

func (this *KdcRep) Init() {
	this.fieldInfos = []int{0, 1, 2, 3, 4, 5, 6}
	this.position = []byte{255, 255, 255}
	this.taggingList = list.New()
	this.Cname = &PrincipalName{}
	this.Padata = &PAData{}
	this.Ticket = &KdcTicket{}
	this.Ticket.Init()
	this.EncData = &EncryptedData{}
}

func (this *KdcRep) ResetPosition() {
	this.position[1] = 255
	this.position[2] = 255
}

func (this *KdcRep) Decode(parseResult *unmarshal.Asn1ParseResult) {
	if parseResult.Tag.IsNested() {
		this.TaggingDecode(parseResult)
	} else {
		this.UniversalDecode(parseResult, false)
	}

}

func (this *KdcRep) UniversalDecode(parseResult *unmarshal.Asn1ParseResult, hasOptional bool) {
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
	case int(PADATA): // not check
		if kdcDirectChildPos == int(PADATA_TYPE) {
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.Padata.PadataType = int(interger.Value())
		} else if kdcDirectChildPos == int(PADATA_VALUE) {
			fmt.Println()
		}
	case int(CREALM):
		octstring := &Asn1OctetString{}
		octstring.DecodeBody(parseResult)
		this.Crealm = octstring.Value()
	case int(CNAME):
		if kdcDirectChildPos == int(NAME_TYPE) {
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.Cname.NameType = int(interger.Value())
		} else if kdcDirectChildPos == int(NAME_STRING) {
			octstring := &Asn1OctetString{}
			octstring.DecodeBody(parseResult)
			this.Cname.NameString = octstring.Value()
		} else {
			log.Fatal("Can't unmarshal PADATA<type,value>, over 2 items", )
		}
	case int(TICKET):
		switch kdcDirectChildPos {
		case TKT_VNO:
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.Ticket.Tktvno = int(interger.Value())
		case REALM:
			octstring := &Asn1OctetString{}
			octstring.DecodeBody(parseResult)
			this.Ticket.Realm = octstring.Value()
		case SNAME:
			if kdcGrandsonPos == int(NAME_TYPE) {
				interger := &Asn1Integer{}
				interger.DecodeBody(parseResult)
				this.Ticket.Sname.NameType = int(interger.Value())
			} else if kdcGrandsonPos == int(NAME_STRING) {
				octstring := &Asn1OctetString{}
				octstring.DecodeBody(parseResult)
				this.Ticket.Sname.NameString = octstring.Value()
			}
		case TK_ENC_PART:
			if kdcGrandsonPos == int(ETYPE) {
				interger := &Asn1Integer{}
				interger.DecodeBody(parseResult)
				this.Ticket.EncPart.Etype = int(interger.Value())
			} else if kdcGrandsonPos == int(KVNO) {
				interger := &Asn1Integer{}
				interger.DecodeBody(parseResult)
				this.Ticket.EncPart.Kvno = uint32(interger.Value())
			} else if kdcGrandsonPos == int(CIPHER) {
				octstring := &Asn1OctetString{}
				octstring.DecodeBody(parseResult)
				this.Ticket.EncPart.Cipher = octstring
			}
		default:
			log.Fatal("Can't unmarshal TICKET<vno,realm,sname,Encry>, over 4 items", )
		}
	case int(ENC_PART):
		if kdcDirectChildPos == int(ETYPE) {
			interger := &Asn1Integer{}
			interger.DecodeBody(parseResult)
			this.EncData.Etype = int(interger.Value())
		} else {
			kdcDirectChildPos += 1
			if kdcDirectChildPos == int(KVNO) {
				interger := &Asn1Integer{}
				interger.DecodeBody(parseResult)
				this.EncData.Kvno = uint32(interger.Value())
			} else if kdcDirectChildPos == int(CIPHER) {
				octstring := &Asn1OctetString{}
				octstring.DecodeBody(parseResult)
				this.EncData.Cipher = octstring
			}
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
	children := parseResult.GetIndex()
	lastPos := -1
	foundPos := -1
	for i := 0; i < children; i++ {
		parseItem := parseResult.Children[i]
		this.ResetPosition()
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
		this.UniversalDecode(body, false)
	} else if length == 1 && isnest {
		for length == 1 && isnest {
			body = body.Children[0]
			length = body.GetIndex()
			isnest = body.Tag.IsNested()
		}
		this.findElement(body, foundPos, 0)
	} else {
		this.findElement(body, foundPos, 0)
	}
}

func (this *KdcRep) findElement(body *unmarshal.Asn1ParseResult, foundPos int, level int) {
	switch foundPos {
	case int(PADATA):
		this.taggingList.PushBack(&Asn1Element{"PADATA_TYPE", PADATA_TYPE, "int"})
		this.taggingList.PushBack(&Asn1Element{"PADATA_VALUE", PADATA_VALUE, "oct string"})
	case int(CNAME):
		this.taggingList.PushBack(&Asn1Element{"NAME_TYPE", NAME_TYPE, "int"})
		this.taggingList.PushBack(&Asn1Element{"NAME_STRING", 16, "string"}) // must be sequence--> HasTag(16, 32)
	case int(TICKET):
		this.taggingList.PushBack(&Asn1Element{"TKT_VNO", TKT_VNO, "int: 5"})
		this.taggingList.PushBack(&Asn1Element{"REALM", REALM, "string"})
		this.taggingList.PushBack(&Asn1Element{"SNAME", 16, "principal"}) // must be sequence--> HasTag(16, 32)
		this.taggingList.PushBack(&Asn1Element{"ENC_PART", int(ENC_PART), "encryptedData"})
	case int(ENC_PART):
		this.taggingList.PushBack(&Asn1Element{"ETYPE", ETYPE, "int: 5"})
		this.taggingList.PushBack(&Asn1Element{"KVNO", KVNO, "uint32"}) // optional
		this.taggingList.PushBack(&Asn1Element{"CIPHER", CIPHER, "oct string"})
	}

	this.DecodeElement(body, level)
}

func (this *KdcRep) DecodeElement(body *unmarshal.Asn1ParseResult, level int) {
	pos := false
	if this.position[0] == byte(TICKET) && this.position[1] == byte(TK_ENC_PART) && this.position[2] == byte(255) && level == 1 && body.GetIndex() < 3 {
		this.Ticket.EncPart.Optional = true
		this.RemoveOptional()
		pos = true
	} else if this.position[0] == byte(ENC_PART) && this.position[1] == byte(255) && level == 0 && body.GetIndex() < 3 {
		this.Ticket.EncPart.Optional = true
		this.RemoveOptional()
		pos = true
	}
	length := this.taggingList.Len()
	index := -1
	for i := 0; i < length; i++ {
		index++
		tmpresult := body.Children[index]
		val := this.taggingList.Front().Value
		t1 := tmpresult.Tag.TagNo
		t2 := val.(*Asn1Element).Tagno
		if this.position[0] == byte(TICKET) {
			if this.position[1] == byte(REALM) && this.position[2] == byte(255) {
				remain := utils.CopyListAfterRemoveHead(this.taggingList)
				this.taggingList.Init()
				this.position[1] = 2
				this.findElement(tmpresult.Children[0], int(CNAME), 1)
				this.taggingList = remain
				this.position[2] = byte(255)
				continue
			} else if this.position[1] == byte(SNAME) && this.position[2] == byte(255) && level == 0 {
				remain := utils.CopyListAfterRemoveHead(this.taggingList)
				this.taggingList.Init()
				this.position[1] = 3
				this.findElement(tmpresult.Children[0], int(ENC_PART), 1)
				this.taggingList = remain
				this.position[2] = byte(255)
				continue
			}
		}
		for {
			t1 = tmpresult.Tag.TagNo
			if t1 == t2 {
				this.position[level+1] = byte(i)
				tmpresult2 := tmpresult.Children[0]
				this.UniversalDecode(tmpresult2, pos)
				this.taggingList.Remove(this.taggingList.Front())
				break
			}
			tmpresult = tmpresult.Children[0]
		}
	}
}

func (this *KdcRep) RemoveOptional() { // enc_part remove optional kvno
	head := this.taggingList.Front()
	this.taggingList.Remove(head)
	e2 := this.taggingList.Front()
	this.taggingList.Remove(e2)
	this.taggingList.PushFront(head.Value)
}

func (this *KdcRep) Display() {
	fmt.Println("KdcRep{")
	fmt.Println("\t Pvno: ", this.Pvno)
	fmt.Println("\t MsgType: ", this.MsgType)
	fmt.Printf("\t Padata: { PadataType: %v, PadataType: %v }\n", this.Padata.PadataType, this.Padata.PadataValue)
	fmt.Println("\t Crealm: ", this.Crealm)
	fmt.Printf("\t Cname: { NameType: %v, NameString: %v }\n", this.Cname.NameType, this.Cname.NameString)
	fmt.Printf("\t Ticket: { Tktvno: %v, Realm: %v", this.Ticket.Tktvno, this.Ticket.Realm)
	fmt.Printf(", Sname: { NameType: %v, NameString: %v }", this.Ticket.Sname.NameType, this.Ticket.Sname.NameString)
	fmt.Printf(", EncPart: { Etype: %v, Kvno: %v, Cipher: %v }}\n", this.Ticket.EncPart.Etype, this.Ticket.EncPart.Kvno, this.Ticket.EncPart.Cipher.ValueBytes)
	fmt.Printf("\t EncData: { Etype: %v, Kvno: %v, Cipher: %v }\n", this.EncData.Etype, this.EncData.Kvno, this.EncData.Cipher.ValueBytes)
	fmt.Println("}")
}
