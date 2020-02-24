package assign

import (
	"github.com/xzy256/hasparse/common"
	"github.com/xzy256/hasparse/unmarshal"
)

type TaggingOption struct {
	TagNo         int
	IsImplicit    bool
	IsAppSpecific bool
}

func NewTaggingOption(tagNo int, isImplicit, isAppSpecific bool) *TaggingOption {
	return &TaggingOption{
		TagNo:         tagNo,
		IsImplicit:    isImplicit,
		IsAppSpecific: isAppSpecific,
	}
}

func NewImplicitAppSpecific(tagNo int) *TaggingOption {
	return &TaggingOption{
		TagNo:         tagNo,
		IsImplicit:    true,
		IsAppSpecific: true,
	}
}

func NewImplicitContextSpecific(tagNo int) *TaggingOption {
	return &TaggingOption{
		TagNo:         tagNo,
		IsImplicit:    true,
		IsAppSpecific: false,
	}
}

func NewExplicitContextSpecific(tagNo int) *TaggingOption {
	return &TaggingOption{
		TagNo:         tagNo,
		IsImplicit:    false,
		IsAppSpecific: false,
	}
}

func (this *TaggingOption) GetTag(isTaggedConstructed bool) *unmarshal.HasTag {
	isConstructed := !this.IsImplicit || isTaggedConstructed
	tag := common.CONTEXT_SPECIFIC
	if this.IsAppSpecific {
		tag = common.APPLICATION
	}
	tagClass := common.NewTagClass(tag)
	flag := tagClass.TCValue
	if isConstructed {
		flag = flag | 0x20
	}else{
		flag = flag | 0x00
	}

	return unmarshal.NewHasTagFlag(flag, this.TagNo)
}


