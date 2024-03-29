// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0x6d8.proto

package oidb

type D6D8ReqBody struct {
	FileInfoReq       *GetFileInfoReqBody  `protobuf:"bytes,1,opt"`
	FileListInfoReq   *GetFileListReqBody  `protobuf:"bytes,2,opt"`
	GroupFileCountReq *GetFileCountReqBody `protobuf:"bytes,3,opt"`
	GroupSpaceReq     *GetSpaceReqBody     `protobuf:"bytes,4,opt"`
}

type D6D8RspBody struct {
	FileInfoRsp     *GetFileInfoRspBody  `protobuf:"bytes,1,opt"`
	FileListInfoRsp *GetFileListRspBody  `protobuf:"bytes,2,opt"`
	FileCountRsp    *GetFileCountRspBody `protobuf:"bytes,3,opt"`
	GroupSpaceRsp   *GetSpaceRspBody     `protobuf:"bytes,4,opt"`
}

type GetFileInfoReqBody struct {
	GroupCode *uint64 `protobuf:"varint,1,opt"`
	AppId     *uint32 `protobuf:"varint,2,opt"`
	BusId     *uint32 `protobuf:"varint,3,opt"`
	FileId    *string `protobuf:"bytes,4,opt"`
	FieldFlag *uint32 `protobuf:"varint,5,opt"`
}

func (x *GetFileInfoReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *GetFileInfoReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *GetFileInfoReqBody) GetBusId() uint32 {
	if x != nil && x.BusId != nil {
		return *x.BusId
	}
	return 0
}

func (x *GetFileInfoReqBody) GetFileId() string {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return ""
}

func (x *GetFileInfoReqBody) GetFieldFlag() uint32 {
	if x != nil && x.FieldFlag != nil {
		return *x.FieldFlag
	}
	return 0
}

type GetFileInfoRspBody struct {
	RetCode       *int32         `protobuf:"varint,1,opt"`
	RetMsg        *string        `protobuf:"bytes,2,opt"`
	ClientWording *string        `protobuf:"bytes,3,opt"`
	FileInfo      *GroupFileInfo `protobuf:"bytes,4,opt"`
}

