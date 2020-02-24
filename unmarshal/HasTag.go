package unmarshal

import "github.com/xzy256/hasparse/common"

type HasTag struct {
	TagNo    int
	TagFlags int
}

func NewHasTag(tag int) *HasTag {
	return &HasTag{
		TagNo:    tag & 0x1F,
		TagFlags: tag & 0xE0,
	}
}

func NewHasUniversalTag(tag common.UniversalTag) *HasTag {
	return &HasTag{
		TagNo:    int(tag),
		TagFlags: common.UNIVERSAL,
	}
}

func NewHasTagFlag(flag, tno int) *HasTag {
	return &HasTag{
		TagNo:    tno,
		TagFlags: flag & 0xE0,
	}
}

func NewAppTag(tagNo int) *HasTag {
	return NewHasTagFlag(common.APPLICATION, tagNo)
}

func NewCtxTag(tagNo int) *HasTag {
	return NewHasTagFlag(common.CONTEXT_SPECIFIC, tagNo)
}

func tagClass(tagFlags int) int {
	tag := common.NewTagClass(tagFlags)
	return tag.TagClassFromTag()
}

func (this *HasTag) UsePrimitive(isPrimitive bool) {
	if isPrimitive {
		this.TagFlags &= ^0x20
	} else {
		this.TagFlags |= 0x20
	}
}

func (this *HasTag) IsPrimitive() bool {
	return (this.TagFlags & 0x20) == 0
}

func universalTag(tflag, tno int) int {
	if tagClass(tflag) == common.UNIVERSAL {
		return int(common.UniversalTagFromValue(tno))
	}
	return int(common.UNKNOWN)
}

func (this *HasTag) IsEOC() bool {
	return universalTag(this.TagFlags, this.TagNo) == int(common.EOC)
}

func (this *HasTag) IsNull() bool {
	return universalTag(this.TagFlags, this.TagNo) == int(common.NULL)
}

func (this *HasTag) IsAppSpecific() bool {
	return tagClass(this.TagFlags) == common.APPLICATION
}

func (this *HasTag) IsContextSpecific() bool {
	return tagClass(this.TagFlags) == common.CONTEXT_SPECIFIC
}

func (this *HasTag) IsSpecific() bool {
	tag := tagClass(this.TagFlags)
	return tag == common.APPLICATION || tag == common.CONTEXT_SPECIFIC
}

func (this *HasTag) TagByte() byte {
	t := this.TagNo
	if this.TagNo >= 0x1F {
		t = 0x1F
	}
	n := this.TagFlags | t
	return (byte)(n & 0xFF)
}

func (this *HasTag) HashCode() int {
	result := this.TagFlags
	result = 31*result + this.TagNo
	return result
}

func (this *HasTag) Equal (cmpTag *HasTag) bool {
	if this.TagFlags != cmpTag.TagFlags {
		return false
	}
	return this.TagNo == cmpTag.TagNo
}

func (this *HasTag) IsNested() bool {
	val := this.TagNo + this.TagFlags
	val = val & 0x20
	flag := val >> 5
	if flag == 1 {
		return true
	}
	return false
}

