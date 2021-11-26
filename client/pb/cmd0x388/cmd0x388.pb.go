// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: cmd0x388.proto

package cmd0x388

type DelImgReq struct {
	SrcUin          *uint64 `protobuf:"varint,1,opt"`
	DstUin          *uint64 `protobuf:"varint,2,opt"`
	ReqTerm         *uint32 `protobuf:"varint,3,opt"`
	ReqPlatformType *uint32 `protobuf:"varint,4,opt"`
	BuType          *uint32 `protobuf:"varint,5,opt"`
	BuildVer        []byte  `protobuf:"bytes,6,opt"`
	FileResid       []byte  `protobuf:"bytes,7,opt"`
	PicWidth        *uint32 `protobuf:"varint,8,opt"`
	PicHeight       *uint32 `protobuf:"varint,9,opt"`
}

func (x *DelImgReq) GetSrcUin() uint64 {
	if x != nil && x.SrcUin != nil {
		return *x.SrcUin
	}
	return 0
}

func (x *DelImgReq) GetDstUin() uint64 {
	if x != nil && x.DstUin != nil {
		return *x.DstUin
	}
	return 0
}

func (x *DelImgReq) GetReqTerm() uint32 {
	if x != nil && x.ReqTerm != nil {
		return *x.ReqTerm
	}
	return 0
}

func (x *DelImgReq) GetReqPlatformType() uint32 {
	if x != nil && x.ReqPlatformType != nil {
		return *x.ReqPlatformType
	}
	return 0
}

func (x *DelImgReq) GetBuType() uint32 {
	if x != nil && x.BuType != nil {
		return *x.BuType
	}
	return 0
}

func (x *DelImgReq) GetBuildVer() []byte {
	if x != nil {
		return x.BuildVer
	}
	return nil
}

func (x *DelImgReq) GetFileResid() []byte {
	if x != nil {
		return x.FileResid
	}
	return nil
}

func (x *DelImgReq) GetPicWidth() uint32 {
	if x != nil && x.PicWidth != nil {
		return *x.PicWidth
	}
	return 0
}

func (x *DelImgReq) GetPicHeight() uint32 {
	if x != nil && x.PicHeight != nil {
		return *x.PicHeight
	}
	return 0
}

type DelImgRsp struct {
	Result    *uint32 `protobuf:"varint,1,opt"`
	FailMsg   []byte  `protobuf:"bytes,2,opt"`
	FileResid []byte  `protobuf:"bytes,3,opt"`
}

func (x *DelImgRsp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *DelImgRsp) GetFailMsg() []byte {
	if x != nil {
		return x.FailMsg
	}
	return nil
}

func (x *DelImgRsp) GetFileResid() []byte {
	if x != nil {
		return x.FileResid
	}
	return nil
}

type ExpRoamExtendInfo struct {
	Resid []byte `protobuf:"bytes,1,opt"`
}

func (x *ExpRoamExtendInfo) GetResid() []byte {
	if x != nil {
		return x.Resid
	}
	return nil
}

type ExpRoamPicInfo struct {
	ShopFlag *uint32 `protobuf:"varint,1,opt"`
	PkgId    *uint32 `protobuf:"varint,2,opt"`
	PicId    []byte  `protobuf:"bytes,3,opt"`
}

func (x *ExpRoamPicInfo) GetShopFlag() uint32 {
	if x != nil && x.ShopFlag != nil {
		return *x.ShopFlag
	}
	return 0
}

func (x *ExpRoamPicInfo) GetPkgId() uint32 {
	if x != nil && x.PkgId != nil {
		return *x.PkgId
	}
	return 0
}

func (x *ExpRoamPicInfo) GetPicId() []byte {
	if x != nil {
		return x.PicId
	}
	return nil
}

type ExtensionCommPicTryUp struct {
	Extinfo [][]byte `protobuf:"bytes,1,rep"`
}

func (x *ExtensionCommPicTryUp) GetExtinfo() [][]byte {
	if x != nil {
		return x.Extinfo
	}
	return nil
}

type ExtensionExpRoamTryUp struct {
	ExproamPicInfo []*ExpRoamPicInfo `protobuf:"bytes,1,rep"`
}

func (x *ExtensionExpRoamTryUp) GetExproamPicInfo() []*ExpRoamPicInfo {
	if x != nil {
		return x.ExproamPicInfo
	}
	return nil
}