func (x *GetFileInfoRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *GetFileInfoRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *GetFileInfoRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

type GetFileListRspBody struct {
	RetCode       *int32                     `protobuf:"varint,1,opt"`
	RetMsg        *string                    `protobuf:"bytes,2,opt"`
	ClientWording *string                    `protobuf:"bytes,3,opt"`
	IsEnd         *bool                      `protobuf:"varint,4,opt"`
	ItemList      []*GetFileListRspBody_Item `protobuf:"bytes,5,rep"`
	MaxTimestamp  *FileTimeStamp             `protobuf:"bytes,6,opt"`
	AllFileCount  *uint32                    `protobuf:"varint,7,opt"`
	FilterCode    *uint32                    `protobuf:"varint,8,opt"`
	SafeCheckFlag *bool                      `protobuf:"varint,11,opt"`
	SafeCheckRes  *uint32                    `protobuf:"varint,12,opt"`
	NextIndex     *uint32                    `protobuf:"varint,13,opt"`
	Context       []byte                     `protobuf:"bytes,14,opt"`
	Role          *uint32                    `protobuf:"varint,15,opt"`
	OpenFlag      *uint32                    `protobuf:"varint,16,opt"`
}

func (x *GetFileListRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *GetFileListRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *GetFileListRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *GetFileListRspBody) GetIsEnd() bool {
	if x != nil && x.IsEnd != nil {
		return *x.IsEnd
	}
	return false
}

func (x *GetFileListRspBody) GetAllFileCount() uint32 {
	if x != nil && x.AllFileCount != nil {
		return *x.AllFileCount
	}
	return 0
}

func (x *GetFileListRspBody) GetFilterCode() uint32 {
	if x != nil && x.FilterCode != nil {
		return *x.FilterCode
	}
	return 0
}

func (x *GetFileListRspBody) GetSafeCheckFlag() bool {
	if x != nil && x.SafeCheckFlag != nil {
		return *x.SafeCheckFlag
	}
	return false
}

func (x *GetFileListRspBody) GetSafeCheckRes() uint32 {
	if x != nil && x.SafeCheckRes != nil {
		return *x.SafeCheckRes
	}
	return 0
}

func (x *GetFileListRspBody) GetNextIndex() uint32 {
	if x != nil && x.NextIndex != nil {
		return *x.NextIndex
	}
	return 0
}

func (x *GetFileListRspBody) GetRole() uint32 {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return 0
}

func (x *GetFileListRspBody) GetOpenFlag() uint32 {
	if x != nil && x.OpenFlag != nil {
		return *x.OpenFlag
	}
	return 0
}

type GroupFileInfo struct {
	FileId         *string `protobuf:"bytes,1,opt"`
	FileName       *string `protobuf:"bytes,2,opt"`
	FileSize       *uint64 `protobuf:"varint,3,opt"`
	BusId          *uint32 `protobuf:"varint,4,opt"`
	UploadedSize   *uint64 `protobuf:"varint,5,opt"`
	UploadTime     *uint32 `protobuf:"varint,6,opt"`
	DeadTime       *uint32 `protobuf:"varint,7,opt"`
	ModifyTime     *uint32 `protobuf:"varint,8,opt"`
	DownloadTimes  *uint32 `protobuf:"varint,9,opt"`
	Sha            []byte  `protobuf:"bytes,10,opt"`
	Sha3           []byte  `protobuf:"bytes,11,opt"`
	Md5            []byte  `protobuf:"bytes,12,opt"`
	LocalPath      *string `protobuf:"bytes,13,opt"`
	UploaderName   *string `protobuf:"bytes,14,opt"`
	UploaderUin    *uint64 `protobuf:"varint,15,opt"`
	ParentFolderId *string `protobuf:"bytes,16,opt"`
}

func (x *GroupFileInfo) GetFileId() string {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return ""
}

func (x *GroupFileInfo) GetFileName() string {
	if x != nil && x.FileName != nil {
		return *x.FileName
	}
	return ""
}

func (x *GroupFileInfo) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *GroupFileInfo) GetBusId() uint32 {
	if x != nil && x.BusId != nil {
		return *x.BusId
	}
	return 0
}

func (x *GroupFileInfo) GetUploadedSize() uint64 {
	if x != nil && x.UploadedSize != nil {
		return *x.UploadedSize
	}
	return 0
}

func (x *GroupFileInfo) GetUploadTime() uint32 {
	if x != nil && x.UploadTime != nil {
		return *x.UploadTime
	}
	return 0
}

func (x *GroupFileInfo) GetDeadTime() uint32 {
	if x != nil && x.DeadTime != nil {
		return *x.DeadTime
	}
	return 0
}

func (x *GroupFileInfo) GetModifyTime() uint32 {
	if x != nil && x.ModifyTime != nil {
		return *x.ModifyTime
	}
	return 0
}

func (x *GroupFileInfo) GetDownloadTimes() uint32 {
	if x != nil && x.DownloadTimes != nil {
		return *x.DownloadTimes
	}
	return 0
}

func (x *GroupFileInfo) GetLocalPath() string {
	if x != nil && x.LocalPath != nil {
		return *x.LocalPath
	}
	return ""
}

func (x *GroupFileInfo) GetUploaderName() string {
	if x != nil && x.UploaderName != nil {
		return *x.UploaderName
	}
	return ""
}

func (x *GroupFileInfo) GetUploaderUin() uint64 {
	if x != nil && x.UploaderUin != nil {
		return *x.UploaderUin
	}
	return 0
}

func (x *GroupFileInfo) GetParentFolderId() string {
	if x != nil && x.ParentFolderId != nil {
		return *x.ParentFolderId
	}
	return ""
}

