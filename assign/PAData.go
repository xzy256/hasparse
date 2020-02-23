package assign

const (
	PADATA_TYPE  int = 1
	PADATA_VALUE int = 2
)

/*
PAData   		::= SEQUENCE {
	-- NOTE: first tag is [1], not [0]
	padata-type     [1] Int32,
	padata-value    [2] OCTET STRING -- might be encoded AP-REQ
}
*/
type PAData struct {
	PadataType  int
	PadataValue *Asn1String
}