type GetImgUrlReq struct {
	GroupCode       *uint64 `protobuf:"varint,1,opt"`
	DstUin          *uint64 `protobuf:"varint,2,opt"`
	Fileid          *uint64 `protobuf:"varint,3,opt"`
	FileMd5         []byte  `protobuf:"bytes,4,opt"`
	UrlFlag         *uint32 `protobuf:"varint,5,opt"`
	UrlType         *uint32 `protobuf:"varint,6,opt"`
	ReqTerm         *uint32 `protobuf:"varint,7,opt"`
	ReqPlatformType *uint32 `protobuf:"varint,8,opt"`
	InnerIp         *uint32 `protobuf:"varint,9,opt"`
	BuType          *uint32 `protobuf:"varint,10,opt"`
	BuildVer        []byte  `protobuf:"bytes,11,opt"`
	FileId          *uint64 `protobuf:"varint,12,opt"`
	FileSize        *uint64 `protobuf:"varint,13,opt"`
	OriginalPic     *uint32 `protobuf:"varint,14,opt"`
	RetryReq        *uint32 `protobuf:"varint,15,opt"`
	FileHeight      *uint32 `protobuf:"varint,16,opt"`
	FileWidth       *uint32 `protobuf:"varint,17,opt"`
	PicType         *uint32 `protobuf:"varint,18,opt"`
	PicUpTimestamp  *uint32 `protobuf:"varint,19,opt"`
	ReqTransferType *uint32 `protobuf:"varint,20,opt"`
	QqmeetGuildId   *uint64 `protobuf:"varint,21,opt"`
	QqmeetChannelId *uint64 `protobuf:"varint,22,opt"`
	DownloadIndex   []byte  `protobuf:"bytes,23,opt"`
}

func (x *GetImgUrlReq) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *GetImgUrlReq) GetDstUin() uint64 {
	if x != nil && x.DstUin != nil {
		return *x.DstUin
	}
	return 0
}

func (x *GetImgUrlReq) GetFileid() uint64 {
	if x != nil && x.Fileid != nil {
		return *x.Fileid
	}
	return 0
}

func (x *GetImgUrlReq) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *GetImgUrlReq) GetUrlFlag() uint32 {
	if x != nil && x.UrlFlag != nil {
		return *x.UrlFlag
	}
	return 0
}

func (x *GetImgUrlReq) GetUrlType() uint32 {
	if x != nil && x.UrlType != nil {
		return *x.UrlType
	}
	return 0
}

func (x *GetImgUrlReq) GetReqTerm() uint32 {
	if x != nil && x.ReqTerm != nil {
		return *x.ReqTerm
	}
	return 0
}

func (x *GetImgUrlReq) GetReqPlatformType() uint32 {
	if x != nil && x.ReqPlatformType != nil {
		return *x.ReqPlatformType
	}
	return 0
}

func (x *GetImgUrlReq) GetInnerIp() uint32 {
	if x != nil && x.InnerIp != nil {
		return *x.InnerIp
	}
	return 0
}

func (x *GetImgUrlReq) GetBuType() uint32 {
	if x != nil && x.BuType != nil {
		return *x.BuType
	}
	return 0
}

func (x *GetImgUrlReq) GetBuildVer() []byte {
	if x != nil {
		return x.BuildVer
	}
	return nil
}

func (x *GetImgUrlReq) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *GetImgUrlReq) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *GetImgUrlReq) GetOriginalPic() uint32 {
	if x != nil && x.OriginalPic != nil {
		return *x.OriginalPic
	}
	return 0
}

func (x *GetImgUrlReq) GetRetryReq() uint32 {
	if x != nil && x.RetryReq != nil {
		return *x.RetryReq
	}
	return 0
}

func (x *GetImgUrlReq) GetFileHeight() uint32 {
	if x != nil && x.FileHeight != nil {
		return *x.FileHeight
	}
	return 0
}

func (x *GetImgUrlReq) GetFileWidth() uint32 {
	if x != nil && x.FileWidth != nil {
		return *x.FileWidth
	}
	return 0
}

func (x *GetImgUrlReq) GetPicType() uint32 {
	if x != nil && x.PicType != nil {
		return *x.PicType
	}
	return 0
}

func (x *GetImgUrlReq) GetPicUpTimestamp() uint32 {
	if x != nil && x.PicUpTimestamp != nil {
		return *x.PicUpTimestamp
	}
	return 0
}

func (x *GetImgUrlReq) GetReqTransferType() uint32 {
	if x != nil && x.ReqTransferType != nil {
		return *x.ReqTransferType
	}
	return 0
}

func (x *GetImgUrlReq) GetQqmeetGuildId() uint64 {
	if x != nil && x.QqmeetGuildId != nil {
		return *x.QqmeetGuildId
	}
	return 0
}

func (x *GetImgUrlReq) GetQqmeetChannelId() uint64 {
	if x != nil && x.QqmeetChannelId != nil {
		return *x.QqmeetChannelId
	}
	return 0
}

func (x *GetImgUrlReq) GetDownloadIndex() []byte {
	if x != nil {
		return x.DownloadIndex
	}
	return nil
}

type GetImgUrlRsp struct {
	Fileid           *uint64     `protobuf:"varint,1,opt"`
	FileMd5          []byte      `protobuf:"bytes,2,opt"`
	Result           *uint32     `protobuf:"varint,3,opt"`
	FailMsg          []byte      `protobuf:"bytes,4,opt"`
	ImgInfo          *ImgInfo    `protobuf:"bytes,5,opt"`
	ThumbDownUrl     [][]byte    `protobuf:"bytes,6,rep"`
	OriginalDownUrl  [][]byte    `protobuf:"bytes,7,rep"`
	BigDownUrl       [][]byte    `protobuf:"bytes,8,rep"`
	DownIp           []uint32    `protobuf:"varint,9,rep"`
	DownPort         []uint32    `protobuf:"varint,10,rep"`
	DownDomain       []byte      `protobuf:"bytes,11,opt"`
	ThumbDownPara    []byte      `protobuf:"bytes,12,opt"`
	OriginalDownPara []byte      `protobuf:"bytes,13,opt"`
	BigDownPara      []byte      `protobuf:"bytes,14,opt"`
	FileId           *uint64     `protobuf:"varint,15,opt"`
	AutoDownType     *uint32     `protobuf:"varint,16,opt"`
	OrderDownType    []uint32    `protobuf:"varint,17,rep"`
	BigThumbDownPara []byte      `protobuf:"bytes,19,opt"`
	HttpsUrlFlag     *uint32     `protobuf:"varint,20,opt"`
	DownIp6          []*IPv6Info `protobuf:"bytes,26,rep"`
	ClientIp6        []byte      `protobuf:"bytes,27,opt"`
}

