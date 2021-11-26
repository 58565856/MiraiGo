// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: data.proto

package pb

type DeviceInfo struct {
	Bootloader   string `protobuf:"bytes,1,opt"`
	ProcVersion  string `protobuf:"bytes,2,opt"`
	Codename     string `protobuf:"bytes,3,opt"`
	Incremental  string `protobuf:"bytes,4,opt"`
	Fingerprint  string `protobuf:"bytes,5,opt"`
	BootId       string `protobuf:"bytes,6,opt"`
	AndroidId    string `protobuf:"bytes,7,opt"`
	BaseBand     string `protobuf:"bytes,8,opt"`
	InnerVersion string `protobuf:"bytes,9,opt"`
}

func (x *DeviceInfo) GetBootloader() string {
	if x != nil {
		return x.Bootloader
	}
	return ""
}

func (x *DeviceInfo) GetProcVersion() string {
	if x != nil {
		return x.ProcVersion
	}
	return ""
}

func (x *DeviceInfo) GetCodename() string {
	if x != nil {
		return x.Codename
	}
	return ""
}

func (x *DeviceInfo) GetIncremental() string {
	if x != nil {
		return x.Incremental
	}
	return ""
}

func (x *DeviceInfo) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *DeviceInfo) GetBootId() string {
	if x != nil {
		return x.BootId
	}
	return ""
}

func (x *DeviceInfo) GetAndroidId() string {
	if x != nil {
		return x.AndroidId
	}
	return ""
}

func (x *DeviceInfo) GetBaseBand() string {
	if x != nil {
		return x.BaseBand
	}
	return ""
}

func (x *DeviceInfo) GetInnerVersion() string {
	if x != nil {
		return x.InnerVersion
	}
	return ""
}

type RequestBody struct {
	RptConfigList []*ConfigSeq `protobuf:"bytes,1,rep"`
}

func (x *RequestBody) GetRptConfigList() []*ConfigSeq {
	if x != nil {
		return x.RptConfigList
	}
	return nil
}

type ConfigSeq struct {
	Type    int32 `protobuf:"varint,1,opt"`
	Version int32 `protobuf:"varint,2,opt"`
}

func (x *ConfigSeq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ConfigSeq) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type D50ReqBody struct {
	Appid                   int64   `protobuf:"varint,1,opt"`
	MaxPkgSize              int32   `protobuf:"varint,2,opt"`
	StartTime               int32   `protobuf:"varint,3,opt"`
	StartIndex              int32   `protobuf:"varint,4,opt"`
	ReqNum                  int32   `protobuf:"varint,5,opt"`
	UinList                 []int64 `protobuf:"varint,6,rep"`
	ReqMusicSwitch          int32   `protobuf:"varint,91001,opt"`
	ReqMutualmarkAlienation int32   `protobuf:"varint,101001,opt"`
	ReqMutualmarkScore      int32   `protobuf:"varint,141001,opt"`
	ReqKsingSwitch          int32   `protobuf:"varint,151001,opt"`
	ReqMutualmarkLbsshare   int32   `protobuf:"varint,181001,opt"`
}

func (x *D50ReqBody) GetAppid() int64 {
	if x != nil {
		return x.Appid
	}
	return 0
}

func (x *D50ReqBody) GetMaxPkgSize() int32 {
	if x != nil {
		return x.MaxPkgSize
	}
	return 0
}

func (x *D50ReqBody) GetStartTime() int32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *D50ReqBody) GetStartIndex() int32 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *D50ReqBody) GetReqNum() int32 {
	if x != nil {
		return x.ReqNum
	}
	return 0
}

func (x *D50ReqBody) GetUinList() []int64 {
	if x != nil {
		return x.UinList
	}
	return nil
}

func (x *D50ReqBody) GetReqMusicSwitch() int32 {
	if x != nil {
		return x.ReqMusicSwitch
	}
	return 0
}

func (x *D50ReqBody) GetReqMutualmarkAlienation() int32 {
	if x != nil {
		return x.ReqMutualmarkAlienation
	}
	return 0
}

