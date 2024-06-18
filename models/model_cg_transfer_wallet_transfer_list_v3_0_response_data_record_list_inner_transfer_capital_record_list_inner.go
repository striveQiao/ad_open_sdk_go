/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferListV30ResponseDataRecordListInnerTransferCapitalRecordListInner struct for CgTransferWalletTransferListV30ResponseDataRecordListInnerTransferCapitalRecordListInner
type CgTransferWalletTransferListV30ResponseDataRecordListInnerTransferCapitalRecordListInner struct {
	// 转账资金创建时间yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"create_time,omitempty"`
	// 转账资金完成时间yyyy-MM-dd HH:mm:ss
	FinishTime *string                                                                         `json:"finish_time,omitempty"`
	Platform   *CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListPlatform `json:"platform,omitempty"`
	// 转账金额(单位：分)
	TransferAmount  *int64                                                                                 `json:"transfer_amount,omitempty"`
	TransferCapital *CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferCapital `json:"transfer_capital,omitempty"`
	TransferStatus  *CgTransferWalletTransferListV30DataRecordListTransferCapitalRecordListTransferStatus  `json:"transfer_status,omitempty"`
}