func (x *GetImgUrlRsp) GetFileid() uint64 {
	if x != nil && x.Fileid != nil {
		return *x.Fileid
	}
	return 0
}

func (x *GetImgUrlRsp) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *GetImgUrlRsp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *GetImgUrlRsp) GetFailMsg() []byte {
	if x != nil {
		return x.FailMsg
	}
	return nil
}

func (x *GetImgUrlRsp) GetImgInfo() *ImgInfo {
	if x != nil {
		return x.ImgInfo
	}
	return nil
}

func (x *GetImgUrlRsp) GetThumbDownUrl() [][]byte {
	if x != nil {
		return x.ThumbDownUrl
	}
	return nil
}

func (x *GetImgUrlRsp) GetOriginalDownUrl() [][]byte {
	if x != nil {
		return x.OriginalDownUrl
	}
	return nil
}

func (x *GetImgUrlRsp) GetBigDownUrl() [][]byte {
	if x != nil {
		return x.BigDownUrl
	}
	return nil
}

func (x *GetImgUrlRsp) GetDownIp() []uint32 {
	if x != nil {
		return x.DownIp
	}
	return nil
}

func (x *GetImgUrlRsp) GetDownPort() []uint32 {
	if x != nil {
		return x.DownPort
	}
	return nil
}

func (x *GetImgUrlRsp) GetDownDomain() []byte {
	if x != nil {
		return x.DownDomain
	}
	return nil
}

func (x *GetImgUrlRsp) GetThumbDownPara() []byte {
	if x != nil {
		return x.ThumbDownPara
	}
	return nil
}

func (x *GetImgUrlRsp) GetOriginalDownPara() []byte {
	if x != nil {
		return x.OriginalDownPara
	}
	return nil
}

func (x *GetImgUrlRsp) GetBigDownPara() []byte {
	if x != nil {
		return x.BigDownPara
	}
	return nil
}

func (x *GetImgUrlRsp) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *GetImgUrlRsp) GetAutoDownType() uint32 {
	if x != nil && x.AutoDownType != nil {
		return *x.AutoDownType
	}
	return 0
}

func (x *GetImgUrlRsp) GetOrderDownType() []uint32 {
	if x != nil {
		return x.OrderDownType
	}
	return nil
}

func (x *GetImgUrlRsp) GetBigThumbDownPara() []byte {
	if x != nil {
		return x.BigThumbDownPara
	}
	return nil
}

func (x *GetImgUrlRsp) GetHttpsUrlFlag() uint32 {
	if x != nil && x.HttpsUrlFlag != nil {
		return *x.HttpsUrlFlag
	}
	return 0
}

func (x *GetImgUrlRsp) GetDownIp6() []*IPv6Info {
	if x != nil {
		return x.DownIp6
	}
	return nil
}

func (x *GetImgUrlRsp) GetClientIp6() []byte {
	if x != nil {
		return x.ClientIp6
	}
	return nil
}

type GetPttUrlReq struct {
	GroupCode       *uint64 `protobuf:"varint,1,opt"`
	DstUin          *uint64 `protobuf:"varint,2,opt"`
	Fileid          *uint64 `protobuf:"varint,3,opt"`
	FileMd5         []byte  `protobuf:"bytes,4,opt"`
	ReqTerm         *uint32 `protobuf:"varint,5,opt"`
	ReqPlatformType *uint32 `protobuf:"varint,6,opt"`
	InnerIp         *uint32 `protobuf:"varint,7,opt"`
	BuType          *uint32 `protobuf:"varint,8,opt"`
	BuildVer        []byte  `protobuf:"bytes,9,opt"`
	FileId          *uint64 `protobuf:"varint,10,opt"`
	FileKey         []byte  `protobuf:"bytes,11,opt"`
	Codec           *uint32 `protobuf:"varint,12,opt"`
	BuId            *uint32 `protobuf:"varint,13,opt"`
	ReqTransferType *uint32 `protobuf:"varint,14,opt"`
	IsAuto          *uint32 `protobuf:"varint,15,opt"`
}

func (x *GetPttUrlReq) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *GetPttUrlReq) GetDstUin() uint64 {
	if x != nil && x.DstUin != nil {
		return *x.DstUin
	}
	return 0
}

func (x *GetPttUrlReq) GetFileid() uint64 {
	if x != nil && x.Fileid != nil {
		return *x.Fileid
	}
	return 0
}

func (x *GetPttUrlReq) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *GetPttUrlReq) GetReqTerm() uint32 {
	if x != nil && x.ReqTerm != nil {
		return *x.ReqTerm
	}
	return 0
}

