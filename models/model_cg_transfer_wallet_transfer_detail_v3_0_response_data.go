/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferDetailV30ResponseData
type CgTransferWalletTransferDetailV30ResponseData struct {
	// 幂等id
	BizRequestNo *string `json:"biz_request_no,omitempty"`
	// 转账总金额(单位：分)
	TransferAmount *int64 `json:"transfer_amount,omitempty"`
	// 转账创建时间yyyy-MM-dd HH:mm:ss
	TransferCreateTime *string                                                 `json:"transfer_create_time,omitempty"`
	TransferDirection  *CgTransferWalletTransferDetailV30DataTransferDirection `json:"transfer_direction,omitempty"`
	// 转账完成时间 yyyy-MM-dd HH:mm:ss
	TransferFinishTime *string `json:"transfer_finish_time,omitempty"`
	// 转账单号
	TransferSerial *string                                              `json:"transfer_serial,omitempty"`
	TransferStatus *CgTransferWalletTransferDetailV30DataTransferStatus `json:"transfer_status,omitempty"`
	// 账户信息列表
	TransferWalletRecordList []*CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInner `json:"transfer_wallet_record_list,omitempty"`
}
