// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0x990.proto

package oidb

type TranslateReqBody struct {
	// TranslateReq translate_req = 1;
	BatchTranslateReq *BatchTranslateReq `protobuf:"bytes,2,opt"`
}

type TranslateRspBody struct {
	// TranslateRsp translate_rsp = 1;
	BatchTranslateRsp *BatchTranslateRsp `protobuf:"bytes,2,opt"`
}

type BatchTranslateReq struct {
	SrcLanguage string   `protobuf:"bytes,1,opt"`
	DstLanguage string   `protobuf:"bytes,2,opt"`
	SrcTextList []string `protobuf:"bytes,3,rep"`
}

type BatchTranslateRsp struct {
	ErrorCode   int32    `protobuf:"varint,1,opt"`
	ErrorMsg    []byte   `protobuf:"bytes,2,opt"`
	SrcLanguage string   `protobuf:"bytes,3,opt"`
	DstLanguage string   `protobuf:"bytes,4,opt"`
	SrcTextList []string `protobuf:"bytes,5,rep"`
	DstTextList []string `protobuf:"bytes,6,rep"`
}