func (x *D50ReqBody) GetReqMutualmarkScore() int32 {
	if x != nil {
		return x.ReqMutualmarkScore
	}
	return 0
}

func (x *D50ReqBody) GetReqKsingSwitch() int32 {
	if x != nil {
		return x.ReqKsingSwitch
	}
	return 0
}

func (x *D50ReqBody) GetReqMutualmarkLbsshare() int32 {
	if x != nil {
		return x.ReqMutualmarkLbsshare
	}
	return 0
}

type ReqDataHighwayHead struct {
	MsgBasehead   *DataHighwayHead `protobuf:"bytes,1,opt"`
	MsgSeghead    *SegHead         `protobuf:"bytes,2,opt"`
	ReqExtendinfo []byte           `protobuf:"bytes,3,opt"`
	Timestamp     int64            `protobuf:"varint,4,opt"` //LoginSigHead? msgLoginSigHead = 5;
}

func (x *ReqDataHighwayHead) GetMsgBasehead() *DataHighwayHead {
	if x != nil {
		return x.MsgBasehead
	}
	return nil
}

func (x *ReqDataHighwayHead) GetMsgSeghead() *SegHead {
	if x != nil {
		return x.MsgSeghead
	}
	return nil
}

func (x *ReqDataHighwayHead) GetReqExtendinfo() []byte {
	if x != nil {
		return x.ReqExtendinfo
	}
	return nil
}

func (x *ReqDataHighwayHead) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type RspDataHighwayHead struct {
	MsgBasehead   *DataHighwayHead `protobuf:"bytes,1,opt"`
	MsgSeghead    *SegHead         `protobuf:"bytes,2,opt"`
	ErrorCode     int32            `protobuf:"varint,3,opt"`
	AllowRetry    int32            `protobuf:"varint,4,opt"`
	Cachecost     int32            `protobuf:"varint,5,opt"`
	Htcost        int32            `protobuf:"varint,6,opt"`
	RspExtendinfo []byte           `protobuf:"bytes,7,opt"`
	Timestamp     int64            `protobuf:"varint,8,opt"`
	Range         int64            `protobuf:"varint,9,opt"`
	IsReset       int32            `protobuf:"varint,10,opt"`
}

func (x *RspDataHighwayHead) GetMsgBasehead() *DataHighwayHead {
	if x != nil {
		return x.MsgBasehead
	}
	return nil
}

func (x *RspDataHighwayHead) GetMsgSeghead() *SegHead {
	if x != nil {
		return x.MsgSeghead
	}
	return nil
}

func (x *RspDataHighwayHead) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *RspDataHighwayHead) GetAllowRetry() int32 {
	if x != nil {
		return x.AllowRetry
	}
	return 0
}

func (x *RspDataHighwayHead) GetCachecost() int32 {
	if x != nil {
		return x.Cachecost
	}
	return 0
}

func (x *RspDataHighwayHead) GetHtcost() int32 {
	if x != nil {
		return x.Htcost
	}
	return 0
}

func (x *RspDataHighwayHead) GetRspExtendinfo() []byte {
	if x != nil {
		return x.RspExtendinfo
	}
	return nil
}

func (x *RspDataHighwayHead) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RspDataHighwayHead) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *RspDataHighwayHead) GetIsReset() int32 {
	if x != nil {
		return x.IsReset
	}
	return 0
}

type DataHighwayHead struct {
	Version    int32  `protobuf:"varint,1,opt"`
	Uin        string `protobuf:"bytes,2,opt"`
	Command    string `protobuf:"bytes,3,opt"`
	Seq        int32  `protobuf:"varint,4,opt"`
	RetryTimes int32  `protobuf:"varint,5,opt"`
	Appid      int32  `protobuf:"varint,6,opt"`
	Dataflag   int32  `protobuf:"varint,7,opt"`
	CommandId  int32  `protobuf:"varint,8,opt"`
	BuildVer   string `protobuf:"bytes,9,opt"`
	LocaleId   int32  `protobuf:"varint,10,opt"`
}

func (x *DataHighwayHead) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DataHighwayHead) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

func (x *DataHighwayHead) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *DataHighwayHead) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *DataHighwayHead) GetRetryTimes() int32 {
	if x != nil {
		return x.RetryTimes
	}
	return 0
}

