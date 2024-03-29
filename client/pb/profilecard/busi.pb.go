// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/profilecard/busi.proto

package profilecard

type BusiColor struct {
	R *int32 `protobuf:"varint,1,opt"`
	G *int32 `protobuf:"varint,2,opt"`
	B *int32 `protobuf:"varint,3,opt"`
}

func (x *BusiColor) GetR() int32 {
	if x != nil && x.R != nil {
		return *x.R
	}
	return 0
}

func (x *BusiColor) GetG() int32 {
	if x != nil && x.G != nil {
		return *x.G
	}
	return 0
}

func (x *BusiColor) GetB() int32 {
	if x != nil && x.B != nil {
		return *x.B
	}
	return 0
}

type BusiComm struct {
	Ver            *int32        `protobuf:"varint,1,opt"`
	Seq            *int32        `protobuf:"varint,2,opt"`
	Fromuin        *int64        `protobuf:"varint,3,opt"`
	Touin          *int64        `protobuf:"varint,4,opt"`
	Service        *int32        `protobuf:"varint,5,opt"`
	SessionType    *int32        `protobuf:"varint,6,opt"`
	SessionKey     []byte        `protobuf:"bytes,7,opt"`
	ClientIp       *int32        `protobuf:"varint,8,opt"`
	Display        *BusiUi       `protobuf:"bytes,9,opt"`
	Result         *int32        `protobuf:"varint,10,opt"`
	ErrMsg         *string       `protobuf:"bytes,11,opt"`
	Platform       *int32        `protobuf:"varint,12,opt"`
	Qqver          *string       `protobuf:"bytes,13,opt"`
	Build          *int32        `protobuf:"varint,14,opt"`
	MsgLoginSig    *BusiLoginSig `protobuf:"bytes,15,opt"`
	Version        *int32        `protobuf:"varint,17,opt"`
	MsgUinInfo     *BusiUinInfo  `protobuf:"bytes,18,opt"`
	MsgRichDisplay *BusiRichUi   `protobuf:"bytes,19,opt"`
}

func (x *BusiComm) GetVer() int32 {
	if x != nil && x.Ver != nil {
		return *x.Ver
	}
	return 0
}

func (x *BusiComm) GetSeq() int32 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *BusiComm) GetFromuin() int64 {
	if x != nil && x.Fromuin != nil {
		return *x.Fromuin
	}
	return 0
}

func (x *BusiComm) GetTouin() int64 {
	if x != nil && x.Touin != nil {
		return *x.Touin
	}
	return 0
}

func (x *BusiComm) GetService() int32 {
	if x != nil && x.Service != nil {
		return *x.Service
	}
	return 0
}

func (x *BusiComm) GetSessionType() int32 {
	if x != nil && x.SessionType != nil {
		return *x.SessionType
	}
	return 0
}

func (x *BusiComm) GetClientIp() int32 {
	if x != nil && x.ClientIp != nil {
		return *x.ClientIp
	}
	return 0
}

func (x *BusiComm) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *BusiComm) GetErrMsg() string {
	if x != nil && x.ErrMsg != nil {
		return *x.ErrMsg
	}
	return ""
}

func (x *BusiComm) GetPlatform() int32 {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return 0
}

func (x *BusiComm) GetQqver() string {
	if x != nil && x.Qqver != nil {
		return *x.Qqver
	}
	return ""
}

func (x *BusiComm) GetBuild() int32 {
	if x != nil && x.Build != nil {
		return *x.Build
	}
	return 0
}

func (x *BusiComm) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type BusiCommonReq struct {
	ServiceCmd *string              `protobuf:"bytes,1,opt"`
	VcReq      *BusiVisitorCountReq `protobuf:"bytes,2,opt"`
	HrReq      *BusiHideRecordsReq  `protobuf:"bytes,3,opt"`
}

func (x *BusiCommonReq) GetServiceCmd() string {
	if x != nil && x.ServiceCmd != nil {
		return *x.ServiceCmd
	}
	return ""
}

type BusiDetailRecord struct {
	Fuin     *int32 `protobuf:"varint,1,opt"`
	Source   *int32 `protobuf:"varint,2,opt"`
	Vtime    *int32 `protobuf:"varint,3,opt"`
	Mod      *int32 `protobuf:"varint,4,opt"`
	HideFlag *int32 `protobuf:"varint,5,opt"`
}

func (x *BusiDetailRecord) GetFuin() int32 {
	if x != nil && x.Fuin != nil {
		return *x.Fuin
	}
	return 0
}

func (x *BusiDetailRecord) GetSource() int32 {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return 0
}

func (x *BusiDetailRecord) GetVtime() int32 {
	if x != nil && x.Vtime != nil {
		return *x.Vtime
	}
	return 0
}

func (x *BusiDetailRecord) GetMod() int32 {
	if x != nil && x.Mod != nil {
		return *x.Mod
	}
	return 0
}

func (x *BusiDetailRecord) GetHideFlag() int32 {
	if x != nil && x.HideFlag != nil {
		return *x.HideFlag
	}
	return 0
}

type BusiHideRecordsReq struct {
	Huin    *int32              `protobuf:"varint,1,opt"`
	Fuin    *int32              `protobuf:"varint,2,opt"`
	Records []*BusiDetailRecord `protobuf:"bytes,3,rep"`
}

func (x *BusiHideRecordsReq) GetHuin() int32 {
	if x != nil && x.Huin != nil {
		return *x.Huin
	}
	return 0
}

