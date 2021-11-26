// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: app.proto

package qweb

type GetAppInfoByIdReq struct {
	//CommonExt ExtInfo = 1;
	AppId           string `protobuf:"bytes,2,opt"`
	NeedVersionInfo int32  `protobuf:"varint,3,opt"`
}

func (x *GetAppInfoByIdReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *GetAppInfoByIdReq) GetNeedVersionInfo() int32 {
	if x != nil {
		return x.NeedVersionInfo
	}
	return 0
}

type GetAppInfoByIdRsp struct {
	AppInfo *ApiAppInfo `protobuf:"bytes,2,opt"`
}

func (x *GetAppInfoByIdRsp) GetAppInfo() *ApiAppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

type ApiAppInfo struct {
	AppId       string `protobuf:"bytes,1,opt"`
	AppName     string `protobuf:"bytes,2,opt"`
	Icon        string `protobuf:"bytes,3,opt"`
	DownloadUrl string `protobuf:"bytes,4,opt"`
	Version     string `protobuf:"bytes,5,opt"`
	Desc        string `protobuf:"bytes,6,opt"`
	// pub accts = 7;
	Type               int32            `protobuf:"varint,8,opt"`
	BaseLibMiniVersion string           `protobuf:"bytes,9,opt"`
	SubPkgs            []*AppSubPkgInfo `protobuf:"bytes,10,rep"`
	// first = 11;
	Domain *DomainConfig `protobuf:"bytes,12,opt"`
}

func (x *ApiAppInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ApiAppInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *ApiAppInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ApiAppInfo) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *ApiAppInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ApiAppInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ApiAppInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ApiAppInfo) GetBaseLibMiniVersion() string {
	if x != nil {
		return x.BaseLibMiniVersion
	}
	return ""
}

func (x *ApiAppInfo) GetSubPkgs() []*AppSubPkgInfo {
	if x != nil {
		return x.SubPkgs
	}
	return nil
}

func (x *ApiAppInfo) GetDomain() *DomainConfig {
	if x != nil {
		return x.Domain
	}
	return nil
}

type AppSubPkgInfo struct {
	SubPkgName  string `protobuf:"bytes,1,opt"`
	DownloadUrl string `protobuf:"bytes,2,opt"`
	Independent int32  `protobuf:"varint,3,opt"`
	FileSize    int32  `protobuf:"varint,4,opt"`
}

func (x *AppSubPkgInfo) GetSubPkgName() string {
	if x != nil {
		return x.SubPkgName
	}
	return ""
}

func (x *AppSubPkgInfo) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *AppSubPkgInfo) GetIndependent() int32 {
	if x != nil {
		return x.Independent
	}
	return 0
}

func (x *AppSubPkgInfo) GetFileSize() int32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type DomainConfig struct {
	RequestDomain      []string `protobuf:"bytes,1,rep"`
	SocketDomain       []string `protobuf:"bytes,2,rep"`
	UploadFileDomain   []string `protobuf:"bytes,3,rep"`
	DownloadFileDomain []string `protobuf:"bytes,4,rep"`
	BusinessDomain     []string `protobuf:"bytes,5,rep"`
	UdpIpList          []string `protobuf:"bytes,6,rep"`
}

func (x *DomainConfig) GetRequestDomain() []string {
	if x != nil {
		return x.RequestDomain
	}
	return nil
}

func (x *DomainConfig) GetSocketDomain() []string {
	if x != nil {
		return x.SocketDomain
	}
	return nil
}

func (x *DomainConfig) GetUploadFileDomain() []string {
	if x != nil {
		return x.UploadFileDomain
	}
	return nil
}

func (x *DomainConfig) GetDownloadFileDomain() []string {
	if x != nil {
		return x.DownloadFileDomain
	}
	return nil
}

func (x *DomainConfig) GetBusinessDomain() []string {
	if x != nil {
		return x.BusinessDomain
	}
	return nil
}

func (x *DomainConfig) GetUdpIpList() []string {
	if x != nil {
		return x.UdpIpList
	}
	return nil
}