type GroupFolderInfo struct {
	FolderId       *string `protobuf:"bytes,1,opt"`
	ParentFolderId *string `protobuf:"bytes,2,opt"`
	FolderName     *string `protobuf:"bytes,3,opt"`
	CreateTime     *uint32 `protobuf:"varint,4,opt"`
	ModifyTime     *uint32 `protobuf:"varint,5,opt"`
	CreateUin      *uint64 `protobuf:"varint,6,opt"`
	CreatorName    *string `protobuf:"bytes,7,opt"`
	TotalFileCount *uint32 `protobuf:"varint,8,opt"`
}

func (x *GroupFolderInfo) GetFolderId() string {
	if x != nil && x.FolderId != nil {
		return *x.FolderId
	}
	return ""
}

func (x *GroupFolderInfo) GetParentFolderId() string {
	if x != nil && x.ParentFolderId != nil {
		return *x.ParentFolderId
	}
	return ""
}

func (x *GroupFolderInfo) GetFolderName() string {
	if x != nil && x.FolderName != nil {
		return *x.FolderName
	}
	return ""
}

func (x *GroupFolderInfo) GetCreateTime() uint32 {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return 0
}

func (x *GroupFolderInfo) GetModifyTime() uint32 {
	if x != nil && x.ModifyTime != nil {
		return *x.ModifyTime
	}
	return 0
}

func (x *GroupFolderInfo) GetCreateUin() uint64 {
	if x != nil && x.CreateUin != nil {
		return *x.CreateUin
	}
	return 0
}

func (x *GroupFolderInfo) GetCreatorName() string {
	if x != nil && x.CreatorName != nil {
		return *x.CreatorName
	}
	return ""
}

func (x *GroupFolderInfo) GetTotalFileCount() uint32 {
	if x != nil && x.TotalFileCount != nil {
		return *x.TotalFileCount
	}
	return 0
}

type GetFileListReqBody struct {
	GroupCode      *uint64        `protobuf:"varint,1,opt"`
	AppId          *uint32        `protobuf:"varint,2,opt"`
	FolderId       *string        `protobuf:"bytes,3,opt"`
	StartTimestamp *FileTimeStamp `protobuf:"bytes,4,opt"`
	FileCount      *uint32        `protobuf:"varint,5,opt"`
	MaxTimestamp   *FileTimeStamp `protobuf:"bytes,6,opt"`
	AllFileCount   *uint32        `protobuf:"varint,7,opt"`
	ReqFrom        *uint32        `protobuf:"varint,8,opt"`
	SortBy         *uint32        `protobuf:"varint,9,opt"`
	FilterCode     *uint32        `protobuf:"varint,10,opt"`
	Uin            *uint64        `protobuf:"varint,11,opt"`
	FieldFlag      *uint32        `protobuf:"varint,12,opt"`
	StartIndex     *uint32        `protobuf:"varint,13,opt"`
	Context        []byte         `protobuf:"bytes,14,opt"`
	ClientVersion  *uint32        `protobuf:"varint,15,opt"`
}

func (x *GetFileListReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *GetFileListReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *GetFileListReqBody) GetFolderId() string {
	if x != nil && x.FolderId != nil {
		return *x.FolderId
	}
	return ""
}

func (x *GetFileListReqBody) GetFileCount() uint32 {
	if x != nil && x.FileCount != nil {
		return *x.FileCount
	}
	return 0
}

func (x *GetFileListReqBody) GetAllFileCount() uint32 {
	if x != nil && x.AllFileCount != nil {
		return *x.AllFileCount
	}
	return 0
}

func (x *GetFileListReqBody) GetReqFrom() uint32 {
	if x != nil && x.ReqFrom != nil {
		return *x.ReqFrom
	}
	return 0
}

func (x *GetFileListReqBody) GetSortBy() uint32 {
	if x != nil && x.SortBy != nil {
		return *x.SortBy
	}
	return 0
}

func (x *GetFileListReqBody) GetFilterCode() uint32 {
	if x != nil && x.FilterCode != nil {
		return *x.FilterCode
	}
	return 0
}

func (x *GetFileListReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *GetFileListReqBody) GetFieldFlag() uint32 {
	if x != nil && x.FieldFlag != nil {
		return *x.FieldFlag
	}
	return 0
}

