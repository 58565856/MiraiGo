// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0x8a7.proto

package oidb

type D8A7ReqBody struct {
	SubCmd                    *uint32 `protobuf:"varint,1,opt"`
	LimitIntervalTypeForUin   *uint32 `protobuf:"varint,2,opt"`
	LimitIntervalTypeForGroup *uint32 `protobuf:"varint,3,opt"`
	Uin                       *uint64 `protobuf:"varint,4,opt"`
	GroupCode                 *uint64 `protobuf:"varint,5,opt"`
}

func (x *D8A7ReqBody) GetSubCmd() uint32 {
	if x != nil && x.SubCmd != nil {
		return *x.SubCmd
	}
	return 0
}

func (x *D8A7ReqBody) GetLimitIntervalTypeForUin() uint32 {
	if x != nil && x.LimitIntervalTypeForUin != nil {
		return *x.LimitIntervalTypeForUin
	}
	return 0
}

func (x *D8A7ReqBody) GetLimitIntervalTypeForGroup() uint32 {
	if x != nil && x.LimitIntervalTypeForGroup != nil {
		return *x.LimitIntervalTypeForGroup
	}
	return 0
}

func (x *D8A7ReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *D8A7ReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

type D8A7RspBody struct {
	CanAtAll                 *bool   `protobuf:"varint,1,opt"`
	RemainAtAllCountForUin   *uint32 `protobuf:"varint,2,opt"`
	RemainAtAllCountForGroup *uint32 `protobuf:"varint,3,opt"`
	PromptMsg1               []byte  `protobuf:"bytes,4,opt"`
	PromptMsg2               []byte  `protobuf:"bytes,5,opt"`
}

func (x *D8A7RspBody) GetCanAtAll() bool {
	if x != nil && x.CanAtAll != nil {
		return *x.CanAtAll
	}
	return false
}

func (x *D8A7RspBody) GetRemainAtAllCountForUin() uint32 {
	if x != nil && x.RemainAtAllCountForUin != nil {
		return *x.RemainAtAllCountForUin
	}
	return 0
}

func (x *D8A7RspBody) GetRemainAtAllCountForGroup() uint32 {
	if x != nil && x.RemainAtAllCountForGroup != nil {
		return *x.RemainAtAllCountForGroup
	}
	return 0
}