func (x *DataHighwayHead) GetAppid() int32 {
	if x != nil {
		return x.Appid
	}
	return 0
}

func (x *DataHighwayHead) GetDataflag() int32 {
	if x != nil {
		return x.Dataflag
	}
	return 0
}

func (x *DataHighwayHead) GetCommandId() int32 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

func (x *DataHighwayHead) GetBuildVer() string {
	if x != nil {
		return x.BuildVer
	}
	return ""
}

func (x *DataHighwayHead) GetLocaleId() int32 {
	if x != nil {
		return x.LocaleId
	}
	return 0
}

type SegHead struct {
	Serviceid     int32  `protobuf:"varint,1,opt"`
	Filesize      int64  `protobuf:"varint,2,opt"`
	Dataoffset    int64  `protobuf:"varint,3,opt"`
	Datalength    int32  `protobuf:"varint,4,opt"`
	Rtcode        int32  `protobuf:"varint,5,opt"`
	Serviceticket []byte `protobuf:"bytes,6,opt"`
	Flag          int32  `protobuf:"varint,7,opt"`
	Md5           []byte `protobuf:"bytes,8,opt"`
	FileMd5       []byte `protobuf:"bytes,9,opt"`
	CacheAddr     int32  `protobuf:"varint,10,opt"`
	QueryTimes    int32  `protobuf:"varint,11,opt"`
	UpdateCacheip int32  `protobuf:"varint,12,opt"`
}

func (x *SegHead) GetServiceid() int32 {
	if x != nil {
		return x.Serviceid
	}
	return 0
}

func (x *SegHead) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

func (x *SegHead) GetDataoffset() int64 {
	if x != nil {
		return x.Dataoffset
	}
	return 0
}

func (x *SegHead) GetDatalength() int32 {
	if x != nil {
		return x.Datalength
	}
	return 0
}

func (x *SegHead) GetRtcode() int32 {
	if x != nil {
		return x.Rtcode
	}
	return 0
}

func (x *SegHead) GetServiceticket() []byte {
	if x != nil {
		return x.Serviceticket
	}
	return nil
}

func (x *SegHead) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *SegHead) GetMd5() []byte {
	if x != nil {
		return x.Md5
	}
	return nil
}

func (x *SegHead) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *SegHead) GetCacheAddr() int32 {
	if x != nil {
		return x.CacheAddr
	}
	return 0
}

func (x *SegHead) GetQueryTimes() int32 {
	if x != nil {
		return x.QueryTimes
	}
	return 0
}

func (x *SegHead) GetUpdateCacheip() int32 {
	if x != nil {
		return x.UpdateCacheip
	}
	return 0
}

type DeleteMessageRequest struct {
	Items []*MessageItem `protobuf:"bytes,1,rep"`
}

func (x *DeleteMessageRequest) GetItems() []*MessageItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type MessageItem struct {
	FromUin int64  `protobuf:"varint,1,opt"`
	ToUin   int64  `protobuf:"varint,2,opt"`
	MsgType int32  `protobuf:"varint,3,opt"`
	MsgSeq  int32  `protobuf:"varint,4,opt"`
	MsgUid  int64  `protobuf:"varint,5,opt"`
	Sig     []byte `protobuf:"bytes,7,opt"`
}

func (x *MessageItem) GetFromUin() int64 {
	if x != nil {
		return x.FromUin
	}
	return 0
}

func (x *MessageItem) GetToUin() int64 {
	if x != nil {
		return x.ToUin
	}
	return 0
}

func (x *MessageItem) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MessageItem) GetMsgSeq() int32 {
	if x != nil {
		return x.MsgSeq
	}
	return 0
}

func (x *MessageItem) GetMsgUid() int64 {
	if x != nil {
		return x.MsgUid
	}
	return 0
}

func (x *MessageItem) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type SubD4 struct {
	Uin int64 `protobuf:"varint,1,opt"`
}

func (x *SubD4) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

