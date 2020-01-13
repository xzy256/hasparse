package assign

const (
	NAME_TYPE   int = 0
	NAME_STRING int = 1
)

type PrincipalName struct {
	NameType   int
	NameString *Asn1OctetString
	filedInfo  []int
	position   []byte
}

func (this *PrincipalName) Init() {
	this.position = []byte{4, 0, 0}
	this.filedInfo = []int{NAME_TYPE, NAME_STRING}
}