func (x *GetPttUrlReq) GetReqPlatformType() uint32 {
	if x != nil && x.ReqPlatformType != nil {
		return *x.ReqPlatformType
	}
	return 0
}

func (x *GetPttUrlReq) GetInnerIp() uint32 {
	if x != nil && x.InnerIp != nil {
		return *x.InnerIp
	}
	return 0
}

func (x *GetPttUrlReq) GetBuType() uint32 {
	if x != nil && x.BuType != nil {
		return *x.BuType
	}
	return 0
}

func (x *GetPttUrlReq) GetBuildVer() []byte {
	if x != nil {
		return x.BuildVer
	}
	return nil
}

func (x *GetPttUrlReq) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *GetPttUrlReq) GetFileKey() []byte {
	if x != nil {
		return x.FileKey
	}
	return nil
}

func (x *GetPttUrlReq) GetCodec() uint32 {
	if x != nil && x.Codec != nil {
		return *x.Codec
	}
	return 0
}

func (x *GetPttUrlReq) GetBuId() uint32 {
	if x != nil && x.BuId != nil {
		return *x.BuId
	}
	return 0
}

func (x *GetPttUrlReq) GetReqTransferType() uint32 {
	if x != nil && x.ReqTransferType != nil {
		return *x.ReqTransferType
	}
	return 0
}

func (x *GetPttUrlReq) GetIsAuto() uint32 {
	if x != nil && x.IsAuto != nil {
		return *x.IsAuto
	}
	return 0
}

type GetPttUrlRsp struct {
	Fileid       *uint64     `protobuf:"varint,1,opt"`
	FileMd5      []byte      `protobuf:"bytes,2,opt"`
	Result       *uint32     `protobuf:"varint,3,opt"`
	FailMsg      []byte      `protobuf:"bytes,4,opt"`
	DownUrl      [][]byte    `protobuf:"bytes,5,rep"`
	DownIp       []uint32    `protobuf:"varint,6,rep"`
	DownPort     []uint32    `protobuf:"varint,7,rep"`
	DownDomain   []byte      `protobuf:"bytes,8,opt"`
	DownPara     []byte      `protobuf:"bytes,9,opt"`
	FileId       *uint64     `protobuf:"varint,10,opt"`
	TransferType *uint32     `protobuf:"varint,11,opt"`
	AllowRetry   *uint32     `protobuf:"varint,12,opt"`
	DownIp6      []*IPv6Info `protobuf:"bytes,26,rep"`
	ClientIp6    []byte      `protobuf:"bytes,27,opt"`
	Domain       *string     `protobuf:"bytes,28,opt"`
}

func (x *GetPttUrlRsp) GetFileid() uint64 {
	if x != nil && x.Fileid != nil {
		return *x.Fileid
	}
	return 0
}

func (x *GetPttUrlRsp) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *GetPttUrlRsp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *GetPttUrlRsp) GetFailMsg() []byte {
	if x != nil {
		return x.FailMsg
	}
	return nil
}

func (x *GetPttUrlRsp) GetDownUrl() [][]byte {
	if x != nil {
		return x.DownUrl
	}
	return nil
}

func (x *GetPttUrlRsp) GetDownIp() []uint32 {
	if x != nil {
		return x.DownIp
	}
	return nil
}

func (x *GetPttUrlRsp) GetDownPort() []uint32 {
	if x != nil {
		return x.DownPort
	}
	return nil
}

func (x *GetPttUrlRsp) GetDownDomain() []byte {
	if x != nil {
		return x.DownDomain
	}
	return nil
}

func (x *GetPttUrlRsp) GetDownPara() []byte {
	if x != nil {
		return x.DownPara
	}
	return nil
}

func (x *GetPttUrlRsp) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *GetPttUrlRsp) GetTransferType() uint32 {
	if x != nil && x.TransferType != nil {
		return *x.TransferType
	}
	return 0
}

func (x *GetPttUrlRsp) GetAllowRetry() uint32 {
	if x != nil && x.AllowRetry != nil {
		return *x.AllowRetry
	}
	return 0
}

func (x *GetPttUrlRsp) GetDownIp6() []*IPv6Info {
	if x != nil {
		return x.DownIp6
	}
	return nil
}

func (x *GetPttUrlRsp) GetClientIp6() []byte {
	if x != nil {
		return x.ClientIp6
	}
	return nil
}

func (x *GetPttUrlRsp) GetDomain() string {
	if x != nil && x.Domain != nil {
		return *x.Domain
	}
	return ""
}

type IPv6Info struct {
	Ip6  []byte  `protobuf:"bytes,1,opt"`
	Port *uint32 `protobuf:"varint,2,opt"`
}

func (x *IPv6Info) GetIp6() []byte {
	if x != nil {
		return x.Ip6
	}
	return nil
}

func (x *IPv6Info) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

type ImgInfo struct {
	FileMd5    []byte  `protobuf:"bytes,1,opt"`
	FileType   *uint32 `protobuf:"varint,2,opt"`
	FileSize   *uint64 `protobuf:"varint,3,opt"`
	FileWidth  *uint32 `protobuf:"varint,4,opt"`
	FileHeight *uint32 `protobuf:"varint,5,opt"`
}