type Sub8A struct {
	MsgInfo         []*Sub8AMsgInfo `protobuf:"bytes,1,rep"`
	AppId           int32           `protobuf:"varint,2,opt"`
	InstId          int32           `protobuf:"varint,3,opt"`
	LongMessageFlag int32           `protobuf:"varint,4,opt"`
	Reserved        []byte          `protobuf:"bytes,5,opt"`
}

func (x *Sub8A) GetMsgInfo() []*Sub8AMsgInfo {
	if x != nil {
		return x.MsgInfo
	}
	return nil
}

func (x *Sub8A) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *Sub8A) GetInstId() int32 {
	if x != nil {
		return x.InstId
	}
	return 0
}

func (x *Sub8A) GetLongMessageFlag() int32 {
	if x != nil {
		return x.LongMessageFlag
	}
	return 0
}

func (x *Sub8A) GetReserved() []byte {
	if x != nil {
		return x.Reserved
	}
	return nil
}

type Sub8AMsgInfo struct {
	FromUin   int64 `protobuf:"varint,1,opt"`
	ToUin     int64 `protobuf:"varint,2,opt"`
	MsgSeq    int32 `protobuf:"varint,3,opt"`
	MsgUid    int64 `protobuf:"varint,4,opt"`
	MsgTime   int64 `protobuf:"varint,5,opt"`
	MsgRandom int32 `protobuf:"varint,6,opt"`
	PkgNum    int32 `protobuf:"varint,7,opt"`
	PkgIndex  int32 `protobuf:"varint,8,opt"`
	DevSeq    int32 `protobuf:"varint,9,opt"`
}

func (x *Sub8AMsgInfo) GetFromUin() int64 {
	if x != nil {
		return x.FromUin
	}
	return 0
}

func (x *Sub8AMsgInfo) GetToUin() int64 {
	if x != nil {
		return x.ToUin
	}
	return 0
}

func (x *Sub8AMsgInfo) GetMsgSeq() int32 {
	if x != nil {
		return x.MsgSeq
	}
	return 0
}

func (x *Sub8AMsgInfo) GetMsgUid() int64 {
	if x != nil {
		return x.MsgUid
	}
	return 0
}

func (x *Sub8AMsgInfo) GetMsgTime() int64 {
	if x != nil {
		return x.MsgTime
	}
	return 0
}

func (x *Sub8AMsgInfo) GetMsgRandom() int32 {
	if x != nil {
		return x.MsgRandom
	}
	return 0
}

func (x *Sub8AMsgInfo) GetPkgNum() int32 {
	if x != nil {
		return x.PkgNum
	}
	return 0
}

func (x *Sub8AMsgInfo) GetPkgIndex() int32 {
	if x != nil {
		return x.PkgIndex
	}
	return 0
}

func (x *Sub8AMsgInfo) GetDevSeq() int32 {
	if x != nil {
		return x.DevSeq
	}
	return 0
}

type SubB3 struct {
	Type            int32              `protobuf:"varint,1,opt"`
	MsgAddFrdNotify *SubB3AddFrdNotify `protobuf:"bytes,2,opt"`
}

func (x *SubB3) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SubB3) GetMsgAddFrdNotify() *SubB3AddFrdNotify {
	if x != nil {
		return x.MsgAddFrdNotify
	}
	return nil
}

type SubB3AddFrdNotify struct {
	Uin  int64  `protobuf:"varint,1,opt"`
	Nick string `protobuf:"bytes,5,opt"`
}

func (x *SubB3AddFrdNotify) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *SubB3AddFrdNotify) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

type Sub44 struct {
	FriendSyncMsg *Sub44FriendSyncMsg `protobuf:"bytes,1,opt"`
	GroupSyncMsg  *Sub44GroupSyncMsg  `protobuf:"bytes,2,opt"`
}

func (x *Sub44) GetFriendSyncMsg() *Sub44FriendSyncMsg {
	if x != nil {
		return x.FriendSyncMsg
	}
	return nil
}

func (x *Sub44) GetGroupSyncMsg() *Sub44GroupSyncMsg {
	if x != nil {
		return x.GroupSyncMsg
	}
	return nil
}

