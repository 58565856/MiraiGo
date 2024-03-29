// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0xec4.proto

package oidb

type Comment struct {
	Id       *string `protobuf:"bytes,1,opt"`
	Comment  *string `protobuf:"bytes,2,opt"`
	Time     *uint64 `protobuf:"varint,3,opt"`
	FromUin  *uint64 `protobuf:"varint,4,opt"`
	ToUin    *uint64 `protobuf:"varint,5,opt"`
	ReplyId  *string `protobuf:"bytes,6,opt"`
	FromNick *string `protobuf:"bytes,7,opt"`
}

func (x *Comment) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Comment) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

func (x *Comment) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Comment) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *Comment) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *Comment) GetReplyId() string {
	if x != nil && x.ReplyId != nil {
		return *x.ReplyId
	}
	return ""
}

func (x *Comment) GetFromNick() string {
	if x != nil && x.FromNick != nil {
		return *x.FromNick
	}
	return ""
}

type Praise struct {
	FromUin  *uint64 `protobuf:"varint,1,opt"`
	ToUin    *uint64 `protobuf:"varint,2,opt"`
	Time     *uint64 `protobuf:"varint,3,opt"`
	FromNick *string `protobuf:"bytes,4,opt"`
}

func (x *Praise) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *Praise) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *Praise) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Praise) GetFromNick() string {
	if x != nil && x.FromNick != nil {
		return *x.FromNick
	}
	return ""
}

type Quest struct {
	Id          *string    `protobuf:"bytes,1,opt"`
	Quest       *string    `protobuf:"bytes,2,opt"`
	QuestUin    *uint64    `protobuf:"varint,3,opt"`
	Time        *uint64    `protobuf:"varint,4,opt"`
	Ans         *string    `protobuf:"bytes,5,opt"`
	AnsTime     *uint64    `protobuf:"varint,6,opt"`
	Comment     []*Comment `protobuf:"bytes,7,rep"`
	Praise      []*Praise  `protobuf:"bytes,8,rep"`
	PraiseNum   *uint64    `protobuf:"varint,9,opt"`
	LikeKey     *string    `protobuf:"bytes,10,opt"`
	SystemId    *uint64    `protobuf:"varint,11,opt"`
	CommentNum  *uint64    `protobuf:"varint,12,opt"`
	ShowType    *uint64    `protobuf:"varint,13,opt"`
	ShowTimes   *uint64    `protobuf:"varint,14,opt"`
	BeenPraised *uint64    `protobuf:"varint,15,opt"`
	QuestRead   *bool      `protobuf:"varint,16,opt"`
	AnsShowType *uint64    `protobuf:"varint,17,opt"`
}

func (x *Quest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Quest) GetQuest() string {
	if x != nil && x.Quest != nil {
		return *x.Quest
	}
	return ""
}

func (x *Quest) GetQuestUin() uint64 {
	if x != nil && x.QuestUin != nil {
		return *x.QuestUin
	}
	return 0
}

func (x *Quest) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Quest) GetAns() string {
	if x != nil && x.Ans != nil {
		return *x.Ans
	}
	return ""
}

func (x *Quest) GetAnsTime() uint64 {
	if x != nil && x.AnsTime != nil {
		return *x.AnsTime
	}
	return 0
}

func (x *Quest) GetPraiseNum() uint64 {
	if x != nil && x.PraiseNum != nil {
		return *x.PraiseNum
	}
	return 0
}

func (x *Quest) GetLikeKey() string {
	if x != nil && x.LikeKey != nil {
		return *x.LikeKey
	}
	return ""
}

func (x *Quest) GetSystemId() uint64 {
	if x != nil && x.SystemId != nil {
		return *x.SystemId
	}
	return 0
}

func (x *Quest) GetCommentNum() uint64 {
	if x != nil && x.CommentNum != nil {
		return *x.CommentNum
	}
	return 0
}

func (x *Quest) GetShowType() uint64 {
	if x != nil && x.ShowType != nil {
		return *x.ShowType
	}
	return 0
}

func (x *Quest) GetShowTimes() uint64 {
	if x != nil && x.ShowTimes != nil {
		return *x.ShowTimes
	}
	return 0
}

func (x *Quest) GetBeenPraised() uint64 {
	if x != nil && x.BeenPraised != nil {
		return *x.BeenPraised
	}
	return 0
}

func (x *Quest) GetQuestRead() bool {
	if x != nil && x.QuestRead != nil {
		return *x.QuestRead
	}
	return false
}

func (x *Quest) GetAnsShowType() uint64 {
	if x != nil && x.AnsShowType != nil {
		return *x.AnsShowType
	}
	return 0
}

type DEC4ReqBody struct {
	Uin        *uint64 `protobuf:"varint,1,opt"`
	QuestNum   *uint64 `protobuf:"varint,2,opt"`
	CommentNum *uint64 `protobuf:"varint,3,opt"`
	Cookie     []byte  `protobuf:"bytes,4,opt"`
	FetchType  *uint32 `protobuf:"varint,5,opt"`
}

func (x *DEC4ReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *DEC4ReqBody) GetQuestNum() uint64 {
	if x != nil && x.QuestNum != nil {
		return *x.QuestNum
	}
	return 0
}

func (x *DEC4ReqBody) GetCommentNum() uint64 {
	if x != nil && x.CommentNum != nil {
		return *x.CommentNum
	}
	return 0
}

func (x *DEC4ReqBody) GetFetchType() uint32 {
	if x != nil && x.FetchType != nil {
		return *x.FetchType
	}
	return 0
}

type DEC4RspBody struct {
	Quest            []*Quest `protobuf:"bytes,1,rep"`
	IsFetchOver      *bool    `protobuf:"varint,2,opt"`
	TotalQuestNum    *uint32  `protobuf:"varint,3,opt"`
	Cookie           []byte   `protobuf:"bytes,4,opt"`
	Ret              *uint32  `protobuf:"varint,5,opt"`
	AnsweredQuestNum *uint32  `protobuf:"varint,6,opt"`
}

func (x *DEC4RspBody) GetIsFetchOver() bool {
	if x != nil && x.IsFetchOver != nil {
		return *x.IsFetchOver
	}
	return false
}

func (x *DEC4RspBody) GetTotalQuestNum() uint32 {
	if x != nil && x.TotalQuestNum != nil {
		return *x.TotalQuestNum
	}
	return 0
}

func (x *DEC4RspBody) GetRet() uint32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *DEC4RspBody) GetAnsweredQuestNum() uint32 {
	if x != nil && x.AnsweredQuestNum != nil {
		return *x.AnsweredQuestNum
	}
	return 0
}