func (x *ImgInfo) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *ImgInfo) GetFileType() uint32 {
	if x != nil && x.FileType != nil {
		return *x.FileType
	}
	return 0
}

func (x *ImgInfo) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *ImgInfo) GetFileWidth() uint32 {
	if x != nil && x.FileWidth != nil {
		return *x.FileWidth
	}
	return 0
}

func (x *ImgInfo) GetFileHeight() uint32 {
	if x != nil && x.FileHeight != nil {
		return *x.FileHeight
	}
	return 0
}

type PicSize struct {
	Original *uint32 `protobuf:"varint,1,opt"`
	Thumb    *uint32 `protobuf:"varint,2,opt"`
	High     *uint32 `protobuf:"varint,3,opt"`
}

func (x *PicSize) GetOriginal() uint32 {
	if x != nil && x.Original != nil {
		return *x.Original
	}
	return 0
}

func (x *PicSize) GetThumb() uint32 {
	if x != nil && x.Thumb != nil {
		return *x.Thumb
	}
	return 0
}

func (x *PicSize) GetHigh() uint32 {
	if x != nil && x.High != nil {
		return *x.High
	}
	return 0
}

type D388ReqBody struct {
	NetType      *uint32         `protobuf:"varint,1,opt"`
	Subcmd       *uint32         `protobuf:"varint,2,opt"`
	TryupImgReq  []*TryUpImgReq  `protobuf:"bytes,3,rep"`
	GetimgUrlReq []*GetImgUrlReq `protobuf:"bytes,4,rep"`
	TryupPttReq  []*TryUpPttReq  `protobuf:"bytes,5,rep"`
	GetpttUrlReq []*GetPttUrlReq `protobuf:"bytes,6,rep"`
	CommandId    *uint32         `protobuf:"varint,7,opt"`
	DelImgReq    []*DelImgReq    `protobuf:"bytes,8,rep"`
	Extension    []byte          `protobuf:"bytes,1001,opt"`
}

func (x *D388ReqBody) GetNetType() uint32 {
	if x != nil && x.NetType != nil {
		return *x.NetType
	}
	return 0
}

func (x *D388ReqBody) GetSubcmd() uint32 {
	if x != nil && x.Subcmd != nil {
		return *x.Subcmd
	}
	return 0
}

func (x *D388ReqBody) GetTryupImgReq() []*TryUpImgReq {
	if x != nil {
		return x.TryupImgReq
	}
	return nil
}

func (x *D388ReqBody) GetGetimgUrlReq() []*GetImgUrlReq {
	if x != nil {
		return x.GetimgUrlReq
	}
	return nil
}

func (x *D388ReqBody) GetTryupPttReq() []*TryUpPttReq {
	if x != nil {
		return x.TryupPttReq
	}
	return nil
}

func (x *D388ReqBody) GetGetpttUrlReq() []*GetPttUrlReq {
	if x != nil {
		return x.GetpttUrlReq
	}
	return nil
}

func (x *D388ReqBody) GetCommandId() uint32 {
	if x != nil && x.CommandId != nil {
		return *x.CommandId
	}
	return 0
}

func (x *D388ReqBody) GetDelImgReq() []*DelImgReq {
	if x != nil {
		return x.DelImgReq
	}
	return nil
}

func (x *D388ReqBody) GetExtension() []byte {
	if x != nil {
		return x.Extension
	}
	return nil
}

type D388RspBody struct {
	ClientIp     *uint32            `protobuf:"varint,1,opt"`
	Subcmd       *uint32            `protobuf:"varint,2,opt"`
	TryupImgRsp  []*D388TryUpImgRsp `protobuf:"bytes,3,rep"`
	GetimgUrlRsp []*GetImgUrlRsp    `protobuf:"bytes,4,rep"`
	TryupPttRsp  []*TryUpPttRsp     `protobuf:"bytes,5,rep"`
	GetpttUrlRsp []*GetPttUrlRsp    `protobuf:"bytes,6,rep"`
	DelImgRsp    []*DelImgRsp       `protobuf:"bytes,7,rep"`
}

func (x *D388RspBody) GetClientIp() uint32 {
	if x != nil && x.ClientIp != nil {
		return *x.ClientIp
	}
	return 0
}

func (x *D388RspBody) GetSubcmd() uint32 {
	if x != nil && x.Subcmd != nil {
		return *x.Subcmd
	}
	return 0
}

func (x *D388RspBody) GetTryupImgRsp() []*D388TryUpImgRsp {
	if x != nil {
		return x.TryupImgRsp
	}
	return nil
}

func (x *D388RspBody) GetGetimgUrlRsp() []*GetImgUrlRsp {
	if x != nil {
		return x.GetimgUrlRsp
	}
	return nil
}

func (x *D388RspBody) GetTryupPttRsp() []*TryUpPttRsp {
	if x != nil {
		return x.TryupPttRsp
	}
	return nil
}

func (x *D388RspBody) GetGetpttUrlRsp() []*GetPttUrlRsp {
	if x != nil {
		return x.GetpttUrlRsp
	}
	return nil
}

