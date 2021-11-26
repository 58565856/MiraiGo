// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: multimsg.proto

package multimsg

type ExternMsg struct {
	ChannelType int32 `protobuf:"varint,1,opt"`
}

func (x *ExternMsg) GetChannelType() int32 {
	if x != nil {
		return x.ChannelType
	}
	return 0
}

type MultiMsgApplyDownReq struct {
	MsgResid []byte `protobuf:"bytes,1,opt"`
	MsgType  int32  `protobuf:"varint,2,opt"`
	SrcUin   int64  `protobuf:"varint,3,opt"`
}

func (x *MultiMsgApplyDownReq) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *MultiMsgApplyDownReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MultiMsgApplyDownReq) GetSrcUin() int64 {
	if x != nil {
		return x.SrcUin
	}
	return 0
}

type MultiMsgApplyDownRsp struct {
	Result           int32      `protobuf:"varint,1,opt"`
	ThumbDownPara    []byte     `protobuf:"bytes,2,opt"`
	MsgKey           []byte     `protobuf:"bytes,3,opt"`
	Uint32DownIp     []int32    `protobuf:"varint,4,rep"`
	Uint32DownPort   []int32    `protobuf:"varint,5,rep"`
	MsgResid         []byte     `protobuf:"bytes,6,opt"`
	MsgExternInfo    *ExternMsg `protobuf:"bytes,7,opt"`
	BytesDownIpV6    [][]byte   `protobuf:"bytes,8,rep"`
	Uint32DownV6Port []int32    `protobuf:"varint,9,rep"`
}

