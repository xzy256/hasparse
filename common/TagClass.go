package common

const (
	UNIVERSAL        = 0x00
	APPLICATION      = 0x40
	CONTEXT_SPECIFIC = 0x80
	PRIVATE          = 0xC0
)

type TagClass struct {
	TCValue int
}

func NewTagClass(value int) *TagClass {
	return &TagClass{
		TCValue: value,
	}
}

func (this *TagClass) IsUniversal() bool {
	return this.TCValue == UNIVERSAL
}

func (this *TagClass) IsAppSpecific() bool {
	return this.TCValue == APPLICATION
}

func (this *TagClass) IsContextSpecific() bool {
	return this.TCValue == CONTEXT_SPECIFIC
}

func (this *TagClass) IsSpecific() bool {
	return this.TCValue == APPLICATION || this.TCValue == CONTEXT_SPECIFIC
}

func (this *TagClass) FromValue(value int) int {
	switch value {
	case 0x00:
		return UNIVERSAL
	case 0x40:
		return APPLICATION
	case 0x80:
		return CONTEXT_SPECIFIC
	case 0xC0:
		return PRIVATE
	default:
		return -1
	}
}

func (this *TagClass) TagClassFromTag() int {
	return this.FromValue(this.TCValue & 0xC0)
}
