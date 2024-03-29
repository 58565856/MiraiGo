// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/pttcenter/shortvideo.proto

package pttcenter

type ShortVideoReqBody struct {
	Cmd                      int32                     `protobuf:"varint,1,opt"`
	Seq                      int32                     `protobuf:"varint,2,opt"`
	PttShortVideoUploadReq   *ShortVideoUploadReq      `protobuf:"bytes,3,opt"`
	PttShortVideoDownloadReq *ShortVideoDownloadReq    `protobuf:"bytes,4,opt"`
	ExtensionReq             []*ShortVideoExtensionReq `protobuf:"bytes,100,rep"`
}

type ShortVideoRspBody struct {
	Cmd                      int32                  `protobuf:"varint,1,opt"`
	Seq                      int32                  `protobuf:"varint,2,opt"`
	PttShortVideoUploadRsp   *ShortVideoUploadRsp   `protobuf:"bytes,3,opt"`
	PttShortVideoDownloadRsp *ShortVideoDownloadRsp `protobuf:"bytes,4,opt"`
}

type ShortVideoUploadReq struct {
	FromUin          int64               `protobuf:"varint,1,opt"`
	ToUin            int64               `protobuf:"varint,2,opt"`
	ChatType         int32               `protobuf:"varint,3,opt"`
	ClientType       int32               `protobuf:"varint,4,opt"`
	Info             *ShortVideoFileInfo `protobuf:"bytes,5,opt"`
	GroupCode        int64               `protobuf:"varint,6,opt"`
	AgentType        int32               `protobuf:"varint,7,opt"`
	BusinessType     int32               `protobuf:"varint,8,opt"`
	SupportLargeSize int32               `protobuf:"varint,20,opt"`
}

type ShortVideoDownloadReq struct {
	FromUin      int64  `protobuf:"varint,1,opt"`
	ToUin        int64  `protobuf:"varint,2,opt"`
	ChatType     int32  `protobuf:"varint,3,opt"`
	ClientType   int32  `protobuf:"varint,4,opt"`
	FileId       string `protobuf:"bytes,5,opt"`
	GroupCode    int64  `protobuf:"varint,6,opt"`
	AgentType    int32  `protobuf:"varint,7,opt"`
	FileMd5      []byte `protobuf:"bytes,8,opt"`
	BusinessType int32  `protobuf:"varint,9,opt"`
	FileType     int32  `protobuf:"varint,10,opt"`
	DownType     int32  `protobuf:"varint,11,opt"`
	SceneType    int32  `protobuf:"varint,12,opt"`
}

type ShortVideoDownloadRsp struct {
	RetCode           int32               `protobuf:"varint,1,opt"`
	RetMsg            string              `protobuf:"bytes,2,opt"`
	SameAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,3,rep"`
	DiffAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,4,rep"`
	DownloadKey       []byte              `protobuf:"bytes,5,opt"`
	FileMd5           []byte              `protobuf:"bytes,6,opt"`
	SameAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,7,rep"`
	DiffAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,8,rep"`
	DownloadAddr      *ShortVideoAddr     `protobuf:"bytes,9,opt"`
	EncryptKey        []byte              `protobuf:"bytes,10,opt"`
}

type ShortVideoUploadRsp struct {
	RetCode           int32               `protobuf:"varint,1,opt"`
	RetMsg            string              `protobuf:"bytes,2,opt"`
	SameAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,3,rep"`
	DiffAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,4,rep"`
	FileId            string              `protobuf:"bytes,5,opt"`
	UKey              []byte              `protobuf:"bytes,6,opt"`
	FileExists        int32               `protobuf:"varint,7,opt"`
	SameAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,8,rep"`
	DiffAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,9,rep"`
	DataHole          []*DataHole         `protobuf:"bytes,10,rep"`
}

type ShortVideoFileInfo struct {
	FileName      string `protobuf:"bytes,1,opt"`
	FileMd5       []byte `protobuf:"bytes,2,opt"`
	ThumbFileMd5  []byte `protobuf:"bytes,3,opt"`
	FileSize      int64  `protobuf:"varint,4,opt"`
	FileResLength int32  `protobuf:"varint,5,opt"`
	FileResWidth  int32  `protobuf:"varint,6,opt"`
	FileFormat    int32  `protobuf:"varint,7,opt"`
	FileTime      int32  `protobuf:"varint,8,opt"`
	ThumbFileSize int64  `protobuf:"varint,9,opt"`
}

type DataHole struct {
	Begin int64 `protobuf:"varint,1,opt"`
	End   int64 `protobuf:"varint,2,opt"`
}

type ShortVideoIpList struct {
	Ip   int32 `protobuf:"varint,1,opt"`
	Port int32 `protobuf:"varint,2,opt"`
}

type ShortVideoAddr struct {
	Host    []string `protobuf:"bytes,10,rep"`
	UrlArgs string   `protobuf:"bytes,11,opt"` //repeated string domain = 13;
}

type ShortVideoExtensionReq struct {
	SubBusiType int32 `protobuf:"varint,1,opt"`
	UserCnt     int32 `protobuf:"varint,2,opt"`
}
