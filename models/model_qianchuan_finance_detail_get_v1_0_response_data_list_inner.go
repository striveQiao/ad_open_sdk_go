/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFinanceDetailGetV10ResponseDataListInner struct for QianchuanFinanceDetailGetV10ResponseDataListInner
type QianchuanFinanceDetailGetV10ResponseDataListInner struct {
	// 非赠款余额
	CashBalance *int64 `json:"cash_balance,omitempty"`
	// 非赠款消耗
	CashCost *int64 `json:"cash_cost,omitempty"`
	// 余额总消耗
	Cost *int64 `json:"cost,omitempty"`
	// 日期，格式 2021-04-05
	Date *string `json:"date,omitempty"`
	// 消返红包消耗
	DeductionCost *int64 `json:"deduction_cost,omitempty"`
	// 赠款余额
	GrantBalance *int64 `json:"grant_balance,omitempty"`
	// 赠款消耗
	GrantCost *int64 `json:"grant_cost,omitempty"`
	// 总存入
	Income *int64 `json:"income,omitempty"`
	// 随心推非赠款消耗
	QcAwemeCashCost *int64 `json:"qc_aweme_cash_cost,omitempty"`
	// 随心推消耗
	QcAwemeCost *int64 `json:"qc_aweme_cost,omitempty"`
	// 随心推赠款消耗
	QcAwemeGrantCost *int64 `json:"qc_aweme_grant_cost,omitempty"`
	// 共享余额消耗
	ShareCost *int64 `json:"share_cost,omitempty"`
	// 共享钱包消耗
	ShareWalletCost *int64 `json:"share_wallet_cost,omitempty"`
	// 总余额
	TotalBalance *int64 `json:"total_balance,omitempty"`
	// 总转入
	TransferIn *int64 `json:"transfer_in,omitempty"`
	// 总转出
	TransferOut      *int64                                                `json:"transfer_out,omitempty"`
	ViewDeliveryType *QianchuanFinanceDetailGetV10DataListViewDeliveryType `json:"view_delivery_type,omitempty"`
}
