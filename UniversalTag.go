package main

type UniversalTag int

const (
	UNKNOWN           UniversalTag = -1
	CHOICE            UniversalTag = -2 // Only for internal using
	ANY               UniversalTag = -3 // Only for internal using
	EOC               UniversalTag = 0  // End of content, use by BER
	BOOLEAN           UniversalTag = 0x01
	INTEGER           UniversalTag = 0x02
	BIT_STRING        UniversalTag = 0x03
	OCTET_STRING      UniversalTag = 0x04
	NULL              UniversalTag = 0x05
	OBJECT_IDENTIFIER UniversalTag = 0x06
	OBJECT_DESCRIPTOR UniversalTag = 0x07
	EXTERNAL          UniversalTag = 0x08
	REAL              UniversalTag = 0x09
	ENUMERATED        UniversalTag = 0x0a
	EMBEDDED_PDV      UniversalTag = 0x0b
	UTF8_STRING       UniversalTag = 0x0c
	RELATIVE_OID      UniversalTag = 0x0d
	RESERVED_14       UniversalTag = 0x0e
	RESERVED_15       UniversalTag = 0x0f
	SEQUENCE          UniversalTag = 0x10
	SEQUENCE_OF       UniversalTag = 0x10
	SET               UniversalTag = 0x11
	SET_OF            UniversalTag = 0x11
	NUMERIC_STRING    UniversalTag = 0x12
	PRINTABLE_STRING  UniversalTag = 0x13
	T61_STRING        UniversalTag = 0x14
	VIDEOTEX_STRING   UniversalTag = 0x15
	IA5_STRING        UniversalTag = 0x16
	UTC_TIME          UniversalTag = 0x17
	GENERALIZED_TIME  UniversalTag = 0x18
	GRAPHIC_STRING    UniversalTag = 0x19
	VISIBLE_STRING    UniversalTag = 0x1a
	GENERAL_STRING    UniversalTag = 0x1b
	UNIVERSAL_STRING  UniversalTag = 0x1c
	CHARACTER_STRING  UniversalTag = 0x1d
	BMP_STRING        UniversalTag = 0x1e
	RESERVED_31       UniversalTag = 0x1f
)

func UniversalTagFromValue(value int) UniversalTag {
	switch value {
	case -2:
		return CHOICE
	case 0x00:
		return EOC
	case 0x01:
		return BOOLEAN
	case 0x02:
		return INTEGER
	case 0x03:
		return BIT_STRING
	case 0x04:
		return OCTET_STRING
	case 0x05:
		return NULL
	case 0x06:
		return OBJECT_IDENTIFIER
	case 0x07:
		return OBJECT_DESCRIPTOR
	case 0x08:
		return EXTERNAL
	case 0x09:
		return REAL
	case 0x0A:
		return ENUMERATED
	case 0x0B:
		return EMBEDDED_PDV
	case 0x0C:
		return UTF8_STRING
	case 0x0D:
		return RELATIVE_OID
	case 0x0E:
		return RESERVED_14
	case 0x0F:
		return RESERVED_15
	case 0x10:
		return SEQUENCE
	case 0x11:
		return SET
	case 0x12:
		return NUMERIC_STRING
	case 0x13:
		return PRINTABLE_STRING
	case 0x14:
		return T61_STRING
	case 0x15:
		return VIDEOTEX_STRING
	case 0x16:
		return IA5_STRING
	case 0x17:
		return UTC_TIME
	case 0x18:
		return GENERALIZED_TIME
	case 0x19:
		return GRAPHIC_STRING
	case 0x1A:
		return VISIBLE_STRING
	case 0x1B:
		return GENERAL_STRING
	case 0x1C:
		return UNIVERSAL_STRING
	case 0x1D:
		return CHARACTER_STRING
	case 0x1E:
		return BMP_STRING
	case 0x1F:
		return RESERVED_31
	default:
		return UNKNOWN
	}
}
