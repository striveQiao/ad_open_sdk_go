/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInnerTransferCapitalRecordListInner struct for CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInnerTransferCapitalRecordListInner
type CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInnerTransferCapitalRecordListInner struct {
	CapitalType *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListCapitalType `json:"capital_type,omitempty"`
	// 失败原因
	FailReason *string                                                                                         `json:"fail_reason,omitempty"`
	Platform   *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListPlatform `json:"platform,omitempty"`
	// 转账资金金额(单位：分)
	TransferAmount *int64                                                                                                `json:"transfer_amount,omitempty"`
	TransferStatus *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus `json:"transfer_status,omitempty"`
}