type Sub44FriendSyncMsg struct {
	Uin         int64    `protobuf:"varint,1,opt"`
	FUin        int64    `protobuf:"varint,2,opt"`
	ProcessType int32    `protobuf:"varint,3,opt"`
	Time        int32    `protobuf:"varint,4,opt"`
	ProcessFlag int32    `protobuf:"varint,5,opt"`
	SourceId    int32    `protobuf:"varint,6,opt"`
	SourceSubId int32    `protobuf:"varint,7,opt"`
	StrWording  []string `protobuf:"bytes,8,rep"`
}

func (x *Sub44FriendSyncMsg) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetFUin() int64 {
	if x != nil {
		return x.FUin
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetProcessType() int32 {
	if x != nil {
		return x.ProcessType
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetProcessFlag() int32 {
	if x != nil {
		return x.ProcessFlag
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetSourceId() int32 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetSourceSubId() int32 {
	if x != nil {
		return x.SourceSubId
	}
	return 0
}

func (x *Sub44FriendSyncMsg) GetStrWording() []string {
	if x != nil {
		return x.StrWording
	}
	return nil
}

type Sub44GroupSyncMsg struct {
	MsgType         int32  `protobuf:"varint,1,opt"`
	MsgSeq          int64  `protobuf:"varint,2,opt"`
	GrpCode         int64  `protobuf:"varint,3,opt"`
	GaCode          int64  `protobuf:"varint,4,opt"`
	OptUin1         int64  `protobuf:"varint,5,opt"`
	OptUin2         int64  `protobuf:"varint,6,opt"`
	MsgBuf          []byte `protobuf:"bytes,7,opt"`
	AuthKey         []byte `protobuf:"bytes,8,opt"`
	MsgStatus       int32  `protobuf:"varint,9,opt"`
	ActionUin       int64  `protobuf:"varint,10,opt"`
	ActionTime      int64  `protobuf:"varint,11,opt"`
	CurMaxMemCount  int32  `protobuf:"varint,12,opt"`
	NextMaxMemCount int32  `protobuf:"varint,13,opt"`
	CurMemCount     int32  `protobuf:"varint,14,opt"`
	ReqSrcId        int32  `protobuf:"varint,15,opt"`
	ReqSrcSubId     int32  `protobuf:"varint,16,opt"`
	InviterRole     int32  `protobuf:"varint,17,opt"`
	ExtAdminNum     int32  `protobuf:"varint,18,opt"`
	ProcessFlag     int32  `protobuf:"varint,19,opt"`
}

func (x *Sub44GroupSyncMsg) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetMsgSeq() int64 {
	if x != nil {
		return x.MsgSeq
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetGrpCode() int64 {
	if x != nil {
		return x.GrpCode
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetGaCode() int64 {
	if x != nil {
		return x.GaCode
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetOptUin1() int64 {
	if x != nil {
		return x.OptUin1
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetOptUin2() int64 {
	if x != nil {
		return x.OptUin2
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetMsgBuf() []byte {
	if x != nil {
		return x.MsgBuf
	}
	return nil
}

func (x *Sub44GroupSyncMsg) GetAuthKey() []byte {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *Sub44GroupSyncMsg) GetMsgStatus() int32 {
	if x != nil {
		return x.MsgStatus
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetActionUin() int64 {
	if x != nil {
		return x.ActionUin
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetActionTime() int64 {
	if x != nil {
		return x.ActionTime
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetCurMaxMemCount() int32 {
	if x != nil {
		return x.CurMaxMemCount
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetNextMaxMemCount() int32 {
	if x != nil {
		return x.NextMaxMemCount
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetCurMemCount() int32 {
	if x != nil {
		return x.CurMemCount
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetReqSrcId() int32 {
	if x != nil {
		return x.ReqSrcId
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetReqSrcSubId() int32 {
	if x != nil {
		return x.ReqSrcSubId
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetInviterRole() int32 {
	if x != nil {
		return x.InviterRole
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetExtAdminNum() int32 {
	if x != nil {
		return x.ExtAdminNum
	}
	return 0
}

func (x *Sub44GroupSyncMsg) GetProcessFlag() int32 {
	if x != nil {
		return x.ProcessFlag
	}
	return 0
}

type GroupMemberReqBody struct {
	GroupCode       int64 `protobuf:"varint,1,opt"`
	Uin             int64 `protobuf:"varint,2,opt"`
	NewClient       bool  `protobuf:"varint,3,opt"`
	ClientType      int32 `protobuf:"varint,4,opt"`
	RichCardNameVer int32 `protobuf:"varint,5,opt"`
}

func (x *GroupMemberReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *GroupMemberReqBody) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *GroupMemberReqBody) GetNewClient() bool {
	if x != nil {
		return x.NewClient
	}
	return false
}

func (x *GroupMemberReqBody) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *GroupMemberReqBody) GetRichCardNameVer() int32 {
	if x != nil {
		return x.RichCardNameVer
	}
	return 0
}

type GroupMemberRspBody struct {
	GroupCode              int64            `protobuf:"varint,1,opt"`
	SelfRole               int32            `protobuf:"varint,2,opt"`
	MemInfo                *GroupMemberInfo `protobuf:"bytes,3,opt"`
	BoolSelfLocationShared bool             `protobuf:"varint,4,opt"`
	GroupType              int32            `protobuf:"varint,5,opt"`
}

func (x *GroupMemberRspBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *GroupMemberRspBody) GetSelfRole() int32 {
	if x != nil {
		return x.SelfRole
	}
	return 0
}

func (x *GroupMemberRspBody) GetMemInfo() *GroupMemberInfo {
	if x != nil {
		return x.MemInfo
	}
	return nil
}

func (x *GroupMemberRspBody) GetBoolSelfLocationShared() bool {
	if x != nil {
		return x.BoolSelfLocationShared
	}
	return false
}

func (x *GroupMemberRspBody) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

type GroupMemberInfo struct {
	Uin         int64  `protobuf:"varint,1,opt"`
	Result      int32  `protobuf:"varint,2,opt"`
	Errmsg      []byte `protobuf:"bytes,3,opt"`
	IsFriend    bool   `protobuf:"varint,4,opt"`
	Remark      []byte `protobuf:"bytes,5,opt"`
	IsConcerned bool   `protobuf:"varint,6,opt"`
	Credit      int32  `protobuf:"varint,7,opt"`
	Card        []byte `protobuf:"bytes,8,opt"`
	Sex         int32  `protobuf:"varint,9,opt"`
	Location    []byte `protobuf:"bytes,10,opt"`
	Nick        []byte `protobuf:"bytes,11,opt"`
	Age         int32  `protobuf:"varint,12,opt"`
	Lev         []byte `protobuf:"bytes,13,opt"`
	Join        int64  `protobuf:"varint,14,opt"`
	LastSpeak   int64  `protobuf:"varint,15,opt"`
	//repeated CustomEntry customEnties = 16;
	//repeated GBarInfo gbarConcerned = 17;
	GbarTitle              []byte `protobuf:"bytes,18,opt"`
	GbarUrl                []byte `protobuf:"bytes,19,opt"`
	GbarCnt                int32  `protobuf:"varint,20,opt"`
	IsAllowModCard         bool   `protobuf:"varint,21,opt"`
	IsVip                  bool   `protobuf:"varint,22,opt"`
	IsYearVip              bool   `protobuf:"varint,23,opt"`
	IsSuperVip             bool   `protobuf:"varint,24,opt"`
	IsSuperQq              bool   `protobuf:"varint,25,opt"`
	VipLev                 int32  `protobuf:"varint,26,opt"`
	Role                   int32  `protobuf:"varint,27,opt"`
	LocationShared         bool   `protobuf:"varint,28,opt"`
	Int64Distance          int64  `protobuf:"varint,29,opt"`
	ConcernType            int32  `protobuf:"varint,30,opt"`
	SpecialTitle           []byte `protobuf:"bytes,31,opt"`
	SpecialTitleExpireTime int32  `protobuf:"varint,32,opt"`
	//FlowersEntry flowerEntry = 33;
	//TeamEntry teamEntry = 34;
	PhoneNum []byte `protobuf:"bytes,35,opt"`
	Job      []byte `protobuf:"bytes,36,opt"`
	MedalId  int32  `protobuf:"varint,37,opt"`
	Level    int32  `protobuf:"varint,39,opt"`
	Honor    string `protobuf:"bytes,41,opt"`
}

func (x *GroupMemberInfo) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *GroupMemberInfo) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *GroupMemberInfo) GetErrmsg() []byte {
	if x != nil {
		return x.Errmsg
	}
	return nil
}

func (x *GroupMemberInfo) GetIsFriend() bool {
	if x != nil {
		return x.IsFriend
	}
	return false
}

func (x *GroupMemberInfo) GetRemark() []byte {
	if x != nil {
		return x.Remark
	}
	return nil
}

func (x *GroupMemberInfo) GetIsConcerned() bool {
	if x != nil {
		return x.IsConcerned
	}
	return false
}

func (x *GroupMemberInfo) GetCredit() int32 {
	if x != nil {
		return x.Credit
	}
	return 0
}

func (x *GroupMemberInfo) GetCard() []byte {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *GroupMemberInfo) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *GroupMemberInfo) GetLocation() []byte {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *GroupMemberInfo) GetNick() []byte {
	if x != nil {
		return x.Nick
	}
	return nil
}

func (x *GroupMemberInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *GroupMemberInfo) GetLev() []byte {
	if x != nil {
		return x.Lev
	}
	return nil
}

func (x *GroupMemberInfo) GetJoin() int64 {
	if x != nil {
		return x.Join
	}
	return 0
}

func (x *GroupMemberInfo) GetLastSpeak() int64 {
	if x != nil {
		return x.LastSpeak
	}
	return 0
}

func (x *GroupMemberInfo) GetGbarTitle() []byte {
	if x != nil {
		return x.GbarTitle
	}
	return nil
}

func (x *GroupMemberInfo) GetGbarUrl() []byte {
	if x != nil {
		return x.GbarUrl
	}
	return nil
}

func (x *GroupMemberInfo) GetGbarCnt() int32 {
	if x != nil {
		return x.GbarCnt
	}
	return 0
}

func (x *GroupMemberInfo) GetIsAllowModCard() bool {
	if x != nil {
		return x.IsAllowModCard
	}
	return false
}

func (x *GroupMemberInfo) GetIsVip() bool {
	if x != nil {
		return x.IsVip
	}
	return false
}

func (x *GroupMemberInfo) GetIsYearVip() bool {
	if x != nil {
		return x.IsYearVip
	}
	return false
}

func (x *GroupMemberInfo) GetIsSuperVip() bool {
	if x != nil {
		return x.IsSuperVip
	}
	return false
}

func (x *GroupMemberInfo) GetIsSuperQq() bool {
	if x != nil {
		return x.IsSuperQq
	}
	return false
}

func (x *GroupMemberInfo) GetVipLev() int32 {
	if x != nil {
		return x.VipLev
	}
	return 0
}

func (x *GroupMemberInfo) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *GroupMemberInfo) GetLocationShared() bool {
	if x != nil {
		return x.LocationShared
	}
	return false
}

func (x *GroupMemberInfo) GetInt64Distance() int64 {
	if x != nil {
		return x.Int64Distance
	}
	return 0
}

func (x *GroupMemberInfo) GetConcernType() int32 {
	if x != nil {
		return x.ConcernType
	}
	return 0
}

func (x *GroupMemberInfo) GetSpecialTitle() []byte {
	if x != nil {
		return x.SpecialTitle
	}
	return nil
}

func (x *GroupMemberInfo) GetSpecialTitleExpireTime() int32 {
	if x != nil {
		return x.SpecialTitleExpireTime
	}
	return 0
}

func (x *GroupMemberInfo) GetPhoneNum() []byte {
	if x != nil {
		return x.PhoneNum
	}
	return nil
}

func (x *GroupMemberInfo) GetJob() []byte {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *GroupMemberInfo) GetMedalId() int32 {
	if x != nil {
		return x.MedalId
	}
	return 0
}

func (x *GroupMemberInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GroupMemberInfo) GetHonor() string {
	if x != nil {
		return x.Honor
	}
	return ""
}
