// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0xe5b.proto

package oidb

type LifeAchievementItem struct {
	AchievementId      *uint32 `protobuf:"varint,1,opt"`
	AchievementTitle   *string `protobuf:"bytes,2,opt"`
	AchievementIcon    *string `protobuf:"bytes,3,opt"`
	HasPraised         *bool   `protobuf:"varint,4,opt"`
	PraiseNum          *uint32 `protobuf:"varint,5,opt"`
	AchievementContent []byte  `protobuf:"bytes,6,opt"`
}

func (x *LifeAchievementItem) GetAchievementId() uint32 {
	if x != nil && x.AchievementId != nil {
		return *x.AchievementId
	}
	return 0
}

func (x *LifeAchievementItem) GetAchievementTitle() string {
	if x != nil && x.AchievementTitle != nil {
		return *x.AchievementTitle
	}
	return ""
}

func (x *LifeAchievementItem) GetAchievementIcon() string {
	if x != nil && x.AchievementIcon != nil {
		return *x.AchievementIcon
	}
	return ""
}

func (x *LifeAchievementItem) GetHasPraised() bool {
	if x != nil && x.HasPraised != nil {
		return *x.HasPraised
	}
	return false
}

func (x *LifeAchievementItem) GetPraiseNum() uint32 {
	if x != nil && x.PraiseNum != nil {
		return *x.PraiseNum
	}
	return 0
}

type DE5BReqBody struct {
	Uin                   *uint64  `protobuf:"varint,1,opt"`
	AchievementId         []uint32 `protobuf:"varint,2,rep"`
	MaxCount              *uint32  `protobuf:"varint,3,opt"`
	ReqAchievementContent *bool    `protobuf:"varint,4,opt"`
}

func (x *DE5BReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *DE5BReqBody) GetMaxCount() uint32 {
	if x != nil && x.MaxCount != nil {
		return *x.MaxCount
	}
	return 0
}

func (x *DE5BReqBody) GetReqAchievementContent() bool {
	if x != nil && x.ReqAchievementContent != nil {
		return *x.ReqAchievementContent
	}
	return false
}

type DE5BRspBody struct {
	AchievementTotalCount *uint32                `protobuf:"varint,1,opt"`
	LifeAchItem           []*LifeAchievementItem `protobuf:"bytes,2,rep"`
	AchievementOpenid     *string                `protobuf:"bytes,3,opt"`
}

func (x *DE5BRspBody) GetAchievementTotalCount() uint32 {
	if x != nil && x.AchievementTotalCount != nil {
		return *x.AchievementTotalCount
	}
	return 0
}

func (x *DE5BRspBody) GetAchievementOpenid() string {
	if x != nil && x.AchievementOpenid != nil {
		return *x.AchievementOpenid
	}
	return ""
}