func (x *GetFileListReqBody) GetStartIndex() uint32 {
	if x != nil && x.StartIndex != nil {
		return *x.StartIndex
	}
	return 0
}

func (x *GetFileListReqBody) GetClientVersion() uint32 {
	if x != nil && x.ClientVersion != nil {
		return *x.ClientVersion
	}
	return 0
}

type GetFileCountReqBody struct {
	GroupCode *uint64 `protobuf:"varint,1,opt"`
	AppId     *uint32 `protobuf:"varint,2,opt"`
	BusId     *uint32 `protobuf:"varint,3,opt"`
}

func (x *GetFileCountReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *GetFileCountReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *GetFileCountReqBody) GetBusId() uint32 {
	if x != nil && x.BusId != nil {
		return *x.BusId
	}
	return 0
}

type GetSpaceReqBody struct {
	GroupCode *uint64 `protobuf:"varint,1,opt"`
	AppId     *uint32 `protobuf:"varint,2,opt"`
}

func (x *GetSpaceReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *GetSpaceReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

type GetFileCountRspBody struct {
	RetCode       *int32  `protobuf:"varint,1,opt"`
	RetMsg        *string `protobuf:"bytes,2,opt"`
	ClientWording *string `protobuf:"bytes,3,opt"`
	AllFileCount  *uint32 `protobuf:"varint,4,opt"`
	FileTooMany   *bool   `protobuf:"varint,5,opt"`
	LimitCount    *uint32 `protobuf:"varint,6,opt"`
	IsFull        *bool   `protobuf:"varint,7,opt"`
}

func (x *GetFileCountRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *GetFileCountRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *GetFileCountRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *GetFileCountRspBody) GetAllFileCount() uint32 {
	if x != nil && x.AllFileCount != nil {
		return *x.AllFileCount
	}
	return 0
}

func (x *GetFileCountRspBody) GetFileTooMany() bool {
	if x != nil && x.FileTooMany != nil {
		return *x.FileTooMany
	}
	return false
}

func (x *GetFileCountRspBody) GetLimitCount() uint32 {
	if x != nil && x.LimitCount != nil {
		return *x.LimitCount
	}
	return 0
}

func (x *GetFileCountRspBody) GetIsFull() bool {
	if x != nil && x.IsFull != nil {
		return *x.IsFull
	}
	return false
}

type GetSpaceRspBody struct {
	RetCode       *int32  `protobuf:"varint,1,opt"`
	RetMsg        *string `protobuf:"bytes,2,opt"`
	ClientWording *string `protobuf:"bytes,3,opt"`
	TotalSpace    *uint64 `protobuf:"varint,4,opt"`
	UsedSpace     *uint64 `protobuf:"varint,5,opt"`
}

func (x *GetSpaceRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *GetSpaceRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *GetSpaceRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *GetSpaceRspBody) GetTotalSpace() uint64 {
	if x != nil && x.TotalSpace != nil {
		return *x.TotalSpace
	}
	return 0
}

func (x *GetSpaceRspBody) GetUsedSpace() uint64 {
	if x != nil && x.UsedSpace != nil {
		return *x.UsedSpace
	}
	return 0
}

type FileTimeStamp struct {
	UploadTime *uint32 `protobuf:"varint,1,opt"`
	FileId     *string `protobuf:"bytes,2,opt"`
}

func (x *FileTimeStamp) GetUploadTime() uint32 {
	if x != nil && x.UploadTime != nil {
		return *x.UploadTime
	}
	return 0
}

func (x *FileTimeStamp) GetFileId() string {
	if x != nil && x.FileId != nil {
		return *x.FileId
	}
	return ""
}

type GetFileListRspBody_Item struct {
	Type       *uint32          `protobuf:"varint,1,opt"`
	FolderInfo *GroupFolderInfo `protobuf:"bytes,2,opt"`
	FileInfo   *GroupFileInfo   `protobuf:"bytes,3,opt"`
}

func (x *GetFileListRspBody_Item) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}