func (x *BusiHideRecordsReq) GetFuin() int32 {
	if x != nil && x.Fuin != nil {
		return *x.Fuin
	}
	return 0
}

type BusiLabel struct {
	Name        []byte     `protobuf:"bytes,1,opt"`
	EnumType    *int32     `protobuf:"varint,2,opt"`
	TextColor   *BusiColor `protobuf:"bytes,3,opt"`
	EdgingColor *BusiColor `protobuf:"bytes,4,opt"`
	LabelAttr   *int32     `protobuf:"varint,5,opt"`
	LabelType   *int32     `protobuf:"varint,6,opt"`
}

func (x *BusiLabel) GetEnumType() int32 {
	if x != nil && x.EnumType != nil {
		return *x.EnumType
	}
	return 0
}

func (x *BusiLabel) GetLabelAttr() int32 {
	if x != nil && x.LabelAttr != nil {
		return *x.LabelAttr
	}
	return 0
}

func (x *BusiLabel) GetLabelType() int32 {
	if x != nil && x.LabelType != nil {
		return *x.LabelType
	}
	return 0
}

type BusiLoginSig struct {
	Type  *int32 `protobuf:"varint,1,opt"`
	Sig   []byte `protobuf:"bytes,2,opt"`
	Appid *int32 `protobuf:"varint,3,opt"`
}

func (x *BusiLoginSig) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *BusiLoginSig) GetAppid() int32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type BusiRichUi struct {
	Name       *string `protobuf:"bytes,1,opt"`
	ServiceUrl *string `protobuf:"bytes,2,opt"` //repeated UiInfo uiList = 3;
}

func (x *BusiRichUi) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BusiRichUi) GetServiceUrl() string {
	if x != nil && x.ServiceUrl != nil {
		return *x.ServiceUrl
	}
	return ""
}

type BusiUi struct {
	Url     *string `protobuf:"bytes,1,opt"`
	Title   *string `protobuf:"bytes,2,opt"`
	Content *string `protobuf:"bytes,3,opt"`
	JumpUrl *string `protobuf:"bytes,4,opt"`
}

func (x *BusiUi) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *BusiUi) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *BusiUi) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *BusiUi) GetJumpUrl() string {
	if x != nil && x.JumpUrl != nil {
		return *x.JumpUrl
	}
	return ""
}

type BusiUinInfo struct {
	Int64Longitude *int64 `protobuf:"varint,1,opt"`
	Int64Latitude  *int64 `protobuf:"varint,2,opt"`
}

func (x *BusiUinInfo) GetInt64Longitude() int64 {
	if x != nil && x.Int64Longitude != nil {
		return *x.Int64Longitude
	}
	return 0
}

func (x *BusiUinInfo) GetInt64Latitude() int64 {
	if x != nil && x.Int64Latitude != nil {
		return *x.Int64Latitude
	}
	return 0
}

type BusiVisitorCountReq struct {
	Requireuin *int32 `protobuf:"varint,1,opt"`
	Operuin    *int32 `protobuf:"varint,2,opt"`
	Mod        *int32 `protobuf:"varint,3,opt"`
	ReportFlag *int32 `protobuf:"varint,4,opt"`
}

func (x *BusiVisitorCountReq) GetRequireuin() int32 {
	if x != nil && x.Requireuin != nil {
		return *x.Requireuin
	}
	return 0
}

func (x *BusiVisitorCountReq) GetOperuin() int32 {
	if x != nil && x.Operuin != nil {
		return *x.Operuin
	}
	return 0
}

func (x *BusiVisitorCountReq) GetMod() int32 {
	if x != nil && x.Mod != nil {
		return *x.Mod
	}
	return 0
}

func (x *BusiVisitorCountReq) GetReportFlag() int32 {
	if x != nil && x.ReportFlag != nil {
		return *x.ReportFlag
	}
	return 0
}

type BusiVisitorCountRsp struct {
	Requireuin *int32 `protobuf:"varint,1,opt"`
	TotalLike  *int32 `protobuf:"varint,2,opt"`
	TotalView  *int32 `protobuf:"varint,3,opt"`
	HotValue   *int32 `protobuf:"varint,4,opt"`
	RedValue   *int32 `protobuf:"varint,5,opt"`
	HotDiff    *int32 `protobuf:"varint,6,opt"`
}

func (x *BusiVisitorCountRsp) GetRequireuin() int32 {
	if x != nil && x.Requireuin != nil {
		return *x.Requireuin
	}
	return 0
}

func (x *BusiVisitorCountRsp) GetTotalLike() int32 {
	if x != nil && x.TotalLike != nil {
		return *x.TotalLike
	}
	return 0
}

func (x *BusiVisitorCountRsp) GetTotalView() int32 {
	if x != nil && x.TotalView != nil {
		return *x.TotalView
	}
	return 0
}

func (x *BusiVisitorCountRsp) GetHotValue() int32 {
	if x != nil && x.HotValue != nil {
		return *x.HotValue
	}
	return 0
}

func (x *BusiVisitorCountRsp) GetRedValue() int32 {
	if x != nil && x.RedValue != nil {
		return *x.RedValue
	}
	return 0
}

func (x *BusiVisitorCountRsp) GetHotDiff() int32 {
	if x != nil && x.HotDiff != nil {
		return *x.HotDiff
	}
	return 0
}
