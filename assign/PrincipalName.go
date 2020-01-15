package assign

const (
	NAME_TYPE   int = 0
	NAME_STRING int = 1
)

/*
PrincipalName   ::= SEQUENCE {
     name-type       [0] Int32,
     name-string     [1] SEQUENCE OF KerberosString
}
*/
type PrincipalName struct {
	NameType   int
	NameString string
}