func (x *D388RspBody) GetDelImgRsp() []*DelImgRsp {
	if x != nil {
		return x.DelImgRsp
	}
	return nil
}

type TryUpImgReq struct {
	GroupCode       *uint64 `protobuf:"varint,1,opt"`
	SrcUin          *uint64 `protobuf:"varint,2,opt"`
	FileId          *uint64 `protobuf:"varint,3,opt"`
	FileMd5         []byte  `protobuf:"bytes,4,opt"`
	FileSize        *uint64 `protobuf:"varint,5,opt"`
	FileName        []byte  `protobuf:"bytes,6,opt"`
	SrcTerm         *uint32 `protobuf:"varint,7,opt"`
	PlatformType    *uint32 `protobuf:"varint,8,opt"`
	BuType          *uint32 `protobuf:"varint,9,opt"`
	PicWidth        *uint32 `protobuf:"varint,10,opt"`
	PicHeight       *uint32 `protobuf:"varint,11,opt"`
	PicType         *uint32 `protobuf:"varint,12,opt"`
	BuildVer        []byte  `protobuf:"bytes,13,opt"`
	InnerIp         *uint32 `protobuf:"varint,14,opt"`
	AppPicType      *uint32 `protobuf:"varint,15,opt"`
	OriginalPic     *uint32 `protobuf:"varint,16,opt"`
	FileIndex       []byte  `protobuf:"bytes,17,opt"`
	DstUin          *uint64 `protobuf:"varint,18,opt"`
	SrvUpload       *uint32 `protobuf:"varint,19,opt"`
	TransferUrl     []byte  `protobuf:"bytes,20,opt"`
	QqmeetGuildId   *uint64 `protobuf:"varint,21,opt"`
	QqmeetChannelId *uint64 `protobuf:"varint,22,opt"`
}

func (x *TryUpImgReq) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *TryUpImgReq) GetSrcUin() uint64 {
	if x != nil && x.SrcUin != nil {
		return *x.SrcUin
	}
	return 0
}

func (x *TryUpImgReq) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *TryUpImgReq) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *TryUpImgReq) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *TryUpImgReq) GetFileName() []byte {
	if x != nil {
		return x.FileName
	}
	return nil
}

func (x *TryUpImgReq) GetSrcTerm() uint32 {
	if x != nil && x.SrcTerm != nil {
		return *x.SrcTerm
	}
	return 0
}

func (x *TryUpImgReq) GetPlatformType() uint32 {
	if x != nil && x.PlatformType != nil {
		return *x.PlatformType
	}
	return 0
}

func (x *TryUpImgReq) GetBuType() uint32 {
	if x != nil && x.BuType != nil {
		return *x.BuType
	}
	return 0
}

func (x *TryUpImgReq) GetPicWidth() uint32 {
	if x != nil && x.PicWidth != nil {
		return *x.PicWidth
	}
	return 0
}

func (x *TryUpImgReq) GetPicHeight() uint32 {
	if x != nil && x.PicHeight != nil {
		return *x.PicHeight
	}
	return 0
}

func (x *TryUpImgReq) GetPicType() uint32 {
	if x != nil && x.PicType != nil {
		return *x.PicType
	}
	return 0
}

func (x *TryUpImgReq) GetBuildVer() []byte {
	if x != nil {
		return x.BuildVer
	}
	return nil
}

func (x *TryUpImgReq) GetInnerIp() uint32 {
	if x != nil && x.InnerIp != nil {
		return *x.InnerIp
	}
	return 0
}

func (x *TryUpImgReq) GetAppPicType() uint32 {
	if x != nil && x.AppPicType != nil {
		return *x.AppPicType
	}
	return 0
}

func (x *TryUpImgReq) GetOriginalPic() uint32 {
	if x != nil && x.OriginalPic != nil {
		return *x.OriginalPic
	}
	return 0
}

func (x *TryUpImgReq) GetFileIndex() []byte {
	if x != nil {
		return x.FileIndex
	}
	return nil
}

func (x *TryUpImgReq) GetDstUin() uint64 {
	if x != nil && x.DstUin != nil {
		return *x.DstUin
	}
	return 0
}

func (x *TryUpImgReq) GetSrvUpload() uint32 {
	if x != nil && x.SrvUpload != nil {
		return *x.SrvUpload
	}
	return 0
}

func (x *TryUpImgReq) GetTransferUrl() []byte {
	if x != nil {
		return x.TransferUrl
	}
	return nil
}

func (x *TryUpImgReq) GetQqmeetGuildId() uint64 {
	if x != nil && x.QqmeetGuildId != nil {
		return *x.QqmeetGuildId
	}
	return 0
}

func (x *TryUpImgReq) GetQqmeetChannelId() uint64 {
	if x != nil && x.QqmeetChannelId != nil {
		return *x.QqmeetChannelId
	}
	return 0
}