func (x *MultiMsgApplyDownRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *MultiMsgApplyDownRsp) GetThumbDownPara() []byte {
	if x != nil {
		return x.ThumbDownPara
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetMsgKey() []byte {
	if x != nil {
		return x.MsgKey
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetUint32DownIp() []int32 {
	if x != nil {
		return x.Uint32DownIp
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetUint32DownPort() []int32 {
	if x != nil {
		return x.Uint32DownPort
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetMsgExternInfo() *ExternMsg {
	if x != nil {
		return x.MsgExternInfo
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetBytesDownIpV6() [][]byte {
	if x != nil {
		return x.BytesDownIpV6
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetUint32DownV6Port() []int32 {
	if x != nil {
		return x.Uint32DownV6Port
	}
	return nil
}

type MultiMsgApplyUpReq struct {
	DstUin  int64  `protobuf:"varint,1,opt"`
	MsgSize int64  `protobuf:"varint,2,opt"`
	MsgMd5  []byte `protobuf:"bytes,3,opt"`
	MsgType int32  `protobuf:"varint,4,opt"`
	ApplyId int32  `protobuf:"varint,5,opt"`
}

func (x *MultiMsgApplyUpReq) GetDstUin() int64 {
	if x != nil {
		return x.DstUin
	}
	return 0
}

func (x *MultiMsgApplyUpReq) GetMsgSize() int64 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *MultiMsgApplyUpReq) GetMsgMd5() []byte {
	if x != nil {
		return x.MsgMd5
	}
	return nil
}

func (x *MultiMsgApplyUpReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MultiMsgApplyUpReq) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

type MultiMsgApplyUpRsp struct {
	Result         int32      `protobuf:"varint,1,opt"`
	MsgResid       string     `protobuf:"bytes,2,opt"`
	MsgUkey        []byte     `protobuf:"bytes,3,opt"`
	Uint32UpIp     []int32    `protobuf:"varint,4,rep"`
	Uint32UpPort   []int32    `protobuf:"varint,5,rep"`
	BlockSize      int64      `protobuf:"varint,6,opt"`
	UpOffset       int64      `protobuf:"varint,7,opt"`
	ApplyId        int32      `protobuf:"varint,8,opt"`
	MsgKey         []byte     `protobuf:"bytes,9,opt"`
	MsgSig         []byte     `protobuf:"bytes,10,opt"`
	MsgExternInfo  *ExternMsg `protobuf:"bytes,11,opt"`
	BytesUpIpV6    [][]byte   `protobuf:"bytes,12,rep"`
	Uint32UpV6Port []int32    `protobuf:"varint,13,rep"`
}

func (x *MultiMsgApplyUpRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetMsgResid() string {
	if x != nil {
		return x.MsgResid
	}
	return ""
}

func (x *MultiMsgApplyUpRsp) GetMsgUkey() []byte {
	if x != nil {
		return x.MsgUkey
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetUint32UpIp() []int32 {
	if x != nil {
		return x.Uint32UpIp
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetUint32UpPort() []int32 {
	if x != nil {
		return x.Uint32UpPort
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetUpOffset() int64 {
	if x != nil {
		return x.UpOffset
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetMsgKey() []byte {
	if x != nil {
		return x.MsgKey
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetMsgSig() []byte {
	if x != nil {
		return x.MsgSig
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetMsgExternInfo() *ExternMsg {
	if x != nil {
		return x.MsgExternInfo
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetBytesUpIpV6() [][]byte {
	if x != nil {
		return x.BytesUpIpV6
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetUint32UpV6Port() []int32 {
	if x != nil {
		return x.Uint32UpV6Port
	}
	return nil
}

type MultiReqBody struct {
	Subcmd               int32                   `protobuf:"varint,1,opt"`
	TermType             int32                   `protobuf:"varint,2,opt"`
	PlatformType         int32                   `protobuf:"varint,3,opt"`
	NetType              int32                   `protobuf:"varint,4,opt"`
	BuildVer             string                  `protobuf:"bytes,5,opt"`
	MultimsgApplyupReq   []*MultiMsgApplyUpReq   `protobuf:"bytes,6,rep"`
	MultimsgApplydownReq []*MultiMsgApplyDownReq `protobuf:"bytes,7,rep"`
	BuType               int32                   `protobuf:"varint,8,opt"`
	ReqChannelType       int32                   `protobuf:"varint,9,opt"`
}

func (x *MultiReqBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *MultiReqBody) GetTermType() int32 {
	if x != nil {
		return x.TermType
	}
	return 0
}

func (x *MultiReqBody) GetPlatformType() int32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *MultiReqBody) GetNetType() int32 {
	if x != nil {
		return x.NetType
	}
	return 0
}

func (x *MultiReqBody) GetBuildVer() string {
	if x != nil {
		return x.BuildVer
	}
	return ""
}

func (x *MultiReqBody) GetMultimsgApplyupReq() []*MultiMsgApplyUpReq {
	if x != nil {
		return x.MultimsgApplyupReq
	}
	return nil
}

func (x *MultiReqBody) GetMultimsgApplydownReq() []*MultiMsgApplyDownReq {
	if x != nil {
		return x.MultimsgApplydownReq
	}
	return nil
}

func (x *MultiReqBody) GetBuType() int32 {
	if x != nil {
		return x.BuType
	}
	return 0
}

func (x *MultiReqBody) GetReqChannelType() int32 {
	if x != nil {
		return x.ReqChannelType
	}
	return 0
}

type MultiRspBody struct {
	Subcmd               int32                   `protobuf:"varint,1,opt"`
	MultimsgApplyupRsp   []*MultiMsgApplyUpRsp   `protobuf:"bytes,2,rep"`
	MultimsgApplydownRsp []*MultiMsgApplyDownRsp `protobuf:"bytes,3,rep"`
}

func (x *MultiRspBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *MultiRspBody) GetMultimsgApplyupRsp() []*MultiMsgApplyUpRsp {
	if x != nil {
		return x.MultimsgApplyupRsp
	}
	return nil
}

func (x *MultiRspBody) GetMultimsgApplydownRsp() []*MultiMsgApplyDownRsp {
	if x != nil {
		return x.MultimsgApplydownRsp
	}
	return nil
}
