package assign

const (
	PADATA_TYPE  int = 1
	PADATA_VALUE int = 2
)

type PAData struct {
	PadataType  int
	PadataValue *Asn1OctetString
	fileInfo   []int
	position   []byte
}

func (this *PAData) Init() {
	this.position = []byte{2, 0, 0}
	this.fileInfo = []int{PADATA_TYPE, PADATA_VALUE}
}