type D388TryUpImgRsp struct {
	FileId        *uint64         `protobuf:"varint,1,opt"`
	Result        *uint32         `protobuf:"varint,2,opt"`
	FailMsg       []byte          `protobuf:"bytes,3,opt"`
	FileExit      *bool           `protobuf:"varint,4,opt"`
	ImgInfo       *ImgInfo        `protobuf:"bytes,5,opt"`
	UpIp          []uint32        `protobuf:"varint,6,rep"`
	UpPort        []uint32        `protobuf:"varint,7,rep"`
	UpUkey        []byte          `protobuf:"bytes,8,opt"`
	Fileid        *uint64         `protobuf:"varint,9,opt"`
	UpOffset      *uint64         `protobuf:"varint,10,opt"`
	BlockSize     *uint64         `protobuf:"varint,11,opt"`
	NewBigChan    *bool           `protobuf:"varint,12,opt"`
	UpIp6         []*IPv6Info     `protobuf:"bytes,26,rep"`
	ClientIp6     []byte          `protobuf:"bytes,27,opt"`
	DownloadIndex []byte          `protobuf:"bytes,28,opt"`
	Info4Busi     *TryUpInfo4Busi `protobuf:"bytes,1001,opt"`
}

func (x *D388TryUpImgRsp) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *D388TryUpImgRsp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *D388TryUpImgRsp) GetFailMsg() []byte {
	if x != nil {
		return x.FailMsg
	}
	return nil
}

func (x *D388TryUpImgRsp) GetFileExit() bool {
	if x != nil && x.FileExit != nil {
		return *x.FileExit
	}
	return false
}

func (x *D388TryUpImgRsp) GetImgInfo() *ImgInfo {
	if x != nil {
		return x.ImgInfo
	}
	return nil
}

func (x *D388TryUpImgRsp) GetUpIp() []uint32 {
	if x != nil {
		return x.UpIp
	}
	return nil
}

func (x *D388TryUpImgRsp) GetUpPort() []uint32 {
	if x != nil {
		return x.UpPort
	}
	return nil
}

func (x *D388TryUpImgRsp) GetUpUkey() []byte {
	if x != nil {
		return x.UpUkey
	}
	return nil
}

func (x *D388TryUpImgRsp) GetFileid() uint64 {
	if x != nil && x.Fileid != nil {
		return *x.Fileid
	}
	return 0
}

func (x *D388TryUpImgRsp) GetUpOffset() uint64 {
	if x != nil && x.UpOffset != nil {
		return *x.UpOffset
	}
	return 0
}

func (x *D388TryUpImgRsp) GetBlockSize() uint64 {
	if x != nil && x.BlockSize != nil {
		return *x.BlockSize
	}
	return 0
}

func (x *D388TryUpImgRsp) GetNewBigChan() bool {
	if x != nil && x.NewBigChan != nil {
		return *x.NewBigChan
	}
	return false
}

func (x *D388TryUpImgRsp) GetUpIp6() []*IPv6Info {
	if x != nil {
		return x.UpIp6
	}
	return nil
}

func (x *D388TryUpImgRsp) GetClientIp6() []byte {
	if x != nil {
		return x.ClientIp6
	}
	return nil
}

func (x *D388TryUpImgRsp) GetDownloadIndex() []byte {
	if x != nil {
		return x.DownloadIndex
	}
	return nil
}

func (x *D388TryUpImgRsp) GetInfo4Busi() *TryUpInfo4Busi {
	if x != nil {
		return x.Info4Busi
	}
	return nil
}

type TryUpInfo4Busi struct {
	DownDomain      []byte `protobuf:"bytes,1,opt"`
	ThumbDownUrl    []byte `protobuf:"bytes,2,opt"`
	OriginalDownUrl []byte `protobuf:"bytes,3,opt"`
	BigDownUrl      []byte `protobuf:"bytes,4,opt"`
	FileResid       []byte `protobuf:"bytes,5,opt"`
}

func (x *TryUpInfo4Busi) GetDownDomain() []byte {
	if x != nil {
		return x.DownDomain
	}
	return nil
}

func (x *TryUpInfo4Busi) GetThumbDownUrl() []byte {
	if x != nil {
		return x.ThumbDownUrl
	}
	return nil
}

func (x *TryUpInfo4Busi) GetOriginalDownUrl() []byte {
	if x != nil {
		return x.OriginalDownUrl
	}
	return nil
}

func (x *TryUpInfo4Busi) GetBigDownUrl() []byte {
	if x != nil {
		return x.BigDownUrl
	}
	return nil
}

func (x *TryUpInfo4Busi) GetFileResid() []byte {
	if x != nil {
		return x.FileResid
	}
	return nil
}

type TryUpPttReq struct {
	GroupCode    *uint64 `protobuf:"varint,1,opt"`
	SrcUin       *uint64 `protobuf:"varint,2,opt"`
	FileId       *uint64 `protobuf:"varint,3,opt"`
	FileMd5      []byte  `protobuf:"bytes,4,opt"`
	FileSize     *uint64 `protobuf:"varint,5,opt"`
	FileName     []byte  `protobuf:"bytes,6,opt"`
	SrcTerm      *uint32 `protobuf:"varint,7,opt"`
	PlatformType *uint32 `protobuf:"varint,8,opt"`
	BuType       *uint32 `protobuf:"varint,9,opt"`
	BuildVer     []byte  `protobuf:"bytes,10,opt"`
	InnerIp      *uint32 `protobuf:"varint,11,opt"`
	VoiceLength  *uint32 `protobuf:"varint,12,opt"`
	NewUpChan    *bool   `protobuf:"varint,13,opt"`
	Codec        *uint32 `protobuf:"varint,14,opt"`
	VoiceType    *uint32 `protobuf:"varint,15,opt"`
	BuId         *uint32 `protobuf:"varint,16,opt"`
}

