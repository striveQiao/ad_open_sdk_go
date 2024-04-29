/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryTransferDetailV30ResponseData
type CgTransferQueryTransferDetailV30ResponseData struct {
	// 幂等id
	BizRequestNo *string `json:"biz_request_no,omitempty"`
	// 转账总金额(单位：分)
	TransferAmount *int64 `json:"transfer_amount,omitempty"`
	// 转账创建时间
	TransferCreateTime *string                                                `json:"transfer_create_time,omitempty"`
	TransferDirection  *CgTransferQueryTransferDetailV30DataTransferDirection `json:"transfer_direction,omitempty"`
	// 转账完成时间
	TransferFinishTime *string `json:"transfer_finish_time,omitempty"`
	// 转账单号
	TransferSerial *string                                             `json:"transfer_serial,omitempty"`
	TransferStatus *CgTransferQueryTransferDetailV30DataTransferStatus `json:"transfer_status,omitempty"`
	// 账户信息列表
	TransferTargetRecordList []*CgTransferQueryTransferDetailV30ResponseDataTransferTargetRecordListInner `json:"transfer_target_record_list,omitempty"`
}