func (x *TryUpPttReq) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *TryUpPttReq) GetSrcUin() uint64 {
	if x != nil && x.SrcUin != nil {
		return *x.SrcUin
	}
	return 0
}

func (x *TryUpPttReq) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *TryUpPttReq) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *TryUpPttReq) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *TryUpPttReq) GetFileName() []byte {
	if x != nil {
		return x.FileName
	}
	return nil
}

func (x *TryUpPttReq) GetSrcTerm() uint32 {
	if x != nil && x.SrcTerm != nil {
		return *x.SrcTerm
	}
	return 0
}

func (x *TryUpPttReq) GetPlatformType() uint32 {
	if x != nil && x.PlatformType != nil {
		return *x.PlatformType
	}
	return 0
}

func (x *TryUpPttReq) GetBuType() uint32 {
	if x != nil && x.BuType != nil {
		return *x.BuType
	}
	return 0
}

func (x *TryUpPttReq) GetBuildVer() []byte {
	if x != nil {
		return x.BuildVer
	}
	return nil
}

func (x *TryUpPttReq) GetInnerIp() uint32 {
	if x != nil && x.InnerIp != nil {
		return *x.InnerIp
	}
	return 0
}

func (x *TryUpPttReq) GetVoiceLength() uint32 {
	if x != nil && x.VoiceLength != nil {
		return *x.VoiceLength
	}
	return 0
}

func (x *TryUpPttReq) GetNewUpChan() bool {
	if x != nil && x.NewUpChan != nil {
		return *x.NewUpChan
	}
	return false
}

func (x *TryUpPttReq) GetCodec() uint32 {
	if x != nil && x.Codec != nil {
		return *x.Codec
	}
	return 0
}

func (x *TryUpPttReq) GetVoiceType() uint32 {
	if x != nil && x.VoiceType != nil {
		return *x.VoiceType
	}
	return 0
}

func (x *TryUpPttReq) GetBuId() uint32 {
	if x != nil && x.BuId != nil {
		return *x.BuId
	}
	return 0
}

type TryUpPttRsp struct {
	FileId      *uint64     `protobuf:"varint,1,opt"`
	Result      *uint32     `protobuf:"varint,2,opt"`
	FailMsg     []byte      `protobuf:"bytes,3,opt"`
	FileExit    *bool       `protobuf:"varint,4,opt"`
	UpIp        []uint32    `protobuf:"varint,5,rep"`
	UpPort      []uint32    `protobuf:"varint,6,rep"`
	UpUkey      []byte      `protobuf:"bytes,7,opt"`
	Fileid      *uint64     `protobuf:"varint,8,opt"`
	UpOffset    *uint64     `protobuf:"varint,9,opt"`
	BlockSize   *uint64     `protobuf:"varint,10,opt"`
	FileKey     []byte      `protobuf:"bytes,11,opt"`
	ChannelType *uint32     `protobuf:"varint,12,opt"`
	UpIp6       []*IPv6Info `protobuf:"bytes,26,rep"`
	ClientIp6   []byte      `protobuf:"bytes,27,opt"`
}

func (x *TryUpPttRsp) GetFileId() uint64 {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return 0
}

func (x *TryUpPttRsp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *TryUpPttRsp) GetFailMsg() []byte {
	if x != nil {
		return x.FailMsg
	}
	return nil
}

func (x *TryUpPttRsp) GetFileExit() bool {
	if x != nil && x.FileExit != nil {
		return *x.FileExit
	}
	return false
}

func (x *TryUpPttRsp) GetUpIp() []uint32 {
	if x != nil {
		return x.UpIp
	}
	return nil
}

func (x *TryUpPttRsp) GetUpPort() []uint32 {
	if x != nil {
		return x.UpPort
	}
	return nil
}

func (x *TryUpPttRsp) GetUpUkey() []byte {
	if x != nil {
		return x.UpUkey
	}
	return nil
}

func (x *TryUpPttRsp) GetFileid() uint64 {
	if x != nil && x.Fileid != nil {
		return *x.Fileid
	}
	return 0
}

func (x *TryUpPttRsp) GetUpOffset() uint64 {
	if x != nil && x.UpOffset != nil {
		return *x.UpOffset
	}
	return 0
}

func (x *TryUpPttRsp) GetBlockSize() uint64 {
	if x != nil && x.BlockSize != nil {
		return *x.BlockSize
	}
	return 0
}

func (x *TryUpPttRsp) GetFileKey() []byte {
	if x != nil {
		return x.FileKey
	}
	return nil
}

func (x *TryUpPttRsp) GetChannelType() uint32 {
	if x != nil && x.ChannelType != nil {
		return *x.ChannelType
	}
	return 0
}

func (x *TryUpPttRsp) GetUpIp6() []*IPv6Info {
	if x != nil {
		return x.UpIp6
	}
	return nil
}

func (x *TryUpPttRsp) GetClientIp6() []byte {
	if x != nil {
		return x.ClientIp6
	}
	return nil
}
